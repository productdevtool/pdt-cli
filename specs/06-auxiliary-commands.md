## **5. Auxiliary Commands**

These commands support the core workflow and can be used for specialized, one-off tasks.

### **pdt test <spec_file>**

* **Role**: The "Inspector"
* **Description**: Instructs the AI to write comprehensive tests for a given feature based on its specification file.
* **Functionality**: This is useful for generating tests for legacy code that doesn't have a task.md or for adding more tests to an existing feature. It constructs a specialized prompt focused on testing methodologies (unit, integration, e2e) and instructs the AI to create test files that match the project's existing testing patterns.

### **pdt doc <spec_file>**

* **Role**: The "Chronicler"
* **Description**: Instructs the AI to update internal documentation (e.g., a /handbook directory) based on a newly implemented feature.
* **Functionality**: This command provides the AI with the feature's spec file and the file paths of the implemented code. It then prompts the AI to act as a technical writer, explaining how the feature works, its API, and how to use it, and then appends this to the relevant handbook files.

### **pdt write <type> <topic>**

* **Role**: The "Marketer"
* **Description**: A versatile content generation tool for creating external-facing materials.
* **Functionality**: The user specifies a content type (e.g., blog, tweet, landing-page-copy) and a topic. The CLI constructs a prompt tailored to that format and audience, instructing the AI to generate the content. The output is saved to a new file in a /content directory.

### **pdt build & pdt deploy**

* **Role**: The "Packager" and "Shipper"
* **Description**: Simple, convenient wrappers for project-specific build and deploy commands.
* **Functionality**: These commands use Go's os/exec package to call external commands defined in project-description.md or common build tools (e.g., npm run build, vercel deploy --prod). This allows the PDT CLI to be the single interface for the entire development lifecycle, even when orchestrating tools from other ecosystems.