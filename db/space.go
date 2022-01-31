package db

import (
	"harago/entity"
	"log"
)

func (db *DB) SaveSpace(userSpace *entity.UserSpace) error {
	return db.client.Save(userSpace).Error
}

func (db *DB) DeleteSpace(email string) {
	if err := db.client.Where(&entity.UserSpace{Email: email}).Delete(&entity.UserSpace{}).Error; err != nil {
		log.Println(err)
	}
}
