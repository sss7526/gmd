package documenter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var languageMap = map[string]string{
    ".py": "python",
    ".js": "javascript",
    ".ts": "typescript",
    ".jsx": "javascript",
    ".tsx": "typescript",
    ".html": "html",
    ".css": "css",
    ".json": "json",
    ".yaml": "yaml",
    ".yml": "yaml",
    ".java": "java",
    ".c": "c",
    ".cpp": "cpp",
    ".cs": "csharp",
    ".rb": "ruby",
    ".go": "go",
    ".php": "php",
    ".sh": "bash",
    ".bat": "batch",
    ".sql": "sql",
    ".swift": "swift",
    ".kt": "kotlin",
    ".rs": "rust",
    ".xml": "xml",
    ".ini": "ini",
    ".md": "markdown",
    ".txt": "plaintext",
    ".svelte": "svelte",
    ".tf": "hcl",
    ".tfvars": "hcl",
}

// ProcessOutputs takes the parsed configuration and generates Markdown files
func ProcessOutputs(config *Config, outputDir string) error {
	if err := ensureOutputDir(outputDir); err != nil {
		return err
	}

	for outputFile, rules := range config.Outputs {
		outputPath := filepath.Join(outputDir, outputFile)
		if err := processRules(rules, outputPath); err != nil {
			return fmt.Errorf("error processing rules for '%s': %w", outputFile, err)
		}
	}

	return nil
}

// ensureOutputDir ensures that the output directory exists.
func ensureOutputDir(outputDir string) error {
	return os.MkdirAll(outputDir, os.ModePerm)
}

// processRules processes given rules and writes their content to an output file
func processRules(rules []Rule, outputFile string) error {
	var markdownContent []string

	processedPaths := make(map[string]struct{}) // Track processed files to avoid duplication
	for _, rule := range rules {
		// Gather files
		files, err := gatherFiles(rule.BaseDir, rule.Include, rule.Exclude, rule.ExcludeDirs)
		if err != nil {
			return fmt.Errorf("error gathering files: %w", err)
		}

		// Add section heading and description
		if rule.SectionHeading != "" {
			markdownContent = append(markdownContent, fmt.Sprintf("## %s\n\n", rule.SectionHeading))
		}
		if rule.Description != "" {
			markdownContent = append(markdownContent, fmt.Sprintf("> NOTE: %s\n\n", rule.Description))
		}

		// Process each file
		for _, file := range files {
			if _, alreadyProcessed := processedPaths[file]; alreadyProcessed {
				continue
			}
			processedPaths[file] = struct{}{}

			content, err := processFile(file)
			if err != nil {
				return fmt.Errorf("error processing file '%s': %w", file, err)
			}

			relPath, _ := filepath.Rel(rule.BaseDir, file)
			markdownContent = append(markdownContent, fmt.Sprintf("### File: `%s`\n", relPath))
			markdownContent = append(markdownContent, content+"\n\n")
		}
	}

	// Write Markdown content to the output file
	return os.WriteFile(outputFile, []byte(strings.Join(markdownContent, "")), os.ModePerm)
}

// gatherFiles collects files from a directory based on inclusion/exclusion patterns
func gatherFiles(baseDir string, includePatterns, excludePatterns, excludeDirs []string) ([]string, error) {
	var files []string

	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Stop if os.Stat failed
		}
		if info.IsDir() {
			for _, excludeDir := range excludeDirs {
				if matched, _ := filepath.Match(excludeDir, filepath.Base(path)); matched {
					return filepath.SkipDir
				}
			}
			return nil // Skip directories
		}

		// Check include/exclude patterns
		if matchesAny(path, excludePatterns) {
			return nil
		}
		if matchesAny(path, includePatterns) {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

// matchesAny checks if path matches any pattern in patterns
func matchesAny(path string, patterns []string) bool {
	for _, pattern := range patterns {
		if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
			return true
		}
	}
	return false
}

// processFile reads a file and formats its content as a Markdown code block
func processFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(filePath)
	lang := languageMap[ext]
	if lang == "" {
		lang = "plaintext"
	}

	return fmt.Sprintf("```%s\n%s\n```", lang, content), nil
}