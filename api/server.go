package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func startServer(httpUrl string, baseUrl string) {
	// Create an MCP server.
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "kube-mcp",
		Version: "0.0.0",
	}, nil)

	// Add MCP-level logging middleware.
	server.AddReceivingMiddleware(createLoggingMiddleware())

	// Add the tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_server_version",
		Description: "Get the Kubernetes API server version details",
	}, getServerVersion)

	// Create the streamable HTTP handler.
	handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
		return server
	}, nil)

	prmPath := "/.well-known/oauth-protected-resource"

	// Add the authentication middleware.
	bearerAuth := createBearerAuth(baseUrl, prmPath)
	authenticatedHandler := bearerAuth(handler)

	// Setup HTTP routes

	// Health check endpoint.
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "healthy",
			"time":   time.Now().Format(time.RFC3339),
		})
	})
	// Create a wrapper handler that routes to either the metadata endpoint or the MCP handler
	http.HandleFunc(prmPath, getProtectedResourceMetadataHandler(baseUrl))

	// Register the authenticated MCP handler
	http.HandleFunc("/mcp", authenticatedHandler.ServeHTTP)

	log.Printf("MCP server listening on %s", httpUrl)
	log.Printf("Protected Resource Metadata available at %s%s", baseUrl, prmPath)

	// Start the HTTP server.
	if err := http.ListenAndServe(httpUrl, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
