package models

// Model represents a single model in the GitHub Models catalog
type Model struct {
	ID          string   `json:"id"`          // Unique model ID, e.g., "openai/gpt-4.1"
	Name        string   `json:"name"`        // Human-readable name of the model
	Description string   `json:"description"` // Short description of the model
	Tags        []string `json:"tags"`        // Optional tags for categorization
	CreatedAt   string   `json:"created_at"`  // ISO timestamp when the model was added
	UpdatedAt   string   `json:"updated_at"`  // ISO timestamp of last update
	Owner       string   `json:"owner"`       // Owner of the model (user/org)
}
