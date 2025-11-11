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
# Start the server in out of cluster mode
go run . --out-of-cluster
```

### Start the inspector

Start the MCP inspector using npm:

```sh
npx @modelcontextprotocol/inspector --transport http http://localhost:9000
```

Once the inspector opens, ensure the "HTTP" transport is selected and Connection Type is "Via Proxy".