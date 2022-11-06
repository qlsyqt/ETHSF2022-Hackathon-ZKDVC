package database

import "time"

type Template struct {
	Id              int64     `gorm:"column:id;primaryKey;autoIncrement:true"`
	TemplateName    string    `gorm:"column:name;type:varchar(255)"`
	DataCategory    string    `gorm:"column:dataCategory;type:varchar(255)"`
	SubCategory     string    `gorm:"column:subCategory;type:varchar(255)"`
	IsExpirable     bool      `gorm:"column:isExpirable"`
	IsAutoRevokable bool      `gorm:"column:isAutoRevokable"`
	Classfications  string    `gorm:"column:classfications;type:text"`
	CreatedAt       time.Time `gorm:"column:createdAt"`
	UpdatedAt       time.Time `gorm:"column:updatedAt"`
}

func (t Template) TableName() string {
	return "template"
}
