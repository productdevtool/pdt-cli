### **Epic 5: Implement Auxiliary Commands**

*Goal: Build the supporting commands that round out the CLI's functionality.*

* [x] **Task 5.1: Implement `pdt build` & `pdt deploy`**
  * Create simple wrappers in `/cmd/build.go` and `/cmd/deploy.go` that use `os/exec` to call external commands defined in `project-description.md`.
* [x] **Task 5.2: Implement `pdt test`**
  * Build the command logic in `/cmd/test.go`. It will take a spec file as an argument.
  * Create a specialized prompt in `/pkg/prompt` for generating tests based on the spec file and project conventions.
* [x] **Task 5.3: Implement `pdt doc` & `pdt write`**
  * Build the command logic for both, accepting arguments for topics/specs.
  * Create their respective specialized prompts in `/pkg/prompt`.