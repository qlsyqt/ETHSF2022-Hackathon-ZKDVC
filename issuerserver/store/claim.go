package store

import (
	"errors"
	"fmt"
	data "issuerserver/models/database"
	"issuerserver/pkg/db"
	"issuerserver/pkg/orm"
	"math/big"
)

func InsertClaim(claim *data.ClaimData) error {
	if err := orm.GetMySqlDb().Create(claim).Error; err != nil {
		return err
	}
	return nil
}

func SelectClaimByHIndex(id *big.Int) (*data.ClaimData, error) {
	claim := data.ClaimData{}
	db := orm.GetMySqlDb()
	if err := db.Where("claim_hi=?", id.Bytes()).First(&claim).Error; err != nil {
		return nil, err
	}
	return &claim, nil
}

func SelectClaimsByBatch(id int64, limit int) ([]*data.ClaimData, error) {
	table := data.ClaimData{}.TableName()
	sql := fmt.Sprintf("select * from %s where pk_id > ? limit %d", table, limit)

	conn, err := db.NewMySqlDb().GetConn()
	if err != nil {
		return nil, err
	}
	ps, err := conn.Prepare(sql)
	if err != nil {
		return nil, err
	}

	rows, err := ps.Query(id)
	if err != nil {
		return nil, err
	}

	ans := make([]*data.ClaimData, 0)
	for rows.Next() {
		claim := data.ClaimData{}
		err = rows.Scan(&claim.Id, &claim.ClaimHi, &claim.ClaimBinary, &claim.Issuer, &claim.Revoked, &claim.ClaimType, &claim.CreatedAt, &claim.UpdatedAt)
		if err != nil {
			return nil, err
		}
		ans = append(ans, &claim)
	}
	return ans, nil
}

func UpdateClaimRevokeStatus(id int64, status bool) error {
	claim := &data.ClaimData{
		Id: id,
	}
	t := orm.GetMySqlDb().Model(claim).Updates(map[string]interface{}{"revoked": status})
	if t.Error != nil {
		return t.Error
	}
	if t.RowsAffected == 0 {
		return errors.New("No rows affected")
	}
	return nil
}
