package tools

import (
	"context"
	"encoding/json"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ListPersistentVolumeClaimsTool = &mcp.Tool{
	Name:        "list_persistent_volume_claims",
	Description: "List the persistent volume claims in the Kubernetes cluster",
}

func ListPersistentVolumeClaimsHandler(ctx context.Context, req *mcp.CallToolRequest, params any) (*mcp.CallToolResult, any, error) {
	log.Printf("Invoking '%s' tool", req.Params.Name)

	pvcs, err := kubernetesApiClient.CoreV1().PersistentVolumeClaims("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	pvcsJson, err := json.Marshal(pvcs)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(pvcsJson)},
		},
	}, nil, nil
}
