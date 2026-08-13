package model

// Exam represents a multiple-choice exam question.
// Answers[0] is always the correct answer; Answers[1:4] are distractors.
type Exam struct {
	ID       int64     `json:"id"`
	Number   int       `json:"number"`
	Question string    `json:"question"`
	Answers  [4]string `json:"answers"`
}

// ExamInput is the payload accepted when creating a new exam question.
type ExamInput struct {
	Question string    `json:"question"`
	Answers  [4]string `json:"answers"`
}

// PagedExams is a page of exam questions along with pagination metadata.
type PagedExams struct {
	Data       []Exam `json:"data"`
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
	Total      int    `json:"total"`
	TotalPages int    `json:"totalPages"`
}
