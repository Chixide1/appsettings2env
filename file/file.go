package file

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ReadJSONFile reads and parses a JSON file
func ReadJSONFile(path string) (map[string]interface{}, error) {
	// Open the JSON file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	// Read file contents
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read input file: %w", err)
	}

	// Parse JSON
	var data map[string]interface{}
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return data, nil
}

// WriteEnvFile writes the flattened key-value pairs to a .env file
func WriteEnvFile(path string, data map[string]interface{}) error {
	// Create output directory if it doesn't exist
	dir := filepath.Dir(path)
	if dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	// Create output file
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	// Write key-value pairs
	for k, v := range data {
		if _, err := file.WriteString(fmt.Sprintf("%s=%q\n", k, v)); err != nil {
			return fmt.Errorf("failed to write to output file: %w", err)
		}
	}

	return nil
}
