# Generate Markdown Documentation 📝

`gmd` is a customizable command-line tool for generating Markdown documentation automatically from your project's source files. With support for flexible YAML-based configuration, you can define which files to include, exclude, and organize your documentation with headings, descriptions, and more.

Not indended to replace any language specific documentation generation toolchain. Suggested use case is to generate custom markdown code snippets for your project files to share with your AI code assistant, blog post, etc.


---

## ⭐ Features

- **Automatic Documentation**:
  Generate Markdown files based on rules for including, excluding, and formatting source files within your projects.

- **Language-Specific Code Blocks**:
  Formats your source files as Markdown code blocks with language tags for syntax highlighting.

- **Customization via YAML**:
  Configure inclusion patterns, exclusion patterns, directories, file types, and Markdown headings/descriptions.

- **Flexible CLI**:
  Use subcommands to initialize default configurations or generate documentation swiftly.


## 🚀 Installation

### **Step 1: Install Go**
If you don’t have Go installed yet, download and install it from the [official website](https://go.dev/dl/).

### **Step 2: Install `gmd`**
Use the `go install` command to install the latest version of `gmd` directly from the source.

```bash
go install github.com/sss7526/gmd@latest
```

### **Step 3: Ensure $HOME/go/bin is in your $PATH

```bash
export PATH=${PATH}:$HOME/go/bin
```


## 🛠️ Usage

The `gmd` CLI is straightforward and provides commands to initialize and run the documentation generator.

### **Command 1: Generate Markdown Documentation**
This is the default functionality of `gmd`. By default, it looks for `gmd-config.yaml` in the current directory and generates Markdown files in the `gmd_output/` folder.

```bash
gmd [OPTIONS]
```

#### **Options:**
- `--config`: Specify a custom path to the YAML configuration file.
    ```bash
    gmd --config ./path/to/your-config.yaml
    ```
- `--output_dir`: Specify a custom directory where Markdown files will be saved.
    ```bash
    gmd --output_dir ./docs/
    ```

#### **Full Example**:
If you want to use a custom configuration file and output directory:
```bash
gmd --config ./custom-config.yaml --output_dir ./custom_output
```

---

### **Command 2: Initialize Default Configuration**

The `init` subcommand creates a default `gmd-config.yaml` file in the current directory. This gives you an example configuration to get started.

```bash
gmd init
```

#### **Default Configuration**
The `init` command generates this sample configuration file:

```yaml
outputs:
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
  subdir_docs.md:
    - base_dir: "./subdir"
      include:
        - "*.go" # Include Go files in the `subdir/`
      exclude: []
      section_heading: "Subdirectory Go Files"
      description: >
        Documentation for Go files located in the project's subdirectories.
```

Modify this file to fit your needs.

---

### **Command 3: Show Help**

Print the full usage instructions and available options.

```bash
gmd help
```

---

## 📄 Configuration File Reference

The YAML configuration file (`gmd-config.yaml`) defines how files are processed for documentation. It is structured as follows:

```yaml
outputs:
  <output_filename>.md:
    - base_dir: <base directory>
      include:
        - <glob pattern of files to include>
      exclude:
        - <glob pattern of files to exclude>
      exclude_dirs:
        - <directories to exclude>
      section_heading: <Heading for this section>
      description: <Optional description for the section>
```

### **Fields**:
1. **`outputs`**:
   - A map of output Markdown filenames to rules that define their content.
   - Each filename (e.g., `main_docs.md`) corresponds to one Markdown file.

2. **`base_dir`**:
   - The directory to scan for files to process.

3. **`include`**:
   - File patterns to include (e.g., `*.go` for all Go files, `*.js` for JavaScript files).

4. **`exclude`**:
   - File patterns to exclude (e.g., `test_*.go` for test files).

5. **`exclude_dirs`**:
   - Specific directories to omit during processing.

6. **`section_heading`**:
   - A Markdown heading to display for this section.

7. **`description`**:
   - An optional description paragraph, displayed below the heading.


## 🛡️ License

This project is licensed under the [MIT License](LICENSE).

---

## 🙌 Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue or submit a pull request.

---

## 💬 Questions? Suggestions?

If you have any questions, feel free to reach out or open a [GitHub issue](https://github.com/sss7526/gmd/issues).
