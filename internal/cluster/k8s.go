package cluster

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/victor-leee/portal-be/internal/config"
	"github.com/victor-leee/portal-be/internal/model"
	"google.golang.org/protobuf/proto"
	v14 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
	v13 "k8s.io/client-go/applyconfigurations/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type k8sClusterManager struct {
	client *kubernetes.Clientset
}

func init() {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}
	clientSet, err := kubernetes.NewForConfig(cfg)
	registerManager(K8S, &k8sClusterManager{
		client: clientSet,
	})
	logrus.Info("Init kubernetes client succeed")
}

func (k *k8sClusterManager) CreateServiceInternalRecord(ctx context.Context, serviceModel *model.RPCService) error {
	_, err := k.client.CoreV1().Services(config.NamespaceDefault).Apply(ctx,
		&v1.ServiceApplyConfiguration{
			TypeMetaApplyConfiguration: v13.TypeMetaApplyConfiguration{
				Kind:       proto.String("Service"),
				APIVersion: proto.String("v1"),
			},
			ObjectMetaApplyConfiguration: &v13.ObjectMetaApplyConfiguration{
				Name:      proto.String(serviceModel.UniqueCompletePath),
				Namespace: proto.String(config.NamespaceDefault),
			},
			Spec: &v1.ServiceSpecApplyConfiguration{
				Ports: []v1.ServicePortApplyConfiguration{
					{
						Protocol: (*v14.Protocol)(proto.String(string(v14.ProtocolTCP))),
						Port:     proto.Int32(config.GlobalDefaultPort),
						TargetPort: &intstr.IntOrString{
							Type:   intstr.Int,
							IntVal: config.GlobalDefaultPort,
						},
					},
				},
				Selector: map[string]string{
					config.SelectorService:   serviceModel.Name,
					config.SelectorServiceID: serviceModel.UniqueCompletePath,
				},
			},
		},
		v12.ApplyOptions{
			FieldManager: config.FieldManagerApplyPatch,
		})
	return err
}
