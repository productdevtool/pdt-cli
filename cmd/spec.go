package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/productdevtool/pdt-cli/pkg/ai"
	"github.com/productdevtool/pdt-cli/pkg/prompt"
	"github.com/productdevtool/pdt-cli/pkg/types"
	"github.com/spf13/cobra"
)

// extractJSON attempts to find and extract a JSON object from a string that might be wrapped in markdown.
func extractJSON(raw string) string {
	start := strings.Index(raw, "{")
	last := strings.LastIndex(raw, "}")
	if start == -1 || last == -1 || last < start {
		return ""
	}
	return raw[start : last+1]
}

func askQuestions(questions []types.ClarifyingQuestion) ([]types.ClarifyingQuestion, error) {
	var fields []huh.Field
	for i, q := range questions {
		fields = append(fields, huh.NewInput().
			Title(q.Question).
			Value(&questions[i].Answer))
	}

	huhForm := huh.NewForm(huh.NewGroup(fields...))
	err := huhForm.Run()
	if err != nil {
		return nil, fmt.Errorf("error running interactive form: %w", err)
	}

	return questions, nil
}

func generateSpecFilename(specDir, userGoal string) (string, error) {
	files, err := ioutil.ReadDir(specDir)
	if err != nil {
		return "", fmt.Errorf("error reading spec directory: %w", err)
	}

	// Find the highest existing spec number
	highestNum := 0
	for _, file := range files {
		match := regexp.MustCompile(`^s(\d+)-.*\.md$`).FindStringSubmatch(file.Name())
		if len(match) > 1 {
			var num int
			fmt.Sscanf(match[1], "%d", &num)
			if num > highestNum {
				highestNum = num
			}
		}
	}

	// Generate a short name from the user goal
	shortName := strings.ToLower(userGoal)
	shortName = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(shortName, "-")
	shortName = strings.Trim(shortName, "-")
	if len(shortName) > 30 {
		shortName = shortName[:30]
	}

	return fmt.Sprintf("s%03d-%s.md", highestNum+1, shortName), nil
}

var specCmd = &cobra.Command{
	Use:   "spec [high-level goal]",
	Short: "Generates a technical specification from a high-level goal.",
	Long:  `This command initiates a conversation with an AI to transform a high-level goal into a detailed technical specification.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Phase 1: Foundational Changes & Initial AI Call
		// Step 1: Define and Ensure Spec Directory
		specDir := ".pdt/specs"
		if err := os.MkdirAll(specDir, os.ModePerm); err != nil {
			color.Red("Error creating spec directory: %v", err)
			os.Exit(1)
		}

		// Step 2: Update Command Signature & Input Handling
		var userGoal string
		if len(args) > 0 {
			userGoal = strings.Join(args, " ")
		} else {
			huhForm := huh.NewForm(huh.NewGroup(huh.NewInput().Title("What is your high-level goal?").Value(&userGoal)))
			if err := huhForm.Run(); err != nil {
				color.Red("Error getting user input: %v", err)
				os.Exit(1)
			}
		}

		// Step 3: Implement Initial Prompt Construction
		initialPrompt, err := prompt.InitialPlannerPrompt(userGoal)
		if err != nil {
			color.Red("Error building initial planner prompt: %v", err)
			os.Exit(1)
		}

		color.Cyan("Generating initial spec draft and questions with AI...")
		aiOutput, err := ai.Executor("gemini", "-m", "gemini-2.5-flash", "-p", initialPrompt)
		if err != nil {
			color.Red("Error executing AI prompt: %v", err)
			os.Exit(1)
		}

		// Step 4: Update AI Executor and Data Handling
		jsonString := extractJSON(aiOutput)
		if jsonString == "" {
			color.Red("Error: No valid JSON object found in the AI response.")
			os.Exit(1)
		}

		var plannerResponse types.PlannerResponse
		if err := json.Unmarshal([]byte(jsonString), &plannerResponse); err != nil {
			color.Red("Error parsing AI response: %v", err)
			color.Yellow("Raw AI output:\n%s", aiOutput)
			os.Exit(1)
		}

		// Phase 2: Building the Interactive Q&A Loop
		if len(plannerResponse.ClarifyingQuestions) > 0 {
			color.Yellow("AI has clarifying questions:")
			answeredQuestions, err := askQuestions(plannerResponse.ClarifyingQuestions)
			if err != nil {
				color.Red("Error asking questions: %v", err)
				os.Exit(1)
			}

			color.Cyan("Refining spec based on your answers...")
			refinementPrompt, err := prompt.RefineSpecPrompt(userGoal, plannerResponse.DraftSpec, answeredQuestions)
			if err != nil {
				color.Red("Error building refinement prompt: %v", err)
				os.Exit(1)
			}

			aiRefinedOutput, err := ai.Executor("gemini", "-m", "gemini-2.5-flash", "-p", refinementPrompt)
			if err != nil {
				color.Red("Error executing AI refinement prompt: %v", err)
				os.Exit(1)
			}

			jsonString = extractJSON(aiRefinedOutput)
			var refinedResponse types.PlannerResponse
			if err := json.Unmarshal([]byte(jsonString), &refinedResponse); err != nil {
				// If parsing fails, assume the response is just the markdown spec
				plannerResponse.DraftSpec = aiRefinedOutput
			} else {
				plannerResponse.DraftSpec = refinedResponse.DraftSpec
			}
		}

		// Finalize and save the spec
		finalFilename, err := generateSpecFilename(specDir, userGoal)
		if err != nil {
			color.Red("Error generating spec filename: %v", err)
			os.Exit(1)
		}

		finalPath := filepath.Join(specDir, finalFilename)
		if err := ioutil.WriteFile(finalPath, []byte(plannerResponse.DraftSpec), 0644); err != nil {
			color.Red("Error writing final spec: %v", err)
			os.Exit(1)
		}

		color.Green("Specification saved to %s", finalPath)
	},
}

func init() {
	rootCmd.AddCommand(specCmd)
}
