package repository

import (
	"harago/entity"
	"log"

	"gorm.io/gorm/clause"
)

func (db *DB) UpsertComponentMapping(cm *entity.ComponentMapping) error {
	return db.client.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "company"}, {Name: "type"}},
		DoUpdates: clause.AssignmentColumns([]string{"component"}),
	}).Create(cm).Error
}

func (db *DB) ListComponentMappings() ([]*entity.ComponentMapping, error) {
	var results []*entity.ComponentMapping
	if err := db.client.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (db *DB) DeleteComponentMapping(cm *entity.ComponentMapping) error {
	tx := db.client.Where(cm.UniqueFilter()).Delete(&entity.ComponentMapping{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		log.Printf("nothing to delete; %+#v\n", cm.UniqueFilter())
	}
	return nil
}
