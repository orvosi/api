package entity

import "time"

// Auditable represents an auditable entity.
// It contains information of creation, update, and deletion time of the entity.
// It also contains information of actor who does the action.
//
// Every struct that is considered to be auditable must
// compose Auditable.
type Auditable struct {
	CreatedBy string `json:"created_at"`
	UpdatedBy string `json:"updated_at"`

	CreatedAt time.Time `json:"created_by"`
	UpdatedAt time.Time `json:"updated_by"`
}
