package handler

import (
	harborModel "github.com/dukhyungkim/harbor-client/model"
	pbAct "github.com/dukhyungkim/libharago/gen/go/proto/action"
	"harago/repo"
	"harago/stream"
	"log"
)

type HarborEventHandler struct {
	streamClient *stream.Client
	etcdClient   *repo.Etcd
}

func NewHarborEventHandler(streamClient *stream.Client, etcdClient *repo.Etcd) *HarborEventHandler {
	return &HarborEventHandler{streamClient: streamClient, etcdClient: etcdClient}
}

func (h *HarborEventHandler) HandleHarborEvent(event *harborModel.WebhookEvent) {
	name := event.EventData.Repository.Name
	request := &pbAct.ActionRequest{
		Type: pbAct.ActionType_DEPLOY,
		Request_OneOf: &pbAct.ActionRequest_ReqDeploy{
			ReqDeploy: &pbAct.ActionRequest_DeployRequest{
				Name:        name,
				ResourceUrl: event.EventData.Resources[0].ResourceURL,
			},
		},
	}
	log.Println("pbAction:", request.String())

	subject := "harago.shared.action"
	if !h.etcdClient.IsShared(name) {
		subject = "harago.company.action"
	}

	if err := h.streamClient.PublishAction(subject, request); err != nil {
		log.Println(err)
	}
}
