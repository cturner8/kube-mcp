package tools

import (
	"context"
	"encoding/json"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ListSecretsTool = &mcp.Tool{
	Name:        "list_secrets",
	Description: "List the secrets in the Kubernetes cluster",
}

func ListSecretsHandler(ctx context.Context, req *mcp.CallToolRequest, params any) (*mcp.CallToolResult, any, error) {
	log.Printf("Invoking '%s' tool", req.Params.Name)

	secrets, err := kubernetesApiClient.CoreV1().Secrets("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	secretsJson, err := json.Marshal(secrets)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(secretsJson)},
		},
	}, nil, nil
}
