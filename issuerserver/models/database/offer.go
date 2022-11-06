package database

import "time"

type Offer struct {
	Id         int64  `gorm:"column:id;primaryKey;autoIncrement:true"`
	Name       string `gorm:"column:name"`
	TemplateId int64  `gorm:"column:templateId"`
	PreClaims  string `gorm:"column:preClaims;type:text"`
	//Claimer    string    `gorm:"claimer;type:varchar(255)"`
	IssuerDid string    `gorm:"column:issuerDid;type:varchar(255)"`
	Link      string    `gorm:"column:link;type:varchar(255)"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}

func (o Offer) TableName() string {
	return "offer"
}
