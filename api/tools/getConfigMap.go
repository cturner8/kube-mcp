package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GetConfigMapTool = &mcp.Tool{
	Name:        "get_config_map",
	Description: "Get a config map in the Kubernetes cluster",
}

type GetConfigMapToolParams struct {
	Name      string `json:"name" jsonschema:"The name of the config map"`
	Namespace string `json:"namespace" jsonschema:"The namespace of the config map"`
}

func GetConfigMapHandler(ctx context.Context, req *mcp.CallToolRequest, params GetConfigMapToolParams) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	cm, err := getKubernetesApiClient().CoreV1().ConfigMaps(params.Namespace).Get(ctx, params.Name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	cmJson, err := json.Marshal(cm)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(cmJson)},
		},
	}, nil, nil
}
