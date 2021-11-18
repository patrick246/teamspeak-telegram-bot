package telegram

import (
	"bytes"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"html/template"
)

var onlineTemplate = template.Must(template.New("user_joined").Parse(`➡️🎧 <i>{{ . }} joined the server</i>`))

type Client struct {
	bot           *tgbotapi.BotAPI
	targetChannel int64
}

func NewClient(token string, targetChannel int64) (*Client, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Client{
		bot:           bot,
		targetChannel: targetChannel,
	}, nil
}

func (c *Client) ReceiveOnline(user string) {
	target := &bytes.Buffer{}
	err := onlineTemplate.Execute(target, user)
	if err != nil {
		return
	}
	message := tgbotapi.NewMessage(c.targetChannel, target.String())
	message.ParseMode = "html"
	_, err = c.bot.Send(message)
	if err != nil {
		return
	}
}
