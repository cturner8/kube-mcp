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

	// Nodes
	mcp.AddTool(server, tools.ListNodesTool, tools.ListNodesHandler)
	mcp.AddTool(server, tools.GetNodeTool, tools.GetNodeHandler)

	// Namespaces
	mcp.AddTool(server, tools.ListNamespacesTool, tools.ListNamespacesHandler)
	mcp.AddTool(server, tools.GetNamespaceTool, tools.GetNamespaceHandler)

	// Services
	mcp.AddTool(server, tools.ListServicesTool, tools.ListServicesHandler)
	mcp.AddTool(server, tools.GetServiceTool, tools.GetServiceHandler)

	// Deployments
	mcp.AddTool(server, tools.ListDeploymentsTool, tools.ListDeploymentsHandler)
	mcp.AddTool(server, tools.GetDeploymentTool, tools.GetDeploymentHandler)

	// Ingresses
	mcp.AddTool(server, tools.ListIngressesTool, tools.ListIngressesHandler)
	mcp.AddTool(server, tools.GetIngressTool, tools.GetIngressHandler)

	// Persistent Volumes
	mcp.AddTool(server, tools.ListPersistentVolumesTool, tools.ListPersistentVolumesHandler)
	mcp.AddTool(server, tools.ListPersistentVolumeClaimsTool, tools.ListPersistentVolumeClaimsHandler)
	mcp.AddTool(server, tools.GetPersistentVolumeTool, tools.GetPersistentVolumeHandler)
	mcp.AddTool(server, tools.GetPersistentVolumeClaimTool, tools.GetPersistentVolumeClaimHandler)

	mcp.AddTool(server, tools.ListPodsTool, tools.ListPodsHandler)
	mcp.AddTool(server, tools.GetPodTool, tools.GetPodHandler)

	// ConfigMaps
	mcp.AddTool(server, tools.ListConfigMapsTool, tools.ListConfigMapsHandler)
	mcp.AddTool(server, tools.GetConfigMapTool, tools.GetConfigMapHandler)

	// Secrets
	mcp.AddTool(server, tools.ListSecretsTool, tools.ListSecretsHandler)
	mcp.AddTool(server, tools.GetSecretTool, tools.GetSecretHandler)

	// Create the streamable HTTP handler.
	handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
		return server
	}, nil)

	// Start serving
	serve(handler)
}
