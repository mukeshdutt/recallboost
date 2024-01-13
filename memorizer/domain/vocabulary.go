package domain

import "time"

type Vocabulary struct {
	VocabularyID  int
	Vocabulary    string
	Detail        string
	ReferenceFrom string
	UsageType     string
	Antonyms      string
	Synomyms      string
	PartsofSpeech string
	CreatedBy     string
	CreatedAt     time.Time
	ModifiedBy    string
	ModifiedAt    time.Time
}
