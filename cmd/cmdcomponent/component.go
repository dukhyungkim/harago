package cmdcomponent

import (
	"harago/gservice/gchat"
	"harago/repository"
	"strings"

	"github.com/jessevdk/go-flags"
	"google.golang.org/api/chat/v1"
)

type CmdComponent struct {
	name   string
	db     *repository.DB
	parser *flags.Parser
}

const (
	subCmdSet    = "set"
	subCmdList   = "list"
	subCmdRemove = "rm"
	subCmdHelp   = "help"
)

func NewCmdSetComponent(db *repository.DB) *CmdComponent {
	parser := flags.NewParser(nil, flags.Default)

	var setCommand SetCommand
	parser.AddCommand(subCmdSet, "123", "seeeeeet", &setCommand)

	return &CmdComponent{
		name:   "/component",
		db:     db,
		parser: parser,
	}
}

func (c *CmdComponent) GetName() string {
	return c.name
}

type subCmd struct {
	Set *struct {
		Company string `long:"company"`
		Type    string `long:"type"`
	} `command:"set"`
	List *struct{} `command:"list"`
	//Remove string `long:"remove"`
	//Help   string `long:"help"`
}

func (c *CmdComponent) Run(event *gchat.ChatEvent) *chat.Message {
	fields := strings.Fields(event.Message.Text)
	if fields == nil {
		return c.Help()
	}

	//var opts subCmd
	//args, err := flags.ParseArgs(&opts, fields[1:])
	//if err != nil {
	//	return &chat.Message{Text: err.Error()}
	//}
	//
	//return &chat.Message{Text: fmt.Sprintf("opts: %+#v, args: %+#v", opts, args)}

	_, err := c.parser.ParseArgs(fields[1:])
	if err != nil {
		return &chat.Message{Text: err.Error()}
	}
	return &chat.Message{Text: "OK"}
}

func (c *CmdComponent) Help() *chat.Message {
	var sb strings.Builder
	c.parser.WriteHelp(&sb)
	return &chat.Message{Text: sb.String()}
}
