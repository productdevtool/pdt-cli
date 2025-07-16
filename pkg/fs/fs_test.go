package fs

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestExists(t *testing.T) {
	// Test case 1: File exists
	dir := t.TempDir()
	filePath := filepath.Join(dir, "test_file.txt")
	err := ioutil.WriteFile(filePath, []byte("hello"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	exists, err := Exists(filePath)
	if err != nil {
		t.Fatalf("Exists returned an error: %v", err)
	}
	if !exists {
		t.Errorf("Expected file to exist, but it doesn't")
	}

	// Test case 2: File does not exist
	nonExistentPath := filepath.Join(dir, "non_existent_file.txt")
	exists, err = Exists(nonExistentPath)
	if err != nil {
		t.Fatalf("Exists returned an error for non-existent file: %v", err)
	}
	if exists {
		t.Errorf("Expected file not to exist, but it does")
	}
}

func TestReadTodoFile(t *testing.T) {
	// Test case 1: Valid todo.md file
	dir := t.TempDir()
	todoPath := filepath.Join(dir, "todo.md")
	content := `# Todo\n\n- [ ] Task 1\n- [ ] Task 2\n  - Subtask\n- Another Task\n`
	err := ioutil.WriteFile(todoPath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create todo file: %v", err)
	}

	expectedTasks := []string{"Task 1", "Task 2", "Subtask", "Another Task"}
	actualTasks, err := ReadTodoFile(todoPath)
	if err != nil {
		t.Fatalf("ReadTodoFile returned an error: %v", err)
	}
	if !reflect.DeepEqual(expectedTasks, actualTasks) {
		t.Errorf("Expected tasks %v, but got %v", expectedTasks, actualTasks)
	}

	// Test case 2: Empty todo.md file
	emptyTodoPath := filepath.Join(dir, "empty_todo.md")
	err = ioutil.WriteFile(emptyTodoPath, []byte(""), 0644)
	if err != nil {
		t.Fatalf("Failed to create empty todo file: %v", err)
	}

	expectedTasks = []string{}
	actualTasks, err = ReadTodoFile(emptyTodoPath)
	if err != nil {
		t.Fatalf("ReadTodoFile returned an error for empty file: %v", err)
	}
	if !reflect.DeepEqual(expectedTasks, actualTasks) {
		t.Errorf("Expected empty tasks, but got %v", actualTasks)
	}
}

func TestCreateDirs(t *testing.T) {
	// Test case 1: Create single directory
	dir := t.TempDir()
	singleDir := filepath.Join(dir, "new_dir")
	err := CreateDirs([]string{singleDir})
	if err != nil {
		t.Fatalf("CreateDirs returned an error: %v", err)
	}
	if _, err := os.Stat(singleDir); os.IsNotExist(err) {
		t.Errorf("Expected directory %s to be created, but it doesn't exist", singleDir)
	}

	// Test case 2: Create nested directories
	nestedDir := filepath.Join(dir, "parent", "child")
	err = CreateDirs([]string{nestedDir})
	if err != nil {
		t.Fatalf("CreateDirs returned an error for nested dirs: %v", err)
	}
	if _, err := os.Stat(nestedDir); os.IsNotExist(err) {
		t.Errorf("Expected directory %s to be created, but it doesn't exist", nestedDir)
	}
}

func TestRewriteTodoFile(t *testing.T) {
	// Test case 1: Rewrite with tasks
	dir := t.TempDir()
	todoPath := filepath.Join(dir, "todo.md")
	tasks := []string{"Task A", "Task B"}

	err := RewriteTodoFile(todoPath, tasks)
	if err != nil {
		t.Fatalf("RewriteTodoFile returned an error: %v", err)
	}

	content, err := ioutil.ReadFile(todoPath)
	if err != nil {
		t.Fatalf("Failed to read rewritten todo file: %v", err)
	}
	expectedContent := `# Todo\n\n- [ ] Task A\n- [ ] Task B\n`
	if string(content) != expectedContent {
		t.Errorf("Expected content:\n%s\nGot:\n%s", expectedContent, string(content))
	}

	// Test case 2: Rewrite with no tasks
	emptyTodoPath := filepath.Join(dir, "empty_todo.md")
	err = RewriteTodoFile(emptyTodoPath, []string{})
	if err != nil {
		t.Fatalf("RewriteTodoFile returned an error for empty tasks: %v", err)
	}

	content, err = ioutil.ReadFile(emptyTodoPath)
	if err != nil {
		t.Fatalf("Failed to read empty rewritten todo file: %v", err)
	}
	expectedContent = `# Todo\n\n`
	if string(content) != expectedContent {
		t.Errorf("Expected empty content:\n%s\nGot:\n%s", expectedContent, string(content))
	}
}

func TestExtractCodeBlocks(t *testing.T) {
	// Test case 1: Markdown with code blocks and file paths
	markdown := `
# Header

Some text.

```go // main.go
package main

func main() {
	fmt.Println("Hello, Go!")
}
```

More text.

```python // script.py
print("Hello, Python!")
```
`
	expected := []CodeBlock{
	{FilePath: "main.go", Content: "package main\n\nfunc main() {\n\tfmt.Println(\"Hello, Go!\")\n}\n"},
	{FilePath: "script.py", Content: "print(\"Hello, Python!\")\n"},
}

	actual, err := ExtractCodeBlocks(markdown)
	if err != nil {
		t.Fatalf("ExtractCodeBlocks returned an error: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	// Test case 2: Markdown with no code blocks
	markdown = "# Just text\nNo code here."
	expected = []CodeBlock{}
	actual, err = ExtractCodeBlocks(markdown)
	if err != nil {
		t.Fatalf("ExtractCodeBlocks returned an error for no code blocks: %v", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected empty, got %v", actual)
	}

	// Test case 3: Markdown with code block but no file path
	markdown = `
```javascript
console.log("No path");
```
`
	expected = []CodeBlock{
	{FilePath: "", Content: "console.log(\"No path\");\n"},
}
	actual, err = ExtractCodeBlocks(markdown)
	if err != nil {
		t.Fatalf("ExtractCodeBlocks returned an error for no file path: %v", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetProjectCommand(t *testing.T) {
	// Create a dummy project-description.md
	dir := t.TempDir()
	projectDescPath := filepath.Join(dir, "docs", "project-description.md")
	err := os.MkdirAll(filepath.Dir(projectDescPath), 0755)
	if err != nil {
		t.Fatalf("Failed to create docs directory: %v", err)
	}
	projectDescContent := `
# Project Description

## Commands
- build: `npm run build`
- deploy: `firebase deploy`
- test: `npm test`

## Other Section
`
	err = ioutil.WriteFile(projectDescPath, []byte(projectDescContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write project-description.md: %v", err)
	}

	// Temporarily change the working directory for the test
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	defer os.Chdir(originalDir) // Restore original working directory

	// Test case 1: Command exists
	expectedCommand := "npm run build"
	actualCommand, err := GetProjectCommand("build")
	if err != nil {
		t.Fatalf("GetProjectCommand returned an error: %v", err)
	}
	if actualCommand != expectedCommand {
		t.Errorf("Expected command %s, got %s", expectedCommand, actualCommand)
	}

	// Test case 2: Command does not exist
	expectedError := "command 'nonexistent' not found in project-description.md"
	_, err = GetProjectCommand("nonexistent")
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error %s, got %v", expectedError, err)
	}
}

func TestGetValidationCommands(t *testing.T) {
	// Create a dummy project-description.md
	dir := t.TempDir()
	projectDescPath := filepath.Join(dir, "docs", "project-description.md")
	err := os.MkdirAll(filepath.Dir(projectDescPath), 0755)
	if err != nil {
		t.Fatalf("Failed to create docs directory: %v", err)
	}
	projectDescContent := `
# Project Description

## Automated Validation
- `npm run lint`
- `go test ./...`

## Other Section
`
	err = ioutil.WriteFile(projectDescPath, []byte(projectDescContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write project-description.md: %v", err)
	}

	// Temporarily change the working directory for the test
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	defer os.Chdir(originalDir) // Restore original working directory

	// Test case 1: Commands exist
	expectedCommands := []string{"npm run lint", "go test ./..."}
	actualCommands, err := GetValidationCommands()
	if err != nil {
		t.Fatalf("GetValidationCommands returned an error: %v", err)
	}
	if !reflect.DeepEqual(expectedCommands, actualCommands) {
		t.Errorf("Expected commands %v, got %v", expectedCommands, actualCommands)
	}

	// Test case 2: No validation commands
	projectDescContent = `
# Project Description

## Other Section
`
	err = ioutil.WriteFile(projectDescPath, []byte(projectDescContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write project-description.md: %v", err)
	}

	expectedCommands = []string{}
	actualCommands, err = GetValidationCommands()
	if err != nil {
		t.Fatalf("GetValidationCommands returned an error: %v", err)
	}
	if !reflect.DeepEqual(expectedCommands, actualCommands) {
		t.Errorf("Expected empty commands, got %v", actualCommands)
	}
}
