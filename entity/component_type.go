package entity

import (
	"time"

	"google.golang.org/api/chat/v1"
)

type ComponentType struct {
	ID        uint
	Company   string `gorm:"size:32;not null;default:null"`
	Type      string `gorm:"size:16;not null;default:null"`
	Component string `gorm:"size:16;not null;default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *ComponentType) UniqueFilter() map[string]interface{} {
	filter := make(map[string]interface{})
	filter["company"] = c.Company
	filter["type"] = c.Type
	return filter
}

func (c *ComponentType) ToCard() *chat.Card {
	return &chat.Card{
		Sections: []*chat.Section{
			{
				Widgets: []*chat.WidgetMarkup{
					{
						KeyValue: &chat.KeyValue{
							TopLabel:         "company",
							Content:          c.Company,
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
					{
						KeyValue: &chat.KeyValue{
							TopLabel:         "component",
							Content:          c.Component,
							ContentMultiline: true,
						},
					},
					{
						KeyValue: &chat.KeyValue{
							TopLabel:         "created_at",
							Content:          c.CreatedAt.Local().String(),
							ContentMultiline: true,
						},
					},
					{
						KeyValue: &chat.KeyValue{
							TopLabel:         "updated_at",
							Content:          c.UpdatedAt.Local().String(),
							ContentMultiline: true,
						},
					},
				},
			},
		},
	}
}
