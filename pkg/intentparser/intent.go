package intentparser

import (
    "fmt"
    "context"
    // "strings"

    openai "github.com/sashabaranov/go-openai"
)

// Intent represents a parsed user intent
type Intent struct {
    Module string                 `json:"module"`
    Name   string                 `json:"intent"`
    Params map[string]string `json:"params"`
}


var llm_client *openai.Client

func Init(apiKey string) {
    fmt.Println("Parsing intent with LLM: ", apiKey)
    llm_client = openai.NewClient(apiKey)
}

// Parse takes user input and returns an Intent.
// Stub: integrate LLM or deterministic parser here.
func LLM_Parse(ctx context.Context, cmd string) (*Intent, error) {
    if llm_client == nil {
        return nil, fmt.Errorf("LLM client not initialized")
    }

    fmt.Println("Parsing intent with LLM: ", cmd)

    return &Intent{
        Module: "fs", // Default to file system module for now
        Name:   "", // Default intent for now
        Params: map[string]string{
            "directory": ".", // Default to current directory
        },
    }, nil
}
