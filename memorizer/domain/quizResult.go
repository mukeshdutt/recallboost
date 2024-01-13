package domain

import "time"

type QuizResult struct {
	QuizResultID  int
	QuizID        int
	QuestionCount int
	AnswerCount   int
	Result        float32
	QuizComment   string
	CreatedBy     string
	CreatedAt     time.Time
	ModifiedBy    string
	ModifiedAt    time.Time
}
