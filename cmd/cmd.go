package cmd

import (
	"fmt"
	"github.com/dukhyungkim/harbor-client"
	"google.golang.org/api/chat/v1"
	cmddeploy "harago/cmd/cmd_deploy"
	"harago/cmd/cmd_harbor"
	"harago/common"
	"harago/config"
	"harago/gservice/gchat"
	"harago/stream"
	"strings"
)

type Command interface {
	GetName() string
	Run(event *gchat.ChatEvent) *chat.Message
	Help() *chat.Message
}

type Executor struct {
	cmdList map[string]Command
}

var executor *Executor

func NewExecutor() *Executor {
	if executor != nil {
		return executor
	}

	executor = &Executor{cmdList: map[string]Command{}}
	return executor
}

func (e *Executor) AddCommand(name string, cmd Command) error {
	if _, has := e.cmdList[name]; has {
		return common.ErrDuplicateCommand(name)
	}

	e.cmdList[name] = cmd
	return nil
}

func (e *Executor) Run(event *gchat.ChatEvent) *chat.Message {
	fields := strings.Fields(event.Message.Text)
	if len(fields) == 0 {
		return &chat.Message{}
	}

	command, has := e.cmdList[fields[0]]
	if !has {
		return &chat.Message{Text: fmt.Sprintf("cannot find command: %s", fields[0])}
	}

	return command.Run(event)
}

func (e *Executor) LoadCommands(cfg *config.Config, streamClient *stream.Client) error {
	harborClient := harbor.NewClient(&harbor.Config{
		URL:      cfg.Harbor.URL,
		Username: cfg.Harbor.Username,
		Password: cfg.Harbor.Password,
	})

	cmdHarbor := cmdharbor.NewHarborCommand(harborClient)
	if err := e.AddCommand(cmdHarbor.GetName(), cmdHarbor); err != nil {
		return err
	}

	cmdDeploy := cmddeploy.NewDeployCommand(streamClient)
	if err := e.AddCommand(cmdDeploy.GetName(), cmdDeploy); err != nil {
		return err
	}

	return nil
}
