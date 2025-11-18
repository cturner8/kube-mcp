package tools

import (
	"context"
	"encoding/json"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ListPodsTool = &mcp.Tool{
	Name:        "list_pods",
	Description: "List the pods in the Kubernetes cluster",
}

func ListPodsHandler(ctx context.Context, req *mcp.CallToolRequest, params any) (*mcp.CallToolResult, any, error) {
	log.Printf("Invoking '%s' tool", req.Params.Name)

	pods, err := kubernetesApiClient.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	podsJson, err := json.Marshal(pods)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(podsJson)},
		},
	}, nil, nil
}
