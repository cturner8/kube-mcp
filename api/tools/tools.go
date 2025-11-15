package tools

import (
	"github.com/cturner8/kube-mcp/config"
	"github.com/cturner8/kube-mcp/kubernetes"

	k8s "k8s.io/client-go/kubernetes"
)

var kubernetesApiClient *k8s.Clientset = kubernetes.CreateKubernetesApiClient(*config.ServerConfig.OutOfCluster, *config.ServerConfig.Kubeconfig)
