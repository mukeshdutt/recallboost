package domain

import "time"

type Revision struct {
	RevisionID      int
	UserID          int
	RevisionFor     string
	RevisionComment string
	CreatedBy       string
	CreatedAt       time.Time
	ModifiedBy      string
	ModifiedAt      time.Time
}
