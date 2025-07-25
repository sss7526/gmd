package main


import (
	"flag"
	"log"
	"fmt"
	"os"
	"path/filepath"
	"github.com/sss7526/gmd/documenter"
)

func main() {
	if len(os.Args) < 2 {
		handleMarkdownProcessing()
		return
	}
	switch os.Args[1] {
	case "init":
		handleInitSubcommand()
	case "help", "--help", "-h":
		printHelp()
	default:
		handleMarkdownProcessing()
	}
}

// handleMarkdownProcessing performs the primary functionality of processing Markdown
func handleMarkdownProcessing() {
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

// handleInitSubcommand creates a default configuration YAML file in the current directory
func handleInitSubcommand() {
    // Get the current working directory
    currentDir, err := os.Getwd()
    if err != nil {
        log.Fatalf("Failed to determine current directory: %v", err)
    }

    // Define the path of the config file to create
    configFilePath := filepath.Join(currentDir, "gmd-config.yaml")

    // Check if the config file already exists
    if _, err := os.Stat(configFilePath); err == nil {
        log.Fatalf("A configuration file ('%s') already exists in the current directory.", configFilePath)
    }

    // Write default config file template
    defaultConfig := `outputs:
  # First output file
  main_docs.md:
    - base_dir: "."
      include:
        - "*.go" # Include all Go files
      exclude:
        - "test_*.go" # Exclude test files
      exclude_dirs:
        - "gmd_output" # Exclude the markdown output directory
      section_heading: "Source Code"
      description: >
        This section contains source Go files and associated documentation.
  # Second output file
  subdir_docs.md:
    - base_dir: "./subdir"
      include:
        - "*.go" # Include all Go files in the ` + "`subdir`" + `
      exclude: []
      section_heading: "Subdirectory Go Files"
      description: >
        Documentation for Go files located in the project's subdirectories.
`

    // Create the file
    err = os.WriteFile(configFilePath, []byte(defaultConfig), os.ModePerm)
    if err != nil {
        log.Fatalf("Failed to create configuration file: %v", err)
    }

    // Success message
    fmt.Printf("Default configuration file created: %s\n", configFilePath)
}


// printHelp prints the help/usage message
func printHelp() {
    helpMessage := `
gmd: Generate Markdown Documentation

Usage:
  gmd [OPTIONS]      Processes files and generates Markdown documentation based on gmd-config.yaml.
      --config       Path to config file (defaults to ./gmd-config.yaml).
      --output_dir   Path to write markdown output (defaults to ./gmd_output/).
  gmd init           Creates a default gmd-config.yaml file in the current directory.
  gmd help           Displays this help message.

Configuration File:
  - gmd-config.yaml must exist in the current directory to run 'gmd' if --config option is not passed.

Examples:
  gmd
  gmd init
`
    fmt.Println(helpMessage)
}