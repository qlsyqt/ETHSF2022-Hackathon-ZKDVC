package store

import (
	data "issuerserver/models/database"
	"issuerserver/pkg/orm"
)

func InsertOffer(offer *data.Offer) error {
	if err := orm.GetMySqlDb().Create(offer).Error; err != nil {
		return err
	}
	return nil
}

func SelectOfferById(offerId int64) (*data.Offer, error) {
	offer := data.Offer{}
	db := orm.GetMySqlDb()
	//Where 采用对象中的字段名，而非db字段
	if err := db.Where("Id=?", offerId).First(&offer).Error; err != nil {
		return nil, err
	}
	return &offer, nil
}
