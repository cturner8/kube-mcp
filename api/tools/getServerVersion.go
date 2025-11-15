package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var GetServerVersionTool = &mcp.Tool{
	Name:        "get_server_version",
	Description: "Get the Kubernetes API server version details",
}

func GetServerVersionHandler(ctx context.Context, req *mcp.CallToolRequest, params any) (*mcp.CallToolResult, any, error) {
	version, err := kubernetesApiClient.Discovery().ServerVersion()
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: version.String()},
		},
	}, nil, nil
}
