# Contextor

A tool for generating project context documentation for LLMs (Large Language Models).

## Overview

Contextor is a Go binary that helps you generate comprehensive context files for your projects. It's especially useful when working with LLMs that need to understand your project structure and codebase. The tool:

1. Takes a markdown file with project overview information
2. Generates a directory tree structure for your project
3. Includes the content of relevant source files
4. Creates a consolidated, well-formatted context file

## Installation

### Prerequisites

- Go 1.16 or later

### Building from Source

Clone the repository and build the binary:

```bash
git clone https://github.com//xavierdavidgarcia/contextor.git
cd contextor
make build
```

### Installation Options

#### User Installation (recommended)

This installs the binary to ~/bin:

```bash
make install
```

If ~/bin is not in your PATH, add the following to your shell profile:

```bash
export PATH=$PATH:$HOME/bin
```

#### System Installation

This installs the binary system-wide (requires sudo):

```bash
make install-system
```

## Usage

### Basic Usage

Create a markdown file with your project overview information:

```bash
# Create a project overview file
nano project-overview.md
```

Then run contextor with the markdown file:

```bash
contextor project-overview.md
```

This will generate a file named `project_context_YYYY-MM-DD.txt` containing:
- Your project overview (from the markdown file)
- The project directory structure
- Predefined sections for environment variables, database tables, etc.
- Contents of all relevant source files in your project

### Options

```
contextor [options] <markdown_file>

Options:
  -v, -version   Print version information
```

## File Filtering

Contextor includes the following files by default:
- Python (.py)
- SQL (.sql)
- JSON (.json)
- YAML (.yaml, .yml)
- TOML (.toml)
- Markdown (.md)

It excludes:
- Hidden files and directories
- Virtual environments (venv, .venv)
- Python cache directories (__pycache__)
- Node modules (node_modules)
- Build directories (build, dist)
- Git directories (.git)
- Compiled Python files (.pyc)

## Output Format

The generated context file contains:

1. Generation timestamp
2. Project overview (from your markdown file)
3. Project structure (directory tree)
4. Standard sections:
   - Environment Variables
   - Database Tables
   - Authentication Flow
   - Important Notes
5. Source code from all relevant files with proper syntax highlighting

## Examples

### Minimal Project Overview

Create a simple `overview.md` file:

```markdown
# My Project

A simple REST API built with FastAPI.

## Features

- User authentication
- Database integration
- API versioning
```

Generate context:

```bash
contextor overview.md
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Version History

- 0.1.0: Initial release with basic functionality
