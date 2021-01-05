package entity

import "github.com/indrasaputra/hashids"

// MedicalRecord holds the user's medical record data
type MedicalRecord struct {
	ID        hashids.ID `json:"id"`
	User      *User      `json:"user"`
	Symptom   string     `json:"symptom"`
	Diagnosis string     `json:"diagnosis"`
	Therapy   string     `json:"therapy"`
	Result    string     `json:"result"`
	Auditable
}

// User holds user's information.
type User struct {
	ID       hashids.ID `json:"id"`
	Email    string     `json:"email"`
	Name     string     `json:"name"`
	GoogleID string     `json:"google_id"`
	Auditable
}
