package tools

import (
	"context"
	"encoding/json"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ListConfigMapsTool = &mcp.Tool{
	Name:        "list_config_maps",
	Description: "List the config maps in the Kubernetes cluster",
}

func ListConfigMapsHandler(ctx context.Context, req *mcp.CallToolRequest, params any) (*mcp.CallToolResult, any, error) {
	log.Printf("Invoking '%s' tool", req.Params.Name)

	configMaps, err := kubernetesApiClient.CoreV1().ConfigMaps("").List(ctx, metav1.ListOptions{})
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
