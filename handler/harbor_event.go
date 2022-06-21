package handler

import (
	"harago/common"
	"harago/repository"
	"harago/stream"
	"log"

	harborModel "github.com/dukhyungkim/harbor-client/model"
	pbAct "github.com/dukhyungkim/libharago/gen/go/proto/action"
)

type HarborEventHandler struct {
	streamClient *stream.Client
	etcdClient   *repository.Etcd
}

func NewHarborEventHandler(streamClient *stream.Client, etcdClient *repository.Etcd) *HarborEventHandler {
	return &HarborEventHandler{streamClient: streamClient, etcdClient: etcdClient}
}

func (h *HarborEventHandler) HandleHarborEvent(event *harborModel.WebhookEvent) {
	name := event.EventData.Repository.Name
	request := &pbAct.ActionRequest{
		Type: pbAct.ActionType_UP,
		Request_OneOf: &pbAct.ActionRequest_ReqDeploy{
			ReqDeploy: &pbAct.ActionRequest_DeployRequest{
				Name:        name,
				ResourceUrl: event.EventData.Resources[0].ResourceURL,
			},
		},
	}
	log.Println("pbAction:", request.String())

	var subject string
	switch {
	case h.etcdClient.IsIgnore(name):
		log.Printf("%s is in ignoredList\n", name)
		return

	case h.etcdClient.IsShared(name):
		subject = common.SharedActionSubject

	case h.etcdClient.IsCompany(name):
		subject = common.CompanyActionSubject

	case h.etcdClient.IsInternal(name):
		subject = common.InternalActionSubject

	default:
		log.Printf("%s is unknown\n", name)
		return
	}

	if err := h.streamClient.PublishAction(subject, request); err != nil {
		log.Println(err)
	}
}
