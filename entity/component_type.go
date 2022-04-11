package entity

import (
	"time"
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
