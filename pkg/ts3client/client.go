package ts3client

import (
	"bufio"
	"context"
	"fmt"
	"github.com/patrick246/teamspeak-telegram-bot/pkg/ts3client/commands"
	"log"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"
)

type Connection struct {
	conn           net.Conn
	connMutex      sync.Mutex
	reader         *bufio.Reader
	done           chan struct{}
	dataErrorLines chan string
	eventListeners map[int64]chan Event
	readErrors     chan error
	state          connectionState
}

type connectionState struct {
	addr     string
	username string
	password string
	events   map[commands.EventType]struct{}
}

func Connect(addr string) (*Connection, error) {
	conn, reader, err := connect(addr)
	if err != nil {
		return nil, err
	}
	ts3conn := &Connection{
		conn:           conn,
		reader:         reader,
		done:           make(chan struct{}),
		dataErrorLines: make(chan string),
		readErrors:     make(chan error),
		eventListeners: make(map[int64]chan Event),
		state: connectionState{
			addr:   addr,
			events: make(map[commands.EventType]struct{}),
		},
	}
	go ts3conn.readResponses()
	go ts3conn.keepalive()
	return ts3conn, nil
}

func (c *Connection) Login(ctx context.Context, username, password string) error {
	_, err := c.Command(ctx, commands.Login(username, password))
	if err != nil {
		return err
	}
	c.state.username = username
	c.state.password = password
	return nil
}

func (c *Connection) Listen(ctx context.Context, event commands.EventType, id *string) error {
	_, err := c.Command(ctx, commands.ServerNotifyRegister(event, id))
	if err != nil {
		return err
	}
	c.state.events[event] = struct{}{}
	return nil
}

func (c *Connection) Command(ctx context.Context, cmd commands.Command) (Result, error) {
	cmdText := cmd.Serialize()

	if deadline, ok := ctx.Deadline(); ok {
		err := c.conn.SetWriteDeadline(deadline)
		if err != nil {
			return Result{}, err
		}
	}

	_, err := c.conn.Write([]byte(cmdText + "\n"))
	if err != nil {
		return Result{}, err
	}

	var firstLine string
	select {
	case <-ctx.Done():
		return Result{}, ctx.Err()
	case err := <-c.readErrors:
		return Result{}, err
	case firstLine = <-c.dataErrorLines:
	}

	if strings.HasPrefix(firstLine, "error") {
		errorData, err := readErrorLine(firstLine)
		if err != nil {
			return Result{}, err
		}
		res := Result{
			Id:        errorData["id"],
			Message:   errorData["msg"],
			Data:      nil,
			ErrorData: errorData,
		}

		if !res.Success() {
			return Result{}, res
		}
		return res, nil
	}

	var secondLine string
	select {
	case <-ctx.Done():
		return Result{}, ctx.Err()
	case err := <-c.readErrors:
		return Result{}, err
	case secondLine = <-c.dataErrorLines:
	}

	data, err := readDataLine(firstLine)
	if err != nil {
		return Result{}, err
	}

	errorData, err := readErrorLine(secondLine)
	if err != nil {
		return Result{}, err
	}

	res := Result{
		Id:        errorData["id"],
		Message:   errorData["msg"],
		Data:      data,
		ErrorData: errorData,
	}

	if !res.Success() {
		return Result{}, res
	}
	return res, nil
}

func (c *Connection) readResponses() {
	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			c.readErrors <- err
			return
		}

		line = strings.Trim(line, "\r\n")
		if strings.HasPrefix(line, "notify") {
			decodedEvent, err := readEvent(line)
			if err != nil {
				select {
				case c.readErrors <- err:
				case <-time.After(500 * time.Millisecond):
					// Skip non-responding event listener
				}
			}
			for _, listener := range c.eventListeners {
				listener <- decodedEvent
			}
		} else {
			c.dataErrorLines <- line
		}
	}
}

func (c *Connection) RegisterListener() (int64, chan Event) {
	token := rand.Int63()
	channel := make(chan Event)
	c.eventListeners[token] = channel
	return token, channel
}

func (c *Connection) UnregisterListener(token int64) {
	channel, ok := c.eventListeners[token]
	if !ok {
		return
	}
	delete(c.eventListeners, token)
	close(channel)
}

func (c *Connection) keepalive() {
	for {
		time.Sleep(30 * time.Second)

		_, err := c.Command(context.Background(), commands.WhoAmI())
		if err != nil {
			resErr, ok := err.(Result)
			if ok {
				log.Printf("Debug: Keepalive failed with result error: %v", resErr)
				continue
			}
			conn, reader, err := connect(c.state.addr)
			if err != nil {
				log.Printf("reconnect failed, trying again: %v", err)
				continue
			}

			c.conn = conn
			c.reader = reader
			go c.readResponses()
		}
		log.Print("Debug: Keepalive successful")
	}

}

func connect(addr string) (net.Conn, *bufio.Reader, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, nil, err
	}

	reader := bufio.NewReader(conn)

	motd, err := reader.ReadString('\n')
	if err != nil {
		return nil, nil, err
	}
	motd = strings.Trim(motd, "\r\n")

	if motd != "TS3" {
		return nil, nil, fmt.Errorf("expected TS3 as MOTD, got %q, connected to the wrong server", motd)
	}

	// Skip second line
	_, err = reader.ReadString('\n')
	if err != nil {
		return nil, nil, fmt.Errorf("expected to read the next welcome line, got err: %v", err)
	}

	return conn, reader, nil
}
