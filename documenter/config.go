package documenter

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config models the structure of the YAML configuration file
type Config struct {
	Outputs map[string][]Rule `yaml:"outputs"` // Maps output filenames to the rules that define their content
}

// Rule defines how to process files in a specific directory
type Rule struct {
	BaseDir			string		`yaml:"base_dir"`			// Root directory to process files from
	Include			[]string 	`yaml:"include"`			// Glob patterns to include files
	Exclude			[]string	`yaml:"exclude"`			// Glob pattersn to exclude files
	ExcludeDirs		[]string	`yaml:"exclude_dirs"`		// Excplicity directories to exclude
	SectionHeading	string 		`yaml:"section_heading"`	// Heading for markdown section
	Description		string		`yaml:"description"`		// Optional description for the section
}

// LoadConfig loads and parses a YAML configuration file into a Config struct
func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	// Ensure the 'outputs' section exists
	if config.Outputs == nil {
		return nil, fmt.Errorf("'outputs' section missing in configuration")
	}

	return &config, nil
}