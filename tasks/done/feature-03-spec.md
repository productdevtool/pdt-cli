### **Epic 3: Implement `pdt spec` (The Architect)**

*Goal: Build the command that refines a task into a detailed plan using AI. This command should be smart enough to find the active task.*

* [x] **Task 3.1: Implement Active Task Detection (/pkg/task)**
  * Create a function to find the current task directory in `docs/todos/work/`.
  * It should gracefully handle cases where there are zero, one, or multiple tasks in the work directory (prompting the user if ambiguous).
* [x] **Task 3.2: Implement AI Executor (/pkg/ai)**
  * Write a function that takes a prompt string and uses `os/exec` to run `gemini-cli`.
  * Ensure the function can stream the output from the AI back to the user's terminal in real-time.
* [x] **Task 3.3: Implement Prompt Generation (/pkg/prompt)**
  * Create a function that generates the prompt for refining a `task.md` file, using `project-description.md` as master context.
* [x] **Task 3.4: Build `pdt spec` Command Logic**
  * In `/cmd/spec.go`, use the `/pkg/task` function to identify the active task.
  * Read the `project-description.md` and the active `task.md`.
  * Use `/pkg/prompt` to build the refinement prompt.
  * Use `/pkg/ai` to execute the prompt and get the AI-generated plan.
  * Update the `task.md` with the new, detailed plan.
* [x] **Task 3.5: Implement Git Integration for Committing the Plan**
  * Use `os/exec` to call `git add` and `git commit` to save the refined `task.md` to version control.