package handler

import (
	"docgo/gservice/gchat"
	"google.golang.org/api/chat/v1"
)

type RoomHandler struct {
}

func NewRoomHandler() gchat.Handler {
	return &RoomHandler{}
}

func (h *RoomHandler) ProcessMessage(event *gchat.ChatEvent) *chat.Message {
	var chatMessage *chat.Message

	switch event.Type {
	case gchat.AddedToSpace:

	case gchat.Message:

	case gchat.RemovedFromSpace:
		chatMessage = &chat.Message{Text: ""}
	}

	return chatMessage
}
