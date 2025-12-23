package config

import (
	"flag"
	"os"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

func getMcpServerCliFlags() McpServerUserConfig {
	var (
		host            = flag.String("host", os.Getenv("HOST"), "host to connect to/listen on")
		port            = flag.String("port", os.Getenv("PORT"), "port number to connect to/listen on")
		outOfCluster    = flag.Bool("out-of-cluster", false, "(optional) indicates the server is running outside of a Kubernetes cluster and should look for a kubeconfig file")
		allowedOrigins  = flag.String("allowed-origins", os.Getenv("ALLOWED_ORIGINS"), "(optional) comma-separated list of allowed CORS origins")
		baseUrl         = flag.String("base-url", os.Getenv("BASE_URL"), "Base URL the application will be accessed from")
		oidcIssuerUrl   = flag.String("oidc-issuer-url", os.Getenv("OIDC_ISSUER_URL"), "URL of the OIDC authentication provider")
		oidcClientId    = flag.String("oidc-client-id", os.Getenv("OIDC_CLIENT_ID"), "ID of the OIDC Client to authenticate against")
		allowedTools    = flag.String("allowed-tools", os.Getenv("ALLOWED_TOOLS"), "(optional) comma-separated list of allowed tools")
		disallowedTools = flag.String("disallowed-tools", os.Getenv("DISALLOWED_TOOLS"), "(optional) comma-separated list of disallowed tools")
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

	return McpServerUserConfig{
		Host:            *host,
		Port:            *port,
		OutOfCluster:    *outOfCluster,
		Kubeconfig:      *kubeconfig,
		AllowedOrigins:  *allowedOrigins,
		BaseURL:         *baseUrl,
		OidcIssuerURL:   *oidcIssuerUrl,
		OidcClientID:    *oidcClientId,
		AllowedTools:    *allowedTools,
		DisallowedTools: *disallowedTools,
	}
}
