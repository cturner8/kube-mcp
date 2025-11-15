package config

import (
	"flag"
	"path/filepath"
	"strings"

	"k8s.io/client-go/util/homedir"
)

type McpServerCliConfig struct {
	Host           *string
	Port           *int
	OutOfCluster   *bool
	Kubeconfig     *string
	AllowedOrigins []string
}

func getMcpServerCliFlags() McpServerCliConfig {
	var (
		host           = flag.String("host", "", "host to connect to/listen on")
		port           = flag.Int("port", 9000, "port number to connect to/listen on")
		outOfCluster   = flag.Bool("out-of-cluster", false, "(optional) indicates the server is running outside of a Kubernetes cluster and should look for a kubeconfig file")
		allowedOrigins = flag.String("allowed-origins", "", "(optional) comma-separated list of allowed CORS origins")
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

	return McpServerCliConfig{
		Host:           host,
		Port:           port,
		OutOfCluster:   outOfCluster,
		Kubeconfig:     kubeconfig,
		AllowedOrigins: strings.Split(*allowedOrigins, ","),
	}
}
