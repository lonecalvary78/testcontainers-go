package vscode

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func writeConfig(configFile string, config *Config) error {
	err := os.MkdirAll(filepath.Dir(configFile), 0o755)
	if err != nil {
		return fmt.Errorf("create directory: %w", err)
	}
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}
	header := "// This file is autogenerated by the 'modulegen' tool.\n"
	data = append([]byte(header), data...)
	return os.WriteFile(configFile, data, 0o644)
}
