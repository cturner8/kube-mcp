// Copyright 2025 The Go MCP SDK Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// createLoggingMiddleware creates an MCP middleware that logs method calls.
func createLoggingMiddleware() mcp.Middleware {
	return func(next mcp.MethodHandler) mcp.MethodHandler {
		return func(
			ctx context.Context,
			method string,
			req mcp.Request,
		) (mcp.Result, error) {
			start := time.Now()
			sessionID := req.GetSession().ID()

			// Log request details.
			slog.Debug("MCP request received",
				"session", sessionID,
				"method", method)

			// Call the actual handler.
			result, err := next(ctx, method, req)

			// Log response details.
			duration := time.Since(start)

			if err != nil {
				slog.Error("MCP request error",
					"session", sessionID,
					"method", method,
					"duration", duration,
					"error", err)
			} else {
				slog.Debug("MCP request completed",
					"session", sessionID,
					"method", method,
					"duration", duration)
			}

			return result, err
		}
	}
}

// isOriginAllowed checks if the given origin is in the allowed list.
// An empty allowed list means all origins are permitted.
func isOriginAllowed(origin string, allowedOrigins []string) bool {
	if len(allowedOrigins) == 0 {
		return true
	}
	for _, allowed := range allowedOrigins {
		if strings.EqualFold(origin, allowed) {
			return true
		}
	}
	return false
}

// createCORSMiddleware returns HTTP middleware that adds CORS headers for the given allowed origins.
func createCORSMiddleware(allowedOrigins []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// If there are no allowed origins, skip CORS
			if len(allowedOrigins) == 0 {
				slog.Debug("No allowed origins configured, skipping CORS headers")
				next.ServeHTTP(w, r)
				return
			}

			slog.Debug("Attempting to apply CORS headers")

			origin := r.Header.Get("Origin")
			if origin != "" && isOriginAllowed(origin, allowedOrigins) {
				slog.Debug("Applying CORS headers for origin", "origin", origin)

				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, mcp-protocol-version")
				w.Header().Set("Access-Control-Max-Age", "3600")
			}

			// Handle preflight requests
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			slog.Debug("CORS origin check complete")

			next.ServeHTTP(w, r)
		})
	}
}
