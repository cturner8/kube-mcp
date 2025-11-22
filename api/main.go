package main

import (
	"log"

	"github.com/cturner8/kube-mcp/config"
	"github.com/cturner8/kube-mcp/kubernetes"
	"github.com/cturner8/kube-mcp/server"
)

func main() {
	kubernetesApiClient := kubernetes.CreateKubernetesApiClient(*config.ServerConfig.OutOfCluster, *config.ServerConfig.Kubeconfig)

	version, err := kubernetesApiClient.Discovery().ServerVersion()
	if err != nil {
		panic(err.Error())
	}

	log.Printf("Connected to Kubernetes API Server. Version %s.", version.String())

	server.StartServer()
}
