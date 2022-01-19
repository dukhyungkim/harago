package handler

import (
	"harago/entity"
	"log"
)

func HandleHarborEvent(event *entity.HarborWebhookEvent) {
	log.Printf("%+v\n", *event)
}
