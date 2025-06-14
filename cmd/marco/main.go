package main

import (
	"fmt"
	"os"
    "context"
    
    "github.com/VanTheBast/marco/pkg/config"
	"github.com/spf13/cobra"
    "github.com/VanTheBast/marco/pkg/mcp/orchestrator"
    "github.com/VanTheBast/marco/pkg/intentparser"
)

func main() {

    // Loading congiguration
    cfg, err := config.Load()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Config error:", err)
        os.Exit(1)
    }

    intentparser.Init(cfg.IntentParser.LLMAPIKey)

	rootCmd := &cobra.Command{Use: "Marco CLI Agent"}
	runCmd := &cobra.Command{
		Use:   "marco [text]",
		Short: "Parse and execute a natural-language command",
		RunE: func(cmd *cobra.Command, args []string) error {
            input := args[0]

            // 4. Parse intent (LLM-backed or fallback)
            ctx := context.Background()
            intent, err := intentparser.LLMParse(ctx, input)
            if err != nil {
                return fmt.Errorf("parse error: %w", err)
            }

            fmt.Println("Parsed intent:", intent)
            fmt.Println("Module:", intent.Module)
            fmt.Println("Name:", intent.Name)
            fmt.Println("Params:", intent.Params)

            for key, value := range intent.Params {
                fmt.Printf("Key: %s, Value: %s\n", key, value)
            }

            // 5. Route & execute
            out, err := orchestrator.RouteIntent(intent)
            if err != nil {
                return fmt.Errorf("exec error: %w", err)
            }
            fmt.Print(out)
            return nil
        },
	}
	rootCmd.AddCommand(runCmd)
	// rootCmd.Execute()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
