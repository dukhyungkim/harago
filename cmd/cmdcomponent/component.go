package cmdcomponent

import (
	"fmt"
	"harago/entity"
	"harago/gservice/gchat"
	"harago/repository"
	"log"
	"strings"

	"github.com/jessevdk/go-flags"
	"google.golang.org/api/chat/v1"
)

type CmdComponent struct {
	name string
	db   *repository.DB
}

func NewCmdSetComponent(db *repository.DB) *CmdComponent {
	return &CmdComponent{
		name: "/component",
		db:   db,
	}
}

func (c *CmdComponent) GetName() string {
	return c.name
}

const (
	subCmdSet    = "set"
	subCmdList   = "list"
	subCmdRemove = "remove"
)

type SubCmdOpts struct {
	Company string `long:"company"`
	Type    string `long:"type"`
}

type SubCmd struct {
	Set    SubCmdOpts `command:"set"`
	List   struct{}   `command:"list" alias:"ls"`
	Remove SubCmdOpts `command:"remove" alias:"rm"`
}

func (c *CmdComponent) Run(event *gchat.ChatEvent) *chat.Message {
	fields := strings.Fields(event.Message.Text)
	if fields == nil {
		return c.Help()
	}

	var subCmd SubCmd
	parser := flags.NewParser(&subCmd, flags.HelpFlag|flags.PassDoubleDash)

	args, err := parser.ParseArgs(fields[1:])
	if err != nil {
		return &chat.Message{Text: err.Error()}
	}

	switch parser.Active.Name {
	case subCmdSet:
		if len(args) == 0 {
			return &chat.Message{Text: "not enough argument"}
		}
		ct := entity.ComponentType{Company: subCmd.Set.Company, Type: subCmd.Set.Type, Component: args[0]}
		err = c.db.UpsertComponentType(&ct)
		if err != nil {
			return &chat.Message{Text: err.Error()}
		}

	case subCmdList:
		var cts []*entity.ComponentType
		cts, err = c.db.ListComponentTypes()
		if err != nil {
			return &chat.Message{Text: err.Error()}
		}
		var sb strings.Builder
		for _, ct := range cts {
			sb.WriteString(fmt.Sprintf("Company: %s, Type: %s, Component: %s, CreatedAt: %s, UpdatedAt: %s\n",
				ct.Company, ct.Type, ct.Component, ct.CreatedAt.Local().String(), ct.UpdatedAt.Local().String()),
			)
		}
		return &chat.Message{Text: sb.String()}

	case subCmdRemove:
		ct := entity.ComponentType{Company: subCmd.Remove.Company, Type: subCmd.Remove.Type}
		err = c.db.DeleteComponentType(&ct)
		if err != nil {
			return &chat.Message{Text: err.Error()}
		}
		log.Println("remove")
	}

	return &chat.Message{Text: "OK"}
}

func (c *CmdComponent) Help() *chat.Message {
	return &chat.Message{Text: "HELP"}
}
