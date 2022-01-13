package cmdharbor

import (
	"docgo/common"
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
	projectsParams := harborModel.NewListProjectsParams()
	if params.Page != 0 {
		projectsParams.Page = params.Page
	}
	if params.Size != 0 {
		projectsParams.PageSize = params.Size
	}

	projects, err := client.ListProjects(projectsParams)
	if err != nil {
		return &chat.Message{Text: common.ErrHarborResponse(err).Error()}
	}
	cards := makeProjectCard(projects)
	return &chat.Message{Text: "list of projects", Cards: cards}
}

func makeProjectCard(projects []*harborModel.Project) []*chat.Card {
	cards := make([]*chat.Card, len(projects))
	for i := range projects {
		cards[i] = &chat.Card{
			Header: &chat.CardHeader{
				Title: projects[i].Name,
			},
			Sections: []*chat.Section{
				{
					Widgets: []*chat.WidgetMarkup{
						{
							KeyValue: &chat.KeyValue{
								TopLabel:         "RepoCount",
								Content:          fmt.Sprint(projects[i].RepoCount),
								ContentMultiline: true,
							},
						},
						{
							KeyValue: &chat.KeyValue{
								TopLabel:         "OwnerName",
								Content:          fmt.Sprint(projects[i].OwnerName),
								ContentMultiline: true,
							},
						},
						{
							KeyValue: &chat.KeyValue{
								TopLabel:         "UpdateTime",
								Content:          projects[i].UpdateTime,
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
								TopLabel:         "PullCount",
								Content:          fmt.Sprint(repositories[i].PullCount),
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
					},
				},
			},
		}
	}
	return cards
}
