package cmd

import (
	"docgo/cmd/cmdharbor"
	"docgo/common"
	"docgo/config"
	"docgo/gservice/gchat"
	"fmt"
	"github.com/dukhyungkim/harbor-client"
	"google.golang.org/api/chat/v1"
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

func (e *Executor) LoadCommands(cfg *config.Config) error {
	harborClient := harbor.NewClient(&harbor.Config{
		URL:      cfg.Harbor.URL,
		Username: cfg.Harbor.Username,
		Password: cfg.Harbor.Password,
	})

	repoCommand := cmdharbor.NewHarborCommand(harborClient)
	if err := e.AddCommand(repoCommand.GetName(), repoCommand); err != nil {
		return err
	}

	return nil
}
