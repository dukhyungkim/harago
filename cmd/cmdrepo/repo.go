package cmdrepo

import (
	"docgo/gservice/gchat"
	"fmt"
	"github.com/dukhyungkim/harbor-client"
	harborModel "github.com/dukhyungkim/harbor-client/model"
	"google.golang.org/api/chat/v1"
	"strconv"
	"strings"
)

type CmdRepo struct {
	name         string
	harborClient harbor.Client
}

const (
	cmdLS   = "ls"
	cmdList = "list"

	minCountList           = 3
	minCountListPagination = 5
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

	switch fields[1] {
	case cmdLS, cmdList:
		if len(fields) < minCountList {
			return c.Help()
		}
		projectName := fields[2]
		repositoriesParams := harborModel.NewListRepositoriesParams()

		if len(fields) >= minCountListPagination {
			page, err := strconv.ParseInt(fields[3], 10, 64)
			if err != nil {
				return &chat.Message{Text: fmt.Sprintf("Error!: %s", err.Error())}
			}
			repositoriesParams.Page = page

			pageSize, err := strconv.ParseInt(fields[4], 10, 64)
			if err != nil {
				return &chat.Message{Text: fmt.Sprintf("Error!: %s", err.Error())}
			}
			repositoriesParams.PageSize = pageSize
		}

		repositories, err := c.harborClient.ListRepositories(projectName, repositoriesParams)
		if err != nil {
			return &chat.Message{Text: fmt.Sprintf("Error!: %s", err.Error())}
		}
		cards := makeRepositoryCard(repositories)
		return &chat.Message{Text: fmt.Sprintf("list of repository in %s", projectName), Cards: cards}
	}

	return c.Help()
}

func (c *CmdRepo) Help() *chat.Message {
	return &chat.Message{Text: "HELP!"}
}

func makeRepositoryCard(repositories []*harborModel.Repository) []*chat.Card {
	cards := make([]*chat.Card, len(repositories))
	for i := range repositories {
		cards[i] = &chat.Card{
			Header: &chat.CardHeader{
				Title: repositories[i].Name,
			},
			Sections: []*chat.Section{
				{
					Widgets: []*chat.WidgetMarkup{
						{
							KeyValue: &chat.KeyValue{
								TopLabel:         "ArtifactCount",
								Content:          fmt.Sprint(repositories[i].ArtifactCount),
								ContentMultiline: true,
							},
						},
						{
							KeyValue: &chat.KeyValue{
								TopLabel:         "CreationTime",
								Content:          repositories[i].CreationTime,
								ContentMultiline: true,
							},
						},
						{
							KeyValue: &chat.KeyValue{
								TopLabel:         "UpdateTime",
								Content:          repositories[i].UpdateTime,
								ContentMultiline: true,
							},
						},
						{
							KeyValue: &chat.KeyValue{
								TopLabel:         "PullCount",
								Content:          fmt.Sprint(repositories[i].PullCount),
								ContentMultiline: true,
							},
						},
					},
				},
			},
		}
	}
	return cards
}
