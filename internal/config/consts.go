package config

import "math/big"

const (
	ServiceHierarchySeparator = "-"
	ServiceKeySeparator       = "."
)

// Below are kubernetes constants

const (
	NamespaceDefault  = "default"
	GlobalDefaultPort = 80
	SelectorService   = "service_name"
	SelectorServiceID = "service_id"

	FieldManagerApplyPatch = "application/apply-patch"
)

var (
	ServiceKeyMaxRandomNumber = big.NewInt(666666)
)
