package domain

import "time"

type Note struct {
	NoteID     int
	UserID     int
	Title      string
	Detail     string
	NoteType   string
	CreatedBy  string
	CreatedAt  time.Time
	ModifiedBy string
	UpdatedAt  time.Time
}
