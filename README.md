# Kubernetes API MCP

## Development

### Create a Minikube cluster

```sh
minikube start --addons=dashboard,metrics-server
```

### Access Kubernetes Dashboard

```sh
minikube dashboard --port 8000
```

This will allow access to the Kubernetes Dashboard app as well as the Kubernetes API itself.

### Run the MCP API

```sh
# Switch to the API directory
cd api
# Install dependencies (if not already present)
go mod tidy
# Set required env variables
## Public URL of the MCP server
export PUBLIC_BASE_URL=""
## URL of your OIDC Issuer. For local dev, can be https://auth.localhost:8443
export OIDC_ISSUER_URL=""
## Client ID of your Kube MCP OAuth Client 
export OIDC_CLIENT_ID=""
# Start the server in out of cluster mode
go run . --out-of-cluster
# Start the server with allowed origins (useful for debugging with inspector, which does not send PRM requests over proxy mode)
go run . --allowed-origins http://localhost:6274
```

### Start the inspector

Start the MCP inspector using npm:

```sh
npx @modelcontextprotocol/inspector --transport http http://localhost:9000/mcp
# (Optional) Disable auto-open in browser when starting the inspector
MCP_AUTO_OPEN_ENABLED=false npx @modelcontextprotocol/inspector --transport http http://localhost:9000/mcp
```

Once the inspector opens, ensure the "HTTP" transport is selected and Connection Type is "Via Proxy".