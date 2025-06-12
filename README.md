# CLI Agent MCP

A CLI-based AI agent with modular MCP architecture, focusing on Canvas LMS integration and file-system operations via natural-language commands.

## Project Structure
```text
/
├── cmd/cli-agent/         # main CLI entrypoint
├── internal/              # non-public packages
├── pkg/
│   ├── intentparser/      # intent parsing logic
│   └── mcp/
│       ├── orchestrator/  # routing logic
│       ├── fs/            # file-system MCP module
│       └── canvas/        # Canvas MCP module
├── configs/               # sample config templates
├── docs/                  # design docs, diagrams
├── scripts/               # helper scripts
├── examples/              # demo workflows
├── tests/                 # integration tests or mocks
├── .github/ISSUE_TEMPLATE/ # issue/PR templates
├── go.mod
├── LICENSE
├── CONTRIBUTING.md
└── CODE_OF_CONDUCT.md
```

# Quickstart
Clone repo, set module path in go.mod.

Build with: 
```
go build -o bin/cli-agent ./cmd/cli-agent
```

Copy ``` configs/config.sample.yaml ``` to ``` ~/.cli-agent/config.yaml ``` and adjust.

Run ``` ./bin/cli-agent --help ```
