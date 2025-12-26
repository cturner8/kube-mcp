package tools

import (
	"slices"

	"github.com/cturner8/kube-mcp/config"
	"github.com/cturner8/kube-mcp/kubernetes"

	k8s "k8s.io/client-go/kubernetes"
)

var kubernetesApiClient *k8s.Clientset = kubernetes.CreateKubernetesApiClient(*config.ServerConfig.OutOfCluster, *config.ServerConfig.Kubeconfig)

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
