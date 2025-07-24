package fs

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/productdevtool/pdt-cli/pkg/types"
)

// Exists checks if a file or directory exists.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ReadTodoFile reads a todo file and returns a list of tasks.
func ReadTodoFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// We'll skip empty lines and lines that are just markdown headers or separators
		if line != "" && !strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "---") {
			// Remove the markdown checkbox prefix if present
			if strings.HasPrefix(line, "- [ ] ") {
				line = strings.TrimPrefix(line, "- [ ] ")
			}
			tasks = append(tasks, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// CreateDirs creates a list of directories if they don't already exist.
func CreateDirs(dirs []string) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}

// RewriteTodoFile rewrites the todo file with the given tasks.
func RewriteTodoFile(path string, tasks []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("# Todo\n\n")
	if err != nil {
		return err
	}

	for _, task := range tasks {
		_, err := file.WriteString(fmt.Sprintf("- [ ] %s\n", task))
		if err != nil {
			return err
		}
	}

	return nil
}

// CodeBlock represents a code block extracted from markdown.
type CodeBlock struct {
	FilePath string
	Content  string
}

// ExtractCodeBlocks extracts code blocks from a markdown string.
func ExtractCodeBlocks(markdown string) ([]CodeBlock, error) {
	var codeBlocks []CodeBlock
	scanner := bufio.NewScanner(strings.NewReader(markdown))

	inCodeBlock := false
	currentFilePath := ""
	currentContent := strings.Builder{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "```") {
			if inCodeBlock {
				// End of a code block
				codeBlocks = append(codeBlocks, CodeBlock{
					FilePath: currentFilePath,
					Content:  currentContent.String(),
				})
				inCodeBlock = false
				currentFilePath = ""
				currentContent.Reset()
			} else {
				// Start of a code block
				inCodeBlock = true
				// Try to extract file path from the line, e.g., ```go // path/to/file.go
				parts := strings.Fields(line)
				if len(parts) > 1 && strings.HasPrefix(parts[1], "//") {
					currentFilePath = strings.TrimSpace(strings.TrimPrefix(parts[1], "//"))
				} else if len(parts) > 1 && !strings.Contains(parts[1], " ") {
					// If it's just a language, assume no path for now
					currentFilePath = ""
				}
			}
		} else if inCodeBlock {
			currentContent.WriteString(line + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return codeBlocks, nil
}

// GetProjectCommand extracts a specific command from project-description.md.
func GetProjectCommand(commandName string) (string, error) {
	content, err := os.ReadFile("docs/project-description.md")
	if err != nil {
		return "", fmt.Errorf("error reading project-description.md: %w", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	inCommandsSection := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "## Commands") {
			inCommandsSection = true
			continue
		}

		if inCommandsSection && strings.HasPrefix(line, fmt.Sprintf("- %s: `", commandName)) {
			parts := strings.Split(line, "`")
			if len(parts) >= 2 {
				return parts[1], nil
			}
		}

		if inCommandsSection && strings.HasPrefix(line, "## ") && !strings.HasPrefix(line, "## Commands") {
			// Exited the commands section
			break
		}
	}

	return "", fmt.Errorf("command '%s' not found in project-description.md", commandName)
}

// GetValidationCommands extracts validation commands from project-description.md.
func GetValidationCommands() ([]string, error) {
	content, err := os.ReadFile("docs/project-description.md")
	if err != nil {
		return nil, fmt.Errorf("error reading project-description.md: %w", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	inValidationSection := false
	var commands []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "## Automated Validation") {
			inValidationSection = true
			continue
		}

		if inValidationSection {
			if strings.HasPrefix(line, "- `") && strings.HasSuffix(line, "`") {
				// Extract command between backticks
				cmd := strings.TrimPrefix(line, "- `")
				cmd = strings.TrimSuffix(cmd, "`")
				commands = append(commands, cmd)
			} else if strings.HasPrefix(line, "## ") {
				// Exited the validation section
				break
			} else if line == "" && len(commands) > 0 {
				// End of commands if there's a blank line after commands have been found
				break
			}
		}
	}

	return commands, nil
}

// ParseSpec parses a markdown spec file and extracts file operations.
func ParseSpec(specContent string) ([]types.FileOperation, error) {
	var operations []types.FileOperation
	// Split the spec by the "### " delimiter to process each operation chunk separately.
	// This avoids the need for complex lookaheads which are not supported in Go's regex engine.
	chunks := strings.Split(specContent, "\n### ")

	// The regex is now simpler, as it only needs to parse the content of each chunk.
	re := regexp.MustCompile("(?s)(CREATE|MODIFY): `([^`]+)`\\n(.*)")

	for _, chunk := range chunks {
		// Skip empty chunks that can result from the split
		if strings.TrimSpace(chunk) == "" {
			continue
		}

		matches := re.FindStringSubmatch(chunk)

		if len(matches) == 4 {
			op := types.FileOperation{
				Type:        matches[1],
				FilePath:    strings.TrimSpace(matches[2]),
				Description: strings.TrimSpace(matches[3]),
			}
			operations = append(operations, op)
		}
	}

	return operations, nil
}
