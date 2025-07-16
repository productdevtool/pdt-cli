## **3.5. The `/pdt_templates` Directory**

A key component of the PDT CLI's effectiveness is the `/pdt_templates` directory. This user-managed directory contains a library of code templates that serve as the high-quality building blocks for new features.

*   **Purpose**: To guide the AI in generating code that is consistent, idiomatic, and adheres to the project's specific architecture and conventions. Instead of asking the AI to generate code from scratch, the CLI instructs it to *adapt* these high-quality templates to the feature's requirements.
*   **Structure**: The directory can be organized by framework or component type (e.g., `/pdt_templates/convex/`, `/pdt_templates/react/`).
*   **Example**: It might contain templates for:
    *   `create_table_schema.ts`: A template for a new database table.
    *   `list_all_query.ts`: A template for a data-fetching query.
    *   `image_grid_selector.tsx`: A pre-built React component.

By using templates, the founder ensures that the AI's output is predictable and follows best practices, dramatically reducing the need for manual refactoring.