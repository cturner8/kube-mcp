package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/homedir"
)

var kubernetesApiClient *kubernetes.Clientset

func main() {
	var (
		host         = flag.String("host", "localhost", "host to connect to/listen on")
		port         = flag.Int("port", 9000, "port number to connect to/listen on")
		outOfCluster = flag.Bool("out-of-cluster", false, "(optional) indicates the server is running outside of a Kubernetes cluster and should look for a kubeconfig file")
	)

	// Attempt to resolve a local kubeconfig path.
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	// Parse command-line flags.
	flag.Parse()

	// Create the Kubernetes API client.
	kubernetesApiClient = createKubernetesApiClient(*outOfCluster, *kubeconfig)

	// Construct the server URL.
	url := fmt.Sprintf("%s:%d", *host, *port)
	// Start the MCP server.
	startServer(url)
}
