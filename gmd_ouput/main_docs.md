## Source Code

> NOTE: This section contains source Go files and documentation Markdown

### File: `main.go`
```go
package main


import (
	"flag"
	"log"
	"fmt"
	"os"
	"path/filepath"
	"gmd/documenter"
)

func main() {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to determine current directory: %v", err)
	}

	configFileDefault := filepath.Join(currentDir, "gmd-config.yaml")
	outputDirDefault := filepath.Join(currentDir, "gmd_ouput")

	// Defint CLI flags
	configPath := flag.String("config", configFileDefault, "Path to the YAML configuration file. Default is 'gmd-config.yaml'.")
	outputDir := flag.String("output_dir", outputDirDefault, "Directory to save the generated Markdown files. Default is 'gmd_output'.")

	// Parse CLI flags
	flag.Parse()

	// Validate paths
	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatalf("Configuration file '%s' does not exist. Use the --config flag to specify the correct file.", *configPath)
	}

	// Load configuration
	config, err := documenter.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	// Process outputs specified in the config
	fmt.Printf("Processing configuration file: %s\n", *configPath)
	fmt.Printf("Output will be written to: %s\n", *outputDir)
	if err := documenter.ProcessOutputs(config, *outputDir); err != nil {
		log.Fatalf("Error processing outputs: %v", err)
	}

	log.Println("Markdown documentation generated successfully.")
}
```

