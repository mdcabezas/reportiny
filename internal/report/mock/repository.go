// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	report "github.com/mdcabezas/reportiny/internal/report"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIRepository) Create(repo report.RepositoryModel) (report.RepositoryModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", repo)
	ret0, _ := ret[0].(report.RepositoryModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIRepositoryMockRecorder) Create(repo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIRepository)(nil).Create), repo)
}

// Delete mocks base method.
func (m *MockIRepository) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIRepository)(nil).Delete), id)
}

// Read mocks base method.
func (m *MockIRepository) Read(id string) (report.RepositoryModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", id)
	ret0, _ := ret[0].(report.RepositoryModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockIRepositoryMockRecorder) Read(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockIRepository)(nil).Read), id)
}

// ReadAll mocks base method.
func (m *MockIRepository) ReadAll() []report.RepositoryModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll")
	ret0, _ := ret[0].([]report.RepositoryModel)
	return ret0
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockIRepositoryMockRecorder) ReadAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockIRepository)(nil).ReadAll))
}

// Update mocks base method.
func (m *MockIRepository) Update(id string, repo report.RepositoryModel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, repo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIRepositoryMockRecorder) Update(id, repo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRepository)(nil).Update), id, repo)
}
