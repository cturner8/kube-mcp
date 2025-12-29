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
	config.Load()
	cfg := config.GetMcpServerConfig()

	initLogger(cfg.LogLevel)

	kubernetesApiClient := kubernetes.CreateKubernetesApiClient(*cfg.OutOfCluster, *cfg.Kubeconfig)

	version, err := kubernetesApiClient.Discovery().ServerVersion()
	if err != nil {
		slog.Error("Failed to get Kubernetes server version", "error", err)
		os.Exit(1)
	}

	slog.Info("Connected to Kubernetes API Server", "version", version.String())

	server.StartServer()
}

func initLogger(logLevel string) {
	level := slog.LevelInfo
	switch strings.ToLower(logLevel) {
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
