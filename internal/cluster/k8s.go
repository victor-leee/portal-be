package cluster

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/victor-leee/portal-be/internal/config"
	"github.com/victor-leee/portal-be/internal/model"
	"google.golang.org/protobuf/proto"
	v14 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	v17 "k8s.io/api/networking/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	v15 "k8s.io/client-go/applyconfigurations/apps/v1"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
	v13 "k8s.io/client-go/applyconfigurations/meta/v1"
	v16 "k8s.io/client-go/applyconfigurations/networking/v1"
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

func (k *k8sClusterManager) ApplyServiceIngress(ctx context.Context, cfg *IngressConfig) error {
	_, err := k.client.NetworkingV1().Ingresses(config.NamespaceDefault).Apply(ctx, &v16.IngressApplyConfiguration{
		TypeMetaApplyConfiguration: v13.TypeMetaApplyConfiguration{
			Kind:       proto.String("Ingress"),
			APIVersion: proto.String("networking.k8s.io/v1"),
		},
		ObjectMetaApplyConfiguration: &v13.ObjectMetaApplyConfiguration{
			Name: proto.String(fmt.Sprintf("%s-%s", "ingress", cfg.ServiceUniquePath)),
		},
		Spec: &v16.IngressSpecApplyConfiguration{
			IngressClassName: proto.String("nginx"),
			Rules: []v16.IngressRuleApplyConfiguration{
				{
					IngressRuleValueApplyConfiguration: v16.IngressRuleValueApplyConfiguration{
						HTTP: &v16.HTTPIngressRuleValueApplyConfiguration{
							Paths: []v16.HTTPIngressPathApplyConfiguration{
								{
									Path:     proto.String(cfg.PrefixMappingPath),
									PathType: (*v17.PathType)(proto.String(string(v1beta1.PathTypePrefix))),
									Backend: &v16.IngressBackendApplyConfiguration{
										Service: &v16.IngressServiceBackendApplyConfiguration{
											Name: proto.String(cfg.ServiceUniquePath),
											Port: &v16.ServiceBackendPortApplyConfiguration{
												Number: proto.Int32(config.GlobalDefaultPort),
											},
										},
									},
								},
							},
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

func (k *k8sClusterManager) ApplyServiceDeployment(ctx context.Context, cfg *DeploymentConfig) error {
	deploymentApplyCfg := &v15.DeploymentApplyConfiguration{
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
							Name:            proto.String(cfg.Service.Name),
							Image:           proto.String(cfg.ImageTag),
							ImagePullPolicy: (*v14.PullPolicy)(proto.String(string(v14.PullNever))),
						},
					},
				},
			},
		},
	}

	// if it's scrpc app, attach a side-car container and create a volume for the 2 containers
	if cfg.Service.Type == config.AppTypeSCRPC {
		volumeMounts := []v1.VolumeMountApplyConfiguration{
			{
				Name:      proto.String("shared-data"),
				MountPath: proto.String("/tmp"),
			},
		}
		deploymentApplyCfg.Spec.Template.Spec.Containers = append(deploymentApplyCfg.Spec.Template.Spec.Containers,
			v1.ContainerApplyConfiguration{
				Name:            proto.String("side-car"),
				Image:           proto.String("github.com/victor-leee/side-car:latest"),
				ImagePullPolicy: (*v14.PullPolicy)(proto.String(string(v14.PullNever))),
				VolumeMounts:    volumeMounts,
				Env: []v1.EnvVarApplyConfiguration{
					{
						Name:  proto.String("SC_SERVICE_NAME"),
						Value: proto.String(cfg.Service.Name),
					},
					{
						Name:  proto.String("SC_SERVICE_KEY"),
						Value: proto.String(cfg.Service.ServiceKey),
					},
				},
			})
		deploymentApplyCfg.Spec.Template.Spec.Containers[0].VolumeMounts = volumeMounts
		deploymentApplyCfg.Spec.Template.Spec.Volumes = []v1.VolumeApplyConfiguration{
			{
				Name: proto.String("shared-data"),
				VolumeSourceApplyConfiguration: v1.VolumeSourceApplyConfiguration{
					EmptyDir: &v1.EmptyDirVolumeSourceApplyConfiguration{},
				},
			},
		}
	} else if cfg.Service.Type == config.AppTypeHTTP {
		// for http app, create an ingress pointing to Service:80
		if err := k.ApplyServiceIngress(ctx, &IngressConfig{
			ServiceUniquePath: cfg.Service.UniqueCompletePath,
			PrefixMappingPath: cfg.Service.PrefixMapping,
		}); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("invalid service type:%s", cfg.Service.Type)
	}

	_, err := k.client.AppsV1().Deployments(config.NamespaceDefault).Apply(ctx, deploymentApplyCfg, v12.ApplyOptions{
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
							IntVal: int32(serviceModel.CustomPort),
						},
					},
				},
				Selector: map[string]string{
					config.SelectorServiceID: serviceModel.UniqueCompletePath,
				},
			},
		},
		v12.ApplyOptions{
			FieldManager: config.FieldManagerApplyPatch,
		})
	return err
}
