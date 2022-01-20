package cmdharbor

import (
	"github.com/dukhyungkim/harbor-client"
	"google.golang.org/api/chat/v1"
	"harago/gservice/gchat"
	"strings"
)

type CmdHarbor struct {
	name         string
	harborClient harbor.Client
}

const (
	cmdInfo = "info"
	cmdList = "list"
)

func NewHarborCommand(harborClient harbor.Client) *CmdHarbor {
	return &CmdHarbor{
		name:         "/harbor",
		harborClient: harborClient,
	}
}

func (c *CmdHarbor) GetName() string {
	return c.name
}

func (c *CmdHarbor) Run(event *gchat.ChatEvent) *chat.Message {
	fields := strings.Fields(event.Message.Text)
	if fields == nil {
		return c.Help()
	}

	params, err := newCmdParams(fields)
	if err != nil {
		return &chat.Message{Text: err.Error()}
	}

	switch params.SubCmd {
	case cmdList:
		return c.handleList(params)
	case cmdInfo:
		return c.handleInfo(params)
	default:
		return c.Help()
	}
}

func (c *CmdHarbor) Help() *chat.Message {
	return &chat.Message{Text: "HELP!"}
}