package entity

import "github.com/indrasaputra/hashids"

// MedicalRecord holds the user's medical record data
type MedicalRecord struct {
	ID        hashids.ID
	User      *User
	Symptom   string
	Diagnosis string
	Therapy   string
	Result    string
	Auditable
}

// User holds user's information.
type User struct {
	ID       hashids.ID
	Email    string
	Name     string
	GoogleID string
	Auditable
}
