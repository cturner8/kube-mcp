// Copyright 2025 The Go MCP SDK Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	tools "github.com/cturner8/kube-mcp/tools"
)

func StartServer() {
	// Create an MCP server.
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "kube-mcp",
		Version: "0.0.0",
		Title:   "Kubernetes API MCP",
	}, nil)

	// Add MCP middlewares.
	server.AddReceivingMiddleware(createLoggingMiddleware())

	// Add the tools
	mcp.AddTool(server, tools.GetServerVersionTool, tools.GetServerVersionHandler)
	// API resource tools
	mcp.AddTool(server, tools.ListNodesTool, tools.ListNodesHandler)
	mcp.AddTool(server, tools.ListNamespacesTool, tools.ListNamespacesHandler)
	mcp.AddTool(server, tools.ListServicesTool, tools.ListServicesHandler)
	mcp.AddTool(server, tools.ListDeploymentsTool, tools.ListDeploymentsHandler)
	mcp.AddTool(server, tools.ListIngressesTool, tools.ListIngressesHandler)
	mcp.AddTool(server, tools.ListPersistentVolumesTool, tools.ListPersistentVolumesHandler)
	mcp.AddTool(server, tools.ListPersistentVolumeClaimsTool, tools.ListPersistentVolumeClaimsHandler)

	// Create the streamable HTTP handler.
	handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
		return server
	}, nil)

	// Start serving
	serve(handler)
}
