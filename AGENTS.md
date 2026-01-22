# AGENTS.md - Agentic Coding Assistant Guidelines

This file documents conventions and guidelines for agentic coding assistants working in the kube-mcp repository.

## Build/Test/Lint Commands

### Building
```bash
go build -v ./...
```
Builds all packages in the repository with verbose output.

### Testing
```bash
go test -v ./...
```
Runs all tests in the repository with verbose output.

```bash
go test -run TestName -v ./path/to/package
```
Runs a single named test with verbose output.

### Linting
No dedicated linting tool is configured. The project uses VS Code's default `gofmt` and `goimports` formatting standards.

### Docker Build
```bash
docker build --platform linux/amd64,linux/arm64 -f api/Dockerfile -t kube-mcp:latest .
```
Build context is the `api/` directory. Multi-platform builds target `linux/amd64` and `linux/arm64`.

---

## Go Code Style Guidelines

### Functions
- **Exported functions**: Use `PascalCase` (e.g., `HandleToolCall`, `NewClient`)
- **Unexported functions**: Use `camelCase` (e.g., `processRequest`, `validateInput`)

### Imports
Group imports in the following order with blank lines between groups:
1. Standard library imports
2. External package imports
3. Internal package imports

Example:
```go
import (
	"context"
	"fmt"
	"log/slog"

	"github.com/some/external"
	"github.com/another/package"

	"kube-mcp/internal/config"
	"kube-mcp/internal/tools"
)
```

### Error Handling
Use explicit `if err != nil` checks. Log errors with structured logging using `log/slog` at appropriate levels (Info, Warn, Error).

### Types and Structs
- **Exported types**: Use `PascalCase` (e.g., `ToolDefinition`, `Config`)
- **JSON tags**: Always include JSON tags with `omitempty` for optional fields

Example:
```go
type ToolParams struct {
	Name     string `json:"name"`
	Optional string `json:"optional,omitempty"`
}
```

### Variables
- Use descriptive `camelCase` names
- Use pointers for optional values (e.g., `*string`, `*int`)
- Avoid single-letter variable names except in tight loops

### Logging
Use `log/slog` with key-value pairs:
```go
slog.Info("processing request", "userID", userID, "action", action)
slog.Error("failed to process", "error", err, "retries", attempts)
```

### Configuration
- Environment variables use the `KUBE_MCP_` prefix (e.g., `KUBE_MCP_PORT`, `KUBE_MCP_LOG_LEVEL`)
- CLI flags follow the same pattern with hyphens (e.g., `--kube-mcp-port`, `--kube-mcp-log-level`)

### Tools Pattern
Follow this consistent pattern when implementing MCP tools:

1. **Tool definition variable**: Define the tool specification
2. **Params struct**: Define parameters with JSON tags
3. **Handler function**: Implement the tool logic

Example:
```go
var MyTool = &mcp.Tool{
	Name: "my_tool",
	// ... specification
}

type MyToolParams struct {
	Input string `json:"input"`
}

func HandleMyTool(ctx context.Context, params MyToolParams) (string, error) {
	// Implementation
}
```

### Context
Always use `ctx` as the first parameter in handler functions that require context.

---

## Docker Guidelines

### Multi-Stage Build
Use multi-stage builds to minimise final image size:
- **Build stage**: `golang:1.25`
- **Runtime stage**: `alpine:3.23`

### Build Context
The build context is the `api/` directory.

### Non-Privileged User
Run containers as a non-privileged user (UID 10001):
```dockerfile
RUN adduser -D -u 10001 appuser
USER appuser
```

### Health Check
Implement health checks for orchestration:
```dockerfile
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
	CMD curl -f http://localhost:9000/health || exit 1
```

### Build Optimisations
- **Cache mounts**: Use cache mounts for `/go/pkg/mod/` to speed up builds
- **Bind mounts**: Use bind mounts for `go.mod` and `go.sum` to avoid unnecessary copies

### Multi-Platform Builds
Build for `linux/amd64` and `linux/arm64` using BuildKit:
```bash
docker buildx build --platform linux/amd64,linux/arm64 -t registry/kube-mcp:latest .
```

### Container Signing
Sign built images using `cosign` after pushing to registry:
```bash
cosign sign --key cosign.key registry/kube-mcp:latest
```

---

## Helm Guidelines

### Chart Location
Helm charts are located in `charts/kube-mcp/`.

### Values Override Pattern
- `values.yaml`: Default values for all environments
- `values.dev.yaml`: Development environment overrides
- `values.local.yaml`: Local development overrides

### Template Conventions
Use named includes for common labels and selectors:
```yaml
{{ include "kube-mcp.labels" . }}
{{ include "kube-mcp.selectorLabels" . }}
```

### Key Value Sections
Organise values in this order:
- `mcp.*`: MCP-specific configuration
- `replicaCount`: Deployment replica count
- `image`: Image repository, tag, pull policy
- `service`: Service type, port, annotations
- `ingress`: Ingress configuration
- `rbac`: RBAC and service account settings

### Publishing Charts
Publish charts using the OCI registry pattern:
```bash
helm package charts/kube-mcp/
helm push kube-mcp-1.0.0.tgz oci://registry/helm-charts/
```

**Note**: Tagging releases as `chart-v*` automatically triggers the Helm publishing workflow.

---

## Key Workspace Notes

### Environment Variables
All environment variables use the `KUBE_MCP_` prefix for clarity and namespace isolation.

### MCP Transport
The project uses streamable HTTP for MCP transport, not legacy SSE.

### Documentation
Documentation updates should be made to the wiki repository (`kube-mcp.wiki`), not the main codebase.

### Default Port
The application runs on port `9000` by default.

---

## File Locations

| Component | Location |
|-----------|----------|
| Go source code | `api/` |
| Internal packages | `api/internal/` |
| Configuration | `api/cmd/` |
| Helm charts | `charts/kube-mcp/` |
| Chart values | `charts/kube-mcp/values.yaml` |
| Docker build | `api/Dockerfile` |
| Go module files | `api/go.mod`, `api/go.sum` |
