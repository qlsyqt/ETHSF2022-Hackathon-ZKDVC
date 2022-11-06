package database

import (
	"time"
)

type ClaimData struct {
	Id          int64     `gorm:"column:pk_id;primaryKey;auto_increment:true"`
	ClaimHi     []byte    `gorm:"column:claim_hi;type:varbinary(32);index:claim_hi,unique"`
	ClaimBinary []byte    `gorm:"column:claim_binary"`
	Issuer      string    `gorm:"column:issuer;index:idx_issuer;type:varchar(255)"`
	Revoked     bool      `gorm:"column:revoked"`
	ClaimType   string    `gorm:"column:claim_type;type:varchar(255)"`
	CreatedAt   time.Time `gorm:"column:created_at""`
	UpdatedAt   time.Time `gorm:"column:updated_at""`
}

func (c ClaimData) TableName() string {
	return "t_claim"
}
