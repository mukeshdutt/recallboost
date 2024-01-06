package models

type NoteType struct {
	TypeID   int    `json:"type_id"`
	TypeName string `json:"type_name"`
	IsActive bool   `json:"is_active"`
}
