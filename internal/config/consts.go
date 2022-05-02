package config

import (
	"math/big"
)

const (
	ServiceHierarchySeparator = "-"
	ServiceKeySeparator       = "."
)

const (
	EnvKubernetesRegistryIPKey = "MINIKUBE_REGISTRY_IP"
)

// Below are kubernetes constants

const (
	NamespaceDefault  = "default"
	GlobalDefaultPort = 80
	SelectorService   = "service_name"
	SelectorServiceID = "service_id"

	FieldManagerApplyPatch = "application/apply-patch"
)

// some github dev api

const (
	APIRetrieveBranches = "https://api.github.com/repos/%s/%s/branches"
)

const (
	AppTypeSCRPC = "sc_rpc"
	AppTypeHTTP  = "http"
)

var (
	ServiceKeyMaxRandomNumber = big.NewInt(666666)
)
