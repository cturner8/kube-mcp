package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ListPodsTool = &mcp.Tool{
	Name:        "list_pods",
	Description: "List the pods in the Kubernetes cluster",
}

type ListPodsToolParams struct {
	Namespace *string `json:"namespace,omitempty" jsonschema:"The namespace of the pods"`
}

func ListPodsHandler(ctx context.Context, req *mcp.CallToolRequest, params ListPodsToolParams) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	namespace := ""
	if params.Namespace != nil {
		namespace = *params.Namespace
	}

	pods, err := kubernetesApiClient.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		slog.Error("failed to list pods from Kubernetes API", "tool", req.Params.Name, "namespace", namespace, "error", err)
		return nil, nil, err
	}

	podsJson, err := json.Marshal(pods)
	if err != nil {
		slog.Error("failed to marshal pods list", "namespace", namespace, "error", err)
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(podsJson)},
		},
	}, nil, nil
}
