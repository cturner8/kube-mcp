package tools

import (
	"context"
	"encoding/json"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ListServicesTool = &mcp.Tool{
	Name:        "list_services",
	Description: "List the services in the Kubernetes cluster",
}

func ListServicesHandler(ctx context.Context, req *mcp.CallToolRequest, params any) (*mcp.CallToolResult, any, error) {
	log.Printf("Invoking '%s' tool", req.Params.Name)

	services, err := kubernetesApiClient.CoreV1().Services("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	servicesJson, err := json.Marshal(services)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(servicesJson)},
		},
	}, nil, nil
}
