package config

import (
	"flag"
	"os"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

func getMcpServerCliFlags() McpServerUserConfig {
	var (
		host            = flag.String("host", os.Getenv("KUBE_MCP_HOST"), "host to connect to/listen on")
		port            = flag.String("port", os.Getenv("KUBE_MCP_PORT"), "port number to connect to/listen on")
		outOfCluster    = flag.Bool("out-of-cluster", os.Getenv("KUBE_MCP_OUT_OF_CLUSTER") == "true", "(optional) indicates the server is running outside of a Kubernetes cluster and should look for a kubeconfig file")
		allowedOrigins  = flag.String("allowed-origins", os.Getenv("KUBE_MCP_ALLOWED_ORIGINS"), "(optional) comma-separated list of allowed CORS origins")
		baseUrl         = flag.String("base-url", os.Getenv("KUBE_MCP_BASE_URL"), "Base URL the application will be accessed from")
		oidcIssuerUrl   = flag.String("oidc-issuer-url", os.Getenv("KUBE_MCP_OIDC_ISSUER_URL"), "URL of the OIDC authentication provider")
		oidcClientId    = flag.String("oidc-client-id", os.Getenv("KUBE_MCP_OIDC_CLIENT_ID"), "ID of the OIDC Client to authenticate against")
		signingMethod   = flag.String("oidc-signing-method", os.Getenv("KUBE_MCP_OIDC_SIGNING_METHOD"), "Signing method for JWTs (HS256 or RS256)")
		scopes          = flag.String("oidc-scopes", os.Getenv("KUBE_MCP_OIDC_SCOPES"), "(optional) comma-separated list of OIDC scopes to request during authentication")
		logLevel        = flag.String("log-level", os.Getenv("KUBE_MCP_LOG_LEVEL"), "Application log level: debug, info, warn, error")
		allowedTools    = flag.String("allowed-tools", os.Getenv("KUBE_MCP_ALLOWED_TOOLS"), "(optional) comma-separated list of allowed tools")
		disallowedTools = flag.String("disallowed-tools", os.Getenv("KUBE_MCP_DISALLOWED_TOOLS"), "(optional) comma-separated list of disallowed tools")
	)

	// Attempt to resolve a local kubeconfig path.
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	// Parse command-line flags if not already loaded
	if !flag.CommandLine.Parsed() {
		flag.Parse()
	}

	return McpServerUserConfig{
		Host:            *host,
		Port:            *port,
		OutOfCluster:    *outOfCluster,
		Kubeconfig:      *kubeconfig,
		AllowedOrigins:  *allowedOrigins,
		BaseURL:         *baseUrl,
		OidcIssuerURL:   *oidcIssuerUrl,
		OidcClientID:    *oidcClientId,
		LogLevel:        *logLevel,
		AllowedTools:    *allowedTools,
		DisallowedTools: *disallowedTools,
		SigningMethod:   *signingMethod,
		Scopes:          *scopes,
	}
}
