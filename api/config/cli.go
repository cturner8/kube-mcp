package config

import (
	"flag"
	"os"
	"path/filepath"
	"sync"

	"k8s.io/client-go/util/homedir"
)

var (
	flagsOnce       sync.Once
	host            *string
	port            *string
	outOfCluster    *bool
	allowedOrigins  *string
	baseUrl         *string
	oidcIssuerUrl   *string
	oidcClientId    *string
	signingMethod   *string
	scopes          *string
	logLevel        *string
	allowedTools    *string
	disallowedTools *string
	kubeconfig      *string
)

func initFlags() {
	flagsOnce.Do(func() {
		host = flag.String("host", "", "host to connect to/listen on")
		port = flag.String("port", "", "port number to connect to/listen on")
		outOfCluster = flag.Bool("out-of-cluster", false, "(optional) indicates the server is running outside of a Kubernetes cluster and should look for a kubeconfig file")
		allowedOrigins = flag.String("allowed-origins", "", "(optional) comma-separated list of allowed CORS origins")
		baseUrl = flag.String("base-url", "", "Base URL the application will be accessed from")
		oidcIssuerUrl = flag.String("oidc-issuer-url", "", "URL of the OIDC authentication provider")
		oidcClientId = flag.String("oidc-client-id", "", "ID of the OIDC Client to authenticate against")
		signingMethod = flag.String("oidc-signing-method", "", "Signing method for JWTs (HS256 or RS256)")
		scopes = flag.String("oidc-scopes", "", "(optional) comma-separated list of OIDC scopes to request during authentication")
		logLevel = flag.String("log-level", "", "Application log level: debug, info, warn, error")
		allowedTools = flag.String("allowed-tools", "", "(optional) comma-separated list of allowed tools")
		disallowedTools = flag.String("disallowed-tools", "", "(optional) comma-separated list of disallowed tools")

		// Attempt to resolve a local kubeconfig path.
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
	})
}

func getMcpServerCliFlags() McpServerUserConfig {
	initFlags()

	// Parse command-line flags if not already loaded
	if !flag.CommandLine.Parsed() {
		flag.Parse()
	}

	// Read values from environment variables if flags weren't set
	hostVal := *host
	if hostVal == "" {
		hostVal = os.Getenv("KUBE_MCP_HOST")
	}
	portVal := *port
	if portVal == "" {
		portVal = os.Getenv("KUBE_MCP_PORT")
	}
	outOfClusterVal := *outOfCluster
	if !outOfClusterVal {
		outOfClusterVal = os.Getenv("KUBE_MCP_OUT_OF_CLUSTER") == "true"
	}
	allowedOriginsVal := *allowedOrigins
	if allowedOriginsVal == "" {
		allowedOriginsVal = os.Getenv("KUBE_MCP_ALLOWED_ORIGINS")
	}
	baseUrlVal := *baseUrl
	if baseUrlVal == "" {
		baseUrlVal = os.Getenv("KUBE_MCP_BASE_URL")
	}
	oidcIssuerUrlVal := *oidcIssuerUrl
	if oidcIssuerUrlVal == "" {
		oidcIssuerUrlVal = os.Getenv("KUBE_MCP_OIDC_ISSUER_URL")
	}
	oidcClientIdVal := *oidcClientId
	if oidcClientIdVal == "" {
		oidcClientIdVal = os.Getenv("KUBE_MCP_OIDC_CLIENT_ID")
	}
	signingMethodVal := *signingMethod
	if signingMethodVal == "" {
		signingMethodVal = os.Getenv("KUBE_MCP_OIDC_SIGNING_METHOD")
	}
	scopesVal := *scopes
	if scopesVal == "" {
		scopesVal = os.Getenv("KUBE_MCP_OIDC_SCOPES")
	}
	logLevelVal := *logLevel
	if logLevelVal == "" {
		logLevelVal = os.Getenv("KUBE_MCP_LOG_LEVEL")
	}
	allowedToolsVal := *allowedTools
	if allowedToolsVal == "" {
		allowedToolsVal = os.Getenv("KUBE_MCP_ALLOWED_TOOLS")
	}
	disallowedToolsVal := *disallowedTools
	if disallowedToolsVal == "" {
		disallowedToolsVal = os.Getenv("KUBE_MCP_DISALLOWED_TOOLS")
	}

	return McpServerUserConfig{
		Host:            hostVal,
		Port:            portVal,
		OutOfCluster:    outOfClusterVal,
		Kubeconfig:      *kubeconfig,
		AllowedOrigins:  allowedOriginsVal,
		BaseURL:         baseUrlVal,
		OidcIssuerURL:   oidcIssuerUrlVal,
		OidcClientID:    oidcClientIdVal,
		LogLevel:        logLevelVal,
		AllowedTools:    allowedToolsVal,
		DisallowedTools: disallowedToolsVal,
		SigningMethod:   signingMethodVal,
		Scopes:          scopesVal,
	}
}
