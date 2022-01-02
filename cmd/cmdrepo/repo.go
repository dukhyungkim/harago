package cmdrepo

import (
	"docgo/gservice/gchat"
	"github.com/dukhyungkim/harbor-client"
	"google.golang.org/api/chat/v1"
)

type CmdRepo struct {
	name         string
	harborClient harbor.Client
}

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
	return &chat.Message{Text: "implement me"}
}

func (c *CmdRepo) Help() *chat.Message {
	return &chat.Message{Text: "HELP!"}
}
