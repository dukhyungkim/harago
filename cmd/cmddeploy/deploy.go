package cmddeploy

import (
	"fmt"
	"harago/common"
	"harago/gservice/gchat"
	"harago/stream"
	"strings"

	pbAct "github.com/dukhyungkim/libharago/gen/go/proto/action"
	"github.com/jessevdk/go-flags"
	"google.golang.org/api/chat/v1"
)

type CmdDeploy struct {
	name         string
	streamClient *stream.Client
}

func NewDeployCommand(streamClient *stream.Client) *CmdDeploy {
	return &CmdDeploy{
		name:         "/deploy",
		streamClient: streamClient,
	}
}

func (c *CmdDeploy) GetName() string {
	return c.name
}

type Opts struct {
	Company string `long:"company" short:"c"`
}

func (c *CmdDeploy) Run(event *gchat.ChatEvent) *chat.Message {
	fields := strings.Fields(event.Message.Text)
	if fields == nil {
		return c.Help()
	}

	var opts Opts
	parser := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash)

	args, err := parser.ParseArgs(fields[1:])
	if err != nil {
		return &chat.Message{Text: err.Error()}
	}

	if len(args) == 0 {
		return &chat.Message{Text: "invalid ResourceURL"}
	}
	resourceURL := args[0]

	subject := common.SharedActionSubject
	if opts.Company != "" {
		subject = fmt.Sprintf(common.SpecificCompanyActionSubject, opts.Company)
	}

	pbAction := &pbAct.ActionRequest{
		Type:  pbAct.ActionType_DEPLOY,
		Space: event.Space.Name,
		Request_OneOf: &pbAct.ActionRequest_ReqDeploy{
			ReqDeploy: &pbAct.ActionRequest_DeployRequest{
				Name:        parseName(resourceURL),
				ResourceUrl: resourceURL,
			},
		},
	}
	if err = c.streamClient.PublishAction(subject, pbAction); err != nil {
		return &chat.Message{Text: err.Error()}
	}

	if subject == common.SharedActionSubject {
		return &chat.Message{Text: fmt.Sprintf("publish to %s, ResourceURL: %s", subject, resourceURL)}
	}
	return &chat.Message{Text: fmt.Sprintf("publish to %s, Company: %s, ResourceURL: %s", subject, opts.Company, resourceURL)}
}

func (c *CmdDeploy) Help() *chat.Message {
	return &chat.Message{Text: "HELP!"}
}

func parseName(resourceURL string) string {
	s1 := strings.Split(resourceURL, ":")
	s2 := strings.Split(s1[0], "/")
	return s2[len(s2)-1]
}
