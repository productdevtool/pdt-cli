### **Epic 2: Implement `pdt todo` (The Initiator)**

*Goal: Build the command that starts the workflow, generates the initial project context, and allows the user to select a task.*

* [x] **Task 2.1: Implement Project Initialization**
  * On first run, check for the existence of `docs/project-description.md`.
  * If not found, use the `/pkg/ai` executor to analyze the codebase and generate the project description file.
  * Use the `survey` library to prompt the user to confirm or edit the generated description.
* [x] **Task 2.2: Implement File System Logic (/pkg/fs)**
  * Write a function to check for the existence of `docs/todo.md`.
  * Write a function to read `docs/todo.md` and parse its contents.
  * Write a a function to create the `docs/todos/work` and `docs/todos/done` directories if they don't exist.
* [x] **Task 2.3: Build Task Selection UI**
  * In `/cmd/todo.go`, use the `/pkg/fs` functions to read the todo list.
  * Use the `survey` library to present the list to the user as a selectable menu.
* [x] **Task 2.4: Implement Workspace Initialization**
  * When a user selects a task, create the unique task folder (e.g., `docs/todos/work/2025-07-12-task-name/`).
  * Create the initial `task.md` file inside the new directory.
  * Rewrite the main `docs/todo.md` file, removing the selected task.
* [ ] **Task 2.5 (Stretch): Implement Orphaned Task Check**
  * Add logic to scan `docs/todos/work/` for directories that don't have a corresponding active process and prompt the user to resume or reset them.