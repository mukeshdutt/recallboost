package domain

import "time"

type Quiz struct {
	QuizID     int
	UserID     int
	WordCount  int
	QuizFor    string
	CreatedBy  string
	CreatedAt  time.Time
	ModifiedBy string
	ModifiedAt time.Time
}
