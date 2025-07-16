### **Epic 8: Implement `pdt upgrade`**

*Goal: Enable the CLI to intelligently update itself and project templates.*

* [ ] **Task 8.1: Design and Implement Repository Comparison Logic**
  * Develop functions to compare the user's current repository against the main PDT template (e.g., using Git commands or file hashing).
* [ ] **Task 8.2: Develop AI Prompt for Intelligent Merging**
  * Create a specialized prompt for the AI to generate non-destructive merge instructions for core file updates.
* [ ] **Task 8.3: Implement the `pdt upgrade` Command**
  * Build the command logic to orchestrate the comparison, AI-driven merging, and application of updates.