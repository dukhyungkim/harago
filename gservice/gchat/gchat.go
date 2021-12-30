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
	service *chat.Service
}

func NewGChat(gService *gservice.GService) (*GChat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	service, err := chat.NewService(ctx, option.WithHTTPClient(gService.GetClient()))
	if err != nil {
		return nil, err
	}

	return &GChat{service: service}, nil
}

func (c *GChat) HandleMessage(event *ChatEvent) *chat.Message {
	log.Printf("%+v\n", *event)
	return &chat.Message{Text: "hello"}
}
