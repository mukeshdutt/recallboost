package domain

import "time"

type UserPhrase struct {
	UserPhraseID  int
	Phrase        string
	ReviseTimes   int
	QuizTimes     int
	QuizActive    bool
	Antonyms      string
	Synomyms      string
	PartsofSpeech string
	CreatedBy     string
	CreatedAt     time.Time
	ModifiedBy    string
	ModifiedAt    time.Time
}
