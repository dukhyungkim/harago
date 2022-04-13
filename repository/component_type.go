package repository

import (
	"harago/entity"
	"log"

	"gorm.io/gorm/clause"
)

func (db *DB) UpsertComponentType(ct *entity.ComponentType) error {
	return db.client.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "component"}},
		DoUpdates: clause.AssignmentColumns([]string{"type"}),
	}).Create(ct).Error
}

func (db *DB) ListComponentTypes() ([]*entity.ComponentType, error) {
	var results []*entity.ComponentType
	if err := db.client.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (db *DB) DeleteComponentType(ct *entity.ComponentType) error {
	tx := db.client.Where(ct.UniqueFilter()).Delete(&entity.ComponentType{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		log.Printf("nothing to delete; %+#v\n", ct.UniqueFilter())
	}
	return nil
}
