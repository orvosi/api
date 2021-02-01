package handler

import "github.com/orvosi/api/usecase"

// UpdateMedicalRecordRequest represents medical record request.
type UpdateMedicalRecordRequest struct {
	Symptom   string `json:"symptom"`
	Diagnosis string `json:"diagnosis"`
	Therapy   string `json:"therapy"`
	Result    string `json:"result"`
}

// MedicalRecordUpdater handles HTTP request and response
// for update medical record.
type MedicalRecordUpdater struct {
	updater usecase.UpdateMedicalRecord
}

// NewMedicalRecordUpdater updates an instance of MedicalUpdater.
func NewMedicalRecordUpdater(updater usecase.UpdateMedicalRecord) *MedicalRecordUpdater {
	return &MedicalRecordUpdater{
		updater: updater,
	}
}
