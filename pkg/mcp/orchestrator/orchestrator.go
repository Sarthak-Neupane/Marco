package orchestrator

import (
    "fmt"

    "github.com/VanTheBast/marco/pkg/intentparser"
    "github.com/VanTheBast/marco/pkg/mcp/fs"
)

// RouteIntent routes the parsed intent to the appropriate MCP module.
func RouteIntent(intent *intentparser.Intent) (string, error) {
    switch intent.Module {
    case "fs":
        return fs.HandleIntent(intent.Name, intent.Params)
    default:
        return "", fmt.Errorf("Unknown module: %s", intent.Module)
    }
}
