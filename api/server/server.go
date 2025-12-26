// Copyright 2025 The Go MCP SDK Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"log/slog"
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
	// Track the active tools based on configuration
	activeTools := []string{}

	// Add MCP middlewares.
	server.AddReceivingMiddleware(createLoggingMiddleware())

	// Add the tools
	if tools.IsToolAllowed(tools.GetServerVersionTool.Name) {
		mcp.AddTool(server, tools.GetServerVersionTool, tools.GetServerVersionHandler)
		activeTools = append(activeTools, tools.GetServerVersionTool.Name)
	}

	// API resource tools

	// Nodes
	if tools.IsToolAllowed(tools.ListNodesTool.Name) {
		mcp.AddTool(server, tools.ListNodesTool, tools.ListNodesHandler)
		activeTools = append(activeTools, tools.ListNodesTool.Name)
	}
	if tools.IsToolAllowed(tools.GetNodeTool.Name) {
		mcp.AddTool(server, tools.GetNodeTool, tools.GetNodeHandler)
		activeTools = append(activeTools, tools.GetNodeTool.Name)
	}

	// Namespaces
	if tools.IsToolAllowed(tools.ListNamespacesTool.Name) {
		mcp.AddTool(server, tools.ListNamespacesTool, tools.ListNamespacesHandler)
		activeTools = append(activeTools, tools.ListNamespacesTool.Name)
	}
	if tools.IsToolAllowed(tools.GetNamespaceTool.Name) {
		mcp.AddTool(server, tools.GetNamespaceTool, tools.GetNamespaceHandler)
		activeTools = append(activeTools, tools.GetNamespaceTool.Name)
	}

	// Services
	if tools.IsToolAllowed(tools.ListServicesTool.Name) {
		mcp.AddTool(server, tools.ListServicesTool, tools.ListServicesHandler)
		activeTools = append(activeTools, tools.ListServicesTool.Name)
	}
	if tools.IsToolAllowed(tools.GetServiceTool.Name) {
		mcp.AddTool(server, tools.GetServiceTool, tools.GetServiceHandler)
		activeTools = append(activeTools, tools.GetServiceTool.Name)
	}

	// Deployments
	if tools.IsToolAllowed(tools.ListDeploymentsTool.Name) {
		mcp.AddTool(server, tools.ListDeploymentsTool, tools.ListDeploymentsHandler)
		activeTools = append(activeTools, tools.ListDeploymentsTool.Name)
	}
	if tools.IsToolAllowed(tools.GetDeploymentTool.Name) {
		mcp.AddTool(server, tools.GetDeploymentTool, tools.GetDeploymentHandler)
		activeTools = append(activeTools, tools.GetDeploymentTool.Name)
	}

	// Ingresses
	if tools.IsToolAllowed(tools.ListIngressesTool.Name) {
		mcp.AddTool(server, tools.ListIngressesTool, tools.ListIngressesHandler)
		activeTools = append(activeTools, tools.ListIngressesTool.Name)
	}
	if tools.IsToolAllowed(tools.GetIngressTool.Name) {
		mcp.AddTool(server, tools.GetIngressTool, tools.GetIngressHandler)
		activeTools = append(activeTools, tools.GetIngressTool.Name)
	}

	// Persistent Volumes
	if tools.IsToolAllowed(tools.ListPersistentVolumesTool.Name) {
		mcp.AddTool(server, tools.ListPersistentVolumesTool, tools.ListPersistentVolumesHandler)
		activeTools = append(activeTools, tools.ListPersistentVolumesTool.Name)
	}
	if tools.IsToolAllowed(tools.ListPersistentVolumeClaimsTool.Name) {
		mcp.AddTool(server, tools.ListPersistentVolumeClaimsTool, tools.ListPersistentVolumeClaimsHandler)
		activeTools = append(activeTools, tools.ListPersistentVolumeClaimsTool.Name)
	}
	if tools.IsToolAllowed(tools.GetPersistentVolumeTool.Name) {
		mcp.AddTool(server, tools.GetPersistentVolumeTool, tools.GetPersistentVolumeHandler)
		activeTools = append(activeTools, tools.GetPersistentVolumeTool.Name)
	}
	if tools.IsToolAllowed(tools.GetPersistentVolumeClaimTool.Name) {
		mcp.AddTool(server, tools.GetPersistentVolumeClaimTool, tools.GetPersistentVolumeClaimHandler)
		activeTools = append(activeTools, tools.GetPersistentVolumeClaimTool.Name)
	}

	// Pods
	if tools.IsToolAllowed(tools.ListPodsTool.Name) {
		mcp.AddTool(server, tools.ListPodsTool, tools.ListPodsHandler)
		activeTools = append(activeTools, tools.ListPodsTool.Name)
	}
	if tools.IsToolAllowed(tools.GetPodTool.Name) {
		mcp.AddTool(server, tools.GetPodTool, tools.GetPodHandler)
		activeTools = append(activeTools, tools.GetPodTool.Name)
	}

	// Events
	if tools.IsToolAllowed(tools.ListEventsTool.Name) {
		mcp.AddTool(server, tools.ListEventsTool, tools.ListEventsHandler)
		activeTools = append(activeTools, tools.ListEventsTool.Name)
	}

	// ConfigMaps
	if tools.IsToolAllowed(tools.ListConfigMapsTool.Name) {
		mcp.AddTool(server, tools.ListConfigMapsTool, tools.ListConfigMapsHandler)
		activeTools = append(activeTools, tools.ListConfigMapsTool.Name)
	}
	if tools.IsToolAllowed(tools.GetConfigMapTool.Name) {
		mcp.AddTool(server, tools.GetConfigMapTool, tools.GetConfigMapHandler)
		activeTools = append(activeTools, tools.GetConfigMapTool.Name)
	}

	// Secrets
	if tools.IsToolAllowed(tools.ListSecretsTool.Name) {
		mcp.AddTool(server, tools.ListSecretsTool, tools.ListSecretsHandler)
		activeTools = append(activeTools, tools.ListSecretsTool.Name)
	}
	if tools.IsToolAllowed(tools.GetSecretTool.Name) {
		mcp.AddTool(server, tools.GetSecretTool, tools.GetSecretHandler)
		activeTools = append(activeTools, tools.GetSecretTool.Name)
	}

	slog.Info("Active tools", "count", len(activeTools))

	// Create the streamable HTTP handler.
	handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
		return server
	}, nil)

	// Start serving
	serve(handler)
}
