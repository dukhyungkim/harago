package handler

import (
	"google.golang.org/api/chat/v1"
	"harago/cmd"
	"harago/db"
	"harago/entity"
	"harago/gservice/gchat"
)

type DMHandler struct {
	cmdExecutor *cmd.Executor
	repo        *db.DB
}

func NewDMHandler(cmdExecutor *cmd.Executor, repo *db.DB) gchat.Handler {
	return &DMHandler{cmdExecutor: cmdExecutor, repo: repo}
}

func (h *DMHandler) ProcessMessage(event *gchat.ChatEvent) *chat.Message {
	switch event.Type {
	case gchat.AddedToSpace:
		userSpace := &entity.UserSpace{
			Name:  event.User.DisplayName,
			Email: event.User.Email,
			Space: event.Space.Name,
		}
		if err := h.repo.SaveSpace(userSpace); err != nil {
			return &chat.Message{Text: err.Error()}
		}
		return &chat.Message{Text: "Save Space"}

	case gchat.Message:
		return h.cmdExecutor.Run(event)

	case gchat.RemovedFromSpace:
		h.repo.DeleteSpace(event.User.Email)
	}

	return &chat.Message{}
}
