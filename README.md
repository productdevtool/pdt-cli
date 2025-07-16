# Product Dev Tool (PDT) CLI

## Overview

The Product Dev Tool (PDT) CLI is designed to streamline the development workflow for AI-native founders. It acts as a "master prompt engineer," abstracting away operational friction and transforming high-level goals into production-ready code.

## Installation

### Prerequisites

*   Go (version 1.22 or higher)
*   `gemini-cli` installed and configured (ensure it's in your PATH)

### Using `go install`

```bash
go install github.com/productdevtool/pdt-cli@latest
```

### From Source

1.  Clone the repository:

    ```bash
    git clone https://github.com/productdevtool/pdt-cli.git
    cd pdt-cli
    ```

2.  Build the executable:

    ```bash
    go build -o pdt
    ```

3.  (Optional) Add the `pdt` executable to your system's PATH.

### From Tarball

1.  Download the latest `pdt-dist.tar.gz` from the [releases page](https://github.com/productdevtool/pdt-cli/releases).

2.  Extract the archive:

    ```bash
    tar -xzvf pdt-dist.tar.gz
    ```

3.  Move the `pdt` and `gemini-cli` executables to a directory in your system's PATH (e.g., `/usr/local/bin`):

    ```bash
    sudo mv pdt-dist/pdt /usr/local/bin/
    sudo mv pdt-dist/gemini-cli /usr/local/bin/
    ```

## Usage

### Core Workflow

The PDT CLI guides you through a structured development workflow:

1.  **`pdt todo`**: The Initiator - Select a task and initialize your workspace.
2.  **`pdt spec`**: The Architect - Refine your task into a detailed technical plan using AI.
3.  **`pdt code`**: The Builder - Generate code based on your plan, with automated validation.
4.  **`pdt commit`**: The Finisher - Review, commit, and clean up your completed task.

### Commands

*   **`pdt todo`**
    *   **Description**: Starts the workflow, generates initial project context, and allows task selection.
    *   **Usage**: `pdt todo`

*   **`pdt spec`**
    *   **Description**: Refines a task into a detailed plan using AI.
    *   **Usage**: `pdt spec`

*   **`pdt code`**
    *   **Description**: Executes the implementation plan, instructing the AI to perform the necessary coding, testing, and validation.
    *   **Usage**: `pdt code`

*   **`pdt commit`**
    *   **Description**: Finalizes the work by reviewing, committing, and cleaning up the completed task.
    *   **Usage**: `pdt commit`

*   **`pdt test [spec_file]`**
    *   **Description**: Instructs the AI to write comprehensive tests for a given feature based on its specification file.
    *   **Usage**: `pdt test path/to/your/spec.md`

*   **`pdt doc [spec_file] [code_paths...]`**
    *   **Description**: Instructs the AI to update internal documentation based on a newly implemented feature.
    *   **Usage**: `pdt doc path/to/your/spec.md src/feature.go src/another_file.go`

*   **`pdt write [content_type] [topic]`**
    *   **Description**: A versatile content generation tool for creating external-facing materials.
    *   **Usage**: `pdt write blog "New Feature X Launch"`

*   **`pdt build`**
    *   **Description**: A convenient wrapper for project-specific build commands.
    *   **Usage**: `pdt build`

*   **`pdt deploy`**
    *   **Description**: A convenient wrapper for project-specific deploy commands.
    *   **Usage**: `pdt deploy`

## Workflow Explanation

(Detailed explanation of the workflow will go here, covering how each command contributes to the overall development process.)

## Contributing

(Information on how to contribute to the project will go here.)

## License

(License information will go here.)
