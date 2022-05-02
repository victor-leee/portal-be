package cluster

import (
	"context"
	"github.com/victor-leee/portal-be/internal/model"
	"sync"
)

// TODO support self-defined port

type DeploymentConfig struct {
	ImageTag string
	Service  *model.RPCService
	Replicas int32
}

type IngressConfig struct {
	PrefixMappingPath string
	ServiceUniquePath string
}

type Manager interface {
	// ApplyServiceInternalDNSRecord adds a cname record to DNS center in the cluster
	// it only functions within the cluster and therefore cannot be accessed externally
	ApplyServiceInternalDNSRecord(ctx context.Context, service *model.RPCService) error
	// ApplyServiceDeployment deploys the image with cfg
	ApplyServiceDeployment(ctx context.Context, cfg *DeploymentConfig) error
	// ApplyServiceIngress maps prefix to a Service
	ApplyServiceIngress(ctx context.Context, cfg *IngressConfig) error
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
