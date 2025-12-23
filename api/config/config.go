package config

import (
	"net/url"
	"strings"
)

type McpServerConfig struct {
	BaseURL        url.URL
	Host           *string
	Port           *string
	OutOfCluster   *bool
	Kubeconfig     *string
	OidcIssuerURL  url.URL
	OidcClientID   string
	AllowedOrigins []string
}

type McpServerUserConfig struct {
	BaseURL        string
	Host           string
	Port           string
	OutOfCluster   bool
	Kubeconfig     string
	OidcIssuerURL  string
	OidcClientID   string
	AllowedOrigins string
}

func parseServerUserConfig(config McpServerUserConfig) {
	baseUrl := config.BaseURL
	oidcIssuerUrl := config.OidcIssuerURL
	oidcClientId := config.OidcClientID

	if baseUrl == "" {
		panic("Base URL is required")
	}
	if oidcIssuerUrl == "" {
		panic("OIDC Issuer URL is required")
	}
	if oidcClientId == "" {
		panic("OIDC Client ID is required")
	}
}

func GetMcpServerConfig() McpServerConfig {
	// Parse CLI flag configuration
	config := getMcpServerCliFlags()
	parseServerUserConfig(config)

	// Filter out empty strings from allowed origins
	allowedOrigins := []string{}
	for origin := range strings.SplitSeq(config.AllowedOrigins, ",") {
		if trimmed := strings.TrimSpace(origin); trimmed != "" {
			allowedOrigins = append(allowedOrigins, trimmed)
		}
	}

	baseUrl, err := url.Parse(config.BaseURL)
	if err != nil {
		panic("Unable to parse Base URL")
	}

	oidcIssuerUrl, err := url.Parse(config.OidcIssuerURL)
	if err != nil {
		panic("Unable to parse OIDC Issuer URL")
	}

	port := config.Port
	if port == "" {
		port = "9000"
	}

	// Build the complete server configuration
	return McpServerConfig{
		BaseURL:        *baseUrl,
		OidcIssuerURL:  *oidcIssuerUrl,
		OidcClientID:   config.OidcClientID,
		Host:           &config.Host,
		Port:           &port,
		OutOfCluster:   &config.OutOfCluster,
		Kubeconfig:     &config.Kubeconfig,
		AllowedOrigins: allowedOrigins,
	}
}

var ServerConfig McpServerConfig = GetMcpServerConfig()
