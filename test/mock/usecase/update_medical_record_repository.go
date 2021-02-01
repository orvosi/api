// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/medical_record_updater.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/orvosi/api/entity"
)

// MockUpdateMedicalRecordRepository is a mock of UpdateMedicalRecordRepository interface
type MockUpdateMedicalRecordRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateMedicalRecordRepositoryMockRecorder
}

// MockUpdateMedicalRecordRepositoryMockRecorder is the mock recorder for MockUpdateMedicalRecordRepository
type MockUpdateMedicalRecordRepositoryMockRecorder struct {
	mock *MockUpdateMedicalRecordRepository
}

// NewMockUpdateMedicalRecordRepository creates a new mock instance
func NewMockUpdateMedicalRecordRepository(ctrl *gomock.Controller) *MockUpdateMedicalRecordRepository {
	mock := &MockUpdateMedicalRecordRepository{ctrl: ctrl}
	mock.recorder = &MockUpdateMedicalRecordRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpdateMedicalRecordRepository) EXPECT() *MockUpdateMedicalRecordRepositoryMockRecorder {
	return m.recorder
}

// DoesRecordExist mocks base method
func (m *MockUpdateMedicalRecordRepository) DoesRecordExist(ctx context.Context, id uint64, email string) (bool, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesRecordExist", ctx, id, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// DoesRecordExist indicates an expected call of DoesRecordExist
func (mr *MockUpdateMedicalRecordRepositoryMockRecorder) DoesRecordExist(ctx, id, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesRecordExist", reflect.TypeOf((*MockUpdateMedicalRecordRepository)(nil).DoesRecordExist), ctx, id, email)
}

// Update mocks base method
func (m *MockUpdateMedicalRecordRepository) Update(ctx context.Context, id uint64, record *entity.MedicalRecord) *entity.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, record)
	ret0, _ := ret[0].(*entity.Error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockUpdateMedicalRecordRepositoryMockRecorder) Update(ctx, id, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUpdateMedicalRecordRepository)(nil).Update), ctx, id, record)
}
