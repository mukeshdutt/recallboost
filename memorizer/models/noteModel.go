package models

import "time"

type Note struct {
	NoteID    int       `json:"note_id"`
	TypeID    int       `json:"type_id"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
