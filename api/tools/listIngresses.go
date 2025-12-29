package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ListIngressesTool = &mcp.Tool{
	Name:        "list_ingresses",
	Description: "List the ingresses in the Kubernetes cluster",
}

type ListIngressesToolParams struct {
	Namespace *string `json:"namespace,omitempty" jsonschema:"The namespace of the ingress(es)"`
}

func ListIngressesHandler(ctx context.Context, req *mcp.CallToolRequest, params ListIngressesToolParams) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	namespace := ""
	if params.Namespace != nil {
		namespace = *params.Namespace
	}

	ingresses, err := getKubernetesApiClient().NetworkingV1().Ingresses(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	ingressesJson, err := json.Marshal(ingresses)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(ingressesJson)},
		},
	}, nil, nil
}
