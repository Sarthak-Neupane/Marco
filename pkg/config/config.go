package config

type IntentParserConfig struct {
    LLMAPIKey string `yaml:"llm_api_key"`
}

type Config struct {
    IntentParser IntentParserConfig `yaml:"intentparser"`
}
