package tools

import (
	"context"
	"encoding/json"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GetDeploymentTool = &mcp.Tool{
	Name:        "get_deployment",
	Description: "Get a deployment in the Kubernetes cluster",
}

type GetDeploymentToolParams struct {
	Name      string `json:"name" jsonschema:"The name of the deployment"`
	Namespace string `json:"namespace" jsonschema:"The namespace of the deployment"`
}

func GetDeploymentHandler(ctx context.Context, req *mcp.CallToolRequest, params GetDeploymentToolParams) (*mcp.CallToolResult, any, error) {
	log.Printf("Invoking '%s' tool", req.Params.Name)

	deployment, err := kubernetesApiClient.AppsV1().Deployments(params.Namespace).Get(ctx, params.Name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	deploymentJson, err := json.Marshal(deployment)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(deploymentJson)},
		},
	}, nil, nil
}
