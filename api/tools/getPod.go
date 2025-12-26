package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GetPodTool = &mcp.Tool{
	Name:        "get_pod",
	Description: "Get a pod in the Kubernetes cluster",
}

type GetPodToolParams struct {
	Name      string `json:"name" jsonschema:"The name of the pod"`
	Namespace string `json:"namespace" jsonschema:"The namespace of the pod"`
}

func GetPodHandler(ctx context.Context, req *mcp.CallToolRequest, params GetPodToolParams) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	pod, err := kubernetesApiClient.CoreV1().Pods(params.Namespace).Get(ctx, params.Name, metav1.GetOptions{})
	if err != nil {
		slog.Error("Failed to get pod from Kubernetes API", "tool", req.Params.Name, "pod", params.Name, "namespace", params.Namespace, "error", err)
		return nil, nil, err
	}

	podJson, err := json.Marshal(pod)
	if err != nil {
		slog.Error("Failed to marshal pod object", "pod", params.Name, "namespace", params.Namespace, "error", err)
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(podJson)},
		},
	}, nil, nil
}
