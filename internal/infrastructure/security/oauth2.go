package security

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"zeusro.com/hermes/internal/core/config"
	"zeusro.com/hermes/internal/core/logprovider"
)

// OAuth2Provider handles OAuth2/OIDC authentication
type OAuth2Provider struct {
	config     *oauth2.Config
	verifier   *oidc.IDTokenVerifier
	provider   *oidc.Provider
	logger     logprovider.Logger
	stateStore map[string]bool // In production, use Redis or database
}

// NewOAuth2Provider creates a new OAuth2 provider
func NewOAuth2Provider(cfg config.Config, log logprovider.Logger) (*OAuth2Provider, error) {
	if !cfg.Security.OAuth2.Enabled {
		return nil, nil
	}

	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, cfg.Security.OAuth2.IssuerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create OIDC provider: %w", err)
	}

	oidcConfig := &oidc.Config{
		ClientID: cfg.Security.OAuth2.ClientID,
	}
	verifier := provider.Verifier(oidcConfig)

	oauth2Config := &oauth2.Config{
		ClientID:     cfg.Security.OAuth2.ClientID,
		ClientSecret: cfg.Security.OAuth2.ClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  cfg.Security.OAuth2.RedirectURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return &OAuth2Provider{
		config:     oauth2Config,
		verifier:   verifier,
		provider:   provider,
		logger:     log,
		stateStore: make(map[string]bool),
	}, nil
}

// GenerateState generates a random state for OAuth2 flow
func (o *OAuth2Provider) GenerateState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	state := base64.URLEncoding.EncodeToString(b)
	o.stateStore[state] = true
	return state, nil
}

// ValidateState validates the OAuth2 state
func (o *OAuth2Provider) ValidateState(state string) bool {
	valid, exists := o.stateStore[state]
	if exists {
		delete(o.stateStore, state)
	}
	return valid && exists
}

// AuthURL returns the OAuth2 authorization URL
func (o *OAuth2Provider) AuthURL(state string) string {
	return o.config.AuthCodeURL(state)
}

// ExchangeCode exchanges authorization code for token
func (o *OAuth2Provider) ExchangeCode(ctx context.Context, code string) (*oidc.IDToken, error) {
	token, err := o.config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %w", err)
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, fmt.Errorf("no id_token in token response")
	}

	idToken, err := o.verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return nil, fmt.Errorf("failed to verify ID token: %w", err)
	}

	return idToken, nil
}

// AuthMiddleware returns a Gin middleware for OAuth2 authentication
func (o *OAuth2Provider) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Authorization header required",
			})
			c.Abort()
			return
		}

		// Verify token
		ctx := c.Request.Context()
		token := authHeader[len("Bearer "):]
		idToken, err := o.verifier.Verify(ctx, token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		// Store user info in context
		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err == nil {
			c.Set("user_id", claims["sub"])
			c.Set("user_email", claims["email"])
		}

		c.Next()
	}
}
