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