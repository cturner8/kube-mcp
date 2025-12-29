package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GetServiceTool = &mcp.Tool{
	Name:        "get_service",
	Description: "Get a service in the Kubernetes cluster",
}

type GetServiceToolParams struct {
	Name      string `json:"name" jsonschema:"The name of the service"`
	Namespace string `json:"namespace" jsonschema:"The namespace of the service"`
}

func GetServiceHandler(ctx context.Context, req *mcp.CallToolRequest, params GetServiceToolParams) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	service, err := getKubernetesApiClient().CoreV1().Services(params.Namespace).Get(ctx, params.Name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	serviceJson, err := json.Marshal(service)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(serviceJson)},
		},
	}, nil, nil
}
