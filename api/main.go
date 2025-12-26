package main

import (
	"log/slog"
	"os"
	"strings"

	"github.com/cturner8/kube-mcp/config"
	"github.com/cturner8/kube-mcp/kubernetes"
	"github.com/cturner8/kube-mcp/server"
)

func main() {
	initLogger()

	kubernetesApiClient := kubernetes.CreateKubernetesApiClient(*config.ServerConfig.OutOfCluster, *config.ServerConfig.Kubeconfig)

	version, err := kubernetesApiClient.Discovery().ServerVersion()
	if err != nil {
		slog.Error("Failed to get Kubernetes server version", "error", err)
		os.Exit(1)
	}

	slog.Info("Connected to Kubernetes API Server", "version", version.String())

	server.StartServer()
}

func initLogger() {
	level := slog.LevelInfo
	switch strings.ToLower(config.ServerConfig.LogLevel) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn", "warning":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level})
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
