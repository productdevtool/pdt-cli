## **2. Core Architecture & Dependencies**

The PDT CLI will be a standalone, statically-linked binary compiled with Golang. This ensures maximum portability and performance, with no runtime dependencies (like Node.js or Python) required on the user's machine.

* **Runtime**: Golang (v1.22+)
* **Execution**: The Go build process will produce a single executable file named pdt. Users can install it via go install, download it from GitHub Releases, or use a package manager like Homebrew.
* **Dependencies (Go Modules)**:
  * **cobra**: The de-facto standard for building powerful, modern CLIs in Go.
  * **AlecAivazis/survey/v2**: A library for creating beautiful, interactive prompts and wizards.
  * **fatih/color**: The most widely used library for colorized terminal output.
  * **briandowns/spinner**: A simple and effective library for displaying spinners during long-running AI operations.

### **Naming Conventions**

*   **Folders and Files**: Use `kebab-case` (e.g., `my-folder`, `my-file.md`).
*   **JavaScript Variables and Database Fields**: Use `camelCase` (e.g., `myVariable`, `databaseFieldName`).