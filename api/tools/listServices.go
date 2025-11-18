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

type ListServicesToolParams struct {
	Namespace *string `json:"namespace,omitempty" jsonschema:"The namespace of the services"`
}

func ListServicesHandler(ctx context.Context, req *mcp.CallToolRequest, params ListServicesToolParams) (*mcp.CallToolResult, any, error) {
	log.Printf("Invoking '%s' tool", req.Params.Name)

	namespace := ""
	if params.Namespace != nil {
		namespace = *params.Namespace
	}

	services, err := kubernetesApiClient.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
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
