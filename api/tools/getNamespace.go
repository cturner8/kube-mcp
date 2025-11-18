package tools

import (
	"context"
	"encoding/json"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GetNamespaceTool = &mcp.Tool{
	Name:        "get_namespace",
	Description: "Get the namespaces in the Kubernetes cluster",
}

type GetNamespaceToolParams struct {
	Name string `json:"name" jsonschema:"The name of the namespace"`
}

func GetNamespaceHandler(ctx context.Context, req *mcp.CallToolRequest, params GetNamespaceToolParams) (*mcp.CallToolResult, any, error) {
	log.Printf("Invoking '%s' tool", req.Params.Name)

	namespace, err := kubernetesApiClient.CoreV1().Namespaces().Get(ctx, params.Name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	namespaceJson, err := json.Marshal(namespace)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(namespaceJson)},
		},
	}, nil, nil
}
