package config

import "math/big"

const (
	ServiceHierarchySeparator = "-"
	ServiceKeySeparator       = "."
)

var (
	ServiceKeyMaxRandomNumber = big.NewInt(666666)
)
