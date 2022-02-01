package handler

import (
	harborModel "github.com/dukhyungkim/harbor-client/model"
	pbAct "github.com/dukhyungkim/libharago/gen/go/proto/action"
	"harago/stream"
	"log"
)

type HarborEventHandler struct {
	stream *stream.Client
}

func NewHarborEventHandler(stream *stream.Client) *HarborEventHandler {
	return &HarborEventHandler{stream: stream}
}

func (h *HarborEventHandler) HandleHarborEvent(event *harborModel.WebhookEvent) {
	request := &pbAct.ActionRequest{
		Type: pbAct.ActionType_DEPLOY,
		Request_OneOf: &pbAct.ActionRequest_ReqDeploy{
			ReqDeploy: &pbAct.ActionRequest_DeployRequest{
				Name:        event.EventData.Repository.Name,
				ResourceUrl: event.EventData.Resources[0].ResourceURL,
			},
		},
	}
	log.Println("pbAction:", request.String())

	if err := h.stream.PublishAction(request); err != nil {
		log.Println(err)
	}
}
