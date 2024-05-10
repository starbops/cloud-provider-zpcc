package ccm

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type loadBalancerManager struct {
	kubeClient *kubernetes.Clientset
	namespace  string
}

func (lbm *loadBalancerManager) GetLoadBalancer(ctx context.Context, clusterName string, service *v1.Service) (status *v1.LoadBalancerStatus, exists bool, err error) {
	return nil, false, nil
}

func (lbm *loadBalancerManager) GetLoadBalancerName(ctx context.Context, clusterName string, service *v1.Service) string {
	return ""
}

func (lbm *loadBalancerManager) EnsureLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) (*v1.LoadBalancerStatus, error) {
	return nil, nil
}

func (lbm *loadBalancerManager) UpdateLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) error {
	return nil
}

func (lbm *loadBalancerManager) EnsureLoadBalancerDeleted(ctx context.Context, clusterName string, service *v1.Service) error {
	return nil
}
