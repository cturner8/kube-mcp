package tools

import (
	"context"
	"encoding/json"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ListIngressesTool = &mcp.Tool{
	Name:        "list_ingresses",
	Description: "List the ingresses in the Kubernetes cluster",
}

func ListIngressesHandler(ctx context.Context, req *mcp.CallToolRequest, params any) (*mcp.CallToolResult, any, error) {
	log.Printf("Invoking '%s' tool", req.Params.Name)

	ingresses, err := kubernetesApiClient.NetworkingV1().Ingresses("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	ingressesJson, err := json.Marshal(ingresses)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(ingressesJson)},
		},
	}, nil, nil
}
