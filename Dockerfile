FROM golang:1.17-alpine AS build-env
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download -x

COPY . .
RUN mkdir bin
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin ./cmd/...

FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=build-env /app/bin/ts3tgbot ./ts3tgbot
USER nonroot:nonroot
ENTRYPOINT ["/app/ts3tgbot"]
