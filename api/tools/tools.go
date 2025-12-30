package tools

import (
	"slices"
	"sync"

	"github.com/cturner8/kube-mcp/config"
	"github.com/cturner8/kube-mcp/kubernetes"

	k8s "k8s.io/client-go/kubernetes"
)

// ClientProvider is an interface for providing access to the Kubernetes API client.
// This allows for dependency injection and makes the tools testable.
type ClientProvider interface {
	GetClient() *k8s.Clientset
}

// defaultClientProvider is the production implementation that uses singleton pattern.
type defaultClientProvider struct {
	once   sync.Once
	client *k8s.Clientset
}

func (p *defaultClientProvider) GetClient() *k8s.Clientset {
	p.once.Do(func() {
		config.Load()
		cfg := config.GetMcpServerConfig()
		p.client = kubernetes.CreateKubernetesApiClient(*cfg.OutOfCluster, *cfg.Kubeconfig)
	})
	return p.client
}

var (
	clientProvider ClientProvider = &defaultClientProvider{}
)

// SetClientProvider sets the client provider for dependency injection.
// This is primarily used for testing to inject mock clients.
func SetClientProvider(provider ClientProvider) {
	clientProvider = provider
}

// getKubernetesApiClient returns the Kubernetes API client from the configured provider.
func getKubernetesApiClient() *k8s.Clientset {
	return clientProvider.GetClient()
}

func IsToolAllowed(toolName string) bool {
	allowedTools := config.ServerConfig.AllowedTools
	disallowedTools := config.ServerConfig.DisallowedTools

	// If no restrictions, all tools are allowed
	if len(allowedTools) == 0 && len(disallowedTools) == 0 {
		return true
	}

	// If allowed tools are specified, check tool has been allowed
	if len(allowedTools) > 0 {
		return slices.Contains(allowedTools, toolName)
	}

	// If disallowed tools are specified, check tool has not been disallowed
	if len(disallowedTools) > 0 {
		return !slices.Contains(disallowedTools, toolName)
	}

	return false
}
