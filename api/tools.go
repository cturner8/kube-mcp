// Copyright 2025 The Go MCP SDK Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var getServerVersionTool = &mcp.Tool{
	Name:        "get_server_version",
	Description: "Get the Kubernetes API server version details",
}

// getServerVersion implements the tool that returns the Kubernetes API server version details
func getServerVersion(ctx context.Context, req *mcp.CallToolRequest, params any) (*mcp.CallToolResult, any, error) {
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

var getNodesTool = &mcp.Tool{
	Name:        "get_nodes",
	Description: "Get the nodes in the Kubernetes cluster",
}

// getNodes implements the tool that returns the nodes registered in the Kubernetes cluster
func getNodes(ctx context.Context, req *mcp.CallToolRequest, params any) (*mcp.CallToolResult, any, error) {
	nodes, err := kubernetesApiClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	nodesJson, err := json.Marshal(nodes)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(nodesJson)},
		},
	}, nil, nil
}
