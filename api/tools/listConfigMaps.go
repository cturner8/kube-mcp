package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ListConfigMapsTool = &mcp.Tool{
	Name:        "list_config_maps",
	Description: "List the config maps in the Kubernetes cluster",
}

type ListConfigMapsToolParams struct {
	Namespace *string `json:"namespace,omitempty" jsonschema:"The namespace of the config maps"`
}

func ListConfigMapsHandler(ctx context.Context, req *mcp.CallToolRequest, params ListConfigMapsToolParams) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	namespace := ""
	if params.Namespace != nil {
		namespace = *params.Namespace
	}

	configMaps, err := kubernetesApiClient.CoreV1().ConfigMaps(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	configMapsJson, err := json.Marshal(configMaps)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(configMapsJson)},
		},
	}, nil, nil
}
