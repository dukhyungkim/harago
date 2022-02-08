package cmd_harbor

import (
	"github.com/dukhyungkim/harbor-client"
	"google.golang.org/api/chat/v1"
	"harago/common"
)

func (c *CmdHarbor) handleInfo(params *cmdParams) *chat.Message {
	if params.ProjectName != "" && params.RepoName != "" && params.ArtifactName != "" {
		return infoArtifact(c.harborClient, params)
	}

	if params.ProjectName != "" && params.RepoName != "" {
		return infoRepository(c.harborClient, params)
	}

	if params.ProjectName != "" {
		return infoProject(c.harborClient, params)
	}

	return c.Help()
}

func infoProject(client *harbor.Client, params *cmdParams) *chat.Message {
	project, err := client.GetProject(params.ProjectName)
	if err != nil {
		return &chat.Message{Text: common.ErrHarborResponse(err).Error()}
	}

	return &chat.Message{Text: "project info", Cards: []*chat.Card{makeProjectCard(project)}}
}

func infoRepository(client *harbor.Client, params *cmdParams) *chat.Message {
	repository, err := client.GetRepository(params.ProjectName, params.RepoName)
	if err != nil {
		return &chat.Message{Text: common.ErrHarborResponse(err).Error()}
	}

	return &chat.Message{Text: "repository info", Cards: []*chat.Card{makeRepositoryCard(repository)}}
}

func infoArtifact(client *harbor.Client, params *cmdParams) *chat.Message {
	artifact, err := client.GetArtifact(params.ProjectName, params.RepoName, params.ArtifactName)
	if err != nil {
		return &chat.Message{Text: common.ErrHarborResponse(err).Error()}
	}

	tags, err := client.ListTags(params.ProjectName, params.RepoName, artifact.Digest, nil)
	if err != nil {
		return &chat.Message{Text: common.ErrHarborResponse(err).Error()}
	}

	return &chat.Message{Text: "artifact info", Cards: []*chat.Card{makeArtifactCard(artifact, tags)}}
}
