package tools

import (
	"context"
	"encoding/json"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ListDeploymentsTool = &mcp.Tool{
	Name:        "list_deployments",
	Description: "List the deployments in the Kubernetes cluster",
}

func ListDeploymentsHandler(ctx context.Context, req *mcp.CallToolRequest, params any) (*mcp.CallToolResult, any, error) {
	log.Printf("Invoking '%s' tool", req.Params.Name)

	deployments, err := kubernetesApiClient.AppsV1().Deployments("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	deploymentsJson, err := json.Marshal(deployments)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(deploymentsJson)},
		},
	}, nil, nil
}
