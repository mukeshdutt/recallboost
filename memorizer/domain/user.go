package domain

import (
	"time"
)

type User struct {
	UserID       int
	Username     string
	FirstName    string
	LastName     string
	Email        string
	Mobile       string
	UserPassword string
	IsActive     bool
	IsVerified   bool
	CreatedBy    string
	CreatedAt    time.Time
	ModifiedBy   int
	ModifiedAt   time.Time
}
