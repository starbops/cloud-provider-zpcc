package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/wait"
	cloudprovider "k8s.io/cloud-provider"
	"k8s.io/cloud-provider/app"
	cloudcontrollerconfig "k8s.io/cloud-provider/app/config"
	"k8s.io/cloud-provider/names"
	"k8s.io/cloud-provider/options"
	cliflag "k8s.io/component-base/cli/flag"

	_ "github.com/starbops/cloud-provider-zpcc/pkg/ccm"
)

func cloudInitializer(config *cloudcontrollerconfig.CompletedConfig) cloudprovider.Interface {
	cloudConfig := config.ComponentConfig.KubeCloudShared.CloudProvider

	cloud, err := cloudprovider.InitCloudProvider(cloudConfig.Name, cloudConfig.CloudConfigFile)
	if err != nil {
		logrus.Fatalf("cloud provider zpcc could not be initialized: %v", err)
	}

	if cloud == nil {
		logrus.Fatal("cloud provider zpcc is nil")
	}

	if !cloud.HasClusterID() {
		if config.ComponentConfig.KubeCloudShared.AllowUntaggedCloud {
			logrus.Warning("detected a cluster without a ClusterID. A ClusterID will be required in the future. Please tag your cluster to avoid any future issues")
		} else {
			logrus.Fatal("no ClusterID found. A ClusterID is required for the cloud zpcc to function properly. This check can be bypassed by setting the allow-untagged-cloud option")
		}
	}

	return cloud
}

func main() {
	ccmOptions, err := options.NewCloudControllerManagerOptions()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"hello": "world",
		}).Fatal("unable to initialize command options")
	}

	controllerInitializers := app.DefaultInitFuncConstructors
	controllerAliases := names.CCMControllerAliases()

	fss := cliflag.NamedFlagSets{}

	command := app.NewCloudControllerManagerCommand(ccmOptions, cloudInitializer, controllerInitializers, controllerAliases, fss, wait.NeverStop)

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
