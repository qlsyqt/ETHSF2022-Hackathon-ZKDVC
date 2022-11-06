package store

import (
	"issuerserver/models/database"
	"issuerserver/pkg/orm"
)

func InsertTemplate(template *database.Template) error {
	if err := orm.GetMySqlDb().Create(template).Error; err != nil {
		return err
	}
	return nil
}

func SelectTemplateById(id int64) (*database.Template, error) {
	template := database.Template{}
	db := orm.GetMySqlDb()
	//Where 采用对象中的字段名，而非db字段
	if err := db.Where("Id=?", id).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}
