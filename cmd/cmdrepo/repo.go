package cmdrepo

import (
	"docgo/gservice/gchat"
	"github.com/dukhyungkim/harbor-client"
	"google.golang.org/api/chat/v1"
	"strings"
)

type CmdRepo struct {
	name         string
	harborClient harbor.Client
}

const (
	cmdLS   = "ls"
	cmdList = "list"
)

func NewRepoCommand(harborClient harbor.Client) *CmdRepo {
	return &CmdRepo{
		name:         "/repo",
		harborClient: harborClient,
	}
}

func (c *CmdRepo) GetName() string {
	return c.name
}

func (c *CmdRepo) Run(event *gchat.ChatEvent) *chat.Message {
	fields := strings.Fields(event.Message.Text)
	if fields == nil {
		return c.Help()
	}

	switch fields[0] {
	case cmdLS, cmdList:
		return &chat.Message{Text: "will show list of repository"}
	}

	return &chat.Message{Text: "implement me"}
}

func (c *CmdRepo) Help() *chat.Message {
	return &chat.Message{Text: "HELP!"}
}
