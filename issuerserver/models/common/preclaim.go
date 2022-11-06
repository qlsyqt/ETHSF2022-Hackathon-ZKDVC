package common

import "time"

type PreClaim struct {
	Name           string    `json:"name"`
	DataCategory   string    `json:"datacategory"`
	LowerBound     []int32   `json:"lowerBound"`
	UpperBound     []int32   `json:"upperBound"`
	CreateDate     time.Time `json:"createDate"`
	ExpirationDate string    `json:"expirationDate"`
}
