// Copyright 2025 The Go MCP SDK Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildServerDefault(t *testing.T) {
	t.Setenv("KUBE_MCP_BASE_URL", "http://localhost:9000")
	t.Setenv("KUBE_MCP_OIDC_ISSUER_URL", "https://auth.localhost:8443")
	t.Setenv("KUBE_MCP_OIDC_CLIENT_ID", "00000000-0000-0000-0000-000000000000")

	_, activeTools := buildServer()

	assert.Greater(t, len(activeTools), 0)
	assert.Contains(t, activeTools, "get_server_version")
}

func TestBuildServerAllowedTools(t *testing.T) {
	t.Setenv("KUBE_MCP_BASE_URL", "http://localhost:9000")
	t.Setenv("KUBE_MCP_OIDC_ISSUER_URL", "https://auth.localhost:8443")
	t.Setenv("KUBE_MCP_OIDC_CLIENT_ID", "00000000-0000-0000-0000-000000000000")
	t.Setenv("KUBE_MCP_ALLOWED_TOOLS", "get_server_version")

	_, activeTools := buildServer()

	assert.Greater(t, len(activeTools), 1)
	assert.Contains(t, activeTools, "get_server_version")
	assert.NotContains(t, activeTools, "list_nodes")
}

func TestBuildServerDisallowedTools(t *testing.T) {
	t.Setenv("KUBE_MCP_BASE_URL", "http://localhost:9000")
	t.Setenv("KUBE_MCP_OIDC_ISSUER_URL", "https://auth.localhost:8443")
	t.Setenv("KUBE_MCP_OIDC_CLIENT_ID", "00000000-0000-0000-0000-000000000000")
	t.Setenv("KUBE_MCP_DISALLOWED_TOOLS", "get_server_version")

	_, activeTools := buildServer()

	assert.Greater(t, len(activeTools), 1)
	assert.NotContains(t, activeTools, "get_server_version")
	assert.Contains(t, activeTools, "list_nodes")
}

func TestStartServer(t *testing.T) {
	t.Skip("fix http listen")
}
