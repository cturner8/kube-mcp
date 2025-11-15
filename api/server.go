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
		Title:   "Kubernetes API MCP",
	}, nil)

	// Add MCP-level middlewares.
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

	// Add HTTP middlewares
	// Add CORS middleware
	corsMiddleware := createCORSMiddleware(config.AllowedOrigins)

	// Add the authentication middleware.
	bearerAuth := createBearerAuth(baseUrl, prmPath)
	authenticatedHandler := bearerAuth(handler)
	authenticatedHandler = corsMiddleware(authenticatedHandler)

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
	// Apply CORS middleware to the metadata endpoint
	metadataHandler := getProtectedResourceMetadataHandler(baseUrl)
	http.HandleFunc(prmPath, corsMiddleware(metadataHandler).ServeHTTP)

	// Register the authenticated MCP handler
	http.HandleFunc("/mcp", authenticatedHandler.ServeHTTP)

	log.Printf("MCP server listening on %s", httpUrl)
	log.Printf("Protected Resource Metadata available at %s%s", baseUrl, prmPath)

	// Start the HTTP server.
	if err := http.ListenAndServe(httpUrl, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
