# Kubernetes MCP Server - AI Coding Agent Instructions

## Project Overview

**kube-mcp** is a Model Context Protocol (MCP) server that exposes Kubernetes cluster operations as tools for AI agents. It bridges Kubernetes API access with MCP clients (like Claude, Cursor, etc.) through OAuth2-secured HTTP endpoints.

### Architecture

```
┌─────────────────────────────────────────────────────────────┐
│ MCP Client (Claude, Cursor, etc.)                           │
└────────────────────┬────────────────────────────────────────┘
                     │ HTTP + MCP Protocol
┌────────────────────▼────────────────────────────────────────┐
│ kube-mcp Server (api/main.go)                               │
├─────────────────────────────────────────────────────────────┤
│ ┌─────────────────────────────────────────────────────────┐ │
│ │ Server Layer (api/server/)                              │ │
│ │ • HTTP server with MCP protocol handling                │ │
│ │ • OAuth2/JWT authentication (auth.go)                   │ │
│ │ • CORS middleware (middleware.go)                       │ │
│ └─────────────────────────────────────────────────────────┘ │
│ ┌─────────────────────────────────────────────────────────┐ │
│ │ Tools Layer (api/tools/)                                │ │
│ │ • 20+ Kubernetes resource tools (get/list operations)   │ │
│ │ • Tool registration & filtering (tools.go)              │ │
│ └─────────────────────────────────────────────────────────┘ │
│ ┌─────────────────────────────────────────────────────────┐ │
│ │ Kubernetes Client (api/kubernetes/)                     │ │
│ │ • k8s.io/client-go wrapper                              │ │
│ │ • In-cluster & out-of-cluster modes                     │ │
│ └─────────────────────────────────────────────────────────┘ │
└────────────────────┬────────────────────────────────────────┘
                     │ Kubernetes API
                     ▼
            Kubernetes Cluster
```

## Key Concepts

### Tool Pattern (Critical for Adding New Features)

Every Kubernetes resource has a consistent pattern in `api/tools/`:

1. **List Tool** (`list{Resource}s.go`): Lists all resources, optionally filtered by namespace
   - Params: `namespace` (optional, nil = all namespaces)
   - Returns: JSON array of resources

2. **Get Tool** (`get{Resource}.go`): Retrieves a single resource
   - Params: `name` (required), `namespace` (required for namespaced resources)
   - Returns: JSON object of the resource

**Example**: `listPods.go` + `getPod.go` follow this pattern exactly.

### Tool Registration & Filtering

- Tools are registered in `api/server/server.go` with conditional checks via `tools.IsToolAllowed()`
- Configuration in `api/config/config.go` supports:
  - `--allowed-tools`: Whitelist specific tools (comma-separated)
  - `--disallowed-tools`: Blacklist specific tools (comma-separated)
  - Cannot specify both simultaneously

### Configuration & Environment

All config flows through `api/config/config.go`:

**Required Environment Variables:**
- `KUBE_MCP_BASE_URL`: Public URL of the MCP server (e.g., `https://mcp.example.com`)
- `KUBE_MCP_OIDC_ISSUER_URL`: OIDC provider URL (e.g., `https://auth.localhost:8443`)
- `KUBE_MCP_OIDC_CLIENT_ID`: OAuth2 client ID

**Optional Flags:**
- `--out-of-cluster`: Connect to Kubernetes outside the cluster (uses kubeconfig)
- `--kubeconfig`: Path to kubeconfig file (default: `~/.kube/config`)
- `--allowed-origins`: CORS origins (comma-separated)
- `--allowed-tools` / `--disallowed-tools`: Tool filtering

## Development Workflow

### Local Development Setup

