// Package intentparser provides functionality to parse free-form user commands
// into structured Intent objects via an LLM backend.
package intentparser

import (
    "context"
    "encoding/json"
    "fmt"

    "github.com/openai/openai-go"
    "github.com/openai/openai-go/option"
)

const (
    // systemPrompt is the system-level instruction to the LLM to emit only JSON.
    systemPrompt = `You parse user commands into JSON intents. Output *only* JSON.`

    // promptTemplate defines the few-shot examples and schema for transforming
    // a raw user command into the desired JSON structure.
    promptTemplate = `
You are an intent parser. You must output *only* JSON matching this schema:

{
  "module": "fs",
  "intent": "<intent name>",
  "params": { ... }
}

Examples:
User: "List all files in src"
{"module":"fs","intent":"list_dir","params":{"path":"src"}}

User: "Find TODO comments in pkg/"
{"module":"fs","intent":"find_pattern","params":{"pattern":"TODO","path":"pkg"}}

Now parse this command into JSON:
---
%s
---
`
)

// Intent represents the parsed output from the LLM.
// Module specifies which MCP module to invoke (e.g., "fs" or "canvas").
// Name is the specific operation within that module.
// Params holds any additional key/value arguments for the intent.
type Intent struct {
    Module string            `json:"module"`
    Name   string            `json:"intent"`
    Params map[string]string `json:"params"`
}

var (
    // client is the shared OpenAI client instance. Must be initialized
    // via Init before calling LLMParse.
    client openai.Client
)

// Init sets up the OpenAI client with the provided API key.
// This function must be called once before invoking LLMParse.
func Init(apiKey string) {
    client = openai.NewClient(
		option.WithAPIKey(apiKey),
	)
}

// LLMParse sends a natural-language command to the LLM, expecting
// a JSON response that matches our Intent struct.
// It returns an Intent on success, or an error if parsing fails.
func LLMParse(ctx context.Context, userCmd string) (*Intent, error) {
    // Build the chat messages by cloning the static base messages
    // and appending the formatted user prompt.
     // Build messages
     messages := []openai.ChatCompletionMessageParamUnion{
        openai.SystemMessage(systemPrompt),
        openai.UserMessage(fmt.Sprintf(promptTemplate, userCmd)),
    }

    // Create the chat completion
    resp, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
        Model:    "gpt-3.5-turbo", // or openai.GPT4
        Messages: messages,
    })
    if err != nil {
        return nil, fmt.Errorf("LLM API error: %w", err)
    }

    // Extract the JSON string
    raw := resp.Choices[0].Message.Content

    // Unmarshal into your Intent struct
    var intent Intent
    if err := json.Unmarshal([]byte(raw), &intent); err != nil {
        return nil, fmt.Errorf("invalid JSON from LLM: %w\nresponse: %s", err, raw)
    }

    // Validate required fields
    if intent.Module == "" || intent.Name == "" {
        return nil, fmt.Errorf("parsed JSON missing fields: %+v", intent)
    }
    return &intent, nil
}
