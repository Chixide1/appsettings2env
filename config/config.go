package config

import (
	"errors"
	"os"
	"path/filepath"
)

// Config holds the configuration for the transformation
type Config struct {
	InputFile  string
	OutputFile string
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.InputFile == "" {
		return errors.New("input file is required")
	}

	// Check if input file exists
	if _, err := os.Stat(c.InputFile); os.IsNotExist(err) {
		return errors.New("input file does not exist")
	}

	// Set default output file if not provided
	if c.OutputFile == "" {
		c.OutputFile = ".env"
	}

	return nil
}

// GetOutputDir returns the directory part of the output file
func (c *Config) GetOutputDir() string {
	return filepath.Dir(c.OutputFile)
}
