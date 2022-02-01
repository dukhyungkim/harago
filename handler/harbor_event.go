package handler

import (
	harborModel "github.com/dukhyungkim/harbor-client/model"
	pb "github.com/dukhyungkim/libharago/gen/go/proto"
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
	request := &pb.ActionRequest{
		Type: pb.ActionType_DEPLOY,
		Request_OneOf: &pb.ActionRequest_DeployReq{
			DeployReq: &pb.ActionRequest_DeployRequest{
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
