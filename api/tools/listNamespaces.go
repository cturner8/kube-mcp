package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ListNamespacesTool = &mcp.Tool{
	Name:        "list_namespaces",
	Description: "List the namespaces in the Kubernetes cluster",
}

func ListNamespacesHandler(ctx context.Context, req *mcp.CallToolRequest, params any) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	namespaces, err := kubernetesApiClient.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	namespacesJson, err := json.Marshal(namespaces)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(namespacesJson)},
		},
	}, nil, nil
}
