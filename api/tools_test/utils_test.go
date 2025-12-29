package tools_test

import (
	"context"
	"log"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func getMcpResultTextContent(result *mcp.CallToolResult) mcp.TextContent {
	if textContent, ok := result.Content[0].(*mcp.TextContent); ok {
		return *textContent
	}
	return mcp.TextContent{}
}

func setupMcpEnv(t *testing.T) {
	t.Setenv("KUBE_MCP_BASE_URL", "http://localhost:9000")
	t.Setenv("KUBE_MCP_OIDC_ISSUER_URL", "https://auth.localhost:8443")
	t.Setenv("KUBE_MCP_OIDC_CLIENT_ID", "00000000-0000-0000-0000-000000000000")
	t.Setenv("KUBE_MCP_OUT_OF_CLUSTER", "true")
}

func setupMcpServerClient[In any](t *testing.T, ctx context.Context, tool *mcp.Tool, handler mcp.ToolHandlerFor[In, any]) *mcp.ClientSession {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "kube-mcp",
		Version: "test",
		Title:   "kube-mcp integration",
	}, nil)
	mcp.AddTool(server, tool, handler)

	serverTransport, clientTransport := mcp.NewInMemoryTransports()

	// Start the server side of the transport.
	go func() {
		conn, err := server.Connect(ctx, serverTransport, nil)
		if err != nil {
			log.Fatalf("server connect failed: %v", err)
		}
		t.Cleanup(func() {
			conn.Close()
		})
	}()

	// Connect the MCP client side of the transport.
	client := mcp.NewClient(&mcp.Implementation{Name: "kube-mcp-client", Version: "test"}, nil)
	cs, err := client.Connect(ctx, clientTransport, nil)
	if err != nil {
		t.Fatalf("client connect failed: %v", err)
	}

	return cs
}
