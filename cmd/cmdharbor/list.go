package cmdharbor

import (
	"docgo/common"
	"docgo/util"
	"fmt"
	"github.com/dukhyungkim/harbor-client"
	harborModel "github.com/dukhyungkim/harbor-client/model"
	"google.golang.org/api/chat/v1"
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
		return &chat.Message{Text: common.ErrHarborResponse(err).Error()}
	}
	cards := makeRepositoryCard(repositories)
	return &chat.Message{Text: fmt.Sprintf("list of repositories in %s", params.ProjectName), Cards: cards}
}

func listArtifacts(client harbor.Client, params *cmdParams) *chat.Message {
	artifactsParams := harborModel.NewListArtifactsParams()
	if params.Page != 0 {
		artifactsParams.Page = params.Page
	}
	if params.Size != 0 {
		artifactsParams.PageSize = params.Size
	}

	artifacts, err := client.ListArtifacts(params.ProjectName, params.RepoName, artifactsParams)
	if err != nil {
		return &chat.Message{Text: common.ErrHarborResponse(err).Error()}
	}
	cards := makeArtifactCard(artifacts)
	return &chat.Message{Text: fmt.Sprintf("list of artifacts in %s/%s", params.ProjectName, params.RepoName), Cards: cards}
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
								Content:          projects[i].OwnerName,
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

func makeArtifactCard(artifacts []*harborModel.Artifact) []*chat.Card {
	cards := make([]*chat.Card, len(artifacts))
	for i := range artifacts {
		cards[i] = &chat.Card{
			Header: &chat.CardHeader{
				Title: artifacts[i].Digest[:15],
			},
			Sections: []*chat.Section{
				{
					Widgets: []*chat.WidgetMarkup{
						//{
						//	KeyValue: &chat.KeyValue{
						//		TopLabel:         "Tags",
						//		Content:          artifacts[i].Tags.,
						//		ContentMultiline: true,
						//	},
						//},
						{
							KeyValue: &chat.KeyValue{
								TopLabel:         "Size",
								Content:          util.ByteCountIEC(int64(artifacts[i].Size)),
								ContentMultiline: true,
							},
						},
						{
							KeyValue: &chat.KeyValue{
								TopLabel:         "PushTime",
								Content:          artifacts[i].PushTime,
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
