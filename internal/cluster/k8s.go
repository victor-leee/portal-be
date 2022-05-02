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
	v15 "k8s.io/client-go/applyconfigurations/apps/v1"
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

func (k *k8sClusterManager) ApplyServiceDeployment(ctx context.Context, cfg *DeploymentConfig) error {
	_, err := k.client.AppsV1().Deployments(config.NamespaceDefault).Apply(ctx,
		&v15.DeploymentApplyConfiguration{
			TypeMetaApplyConfiguration: v13.TypeMetaApplyConfiguration{
				Kind:       proto.String("Deployment"),
				APIVersion: proto.String("apps/v1"),
			},
			ObjectMetaApplyConfiguration: &v13.ObjectMetaApplyConfiguration{
				Name:      proto.String(cfg.Service.UniqueCompletePath),
				Namespace: proto.String(config.NamespaceDefault),
			},
			Spec: &v15.DeploymentSpecApplyConfiguration{
				Replicas: proto.Int32(cfg.Replicas),
				Selector: &v13.LabelSelectorApplyConfiguration{
					MatchLabels: map[string]string{
						config.SelectorServiceID: cfg.Service.UniqueCompletePath,
					},
				},
				Template: &v1.PodTemplateSpecApplyConfiguration{
					ObjectMetaApplyConfiguration: &v13.ObjectMetaApplyConfiguration{
						Labels: map[string]string{
							config.SelectorServiceID: cfg.Service.UniqueCompletePath,
						},
					},
					Spec: &v1.PodSpecApplyConfiguration{
						Containers: []v1.ContainerApplyConfiguration{
							{
								Name:  proto.String(cfg.Service.Name),
								Image: proto.String(cfg.ImageTag),
								Ports: []v1.ContainerPortApplyConfiguration{
									{
										ContainerPort: proto.Int32(config.GlobalDefaultPort),
									},
								},
								ImagePullPolicy: (*v14.PullPolicy)(proto.String(string(v14.PullNever))),
							},
						},
					},
				},
			},
		}, v12.ApplyOptions{
			FieldManager: config.FieldManagerApplyPatch,
		})
	return err
}

func (k *k8sClusterManager) ApplyServiceInternalDNSRecord(ctx context.Context, serviceModel *model.RPCService) error {
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
