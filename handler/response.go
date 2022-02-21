package handler

import (
	"fmt"
	pbAct "github.com/dukhyungkim/libharago/gen/go/proto/action"
	"google.golang.org/api/chat/v1"
	"harago/common"
	"harago/gservice/gchat"
	"harago/repository"
	"log"
	"strings"
)

type ResponseHandler struct {
	gChat *gchat.GChat
	repo  *repository.DB
}

func NewResponseHandler(gChat *gchat.GChat, repo *repository.DB) *ResponseHandler {
	return &ResponseHandler{gChat: gChat, repo: repo}
}

func (h *ResponseHandler) NotifyResponse(response *pbAct.ActionResponse) {
	if response.GetSpace() == "" {
		h.broadcastMessage(response)
		return
	}
	h.sendMessageToSpace(response.GetSpace(), response)
}

func (h *ResponseHandler) broadcastMessage(response *pbAct.ActionResponse) {
	spaces, err := h.repo.FindSpaces()
	if err != nil {
		log.Println(err)
		return
	}

	var message *chat.Message
	switch response.GetType() {
	case pbAct.ActionType_DEPLOY:
		message = formatDeployResponse(response)
	default:
		log.Println(common.ErrUnknownActionType(response.GetType()))
		return
	}

	for _, space := range spaces {
		go h.gChat.SendMessage(space.Space, message)
	}
}

func (h *ResponseHandler) sendMessageToSpace(space string, response *pbAct.ActionResponse) {
	var message *chat.Message
	switch response.GetType() {
	case pbAct.ActionType_DEPLOY:
		message = formatDeployResponse(response)
	default:
		log.Println(common.ErrUnknownActionType(response.GetType()))
		return
	}

	go h.gChat.SendMessage(space, message)
}

func formatDeployResponse(response *pbAct.ActionResponse) *chat.Message {
	deployResp := response.GetRespDeploy()

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%s - Done from %s with %s\n",
		response.GetType().String(), deployResp.GetCompany(), deployResp.GetResourceUrl()))
	sb.WriteString("```")
	sb.WriteString(deployResp.GetText())
	sb.WriteString("```")

	return &chat.Message{Text: sb.String()}
}
