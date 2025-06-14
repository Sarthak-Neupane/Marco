# CLI Agent MCP

A CLI-based AI agent with modular MCP architecture, file-system operations via natural-language commands.

## Project Structure
```text
/
├── cmd/marco/         # main CLI entrypoint
├── internal/              # non-public packages
├── pkg/
│   ├── intentparser/      # intent parsing logic
│   └── mcp/
│       ├── orchestrator/  # routing logic
│       ├── fs/            # file-system MCP module
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

## Quickstart

1. **Clone the repo** and set your module path in `go.mod`:
   ```bash
   git clone github.com/VanTheBast/Marco.git
   cd Marco
   go mod edit -module github.com/VanTheBast/Marco
   go mod tidy
   ```

2. **Build the CLI**:
   ```bash
   go build -o bin/marco ./cmd/marco
   ```

3. **Configure your API key**  
   - Copy the sample into your home directory:
     ```bash
     mkdir -p ~/.marco
     cp configs/config.sample.yaml ~/.marco/config.yaml
     ```
   - Edit `~/.marco/config.yaml` and set your key:
     ```yaml
     intentparser:
       llm_api_key: "sk-…your-key-here…" 
     ```
   - **Ensure** your `.gitignore` includes:
     ```
     # User config
     ~/.marco/config.yaml
     ```

4. **Run the CLI**  
   ```bash
   # Show help
   ./bin/marco --help

   # Example commands:
   ./bin/marco run "list files"
   ./bin/marco run "find TODO in ."
   ```

---

For more details on configuring other modules, contributing, or reporting issues, see the respective files in the repo.
