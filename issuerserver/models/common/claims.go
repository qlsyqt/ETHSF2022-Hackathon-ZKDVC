package common

import "time"

type ClaimRequest struct {
	//发行者did
	IssuerDid string
	//持有者did，亦为数据持有者
	HolderDid string
	//过期日期
	ExpiredAt time.Time
	//数据主类型（NFT/Snapshot等）
	DataCategory string
	//数据子类型（合约id等）
	SubCategory string
	//下界是否包含
	LowerBoundInclude bool
	//下界的值
	LowerBoundValue int32
	//上界是否包含
	UpperBoundInclude bool
	//上界的值
	UpperBoundValue int32
	//持有者钱包地址
	HolderAddress string
}
