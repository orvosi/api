package entity

import "github.com/indrasaputra/hashids"

// MedicalRecord holds the user's medical record data
type MedicalRecord struct {
	ID        hashids.ID `json:"id"`
	Symptom   string     `json:"symptom"`
	Diagnosis string     `json:"diagnosis"`
	Therapy   string     `json:"therapy"`
	Result    string     `json:"result"`
}

// User holds user's information.
type User struct {
	ID    hashids.ID `json:"id"`
	Email string     `json:"email"`
	Name  string     `json:"name"`
}
