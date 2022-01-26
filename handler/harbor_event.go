package handler

import (
	harborModel "github.com/dukhyungkim/harbor-client/model"
	"log"
)

func HandleHarborEvent(event *harborModel.WebhookEvent) {
	log.Printf("%+v\n", *event)
	log.Println(event.EventData.Repository.Name)
	log.Println(event.EventData.Resources[0].ResourceURL)
}
