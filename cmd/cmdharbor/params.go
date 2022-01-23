package cmdharbor

import (
	"errors"
	"strconv"
)

type cmdParams struct {
	SubCmd       string
	ProjectName  string
	RepoName     string
	ArtifactName string
	Page         int64
	Size         int64
}

func newCmdParams(fields []string) (*cmdParams, error) {
	if len(fields) < 2 {
		return nil, errors.New("invalid sub command")
	}

	params := &cmdParams{SubCmd: fields[1]}
	paramsField := fields[2:]
	if len(paramsField)%2 == 1 {
		return nil, errors.New("invalid params")
	}

	for i := 0; i < len(paramsField); i += 2 {
		if paramsField[i] == "proj" {
			if i+1 >= len(paramsField) {
				return nil, errors.New("invalid projectName")
			}
			params.ProjectName = paramsField[i+1]
			continue
		}

		if paramsField[i] == "repo" {
			if i+1 >= len(paramsField) {
				return nil, errors.New("invalid repositoryName")
			}
			params.RepoName = paramsField[i+1]
			continue
		}

		if paramsField[i] == "artifact" {
			if i+1 >= len(paramsField) {
				return nil, errors.New("invalid artifactName")
			}
			params.ArtifactName = paramsField[i+1]
			continue
		}

		if paramsField[i] == "page" {
			page, err := strconv.ParseInt(paramsField[i+1], 10, 64)
			if err != nil {
				return nil, err
			}
			params.Page = page
			continue
		}

		if paramsField[i] == "size" {
			size, err := strconv.ParseInt(paramsField[i+1], 10, 64)
			if err != nil {
				return nil, err
			}
			params.Size = size
			continue
		}
	}

	return params, nil
}
