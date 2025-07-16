## **4. The Core Workflow Commands**

This workflow is designed to guide a feature from a high-level idea to a fully implemented and committed piece of work.

### **pdt spec "<feature description>"**

*   **Role**: The "Architect"
*   **Description**: This is the starting point for any new feature. It transforms a founder's high-level goal into a detailed, actionable technical plan.
*   **Functionality**:
    1.  **Interactive Refinement**: The command initiates an interactive session guided by Gemini. It asks clarifying questions to flesh out the feature's description, user stories, and technical requirements (e.g., "Where will this be displayed?", "What data is needed?").
    2.  **Specification Synthesis**: PDT synthesizes the conversation into a structured technical plan.
    3.  **File Creation**: The plan is saved as a new, sequentially numbered markdown file in the `/specs` directory (e.g., `/specs/003-fabric-selection.md`). This file becomes the single source of truth for the feature.

### **pdt code <spec_file>**

*   **Role**: The "Builder"
*   **Description**: This command orchestrates the entire code generation process, using the specification file as its script and the `/pdt_templates` directory as its library.
*   **Functionality**:
    1.  **Parse Specification**: The CLI reads the specified `.md` file to understand the required components (e.g., database tables, UI components, API endpoints).
    2.  **Select Templates**: It intelligently selects the necessary building blocks from the local `/pdt_templates/` directory based on the spec.
    3.  **Construct Master Prompt**: This is the core IP. The CLI assembles a detailed, context-rich prompt for the Gemini API, which includes:
        *   The full feature specification.
        *   The contents of relevant `/pdt_templates`.
        *   The contents of existing code files that need to be modified.
        *   Architectural rules from a `gemini.md` file.
    4.  **Execute and Apply**: The CLI sends the master prompt to Gemini, receives the generated code (new files and diffs), and automatically applies the changes to the local filesystem.

### **pdt test <spec_file>**

*   **Role**: The "Inspector"
*   **Description**: After the code is generated, this command creates the necessary tests to ensure quality and prevent regressions.
*   **Functionality**: Using a similar process to `pdt code`, it analyzes the spec and the generated code to produce a comprehensive suite of tests (unit, integration, etc.) that match the project's existing testing patterns.

### **pdt commit**

*   **Role**: The "Finisher"
*   **Description**: This command finalizes the work by reviewing, committing, and cleaning up the completed task.
*   **Functionality**:
    1.  **Review**: It presents the user with a `git diff` of all the changes made during the session for a final review.
    2.  **Commit**: Upon user approval, it commits all staged changes with a descriptive, AI-generated commit message based on the spec file.
    3.  **Cleanup**: It can optionally perform cleanup tasks, such as archiving the spec file.