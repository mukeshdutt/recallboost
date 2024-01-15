package domain

import (
	"time"
)

type Vocabulary struct {
	VocabularyID  int       `gorm:"column:vocabulary_id;primaryKey"`
	Vocabulary    string    `gorm:"column:vocabulary"`
	Detail        string    `gorm:"column:detail"`
	ReferenceFrom string    `gorm:"column:reference_from"`
	UsageType     string    `gorm:"column:usage_type"`
	Antonyms      string    `gorm:"column:antonym"`
	Synomyms      string    `gorm:"column:synomyms"`
	PartsofSpeech string    `gorm:"column:parts_of_speech"`
	CreatedBy     string    `gorm:"column:created_by"`
	CreatedAt     time.Time `gorm:"column:create_at"`
	ModifiedBy    string    `gorm:"column:modified_by"`
	ModifiedAt    time.Time `gorm:"column:modified_at"`
}
