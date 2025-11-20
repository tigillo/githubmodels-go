package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tigillo/githubmodels-go/models"
)

// Client is the main GitHub Models API client
type Client struct {
	token   string
	Client  *http.Client
	BaseURL string // exported so tests can override
}

// NewClient creates a new GitHub Models API client
func NewClient(token string) *Client {
	return &Client{
		token:   token,
		Client:  http.DefaultClient,
		BaseURL: "https://models.github.ai", // production default
	}
}

// Model represents a GitHub Models API model
type Model struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

// ListModels returns all available models from the catalog
func (c *Client) ListModels(ctx context.Context) ([]Model, error) {
	url := fmt.Sprintf("%s/catalog/models", c.BaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var models []Model
	if err := json.NewDecoder(resp.Body).Decode(&models); err != nil {
		return nil, err
	}

	return models, nil
}

// ChatCompletion sends a chat completion request to GitHub Models API
func (c *Client) ChatCompletion(ctx context.Context, reqData models.ChatRequest) (*models.ChatResponse, error) {
	url := fmt.Sprintf("%s/inference/chat/completions", c.BaseURL)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("marshal error: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// Parse rate limit headers (do this before checking status so we have them on errors too)
	rateLimit := parseRateLimitHeaders(resp.Header)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// Create a partial response with rate limit info for error cases
		errorResp := &models.ChatResponse{
			RateLimit: rateLimit,
		}
		// Return the partial response so caller can access rate limit info
		// Note: This changes the signature behavior slightly - we return a response even on error
		return errorResp, fmt.Errorf(
			"unexpected status code: %d, response body: %s",
			resp.StatusCode,
			string(body),
		)
	}

	var chatResp models.ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return nil, fmt.Errorf(
			"failed to decode success response: %w (body: %s)",
			err, string(body),
		)
	}

	// Attach rate limit info to response
	chatResp.RateLimit = rateLimit

	return &chatResp, nil
}

// parseRateLimitHeaders extracts rate limit information from HTTP headers
func parseRateLimitHeaders(headers http.Header) models.RateLimitInfo {
	info := models.RateLimitInfo{}

	if limit := headers.Get("X-RateLimit-Limit"); limit != "" {
		fmt.Sscanf(limit, "%d", &info.Limit)
	}

	if remaining := headers.Get("X-RateLimit-Remaining"); remaining != "" {
		fmt.Sscanf(remaining, "%d", &info.Remaining)
	}

	if reset := headers.Get("X-RateLimit-Reset"); reset != "" {
		fmt.Sscanf(reset, "%d", &info.Reset)
	}

	if retryAfter := headers.Get("Retry-After"); retryAfter != "" {
		fmt.Sscanf(retryAfter, "%d", &info.RetryAfter)
	}

	return info
}
