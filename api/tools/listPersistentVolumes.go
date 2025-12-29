package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ListPersistentVolumesTool = &mcp.Tool{
	Name:        "list_persistent_volumes",
	Description: "List the persistent volumes in the Kubernetes cluster",
}

func ListPersistentVolumesHandler(ctx context.Context, req *mcp.CallToolRequest, params any) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	pvs, err := getKubernetesApiClient().CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	pvsJson, err := json.Marshal(pvs)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(pvsJson)},
		},
	}, nil, nil
}
