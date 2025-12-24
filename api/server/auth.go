// Copyright 2025 The Go MCP SDK Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/oauthex"

	"github.com/cturner8/kube-mcp/config"
)

// JWTClaims represents the claims in our JWT tokens.
type JWTClaims struct {
	Scope        string `json:"scope"`
	Sub          string `json:"sub"`
	Username     string `json:"preferred_username"`
	Email        string `json:"email"`
	ShouldReject bool   `json:"shouldReject,omitempty"`
}

// Validate errors out if `ShouldReject` is true.
func (c *JWTClaims) Validate(ctx context.Context) error {
	if c.ShouldReject {
		return errors.New("should reject was set to true")
	}
	return nil
}

func getCustomClaims() validator.CustomClaims {
	return &JWTClaims{}
}

// getProtectedResourceMetadata returns the OAuth2 Protected Resource Metadata
// describing this server's capabilities and requirements.
func getProtectedResourceMetadata() *oauthex.ProtectedResourceMetadata {
	baseUrl := config.ServerConfig.BaseURL.String()
	issuerURL := config.ServerConfig.OidcIssuerURL.String()
	signingMethod := config.ServerConfig.SigningMethod
	scopes := config.ServerConfig.Scopes

	return &oauthex.ProtectedResourceMetadata{
		// Required: The resource identifier for this server
		Resource: baseUrl,
		// Optional: Authorization servers that can issue tokens for this resource
		AuthorizationServers: []string{issuerURL},
		// Optional: Human-readable name for the resource
		ResourceName: "Kubernetes API MCP Server",
		// Optional: Documentation URL for developers
		ResourceDocumentation: "https://github.com/cturner8/kube-mcp",
		// Optional: Scopes supported by this resource
		ScopesSupported: scopes,
		// Optional: Bearer token methods supported
		BearerMethodsSupported: []string{"header"},
		// Optional: JWS signing algorithms supported by the resource
		ResourceSigningAlgValuesSupported: []string{signingMethod},
		// Optional: Support for Authorization Details (RFC 9396)
		AuthorizationDetailsTypesSupported: []string{},
		// Optional: DPoP support
		DPOPBoundAccessTokensRequired: false,
	}
}

func createBearerAuth(baseUrl string, prmPath string) func(http.Handler) http.Handler {
	jwksProvider := jwks.NewCachingProvider(&config.ServerConfig.OidcIssuerURL, time.Minute*5) // Cache JWKS for 5 minutes
	// Determine signing method dynamically
	signingMethod := config.ServerConfig.SigningMethod
	signingValidator := validator.RS256
	if strings.EqualFold(signingMethod, "HS256") {
		signingValidator = validator.HS256
	}
	// Set up the validator using the chosen algorithm
	jwtValidator, err := validator.New(
		jwksProvider.KeyFunc,
		signingValidator,
		config.ServerConfig.OidcIssuerURL.String(),
		[]string{config.ServerConfig.OidcClientID},
		validator.WithCustomClaims(getCustomClaims),
		validator.WithAllowedClockSkew(30*time.Second),
	)
	if err != nil {
		panic("Error setting up JWT validator: " + err.Error())
	}

	scopes := config.ServerConfig.Scopes

	authOptions := &auth.RequireBearerTokenOptions{
		Scopes:              scopes,
		ResourceMetadataURL: fmt.Sprintf("%s%s", baseUrl, prmPath),
	}
	return auth.RequireBearerToken(func(ctx context.Context, tokenString string, _ *http.Request) (*auth.TokenInfo, error) {
		validatedClaims, err := jwtValidator.ValidateToken(ctx, tokenString)
		if err != nil {
			// Return standard error for invalid tokens.
			return nil, fmt.Errorf("%w: %v", auth.ErrInvalidToken, err)
		}

		// Extract claims and verify token validity.
		claims, ok := validatedClaims.(*validator.ValidatedClaims)
		if !ok {
			return nil, fmt.Errorf("%w: invalid token claims", auth.ErrInvalidToken)
		}

		customClaims, ok := claims.CustomClaims.(*JWTClaims)
		if !ok {
			return nil, fmt.Errorf("%w: invalid custom claims type", auth.ErrInvalidToken)
		}

		return &auth.TokenInfo{
			Scopes:     strings.Split(customClaims.Scope, " "),       // User permissions
			Expiration: time.Unix(claims.RegisteredClaims.Expiry, 0), // Token expiration time
		}, nil
	}, authOptions)
}

// getProtectedResourceMetadataHandler returns the Protected Resource Metadata
// as JSON. This endpoint is typically served at /.well-known/oauth-protected-resource
func getProtectedResourceMetadataHandler() http.HandlerFunc {
	metadata := getProtectedResourceMetadata()

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(metadata); err != nil {
			http.Error(w, "Failed to encode metadata", http.StatusInternalServerError)
		}
	}
}
