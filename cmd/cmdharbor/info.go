package cmdharbor

import (
	"google.golang.org/api/chat/v1"
	"log"
)

func (c *CmdHarbor) handleInfo(params *cmdParams) *chat.Message {
	log.Printf("%+v\n", params)
	return &chat.Message{Text: "handleInfo"}
}
