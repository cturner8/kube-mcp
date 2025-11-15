package config

import (
	"net/url"
	"os"
)

type McpServerEnvConfig struct {
	PublicBaseURL url.URL
	IssuerURL     url.URL
	ClientID      string
}

func getMcpServerEnvConfig() McpServerEnvConfig {
	var (
		PUBLIC_BASE_URL, PUBLIC_BASE_URL_SET = os.LookupEnv("PUBLIC_BASE_URL")
		OIDC_ISSUER_URL, OIDC_ISSUER_URL_SET = os.LookupEnv("OIDC_ISSUER_URL")
		OIDC_CLIENT_ID, OIDC_CLIENT_ID_SET   = os.LookupEnv("OIDC_CLIENT_ID")
	)

	if !PUBLIC_BASE_URL_SET {
		panic("[PUBLIC_BASE_URL] environment variable is required")
	}
	if !OIDC_ISSUER_URL_SET {
		panic("[OIDC_ISSUER_URL] environment variable is required")
	}
	if !OIDC_CLIENT_ID_SET {
		panic("[OIDC_CLIENT_ID] environment variable is required")
	}

	issuerURL, err := url.Parse(OIDC_ISSUER_URL)
	if err != nil {
		panic("[OIDC_ISSUER_URL] environment variable is not a valid URL")
	}

	publicBaseURL, err := url.Parse(PUBLIC_BASE_URL)
	if err != nil {
		panic("[PUBLIC_BASE_URL] environment variable is not a valid URL")
	}

	return McpServerEnvConfig{
		PublicBaseURL: *publicBaseURL,
		IssuerURL:     *issuerURL,
		ClientID:      OIDC_CLIENT_ID,
	}
}
