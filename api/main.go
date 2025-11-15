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
	// Start the MCP server.
	startServer(url, config.PublicBaseURL.String())
}
