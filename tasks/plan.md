# **PDT CLI - Development Plan**

This document provides a high-level overview of the development plan for the Product Dev Tool (PDT) CLI. It tracks the major features, their status, and their correlation with the official project specifications located in the `/specs` directory.

---

## **Completed Features**

These features represent the foundational work on the CLI.

*   **[✓] [Feature 01: Project Foundation](./done/feature-01-foundation.md)**
    *   **Goal:** Set up the basic Go project structure and dependencies.
    *   **Related Specs:** `specs/02-core-architecture.md`, `specs/03-project-structure.md`

*   **[✓] [Feature 02: `pdt todo` (Legacy)](./done/feature-02-todo.md)**
    *   **Goal:** Initial implementation of the task initiation command.
    *   **Note:** This workflow is now deprecated and will be removed as part of Feature 09.

*   **[✓] [Feature 03: `pdt spec` (Legacy)](./done/feature-03-spec.md)**
    *   **Goal:** Initial implementation of the AI-driven specification refinement.
    *   **Note:** This workflow is being replaced by the new interactive model in Feature 09.

*   **[✓] [Feature 04: `pdt code` & `pdt commit` (Legacy)](./done/feature-04-code-commit.md)**
    *   **Goal:** Initial implementation of code generation and commit logic.
    *   **Note:** This workflow is being refactored in Feature 09 to be spec-file driven.

*   **[✓] [Feature 05: Auxiliary Commands](./done/feature-05-aux-cmds.md)**
    *   **Goal:** Build the supporting commands (`build`, `deploy`, `test`, `doc`, `write`).
    *   **Related Specs:** `specs/06-auxiliary-commands.md`

*   **[✓] [Feature 06: Polish & Documentation](./done/feature-06-polish.md)**
    *   **Goal:** Prepare the CLI for initial release with a comprehensive README and polished output.

---

## **What's Next**

This is the current focus of our development efforts.

### **In Progress**

*   **[WIP] [Feature 07: Completion & Refinement](./wip/feature-07-completion-and-refinement.md)**
    *   **Goal:** Solidify the existing codebase by adding comprehensive unit tests to core packages.
    *   **Priority:** High. Must be completed before starting the major refactor.

### **Todo**

*   **[ ] [Feature 09: Refactor CLI to Match New Spec](./todo/feature-09-refactor-cli.md)**
    *   **Goal:** Overhaul the core workflow to match the new, more powerful `pdt spec "..."` model.
    *   **Related Specs:** `specs/05-core-workflow.md`, `specs/04-pdt-templates.md`
    *   **Priority:** Critical. This is the most important next step to align the tool with its vision.

*   **[ ] [Feature 08: Implement `pdt upgrade`](./todo/feature-08-upgrade.md)**
    *   **Goal:** Enable the CLI to intelligently update itself and project templates.
    *   **Related Specs:** `specs/07-roadmap.md`
    *   **Priority:** Low. To be addressed after the core CLI is stable and fully refactored.
