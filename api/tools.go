// Copyright 2025 The Go MCP SDK Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// getVersion implements the tool that returns the Kubernetes API server version details
func getVersion(ctx context.Context, req *mcp.CallToolRequest, params any) (*mcp.CallToolResult, any, error) {
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
