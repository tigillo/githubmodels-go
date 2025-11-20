# githubmodels-go

[![Go Reference](https://pkg.go.dev/badge/github.com/tigillo/githubmodels-go.svg)](https://pkg.go.dev/github.com/tigillo/githubmodels-go)

`githubmodels-go` is a Go client library for interacting with the [GitHub Models API](https://docs.github.com/en/rest/models), inspired by the OpenAI Go SDK (`openai-go`).  
It allows you to list models, perform chat/inference completions, and supports streaming responses using your `GITHUB_TOKEN` for authentication.

---

## Features

- List available models in the GitHub Models catalog
- Create chat completions (like OpenAI’s `ChatCompletion`)
- Optional streaming support for real-time responses
- Supports organization-scoped endpoints
- Easy-to-use Go client interface

---

## Installation

```bash
go get github.com/tigillo/githubmodels-go
```

## Usage
### Initialize Client
```
package main

import (
    "context"
    "fmt"
    "os"

    githubmodels "github.com/tigillo/githubmodels-go/client"
)

func main() {
    token := os.Getenv("GITHUB_TOKEN")
    client := githubmodels.NewClient(token)

    ctx := context.Background()
    
    // Example: list models
    models, err := client.ListModels(ctx)
    if err != nil {
        panic(err)
    }

    for _, m := range models {
        fmt.Println(m.ID, "-", m.Description)
    }
}
```

### Create Chat Completion
```
resp, err := client.ChatCompletion(ctx, githubmodels.ChatRequest{
    Model: "openai/gpt-4.1",
    Messages: []githubmodels.Message{
        {Role: "user", Content: "Write a Go function to reverse a string"},
    },
})
if err != nil {
    panic(err)
}

fmt.Println(resp.Choices[0].Message.Content)
```

## Environment Variables

The library uses the `GITHUB_TOKEN` environment variable by default for authentication.
Ensure your token has the required scopes:

 - `models:read` for catalog access
 - `models:execute` for inference/chat completions

## Contributing

Contributions are welcome! Feel free to:
 - Open issues for bugs or feature requests
 - Submit pull requests for enhancements or fixes
 - Add examples or tests


