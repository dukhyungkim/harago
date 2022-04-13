package entity

import (
	"google.golang.org/api/chat/v1"
)

type ComponentType struct {
	ID        uint
	Type      string `gorm:"size:32;not null;default:null"`
	Component string `gorm:"size:32;not null;default:null"`
}

func (c *ComponentType) UniqueFilter() map[string]interface{} {
	filter := make(map[string]interface{})
	filter["component"] = c.Component
	return filter
}

func (c *ComponentType) ToCard() *chat.Card {
	return &chat.Card{
		Sections: []*chat.Section{
			{
				Widgets: []*chat.WidgetMarkup{
					{
						KeyValue: &chat.KeyValue{
							TopLabel:         "component",
							Content:          c.Component,
							ContentMultiline: true,
						},
					},
					{
						KeyValue: &chat.KeyValue{
							TopLabel:         "type",
							Content:          c.Type,
							ContentMultiline: true,
						},
					},
				},
			},
		},
	}
}
