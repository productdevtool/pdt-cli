### **Epic 1: Project Foundation & CLI Scaffolding**

*Goal: Set up the basic Go project structure, install dependencies, and create the main entry point for the CLI application.*

* [x] **Task 1.1: Initialize Go Module**
  * Create a new Git repository.
  * Run `go mod init github.com/your-org/pdt`.
* [x] **Task 1.2: Install Core Dependencies**
  * Run `go get` to install `cobra`, `survey`, `color`, and `spinner`.
* [x] **Task 1.3: Create Basic Cobra Structure**
  * Set up the `main.go` file.
  * Create a root command in `/cmd/root.go`.
  * Add placeholder files for each command (`todo`, `spec`, `code`, `commit`, `test`, `doc`, `write`, `build`, `deploy`) in the `/cmd` directory.
* [x] **Task 1.4: Create Core Packages**
  * Create the directory structure for `/pkg/ai`, `/pkg/prompt`, `/pkg/fs`, and `/pkg/task`.
  * Add placeholder `.go` files in each package to establish the structure.