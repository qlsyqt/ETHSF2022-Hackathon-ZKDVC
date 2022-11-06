package common

type DCP struct {
	DataCategory  string
	SubCategory   string
	LowerBoundary Boundary
	UpperBoundary Boundary
	HolderAddress string
}

type Boundary struct {
	Include bool
	Value   int32
}
