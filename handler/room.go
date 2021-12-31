package handler

import (
	"docgo/gservice/gchat"
	"google.golang.org/api/chat/v1"
)

type DMHandler struct {
}

func NewDMHandler() gchat.Handler {
	return &DMHandler{}
}

func (h *DMHandler) ProcessMessage(event *gchat.ChatEvent) *chat.Message {
	var chatMessage *chat.Message

	switch event.Type {
	case gchat.AddedToSpace:

	case gchat.Message:

	case gchat.RemovedFromSpace:
		chatMessage = &chat.Message{Text: ""}
	}

	return chatMessage
}
