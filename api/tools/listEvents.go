package tools

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ListEventsTool = &mcp.Tool{
	Name:        "list_events",
	Description: "List the events in the Kubernetes cluster",
}

type ListEventsToolParams struct {
	Namespace *string `json:"namespace,omitempty" jsonschema:"The namespace of the events"`
}

func ListEventsHandler(ctx context.Context, req *mcp.CallToolRequest, params ListEventsToolParams) (*mcp.CallToolResult, any, error) {
	slog.Debug("Tool invoked", "tool", req.Params.Name)

	namespace := ""
	if params.Namespace != nil {
		namespace = *params.Namespace
	}

	events, err := kubernetesApiClient.CoreV1().Events(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		slog.Error("Failed to list events from Kubernetes API", "tool", req.Params.Name, "namespace", namespace, "error", err)
		return nil, nil, err
	}

	eventsJson, err := json.Marshal(events)
	if err != nil {
		slog.Error("Failed to marshal events list", "namespace", namespace, "error", err)
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(eventsJson)},
		},
	}, nil, nil
}
