package tools_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/cturner8/kube-mcp/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestGetServerVersion(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// setup env
	t.Setenv("KUBE_MCP_BASE_URL", "http://localhost:9000")
	t.Setenv("KUBE_MCP_OIDC_ISSUER_URL", "https://auth.localhost:8443")
	t.Setenv("KUBE_MCP_OIDC_CLIENT_ID", "00000000-0000-0000-0000-000000000000")
	t.Setenv("KUBE_MCP_OUT_OF_CLUSTER", "true")

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "kube-mcp",
		Version: "test",
		Title:   "kube-mcp integration",
	}, nil)
	mcp.AddTool(server, tools.GetServerVersionTool, tools.GetServerVersionHandler)

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
	client := mcp.NewClient(&mcp.Implementation{Name: "kube-mcp-client", Version: "test"}, nil)
	cs, err := client.Connect(ctx, clientTransport, nil)
	if err != nil {
		t.Fatalf("client connect failed: %v", err)
	}
	defer cs.Close()

	result, err := cs.CallTool(ctx, &mcp.CallToolParams{
		Name: tools.GetServerVersionTool.Name,
	})
	if err != nil {
		t.Fatalf("failed to call tool: %v", err)
	}

	for _, content := range result.Content {
		if textContent, ok := content.(*mcp.TextContent); ok {
			t.Logf("GetServerVersion tool result: %s", textContent.Text)
		}
	}
}
