package database

import "time"

type DidData struct {
	Did            string    `gorm:"column:did;primaryKey"`
	AuthPrivateKey []byte    `gorm:"column:auth_private_key;type:varbinary(32)"`
	AuthClaimHi    []byte    `gorm:"column:auth_claim_hi;type:varbinary(32)"`
	Username       string    `gorm:"uniqueIndex:idx_first_second;column:username;type:varchar(255)"`
	IsGenesis      bool      `gorm:"column:is_genesis"`
	CltId          uint64    `gorm:"column:clt_id"`
	RotId          uint64    `gorm:"column:rot_id"`
	TorId          uint64    `gorm:"column:tor_id"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (c DidData) TableName() string {
	return "t_did"
}
