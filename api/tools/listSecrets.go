package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ListSecretsTool = &mcp.Tool{
	Name:        "list_secrets",
	Description: "List the secrets in the Kubernetes cluster",
}

type ListSecretsToolParams struct {
	Namespace *string `json:"namespace,omitempty" jsonschema:"The namespace of the secrets"`
}

func ListSecretsHandler(ctx context.Context, req *mcp.CallToolRequest, params ListSecretsToolParams) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	namespace := ""
	if params.Namespace != nil {
		namespace = *params.Namespace
	}

	secrets, err := getKubernetesApiClient().CoreV1().Secrets(namespace).List(ctx, metav1.ListOptions{})
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
