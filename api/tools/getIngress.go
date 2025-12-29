package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GetIngressTool = &mcp.Tool{
	Name:        "get_ingress",
	Description: "Get an ingress in the Kubernetes cluster",
}

type GetIngressToolParams struct {
	Name      string `json:"name" jsonschema:"The name of the ingress"`
	Namespace string `json:"namespace" jsonschema:"The namespace of the ingress"`
}

func GetIngressHandler(ctx context.Context, req *mcp.CallToolRequest, params GetIngressToolParams) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	ingress, err := getKubernetesApiClient().NetworkingV1().Ingresses(params.Namespace).Get(ctx, params.Name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}

	ingressJson, err := json.Marshal(ingress)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(ingressJson)},
		},
	}, nil, nil
}
