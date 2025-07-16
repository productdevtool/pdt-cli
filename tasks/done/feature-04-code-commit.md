### **Epic 4: Implement `pdt code` & `pdt commit` (The Builder & Finisher)**

*Goal: Build the commands that execute the plan and finalize the work, automatically detecting the active task.*

* [x] **Task 4.1: Build `pdt code` Command Logic**
  * In `/cmd/code.go`, use the active task detection logic from `/pkg/task`.
  * Read the `project-description.md` and the active `task.md`.
  * Use `/pkg/prompt` to generate the master implementation prompt.
  * Use `/pkg/ai` to execute the code generation, showing a spinner.
* [x] **Task 4.2: Implement Automated Validation**
  * Add logic to `pdt code` to run the test or check commands defined in `project-description.md` after code generation. Loop with the AI to fix issues.
* [x] **Task 4.3: Build `pdt commit` Command Logic**
  * In `/cmd/commit.go`, use the active task detection logic.
  * Use `os/exec` to show the user a `git diff`.
  * Upon confirmation, generate a commit message with the AI.
  * Commit the changes.
* [x] **Task 4.4: Implement Task Cleanup**
  * Add logic to `pdt commit` to move the `task.md` from the `/work` directory to the `/done` directory and remove the empty `/work` folder.