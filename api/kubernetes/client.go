package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func resolveClusterConfig(outOfCluster bool, kubeconfig *string) *rest.Config {
	if !outOfCluster {
		// creates the in-cluster config
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		return config
	}

	// creates the out-of-cluster config
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return config
}

func CreateKubernetesApiClient(outOfCluster bool, kubeconfig string) *kubernetes.Clientset {
	config := resolveClusterConfig(outOfCluster, &kubeconfig)
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return client
}
