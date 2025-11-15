package config

import (
	"net/url"
	"strings"
)

type McpServerConfig struct {
	PublicBaseURL  *url.URL
	Host           *string
	Port           *int
	OutOfCluster   *bool
	Kubeconfig     *string
	IssuerURL      *url.URL
	ClientID       string
	AllowedOrigins []string
}

func GetMcpServerConfig() McpServerConfig {
	// Parse environment variable configuration
	envConfig := getMcpServerEnvConfig()
	// Parse CLI flag configuration
	cliConfig := getMcpServerCliFlags()

	// Filter out empty strings from allowed origins
	allowedOrigins := []string{}
	for _, origin := range cliConfig.AllowedOrigins {
		if trimmed := strings.TrimSpace(origin); trimmed != "" {
			allowedOrigins = append(allowedOrigins, trimmed)
		}
	}

	// Build the complete server configuration
	return McpServerConfig{
		PublicBaseURL:  &envConfig.PublicBaseURL,
		IssuerURL:      &envConfig.IssuerURL,
		ClientID:       envConfig.ClientID,
		Host:           cliConfig.Host,
		Port:           cliConfig.Port,
		OutOfCluster:   cliConfig.OutOfCluster,
		Kubeconfig:     cliConfig.Kubeconfig,
		AllowedOrigins: allowedOrigins,
	}
}

var ServerConfig McpServerConfig = GetMcpServerConfig()
