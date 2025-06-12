package intentparser

import "fmt"

// Intent represents a parsed user intent
type Intent struct {
    Module string                 `json:"module"`
    Name   string                 `json:"intent"`
    Params map[string]interface{} `json:"params"`
}

// Parse takes user input and returns an Intent.
// Stub: integrate LLM or deterministic parser here.
func Parse(input string) (*Intent, error) {
    return nil, fmt.Errorf("Parse not implemented")
}
