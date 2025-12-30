//go:build test
// +build test

package config

import "sync"

func Reset() {
	loadOnce = sync.Once{}
	ServerConfig = McpServerConfig{}
}
