package usecase_test

import (
	mock_usecase "github.com/orvosi/api/test/mock/usecase"
	"github.com/orvosi/api/usecase"
)

type MedicalRecordCreator_Executor struct {
	usecase  *usecase.MedicalRecordCreator
	inserter *mock_usecase.MockMedicalRecordInserter
}
