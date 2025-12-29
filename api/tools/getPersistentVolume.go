package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GetPersistentVolumeTool = &mcp.Tool{
	Name:        "get_persistent_volume",
	Description: "Get a persistent volume in the Kubernetes cluster",
}

type GetPersistentVolumeToolParams struct {
	Name string `json:"name" jsonschema:"The name of the persistent volume"`
}

func GetPersistentVolumeHandler(ctx context.Context, req *mcp.CallToolRequest, params GetPersistentVolumeToolParams) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	pv, err := getKubernetesApiClient().CoreV1().PersistentVolumes().Get(ctx, params.Name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	pvJson, err := json.Marshal(pv)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(pvJson)},
		},
	}, nil, nil
}
