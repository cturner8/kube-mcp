package tools_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/cturner8/kube-mcp/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestListConfigMaps_NoNamespace(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	setupMcpEnv(t)

	session := setupMcpServerClient(t, ctx, tools.ListConfigMapsTool, tools.ListConfigMapsHandler)
	defer session.Close()

	result, err := session.CallTool(ctx, &mcp.CallToolParams{
		Name: tools.ListConfigMapsTool.Name,
	})
	if err != nil {
		t.Fatalf("failed to call tool: %v", err)
	}

	textContent := getMcpResultTextContent(result)
	if textContent.Text == "" {
		t.Fatalf("expected non-empty config map list, got empty string")
	}
	if strings.HasPrefix(strings.ToLower(textContent.Text), "error") {
		t.Fatalf("expected successful response, got error: %s", textContent.Text)
	}
}

func TestListConfigMaps_WithNamespace(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	setupMcpEnv(t)

	session := setupMcpServerClient(t, ctx, tools.ListConfigMapsTool, tools.ListConfigMapsHandler)
	defer session.Close()

	result, err := session.CallTool(ctx, &mcp.CallToolParams{
		Name:      tools.ListConfigMapsTool.Name,
		Arguments: map[string]any{"namespace": "default"},
	})
	if err != nil {
		t.Fatalf("failed to call tool: %v", err)
	}

	textContent := getMcpResultTextContent(result)
	if textContent.Text == "" {
		t.Fatalf("expected non-empty config map list, got empty string")
	}
	if strings.Contains(strings.ToLower(textContent.Text), "error") {
		t.Fatalf("expected successful response, got error: %s", textContent.Text)
	}
}
