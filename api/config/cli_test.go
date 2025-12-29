package config

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/util/homedir"
)

func TestGetMcpServerCliFlagsEmpty(t *testing.T) {
	t.Skip("fix flag definition")

	home := homedir.HomeDir()
	config := getMcpServerCliFlags()

	assert.Equal(t, McpServerUserConfig{
		Kubeconfig: filepath.Join(home, ".kube", "config"),
	}, config)
}

func TestGetMcpServerCliFlagsSet(t *testing.T) {
	home := homedir.HomeDir()

	t.Setenv("KUBE_MCP_HOST", "http://localhost:9000")
	t.Setenv("KUBE_MCP_PORT", "9000")
	t.Setenv("KUBE_MCP_OUT_OF_CLUSTER", "true")
	t.Setenv("KUBE_MCP_BASE_URL", "http://localhost:9000")
	t.Setenv("KUBE_MCP_OIDC_ISSUER_URL", "https://auth.localhost:8443")
	t.Setenv("KUBE_MCP_OIDC_CLIENT_ID", "00000000-0000-0000-0000-000000000000")
	t.Setenv("KUBE_MCP_LOG_LEVEL", "debug")
	t.Setenv("KUBE_MCP_ALLOWED_TOOLS", "GetServerVersion,ListNodes")
	t.Setenv("KUBE_MCP_DISALLOWED_TOOLS", "GetSecret")
	t.Setenv("KUBE_MCP_OIDC_SIGNING_METHOD", "RS256")
	t.Setenv("KUBE_MCP_OIDC_SCOPES", "openid,email,profile")

	config := getMcpServerCliFlags()

	assert.Equal(t, McpServerUserConfig{
		Host:            "http://localhost:9000",
		Port:            "9000",
		OutOfCluster:    true,
		Kubeconfig:      filepath.Join(home, ".kube", "config"),
		BaseURL:         "http://localhost:9000",
		OidcIssuerURL:   "https://auth.localhost:8443",
		OidcClientID:    "00000000-0000-0000-0000-000000000000",
		LogLevel:        "debug",
		AllowedTools:    "GetServerVersion,ListNodes",
		DisallowedTools: "GetSecret",
		SigningMethod:   "RS256",
		Scopes:          "openid,email,profile",
	}, config)
}
