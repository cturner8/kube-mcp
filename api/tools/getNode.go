package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GetNodeTool = &mcp.Tool{
	Name:        "get_node",
	Description: "Get a node in the Kubernetes cluster",
}

type GetNodeToolParams struct {
	Name string `json:"name" jsonschema:"The name of the node"`
}

func GetNodeHandler(ctx context.Context, req *mcp.CallToolRequest, params GetNodeToolParams) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	node, err := getKubernetesApiClient().CoreV1().Nodes().Get(ctx, params.Name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	nodeJson, err := json.Marshal(node)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(nodeJson)},
		},
	}, nil, nil
}
