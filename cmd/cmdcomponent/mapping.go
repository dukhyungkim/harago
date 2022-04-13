package cmdcomponent

import (
	"fmt"
	"harago/entity"

	"google.golang.org/api/chat/v1"
)

type MappingOpts struct {
	Company string `long:"company"`
	Type    string `long:"type"`
}

type Mapping struct {
	Set    MappingOpts `command:"set"`
	List   struct{}    `command:"list" alias:"ls"`
	Remove MappingOpts `command:"remove" alias:"rm"`
}

func (m *Mapping) Run(helper *cmdRunHelper) *chat.Message {
	switch helper.cmd {
	case subCmdSet:
		if len(helper.args) == 0 {
			return &chat.Message{Text: "not enough argument"}
		}
		cm := entity.ComponentMapping{Company: m.Set.Company, Type: m.Set.Type, Component: helper.args[0]}
		err := helper.db.UpsertComponentMapping(&cm)
		if err != nil {
			return &chat.Message{Text: err.Error()}
		}

	case subCmdList:
		var cms []*entity.ComponentMapping
		cms, err := helper.db.ListComponentMappings()
		if err != nil {
			return &chat.Message{Text: err.Error()}
		}

		cards := make([]*chat.Card, len(cms))
		for i := range cms {
			cards[i] = cms[i].ToCard()
		}
		return &chat.Message{Text: "List of ComponentMappings", Cards: cards}

	case subCmdRemove:
		cm := entity.ComponentMapping{Company: m.Remove.Company, Type: m.Remove.Type}
		err := helper.db.DeleteComponentMapping(&cm)
		if err != nil {
			return &chat.Message{Text: err.Error()}
		}

	default:
		return &chat.Message{Text: fmt.Sprintf("sub command not found: %s", helper.cmd)}
	}

	return &chat.Message{Text: "Done"}
}
