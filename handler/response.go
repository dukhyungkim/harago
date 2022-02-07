package handler

import (
	pbAct "github.com/dukhyungkim/libharago/gen/go/proto/action"
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

func (h *ResponseHandler) NotifyResponse(message *pbAct.ActionResponse) {
	if message.GetSpace() == "" {
		h.broadcastMessage(message.GetText())
		return
	}
	h.sendMessageToSpace(message.GetSpace(), message.GetText())
}

func (h *ResponseHandler) broadcastMessage(text string) {
	spaces, err := h.repo.FindSpaces()
	if err != nil {
		log.Println(err)
		return
	}

	msg := formatMessage(text)

	for _, space := range spaces {
		go h.gChat.SendMessage(space.Space, msg)
	}
}

func (h *ResponseHandler) sendMessageToSpace(space string, text string) {
	msg := formatMessage(text)

	go h.gChat.SendMessage(space, msg)
}

func formatMessage(message string) *chat.Message {
	sb := strings.Builder{}

	sb.WriteString("```\n")
	sb.WriteString(message)
	sb.WriteString("```\n")

	return &chat.Message{Text: sb.String()}
}
