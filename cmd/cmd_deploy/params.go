package cmd_deploy

import (
	"errors"
)

type cmdParams struct {
	Company     string
	ResourceURL string
}

func newCmdParams(fields []string) (*cmdParams, error) {
	if len(fields) < 1 {
		return nil, errors.New("invalid params")
	}

	params := &cmdParams{}
	for i := 0; i < len(fields); i += 1 {
		if fields[i] == "company" {
			if i+1 >= len(fields) {
				return nil, errors.New("invalid company")
			}
			params.Company = fields[i+1]
			i++
			continue
		}

		if params.ResourceURL == "" {
			params.ResourceURL = fields[i]
		}
	}

	return params, nil
}
