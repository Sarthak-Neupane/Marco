package orchestrator

import (
    "fmt"

    "github.com/VanTheBast/marco/pkg/intentparser"
)

// RouteIntent routes the parsed intent to the appropriate MCP module.
func RouteIntent(intent *intentparser.Intent) (string, error) {
    switch intent.Module {
    case "fs":
        return "", fmt.Errorf("FS module not implemented")
    case "canvas":
        return "", fmt.Errorf("Canvas module not implemented")
    default:
        return "", fmt.Errorf("Unknown module: %s", intent.Module)
    }
}
