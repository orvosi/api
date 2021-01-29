package entity

import "time"

// Auditable represents an auditable entity.
// It contains information of creation, update, and deletion time of the entity.
// It also contains information of actor who does the action.
//
// Every struct that is considered to be auditable must
// compose Auditable.
type Auditable struct {
	CreatedBy string
	UpdatedBy string

	CreatedAt time.Time
	UpdatedAt time.Time
}
