package cmddeploy

import (
	"fmt"
	pbAct "github.com/dukhyungkim/libharago/gen/go/proto/action"
	"google.golang.org/api/chat/v1"
	"harago/common"
	"harago/gservice/gchat"
	"harago/stream"
	"strings"
)

type CmdHDeploy struct {
	name         string
	streamClient *stream.Client
}

func NewDeployCommand(streamClient *stream.Client) *CmdHDeploy {
	return &CmdHDeploy{
		name:         "/deploy",
		streamClient: streamClient,
	}
}

func (c *CmdHDeploy) GetName() string {
	return c.name
}

func (c *CmdHDeploy) Run(event *gchat.ChatEvent) *chat.Message {
	fields := strings.Fields(event.Message.Text)
	if fields == nil {
		return c.Help()
	}

	params, err := newCmdParams(fields[1:])
	if err != nil {
		return &chat.Message{Text: err.Error()}
	}

	if params.ResourceURL == "" {
		return &chat.Message{Text: "empty ResourceURL"}
	}

	if len(strings.Split(params.ResourceURL, ":")) != 2 {
		return &chat.Message{Text: "invalid ResourceURL"}
	}

	subject := common.SharedActionSubject
	if params.Company != "" {
		subject = common.CompanyActionSubject
	}

	pbAction := &pbAct.ActionRequest{
		Type:  pbAct.ActionType_DEPLOY,
		Space: event.Space.Name,
		Request_OneOf: &pbAct.ActionRequest_ReqDeploy{
			ReqDeploy: &pbAct.ActionRequest_DeployRequest{
				Name:        parseName(params.ResourceURL),
				ResourceUrl: params.ResourceURL,
			},
		},
	}
	if err = c.streamClient.PublishAction(subject, pbAction); err != nil {
		return &chat.Message{Text: err.Error()}
	}

	if subject == common.SharedActionSubject {
		return &chat.Message{Text: fmt.Sprintf("publish to %s, ResourceURL: %s", subject, params.ResourceURL)}
	}
	return &chat.Message{Text: fmt.Sprintf("publish to %s, Company: %s, ResourceURL: %s", subject, params.Company, params.ResourceURL)}
}

func (c *CmdHDeploy) Help() *chat.Message {
	return &chat.Message{Text: "HELP!"}
}

func parseName(resourceURL string) string {
	s1 := strings.Split(resourceURL, ":")
	s2 := strings.Split(s1[0], "/")
	return s2[len(s2)-1]
}
