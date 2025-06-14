// pkg/config/load.go
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// getConfigPath returns the path to the user’s config file.
func getConfigPath() (string, error) {
	// 1) Allow override via environment
	if custom := os.Getenv("CLI_AGENT_CONFIG"); custom != "" {
		return custom, nil
	}
	// 2) Default to ~/.marco/config.yaml
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine home directory: %w", err)
	}
	dir := filepath.Join(home, ".marco")
	return filepath.Join(dir, "config.yaml"), nil
}

func Load() (*Config, error) {
	_, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot find home dir: %w", err)
	}
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config %s: %w", path, err)
	}
	if err != nil {
		return nil, fmt.Errorf("reading config %s: %w", path, err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}
	if cfg.IntentParser.LLMAPIKey == "" {
		return nil, fmt.Errorf("intentparser.llm_api_key is empty in %s", path)
	}
	return &cfg, nil
}
