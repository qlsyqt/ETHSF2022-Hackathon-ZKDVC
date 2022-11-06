package store

import (
	"errors"
	"log"
	"time"
	"wallet-end/models/database"
	"wallet-end/pkgs/orm"
)

func SelectDidById(did string) (*database.DidData, error) {
	didData := database.DidData{}
	db := orm.GetMySqlDb()
	if err := db.Where("Did=?", did).First(&didData).Error; err != nil {
		log.Printf("select did failed, id:%s", did)
		return nil, err
	}
	return &didData, nil
}

func InsertDid(did *database.DidData) error {
	if err := orm.GetMySqlDb().Create(did).Error; err != nil {
		log.Println("insert did failed")
		return err
	}
	return nil
}

func UpdateDid(did *database.DidData) error {
	t := orm.GetMySqlDb().Model(did).Updates(map[string]interface{}{"is_genesis": 0, "updated_at": time.Now()})

	if t.Error != nil {
		return t.Error
	}
	if t.RowsAffected == 0 {
		return errors.New("No rows affected")
	}
	return nil
}
