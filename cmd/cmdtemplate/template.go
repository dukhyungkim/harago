package cmdtemplate

import (
	"harago/gservice/gchat"
	"harago/repository"
	"strings"

	"github.com/jessevdk/go-flags"
	"google.golang.org/api/chat/v1"
)

type CmdTemplate struct {
	name       string
	etcdClient *repository.Etcd
}

func NewTemplateCommand(etcdClient *repository.Etcd) *CmdTemplate {
	return &CmdTemplate{
		name:       "/template",
		etcdClient: etcdClient,
	}
}

func (c *CmdTemplate) GetName() string {
	return c.name
}

type Opts struct {
	List SubCmdOpts `command:"list" alias:"ls"`
	Show SubCmdOpts `command:"show"`
}

type SubCmdOpts struct {
	ProjectName  string `long:"project" alias:"proj"`
	RepoName     string `long:"repository" alias:"repo"`
	ArtifactName string `long:"artifact"`
	Page         int64  `long:"page"`
	Size         int64  `long:"size"`
}

const (
	subCmdList = "list"
	subCmdShow = "show"
)

func (c *CmdTemplate) Run(event *gchat.ChatEvent) *chat.Message {
	fields := strings.Fields(event.Message.Text)
	if fields == nil {
		return c.Help()
	}

	var opts Opts
	parser := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash)

	_, err := parser.ParseArgs(fields[1:])
	if err != nil {
		return &chat.Message{Text: err.Error()}
	}

	switch parser.Active.Name {
	case subCmdList:
		return c.handleList(&opts.List)
	case subCmdShow:
		return c.Help()
	default:
		return c.Help()
	}
}

func (c *CmdTemplate) Help() *chat.Message {
	return &chat.Message{Text: "HELP!"}
}

func (c *CmdTemplate) handleList(s *SubCmdOpts) *chat.Message {
	templates, err := c.etcdClient.ListTemplates()
	if err != nil {
		return &chat.Message{Text: err.Error()}
	}

	return &chat.Message{Text: strings.Join(templates, "\n")}
}
