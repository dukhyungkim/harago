package cmdcomponent

import "log"

type SetCommand struct {
	Company string `long:"company"`
	Type    string `long:"type"`
}

func (c *SetCommand) Execute(args []string) error {
	log.Println(c.Company, c.Type)
	log.Println(args)
	return nil
}
