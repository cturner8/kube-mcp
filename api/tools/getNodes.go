package tools

import (
	"context"
	"encoding/json"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GetNodesTool = &mcp.Tool{
	Name:        "get_nodes",
	Description: "Get the nodes in the Kubernetes cluster",
}

// GetNodes implements the tool that returns the nodes registered in the Kubernetes cluster
func GetNodes(ctx context.Context, req *mcp.CallToolRequest, params any) (*mcp.CallToolResult, any, error) {
	log.Printf("Invoking '%s' tool", req.Params.Name)

	nodes, err := kubernetesApiClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	nodesJson, err := json.Marshal(nodes)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(nodesJson)},
		},
	}, nil, nil
}