```bash
# 1. Start Minikube with required addons
minikube start --addons=dashboard,metrics-server

# 2. Set environment variables
export KUBE_MCP_BASE_URL="http://localhost:9000"
export KUBE_MCP_OIDC_ISSUER_URL="https://auth.localhost:8443"
export KUBE_MCP_OIDC_CLIENT_ID="10742d2e-1f00-4b19-8412-5a2f77b53b4d"

# 3. Run the server in out-of-cluster mode
cd api
go run . --out-of-cluster

# 4. In another terminal, start the MCP inspector
npx @modelcontextprotocol/inspector --transport http --server-url http://localhost:9000/mcp
```

### Build & Deployment

- **Docker**: Multi-stage build in `api/Dockerfile` (Go 1.25 → Alpine 3.22)
- **Compose**: `compose.yml` for local testing with OAuth2 proxy
- **Kubernetes**: Manifests in `manifests/` with Kustomize support
  - Deployment requires ServiceAccount with appropriate RBAC (see `role.yml`)
  - Image: `ghcr.io/cturner8/kube-mcp:dev`

### Testing Tools

- **MCP Inspector**: Web UI for testing tools interactively
  - Disable auto-open: `MCP_AUTO_OPEN_ENABLED=false npx @modelcontextprotocol/inspector ...`
  - Use "HTTP" transport + "Via Proxy" connection type

## Common Tasks

### Adding a New Kubernetes Resource Tool

1. Create `api/tools/get{Resource}.go` following `getPod.go` pattern
2. Create `api/tools/list{Resource}s.go` following `listPods.go` pattern
3. Register both in `api/server/server.go` with `mcp.AddTool()` calls
4. Update tool filtering logic if needed in `api/tools/tools.go`

**Key Pattern**: All tools use `kubernetesApiClient` (initialized in `tools.go`) and return JSON-marshaled Kubernetes objects.

### Modifying Authentication

- JWT validation: `api/server/auth.go` (uses `github.com/auth0/go-jwt-middleware/v2`)
- Custom claims: Extend `JWTClaims` struct in `auth.go`
- Protected resource metadata: Configure in `getProtectedResourceMetadata()`

### Debugging

- Enable logging: Tools log via `log.Printf()` in handlers
- Check Kubernetes connectivity: `GetServerVersionTool` verifies API access
- CORS issues: Use `--allowed-origins` flag for local development

## Dependencies

- **MCP SDK**: `github.com/modelcontextprotocol/go-sdk` (v1.1.0)
- **Kubernetes**: `k8s.io/client-go` (v0.34.1), `k8s.io/apimachinery` (v0.34.1)
- **Auth**: `github.com/auth0/go-jwt-middleware/v2` (v2.3.0)
- **Go**: 1.25.3+

## File Structure Reference

```
api/
├── main.go                 # Entry point: initializes K8s client, starts server
├── config/                 # Configuration parsing & validation
│   ├── config.go          # Main config struct & parsing logic
│   └── cli.go             # CLI flag definitions
├── kubernetes/            # Kubernetes client wrapper
│   ├── kubernetes.go      # Client initialization (in/out-of-cluster)
│   └── client.go          # Client creation logic
├── server/                # HTTP server & MCP protocol
│   ├── server.go          # Tool registration & MCP server setup
│   ├── auth.go            # OAuth2/JWT validation
│   ├── http.go            # HTTP handler setup
│   └── middleware.go      # Logging & CORS middleware
└── tools/                 # Kubernetes resource tools (20+ files)
    ├── tools.go           # Tool filtering & shared client
    ├── get{Resource}.go   # Get single resource (11 files)
    └── list{Resource}s.go # List resources (9 files)
```

## Conventions

- **Error Handling**: Panic on startup config errors; return errors from tool handlers
- **Logging**: Use `log.Printf()` for debugging; log tool invocations
- **JSON Output**: All tools return Kubernetes objects as JSON strings
- **Namespaces**: Optional in list operations (nil = all namespaces); required in get operations
- **Naming**: Tool names use snake_case (e.g., `get_pod`, `list_pods`)
