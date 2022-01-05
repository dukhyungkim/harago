package handler

import (
	"docgo/cmd"
	"docgo/gservice/gchat"
	"google.golang.org/api/chat/v1"
)

type RoomHandler struct {
	cmdExecutor *cmd.Executor
}

func NewRoomHandler(cmdExecutor *cmd.Executor) gchat.Handler {
	return &RoomHandler{cmdExecutor: cmdExecutor}
}

func (h *RoomHandler) ProcessMessage(event *gchat.ChatEvent) *chat.Message {
	var chatMessage *chat.Message

	switch event.Type {
	case gchat.AddedToSpace:

	case gchat.Message:
		h.cmdExecutor.Run(event)

	case gchat.RemovedFromSpace:
		chatMessage = &chat.Message{}
	}

	return chatMessage
}