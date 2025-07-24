package types

// ClarifyingQuestion represents a question asked by the AI to get more information.
type ClarifyingQuestion struct {
	QuestionID string `json:"questionId"`
	Question   string `json:"question"`
	Answer     string `json:"answer,omitempty"`
}

// PlannerResponse represents the AI's response to the initial spec generation prompt.
type PlannerResponse struct {
	DraftSpec           string               `json:"draftSpec"`
	ClarifyingQuestions []ClarifyingQuestion `json:"clarifyingQuestions"`
}

// CodeBlock represents a single block of code extracted from markdown.
type CodeBlock struct {
	Language string
	FilePath string // The file path where the code should be written
	Content  string
}

// FileOperation represents a single action to be taken on a file, as parsed from a spec.
type FileOperation struct {
	Type        string // e.g., "CREATE", "MODIFY"
	FilePath    string
	Description string
}
