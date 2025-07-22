### **Feature 9: Refactor CLI to Match New Spec**

*Goal: Update the CLI's core logic to align with the new, more powerful workflow defined in the `/specs` directory.*

* [ ] **Task 9.1: Refactor `pdt spec`**
  * Modify the `spec` command to accept a feature description string as an argument (e.g., `pdt spec "a new feature"`).
  * Implement the interactive refinement logic using the `survey` library to ask clarifying questions.
  * Implement the logic to synthesize the conversation and save it as a new, numbered file in the `/specs` directory.

* [ ] **Task 9.2: Refactor `pdt code`**
  * Modify the `code` command to accept a spec file path as an argument (e.g., `pdt code specs/001-new-feature.md`).
  * Remove the old "active task detection" logic.
  * Implement the new logic to parse the spec, select templates from `/pdt_templates`, and construct the master prompt.

* [ ] **Task 9.3: Refactor `pdt commit`**
  * Modify the `commit` command to no longer rely on an "active task."
  * It should now simply stage all current changes and use the AI to generate a commit message based on the diff.

* [ ] **Task 9.4: Remove `pdt todo`**
  * Since the new workflow starts with `pdt spec`, the `todo` command is now obsolete.
  * Remove the `cmd/todo.go` file and all related logic.

* [ ] **Task 9.5: Update `README.md`**
  * Update the main `README.md` to reflect the new `pdt spec "..."` workflow and remove any mention of `pdt todo`.