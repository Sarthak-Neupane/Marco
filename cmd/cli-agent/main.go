package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

func main() {
    rootCmd := &cobra.Command{
        Use:   "cli-agent",
        Short: "CLI Agent MCP - AI-driven CLI for Canvas and FS",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Welcome to CLI Agent MCP. Use --help to see commands.")
        },
    }

    // add subcommands here

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
