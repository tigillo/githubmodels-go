package models

// Message represents a single message in a chat request
type Message struct {
	Role    string `json:"role"`    // "user", "system", "assistant"
	Content string `json:"content"` // Message content
}

// ChatRequest represents a request to the chat completion endpoint
type ChatRequest struct {
	Model    string    `json:"model"`    // Model ID, e.g., "openai/gpt-4.1"
	Messages []Message `json:"messages"` // Conversation messages
}

// Choice represents a single choice in the chat response
type Choice struct {
	Message Message `json:"message"` // The generated message from the model
}

// RateLimitInfo contains rate limit information from GitHub API response headers
type RateLimitInfo struct {
	Limit      int   // X-RateLimit-Limit: Maximum requests per hour
	Remaining  int   // X-RateLimit-Remaining: Requests remaining in current window
	Reset      int64 // X-RateLimit-Reset: Unix timestamp when the limit resets
	RetryAfter int   // Retry-After: Seconds to wait before retrying (only on 429)
}

// Usage contains token usage information from the API response
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatResponse represents the response from the chat completion endpoint
type ChatResponse struct {
	ID        string        `json:"id"`      // Response ID
	Object    string        `json:"object"`  // Type of object, e.g., "chat.completion"
	Choices   []Choice      `json:"choices"` // List of choices
	Usage     Usage         `json:"usage"`   // Token usage information
	RateLimit RateLimitInfo // Rate limit information from response headers
}
