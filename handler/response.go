package handler

import (
	"google.golang.org/api/chat/v1"
	"harago/gservice/gchat"
	"harago/repo"
	"log"
	"strings"
)

type ResponseHandler struct {
	gChat *gchat.GChat
	repo  *repo.DB
}

func NewResponseHandler(gChat *gchat.GChat, repo *repo.DB) *ResponseHandler {
	return &ResponseHandler{gChat: gChat, repo: repo}
}

func (h *ResponseHandler) NotifyResponse(message string) {
	spaces, err := h.repo.FindSpaces()
	if err != nil {
		log.Println(err)
		return
	}

	msg := formatMessage(message)

	for _, space := range spaces {
		go h.gChat.SendMessage(space.Space, msg)
	}
}

func formatMessage(message string) *chat.Message {
	sb := strings.Builder{}

	sb.WriteString("```\n")
	sb.WriteString(message)
	sb.WriteString("```\n")

	return &chat.Message{Text: sb.String()}
}
