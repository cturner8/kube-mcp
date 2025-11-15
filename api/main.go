package main

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
)

var kubernetesApiClient *kubernetes.Clientset
var config McpServerConfig

func main() {
	config = getMcpServerConfig()

	// Create the Kubernetes API client.
	kubernetesApiClient = createKubernetesApiClient(*config.OutOfCluster, *config.Kubeconfig)

	// Construct the server URL.
	url := fmt.Sprintf("%s:%d", *config.Host, *config.Port)
	baseUrl := fmt.Sprintf("http://%s", url) // TODO: Fix protocol
	// Start the MCP server.
	startServer(url, baseUrl)
}
