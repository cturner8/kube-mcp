package main

import (
	"log"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func startServer(url string) {
	// Create an MCP server.
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "time-server",
		Version: "1.0.0",
	}, nil)

	// Add MCP-level logging middleware.
	server.AddReceivingMiddleware(createLoggingMiddleware())

	// Add the cityTime tool.
	mcp.AddTool(server, &mcp.Tool{
		Name:        "cityTime",
		Description: "Get the current time in NYC, San Francisco, or Boston",
	}, getTime)

	// Create the streamable HTTP handler.
	handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
		return server
	}, nil)

	log.Printf("MCP server listening on %s", url)

	// Start the HTTP server.
	if err := http.ListenAndServe(url, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
