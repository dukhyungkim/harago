package cmdharbor

import (
	"fmt"
	"github.com/dukhyungkim/harbor-client"
	harborModel "github.com/dukhyungkim/harbor-client/model"
	"google.golang.org/api/chat/v1"
	"log"
)

func (c *CmdHarbor) handleList(params *cmdParams) *chat.Message {
	if params.RepoName != "" {
		return listArtifacts(c.harborClient, params)
	}

	if params.ProjectName != "" {
		return listRepositories(c.harborClient, params)
	}

	return listProjects(c.harborClient, params)
}

func listArtifacts(client harbor.Client, params *cmdParams) *chat.Message {
	log.Printf("%+v\n", params)
	return &chat.Message{Text: "listArtifacts"}
}

func listRepositories(client harbor.Client, params *cmdParams) *chat.Message {
	repositoriesParams := harborModel.NewListRepositoriesParams()
	if params.Page != 0 {
		repositoriesParams.Page = params.Page
	}
	if params.Size != 0 {
		repositoriesParams.PageSize = params.Size
	}

	repositories, err := client.ListRepositories(params.ProjectName, repositoriesParams)
	if err != nil {
		return &chat.Message{Text: fmt.Sprintf("Error!: %s", err.Error())}
	}
	cards := makeRepositoryCard(repositories)
	return &chat.Message{Text: fmt.Sprintf("list of repositories in %s", params.ProjectName), Cards: cards}
}

func listProjects(client harbor.Client, params *cmdParams) *chat.Message {
	log.Printf("%+v\n", params)
	return &chat.Message{Text: "listProjects"}
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
