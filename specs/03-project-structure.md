## **3. Project Structure**

The CLI will follow a standard Go project layout to ensure maintainability and clarity.

* /cmd: The entry point for the application. Each command (todo, spec, code, etc.) will have its own file within this directory, managed by Cobra.
* /pkg: Contains the core, reusable logic of our application.
  * /pkg/ai: A package responsible for executing the gemini-cli command (using Go's os/exec package) and streaming its output.
  * /pkg/prompt: A package containing functions for constructing the master prompts sent to the AI.
  * /pkg/fs: A utility package for handling file system operations (reading specs, writing content, managing task directories).
  * /pkg/task: A new package to manage the task lifecycle, including state changes in the docs/todos/ directories.