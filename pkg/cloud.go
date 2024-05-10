package ccm

import (
	"io"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	cloudprovider "k8s.io/cloud-provider"
)

const ProviderName = "zpcc"

type CloudProvider struct {
	loadBalancer cloudprovider.LoadBalancer
	instances    cloudprovider.InstancesV2
}

var _ cloudprovider.Interface = &CloudProvider{}

func init() {
	cloudprovider.RegisterCloudProvider(ProviderName, newCloudProvider)
}

func newCloudProvider(reader io.Reader) (cloudprovider.Interface, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	config, err := clientcmd.NewClientConfigFromBytes(bytes)
	if err != nil {
		return nil, err
	}

	clientConfig, err := config.ClientConfig()
	if err != nil {
		return nil, err
	}
	rawConfig, err := config.RawConfig()
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}

	namespace := rawConfig.Contexts[rawConfig.CurrentContext].Namespace

	return &CloudProvider{
		loadBalancer: &loadBalancerManager{
			kubeClient: client,
			namespace:  namespace,
		},
	}, nil
}

func (c *CloudProvider) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
	clientset := clientBuilder.ClientOrDie(ProviderName)
	sharedInformer := informers.NewSharedInformerFactory(clientset, 0)

	sharedInformer.Start(stop)
	sharedInformer.WaitForCacheSync(stop)
}

func (c *CloudProvider) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

func (c *CloudProvider) HasClusterID() bool {
	return false
}

func (c *CloudProvider) Instances() (cloudprovider.Instances, bool) {
	return nil, false
}

func (c *CloudProvider) InstancesV2() (cloudprovider.InstancesV2, bool) {
	return c.instances, true
}

func (c *CloudProvider) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return c.loadBalancer, true
}

func (c *CloudProvider) ProviderName() string {
	return ProviderName
}

func (c *CloudProvider) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

func (c *CloudProvider) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}
