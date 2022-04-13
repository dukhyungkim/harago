package cmdcomponent

import (
	"fmt"
	"harago/entity"
	"log"

	"google.golang.org/api/chat/v1"
)

type TypeOpts struct {
	Component string `long:"component"`
}

type Type struct {
	Add    TypeOpts `command:"add"`
	List   struct{} `command:"list" alias:"ls"`
	Remove TypeOpts `command:"remove" alias:"rm"`
}

func (t *Type) Run(helper *cmdRunHelper) *chat.Message {
	switch helper.cmd {
	case subCmdAdd:
		if len(helper.args) == 0 {
			return &chat.Message{Text: "not enough argument"}
		}
		ct := entity.ComponentType{Component: t.Add.Component, Type: helper.args[0]}
		log.Println(ct)
		err := helper.db.UpsertComponentType(&ct)
		if err != nil {
			return &chat.Message{Text: err.Error()}
		}

	case subCmdList:
		var cts []*entity.ComponentType
		cts, err := helper.db.ListComponentTypes()
		if err != nil {
			return &chat.Message{Text: err.Error()}
		}

		cards := make([]*chat.Card, len(cts))
		for i := range cts {
			cards[i] = cts[i].ToCard()
		}
		return &chat.Message{Text: "List of ComponentTypes", Cards: cards}

	case subCmdRemove:
		ct := entity.ComponentType{Component: t.Remove.Component}
		err := helper.db.DeleteComponentType(&ct)
		if err != nil {
			return &chat.Message{Text: err.Error()}
		}

	default:
		return &chat.Message{Text: fmt.Sprintf("sub command not found: %s", helper.cmd)}
	}

	return &chat.Message{Text: "Done"}
}
