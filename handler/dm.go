package handler

import (
	"google.golang.org/api/chat/v1"
	"harago/cmd"
	"harago/gservice/gchat"
)

type DMHandler struct {
	cmdExecutor *cmd.Executor
}

func NewDMHandler(cmdExecutor *cmd.Executor) gchat.Handler {
	return &DMHandler{cmdExecutor: cmdExecutor}
}

func (h *DMHandler) ProcessMessage(event *gchat.ChatEvent) *chat.Message {
	var chatMessage *chat.Message

	switch event.Type {
	case gchat.AddedToSpace:

	case gchat.Message:
		chatMessage = h.cmdExecutor.Run(event)

	case gchat.RemovedFromSpace:
		chatMessage = &chat.Message{}
	}

	return chatMessage
}
