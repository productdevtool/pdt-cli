Let's walk through it step-by-step from the founder's perspective. The founder's job is to define *what* they want; the PDT's job is to handle the *how*.

### The Bow Tie E-Commerce Example

**Founder's Goal:** "On my product page for a custom bow tie, I want to let customers choose from a selection of available fabrics. Each fabric has a name and a picture."

-----

#### Step 1: Define the Feature (`pdt spec`)

The founder doesn't start by thinking about code. They start by describing the feature.

  * **Founder's Action:** They open their terminal and run:
    ```bash
    pdt spec "fabric selection for bow ties"
    ```
  * **What PDT Does:**
    1.  The CLI starts an interactive session, guided by Gemini. It asks clarifying questions to turn the idea into a detailed plan.
          * `PDT: Where will the user select the fabric? (e.g., On the product page)`
          * `Founder: On the product page`
          * `PDT: How should the fabrics be displayed? (e.g., A grid of images, a dropdown)`
          * `Founder: A grid of clickable images`
          * `PDT: What data defines a fabric? (e.g., name, image URL, in_stock_yards)`
          * `Founder: name, image, and a price`
    2.  PDT synthesizes this conversation into a structured plan and saves it as a new file: `/specs/003-fabric-selection.md`.

-----

#### Step 2: Build the Feature (`pdt code`)

This is where the magic happens and our `/pdt_templates` directory comes into play.

  * **Founder's Action:** They review the spec file, are happy with it, and run the command to build:
    ```bash
    pdt code specs/003-fabric-selection.md
    ```
  * **What PDT Does (The Orchestration):**
    1.  **Parse Spec:** The CLI reads the `003-fabric-selection.md` file to understand the required components: a new database table for `fabrics`, a query to fetch them, and a "grid of clickable images" UI component on the product page.
    2.  **Select Templates:** The CLI's logic identifies the necessary building blocks from the local `/pdt_templates/` directory. It knows that for this task, it will need:
          * `/pdt_templates/convex/create_table_schema.ts`: A template for defining a new Convex database table.
          * `/pdt_templates/convex/list_all_query.ts`: A template for a standard data-fetching query.
          * `/pdt_templates/react/image_grid_selector.tsx`: A pre-built, high-quality React component template for displaying a grid of selectable images. It has props for `items`, `onSelect`, etc.
    3.  **Construct Master Prompt:** The CLI now assembles a detailed, context-rich prompt for the Gemini API. This is our core IP. The prompt looks something like this:
        ```
        SYSTEM: You are an expert Next.js and Convex developer. Your task is to implement a new feature based on the user's spec, adhering to the architectural rules and using the provided code templates as your guide.

        ARCHITECTURAL RULES:
        ---
        {...contents of gemini.md...}
        ---

        FEATURE SPECIFICATION:
        ---
        {...contents of specs/003-fabric-selection.md...}
        ---

        RELEVANT EXISTING CODE:
        Here is the code for the product page where the new component should be added:
        `app/products/[productId]/page.tsx`:
        ---
        {...contents of the existing product page file...}
        ---

        CODE TEMPLATES TO USE:
        For the database table, use this structure:
        `pdt_templates/convex/create_table_schema.ts`:
        ---
        {...contents of the schema template...}
        ---
        For the UI, use this React component structure:
        `pdt_templates/react/image_grid_selector.tsx`:
        ---
        {...contents of the React component template...}
        ---
        {...and so on for other templates...}

        INSTRUCTIONS:
        1. Create a new file `convex/fabrics.ts` to define the 'fabrics' table and a query to list all fabrics.
        2. Modify `app/products/[productId]/page.tsx` to fetch the fabrics and render the new `FabricSelector` component.
        3. Create the `components/FabricSelector.tsx` file based on the provided template and adapt it for the fabric data.

        Provide the complete code for new files and diffs for modified files.
        ```
    4.  **Execute & Apply:** The CLI sends this prompt to Gemini. Gemini returns the code. The PDT then automatically creates the new files and applies the changes to existing ones in the founder's local project.

-----

#### Step 3: Test & Document (`pdt test` / `pdt doc`)

  * **Founder's Action:** The founder starts the dev server, sees the new fabric selector working, and is happy. They then run:
    ```bash
    pdt test specs/003-fabric-selection.md
    pdt doc specs/003-fabric-selection.md
    ```
  * **What PDT Does:** Using a similar process, PDT generates tests for the new code and updates the project's internal documentation.

The founder successfully added a complex feature without writing a single line of production code. They acted as the director, using the `spec` file as the script, and PDT was the expert crew that handled the execution using our pre-approved, high-quality `pdt_templates` as its guide.
