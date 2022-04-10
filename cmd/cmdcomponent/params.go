package cmdcomponent

import (
	"errors"
	"log"

	"github.com/jessevdk/go-flags"
)

type cmdParams struct {
	SubCmd    string
	Company   string `long:"company"`
	Component string
}

func newCmdParams(fields []string) (*cmdParams, error) {
	if len(fields) < 1 {
		return nil, errors.New("invalid sub command")
	}

	params := cmdParams{SubCmd: fields[0]}
	args, err := flags.ParseArgs(&params, fields[1:])
	if err != nil {
		return nil, err
	}

	//if len(args) == 0 {
	//	args[0]
	//}
	params.Component = args[0]
	log.Printf("%+#v\n", params)

	return &params, nil
}
