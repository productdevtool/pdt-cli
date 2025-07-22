### **Epic 7: Feature Completion & Refinement**

*Goal: Complete and refine existing features, ensuring robustness and full functionality.*

* [x] **Task 7.1: Complete `pdt todo` Project Initialization User Confirmation**
  * Implement the `survey` prompt for user confirmation/editing of the generated `project-description.md`.
* [x] **Task 7.2: Refine `pkg/ai/ai.go` for Robust `gemini-cli` Wrapping**
  * Enhance `ai.Executor` to handle various `gemini-cli` output formats and potential errors gracefully.
  * Ensure proper error propagation and user-friendly messages for AI interaction.
* [x] **Task 7.3: Implement AI Output Parsing and File Writing**
  * For `pdt code`, `pdt test`, `pdt doc`, and `pdt write`, implement the parsing of AI-generated code blocks and content, and write them to the appropriate files.
* [x] **Task 7.4: Implement Automated Validation Loop for `pdt code`**
  * Integrate the execution of project-specific test/check commands after code generation.
  * Implement a loop to re-prompt the AI for fixes if validation fails.
* [ ] **Task 7.5: Add Comprehensive Unit Tests for Core Packages**
  * Write unit tests for `pkg/fs`, `pkg/prompt`, `pkg/task`, and `pkg/ai` to ensure their reliability.