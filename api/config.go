// Copyright 2025 The Go MCP SDK Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"net/url"
	"os"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

type McpServerConfig struct {
	Host         *string
	Port         *int
	OutOfCluster *bool
	Kubeconfig   *string
	IssuerURL    *url.URL
	ClientID     string
	AllowOrigins []string
}

type McpServerCliConfig struct {
	Host         *string
	Port         *int
	OutOfCluster *bool
	Kubeconfig   *string
}

type McpServerEnvConfig struct {
	IssuerURL url.URL
	ClientID  string
}

func getMcpServerCliFlags() McpServerCliConfig {
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

	return McpServerCliConfig{
		Host:         host,
		Port:         port,
		OutOfCluster: outOfCluster,
		Kubeconfig:   kubeconfig,
	}
}

func getMcpServerEnvConfig() McpServerEnvConfig {
	var (
		OIDC_ISSUER_URL, OIDC_ISSUER_URL_SET = os.LookupEnv("OIDC_ISSUER_URL")
		OIDC_CLIENT_ID, OIDC_CLIENT_ID_SET   = os.LookupEnv("OIDC_CLIENT_ID")
	)

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

	return McpServerEnvConfig{
		IssuerURL: *issuerURL,
		ClientID:  OIDC_CLIENT_ID,
	}
}

func getMcpServerConfig() McpServerConfig {
	// Parse environment variable configuration
	envConfig := getMcpServerEnvConfig()
	// Parse CLI flag configuration
	cliConfig := getMcpServerCliFlags()

	// Build the complete server configuration
	return McpServerConfig{
		IssuerURL:    &envConfig.IssuerURL,
		ClientID:     envConfig.ClientID,
		Host:         cliConfig.Host,
		Port:         cliConfig.Port,
		OutOfCluster: cliConfig.OutOfCluster,
		Kubeconfig:   cliConfig.Kubeconfig,
	}
}
