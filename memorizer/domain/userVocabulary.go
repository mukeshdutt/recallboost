package domain

import "time"

type UserVocabulary struct {
	UserVocabularyID int
	UserID           int
	VocabularyID     int
	ReviseTimes      int
	QuizTimes        int
	ReviseActive     bool
	CreatedBy        string
	CreatedAt        time.Time
	ModifiedBy       string
	ModifiedAt       time.Time
}
