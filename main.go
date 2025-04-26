// cmd/jsontoenv/main.go
package main

import (
	"appsettings2env/config"
	"appsettings2env/converter"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line flags
	var inputFile string
	flag.StringVar(&inputFile, "input", "", "Path to input JSON file")
	outputFile := flag.String("output", ".env", "Path to output .env file")

	// Parse flags
	flag.Parse()

	// Check if input is provided as a flag
	// If not, check for positional arguments
	if inputFile == "" {
		// Get remaining args after flags are parsed
		args := flag.Args()
		if len(args) > 0 {
			// Use first positional argument as input file
			inputFile = args[0]
		}
	}

	// Still no input file? Show usage
	if inputFile == "" {
		fmt.Println("Error: Input file is required")
		fmt.Println("Usage:")
		fmt.Println("  1. Using flags: jsontoenv -input=file.json [options]")
		fmt.Println("  2. Using args: jsontoenv file.json [options]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Create configuration
	cfg := config.Config{
		InputFile:  inputFile,
		OutputFile: *outputFile,
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "Configuration error: %v\n", err)
		os.Exit(1)
	}

	// Process the JSON file
	if err := converter.ConvertAppSettingsToEnv(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Transformation complete! Output written to %s\n", cfg.OutputFile)
}
