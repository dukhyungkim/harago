package gchat

import (
	"context"
	"docgo/gservice"
	"google.golang.org/api/chat/v1"
	"google.golang.org/api/option"
	"log"
	"time"
)

type GChat struct {
	service     *chat.Service
	dmHandler   Handler
	roomHandler Handler
}

func NewGChat(gService *gservice.GService, dmHandler Handler, roomHandler Handler) (*GChat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	service, err := chat.NewService(ctx, option.WithHTTPClient(gService.GetClient()))
	if err != nil {
		return nil, err
	}

	return &GChat{service: service, dmHandler: dmHandler, roomHandler: roomHandler}, nil
}

func (c *GChat) HandleMessage(event *ChatEvent) *chat.Message {
	log.Printf("%+v\n", *event)
	var chatMessage *chat.Message
	if event.Space.Type == DM {
		chatMessage = c.dmHandler.ProcessMessage(event)
	} else {
		chatMessage = c.roomHandler.ProcessMessage(event)
	}
	return chatMessage
}
