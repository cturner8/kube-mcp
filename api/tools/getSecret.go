package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GetSecretTool = &mcp.Tool{
	Name:        "get_secret",
	Description: "Get a secret in the Kubernetes cluster",
}

type GetSecretToolParams struct {
	Name      string `json:"name" jsonschema:"The name of the secret"`
	Namespace string `json:"namespace" jsonschema:"The namespace of the secret"`
}

func GetSecretHandler(ctx context.Context, req *mcp.CallToolRequest, params GetSecretToolParams) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	secret, err := kubernetesApiClient.CoreV1().Secrets(params.Namespace).Get(ctx, params.Name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	secretJson, err := json.Marshal(secret)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(secretJson)},
		},
	}, nil, nil
}
