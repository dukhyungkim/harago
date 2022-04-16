package cmdcomponent

import (
	"harago/gservice/gchat"
	"harago/repository"
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

type SubCmd struct {
	Mapping Mapping `command:"mapping" alias:"m"`
	Type    Type    `command:"type" alias:"t"`
}

type cmdRunHelper struct {
	cmd  string
	args []string
	db   *repository.DB
}

const (
	subCmdMapping = "mapping"
	subCmdType    = "type"

	subCmdSet    = "set"
	subCmdAdd    = "add"
	subCmdList   = "list"
	subCmdRemove = "remove"
)

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

	helper := &cmdRunHelper{
		cmd:  parser.Active.Active.Name,
		args: args,
		db:   c.db,
	}

	switch parser.Active.Name {
	case subCmdMapping:
		return subCmd.Mapping.Run(helper)

	case subCmdType:
		return subCmd.Type.Run(helper)

	default:
		return &chat.Message{Text: "not found command - cannot be here"}
	}
}

func (c *CmdComponent) Help() *chat.Message {
	return &chat.Message{Text: "HELP"}
}
