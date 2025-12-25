# Development Guide

## Create a Minikube cluster

```sh
minikube start --addons=dashboard,metrics-server
```

## Access Kubernetes Dashboard

```sh
minikube dashboard --port 8000
```

This will allow access to the Kubernetes Dashboard app as well as the Kubernetes API itself.

## Run the MCP API

```sh
# Switch to the API directory
cd api
# Install dependencies (if not already present)
go mod tidy
# Set required env variables
## Public URL of the MCP server
export KUBE_MCP_BASE_URL=""
## URL of your OIDC Issuer. For local dev, can be https://auth.localhost:8443
export KUBE_MCP_OIDC_ISSUER_URL=""
## Client ID of your Kube MCP OAuth Client 
export KUBE_MCP_OIDC_CLIENT_ID=""
# or provide required config as cli flag
go run . --base-url "" --oidc-issuer-url "" --oidc-client-id ""
# Start the server in out of cluster mode
go run . --out-of-cluster
# Start the server with allowed origins (useful for debugging with inspector, which does not send PRM requests over proxy mode)
go run . --allowed-origins http://localhost:6274
```

## Start the inspector

Start the MCP inspector using npm:

```sh
npx @modelcontextprotocol/inspector --transport http --server-url http://localhost:9000/mcp
# (Optional) Disable auto-open in browser when starting the inspector
MCP_AUTO_OPEN_ENABLED=false npx @modelcontextprotocol/inspector --transport http --server-url http://localhost:9000/mcp
```

Once the inspector opens, ensure the "HTTP" transport is selected and Connection Type is "Via Proxy".


## Helm chart development

For helm chart development, run a helm install pointing to the local chart directory:

```sh
helm install kube-mcp charts/kube-mcp
```

To override the default `values.yaml`, either:
- create a `values.dev.yaml` within `charts/kube-mcp` (additional `values.*.yaml` files are ignored by git)
- override using `--set` flags (more limited than using the values file)

For example:

```sh
helm install kube-mcp charts/kube-mcp \
    --values charts/kube-mcp/values.dev.yaml \
    --set ingress.enabled=true
```