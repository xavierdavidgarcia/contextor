package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Exclusion patterns for directories and files
var (
	excludeDirs  = []string{".git", ".venv", "venv", "__pycache__", "node_modules", "build", "dist"}
	excludeFiles = []string{".DS_Store", ".gitignore"}
	includeExts  = []string{".py", ".sql", ".json", ".yaml", ".yml", ".toml", ".md"}
)

func main() {
	// Define flags
	versionFlag := flag.Bool("version", false, "Print version information")
	versionShortFlag := flag.Bool("v", false, "Print version information")
	flag.Parse()

	// Handle version flag
	if *versionFlag || *versionShortFlag {
		PrintVersion()
		return
	}

	// Check for markdown file argument
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: contextor [options] <markdown_file>")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Get the markdown file path from arguments
	mdFilePath := args[0]

	// Create output file with date in the filename
	currentTime := time.Now()
	dateStr := currentTime.Format("2006-01-02")
	outputFile := fmt.Sprintf("project_context_%s.txt", dateStr)
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Add generation date at the top
	dateTimeHeader := fmt.Sprintf("# Generated on: %s\n\n", currentTime.Format("2006-01-02 15:04:05"))
	file.WriteString(dateTimeHeader)
	
	// Add markdown content from file
	err = addMarkdownContent(file, mdFilePath)
	if err != nil {
		fmt.Printf("Error adding markdown content: %v\n", err)
		os.Exit(1)
	}

	// Add project structure
	addSection(file, "\n# Project Structure")
	writeHorizontalLine(file)
	file.WriteString("\n\n")
	addProjectStructure(file)

	// Add environment variables section
	addSection(file, "\n# Environment Variables")
	writeHorizontalLine(file)
	file.WriteString(`
Required:
- SUPABASE_URL: Supabase project URL
- SUPABASE_KEY: Service role key for admin operations
`)

	// Add database tables section
	addSection(file, "\n# Database Tables")
	writeHorizontalLine(file)
	file.WriteString(`
Main tables:
1. auth.users (managed by Supabase)
2. users_tenants (mapping table)
3. tenants (tenant information)
`)

	// Add authentication flow section
	addSection(file, "\n# Authentication Flow")
	writeHorizontalLine(file)
	file.WriteString(`
1. User signup → Create auth user → Create tenant mapping
2. Email verification required
3. Login → Receive JWT → Use token for authenticated requests
`)

	// Add important notes section
	addSection(file, "\n# Important Notes")
	writeHorizontalLine(file)
	file.WriteString(`
- Using service_role key for admin operations
- Email verification required by default
- RLS policies must be properly configured
- Tenant isolation through RLS policies
`)

	// Add generation timestamp
	timestamp := fmt.Sprintf("\nGenerated on: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	file.WriteString(timestamp)

	// Add separator for files section
	writeHorizontalLine(file)
	addSection(file, "\n# Project Files Content")
	writeHorizontalLine(file)
	file.WriteString("\n\n")

	// Process project files
	processProjectFiles(file)

	fmt.Printf("Context file generated successfully at: %s\n", outputFile)
}

func addMarkdownContent(file *os.File, mdFilePath string) error {
	mdContent, err := os.ReadFile(mdFilePath)
	if err != nil {
		return fmt.Errorf("failed to read markdown file: %v", err)
	}

	_, err = file.Write(mdContent)
	if err != nil {
		return fmt.Errorf("failed to write markdown content: %v", err)
	}

	return nil
}

func writeHorizontalLine(file *os.File) {
	file.WriteString(strings.Repeat("=", 80) + "\n")
}

func addSection(file *os.File, title string) {
	file.WriteString(title + "\n")
}

func addProjectStructure(file *os.File) {
	// Try to use external tree command if available
	cmd := exec.Command("tree", "-L", "3", "--dirsfirst", "-I", "venv|__pycache__|*.pyc|.git|.env|node_modules|build|dist")
	output, err := cmd.CombinedOutput()
	
	if err == nil {
		file.Write(output)
	} else {
		// If tree command fails, use our own directory tree implementation
		pwd, _ := os.Getwd()
		file.WriteString("External 'tree' command not found. Using internal directory tree implementation.\n\n")
		printDirectoryTree(file, pwd, "", 0, 3)
	}
}

// shouldExclude checks if a file or directory should be excluded
func shouldExclude(name string, isDir bool) bool {
	// Skip hidden files/directories
	if strings.HasPrefix(name, ".") {
		return true
	}

	// Skip Python compiled files
	if strings.HasSuffix(name, ".pyc") {
		return true
	}

	// Check directory exclusion list
	if isDir {
		for _, exclude := range excludeDirs {
			if name == exclude {
				return true
			}
		}
		return false
	}

	// Check file exclusion list
	for _, exclude := range excludeFiles {
		if name == exclude {
			return true
		}
	}

	return false
}

// isIncludedExtension checks if a file has an extension we want to include
func isIncludedExtension(name string) bool {
	ext := filepath.Ext(name)
	for _, include := range includeExts {
		if ext == include {
			return true
		}
	}
	return false
}

// printDirectoryTree prints a directory tree recursively
func printDirectoryTree(file *os.File, path string, prefix string, depth int, maxDepth int) {
	if depth > maxDepth {
		return
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintf(file, "%sError reading directory: %v\n", prefix, err)
		return
	}

	// Filter entries
	var filteredEntries []os.DirEntry
	for _, entry := range entries {
		name := entry.Name()
		if shouldExclude(name, entry.IsDir()) {
			continue
		}
		filteredEntries = append(filteredEntries, entry)
	}

	// Print entries
	for i, entry := range filteredEntries {
		name := entry.Name()
		isLast := i == len(filteredEntries)-1

		// Set correct prefix for the line and for children
		connector := "├── "
		newPrefix := prefix + "│   "
		if isLast {
			connector = "└── "
			newPrefix = prefix + "    "
		}

		fmt.Fprintf(file, "%s%s%s\n", prefix, connector, name)

		if entry.IsDir() {
			printDirectoryTree(file, filepath.Join(path, name), newPrefix, depth+1, maxDepth)
		}
	}
}

// processProjectFiles recursively processes files in project directories
func processProjectFiles(file *os.File) {
	pwd, _ := os.Getwd()
	var files []string

	// First, collect all relevant files
	err := filepath.WalkDir(pwd, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if shouldExclude(d.Name(), true) {
				return filepath.SkipDir
			}
			return nil
		}

		// Process only non-excluded files with included extensions
		if !shouldExclude(d.Name(), false) && isIncludedExtension(d.Name()) {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		fmt.Fprintf(file, "Error traversing directories: %v\n", err)
		return
	}

	// Then process each file
	for _, path := range files {
		relPath, err := filepath.Rel(pwd, path)
		if err != nil {
			relPath = path
		}

		// Write file content with appropriate formatting
		writeHorizontalLine(file)
		addSection(file, fmt.Sprintf("\n# File: %s\n", relPath))
		writeHorizontalLine(file)
		file.WriteString("\n\n")

		// Get language for code block based on file extension
		lang := getLanguageForFile(path)

		// Write file content with code block
		if lang != "" {
			file.WriteString("```" + lang + "\n")
		}

		fileContent, err := os.ReadFile(path)
		if err != nil {
			file.WriteString(fmt.Sprintf("Error reading file: %v\n", err))
		} else {
			file.Write(fileContent)
			if !strings.HasSuffix(string(fileContent), "\n") {
				file.WriteString("\n")
			}
		}

		if lang != "" {
			file.WriteString("```\n\n")
		} else {
			file.WriteString("\n")
		}
	}
}

// getLanguageForFile returns the language identifier for markdown code blocks
func getLanguageForFile(path string) string {
	ext := filepath.Ext(path)
	switch ext {
	case ".py":
		return "python"
	case ".sql":
		return "sql"
	case ".json":
		return "json"
	case ".yaml", ".yml":
		return "yaml"
	case ".toml":
		return "toml"
	case ".md":
		return "markdown"
	default:
		return ""
	}
}