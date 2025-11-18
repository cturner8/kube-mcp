// Copyright 2025 The Go MCP SDK Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cturner8/kube-mcp/config"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func serve(handler *mcp.StreamableHTTPHandler) {
	httpUrl := fmt.Sprintf("%s:%d", *config.ServerConfig.Host, *config.ServerConfig.Port)
	baseUrl := config.ServerConfig.PublicBaseURL.String()

	prmPath := "/.well-known/oauth-protected-resource"

	// Add middlewares
	// Add CORS middleware
	corsMiddleware := createCORSMiddleware(config.ServerConfig.AllowedOrigins)

	// Add the authentication middleware.
	bearerAuth := createBearerAuth(baseUrl, prmPath)
	authenticatedHandler := bearerAuth(handler)
	authenticatedHandler = corsMiddleware(authenticatedHandler)

	// Setup HTTP routes

	// Health check endpoint.
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

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
