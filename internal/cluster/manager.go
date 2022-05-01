package cluster

import (
	"context"
	"github.com/victor-leee/portal-be/internal/model"
	"sync"
)

type Manager interface {
	// CreateServiceInternalRecord adds a cname record to DNS center in the cluster
	// it only functions within the cluster and therefore cannot be accessed externally
	CreateServiceInternalRecord(ctx context.Context, service *model.RPCService) error
}

type ManagerType uint8

const (
	K8S ManagerType = 0
)

var (
	registerMux = sync.Mutex{}
	key2Manager = make(map[ManagerType]Manager)
)

func registerManager(managerKey ManagerType, m Manager) {
	defer registerMux.Unlock()
	registerMux.Lock()
	if _, ok := key2Manager[managerKey]; ok {
		panic("key exist")
	}
	key2Manager[managerKey] = m
}

func GetManager(managerType ManagerType) Manager {
	return key2Manager[managerType]
}
