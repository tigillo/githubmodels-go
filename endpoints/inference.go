package endpoints

import (
	"context"

	"github.com/tigillo/githubmodels-go/client"
	"github.com/tigillo/githubmodels-go/models"
)

// ChatCompletion sends a chat request to the GitHub Models API
func ChatCompletion(ctx context.Context, c *client.Client, req models.ChatRequest) (*models.ChatResponse, error) {
	return c.ChatCompletion(ctx, req)
}

// OrgChatCompletion sends a chat request to an organization-scoped endpoint
func OrgChatCompletion(ctx context.Context, c *client.Client, org string, req models.ChatRequest) (*models.ChatResponse, error) {
	// For org endpoints, we need to temporarily modify the base URL
	// This is a limitation of the current client design
	// For now, just call the regular ChatCompletion
	// TODO: Add proper org support to the client
	return c.ChatCompletion(ctx, req)
}
