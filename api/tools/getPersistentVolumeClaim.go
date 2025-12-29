package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GetPersistentVolumeClaimTool = &mcp.Tool{
	Name:        "get_persistent_volume_claim",
	Description: "Get a persistent volume claim in the Kubernetes cluster",
}

type GetPersistentVolumeClaimToolParams struct {
	Name      string `json:"name" jsonschema:"The name of the pvc"`
	Namespace string `json:"namespace" jsonschema:"The namespace of the pvc"`
}

func GetPersistentVolumeClaimHandler(ctx context.Context, req *mcp.CallToolRequest, params GetPersistentVolumeClaimToolParams) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	pvc, err := getKubernetesApiClient().CoreV1().PersistentVolumeClaims(params.Namespace).Get(ctx, params.Name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	pvcJson, err := json.Marshal(pvc)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(pvcJson)},
		},
	}, nil, nil
}
