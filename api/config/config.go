package config

import (
	"net/url"
	"strings"
)

type McpServerConfig struct {
	BaseURL         url.URL
	Host            *string
	Port            *string
	OutOfCluster    *bool
	Kubeconfig      *string
	OidcIssuerURL   url.URL
	OidcClientID    string
	AllowedOrigins  []string
	AllowedTools    []string
	DisallowedTools []string
	Scopes          []string
	SigningMethod   string
}

type McpServerUserConfig struct {
	BaseURL         string
	Host            string
	Port            string
	OutOfCluster    bool
	Kubeconfig      string
	OidcIssuerURL   string
	OidcClientID    string
	AllowedOrigins  string
	AllowedTools    string
	DisallowedTools string
	Scopes          string
	SigningMethod   string
}

func parseServerUserConfig(config McpServerUserConfig) {
	if config.BaseURL == "" {
		panic("Base URL is required")
	}
	if config.OidcIssuerURL == "" {
		panic("OIDC Issuer URL is required")
	}
	if config.OidcClientID == "" {
		panic("OIDC Client ID is required")
	}
	if config.AllowedTools != "" && config.DisallowedTools != "" {
		panic("Cannot specify both allowed-tools and disallowed-tools")
	}
	// Validate signing method if provided
	if config.SigningMethod != "HS256" && config.SigningMethod != "RS256" {
		panic("Invalid signing method. Supported values are HS256 or RS256")
	}
}

func splitStringArg(input string) []string {
	// Filter out empty strings from value
	output := []string{}
	for value := range strings.SplitSeq(input, ",") {
		if trimmed := strings.TrimSpace(value); trimmed != "" {
			output = append(output, trimmed)
		}
	}
	return output
}

func GetMcpServerConfig() McpServerConfig {
	// Parse CLI flag configuration
	config := getMcpServerCliFlags()
	parseServerUserConfig(config)

	allowedOrigins := splitStringArg(config.AllowedOrigins)
	allowedTools := splitStringArg(config.AllowedTools)
	disallowedTools := splitStringArg(config.DisallowedTools)
	scopes := splitStringArg(config.Scopes)

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
		BaseURL:         *baseUrl,
		OidcIssuerURL:   *oidcIssuerUrl,
		OidcClientID:    config.OidcClientID,
		Host:            &config.Host,
		Port:            &port,
		OutOfCluster:    &config.OutOfCluster,
		Kubeconfig:      &config.Kubeconfig,
		AllowedOrigins:  allowedOrigins,
		AllowedTools:    allowedTools,
		DisallowedTools: disallowedTools,
		SigningMethod:   config.SigningMethod,
		Scopes:          scopes,
	}
}

var ServerConfig McpServerConfig = GetMcpServerConfig()
