//go:build test
// +build test

package tools

func Reset() {
	clientProvider = &defaultClientProvider{}
}
