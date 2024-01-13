package domain

import "time"

type Phrase struct {
	PhraseID      int
	Phrase        string
	Detail        string
	ReferenceFrom string
	UsageType     string
	PhraseType    string
	CreatedBy     string
	CreatedAt     time.Time
	ModifiedBy    string
	ModifiedAt    time.Time
}
