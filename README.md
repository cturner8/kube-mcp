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