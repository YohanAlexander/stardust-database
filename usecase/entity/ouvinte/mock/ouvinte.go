// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/entity/ouvinte/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/yohanalexander/deezefy-music/entity"
	reflect "reflect"
)

// MockRead is a mock of Read interface
type MockRead struct {
	ctrl     *gomock.Controller
	recorder *MockReadMockRecorder
}

// MockReadMockRecorder is the mock recorder for MockRead
type MockReadMockRecorder struct {
	mock *MockRead
}

// NewMockRead creates a new mock instance
func NewMockRead(ctrl *gomock.Controller) *MockRead {
	mock := &MockRead{ctrl: ctrl}
	mock.recorder = &MockReadMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRead) EXPECT() *MockReadMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockRead) Get(email string) (*entity.Ouvinte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", email)
	ret0, _ := ret[0].(*entity.Ouvinte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockReadMockRecorder) Get(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRead)(nil).Get), email)
}

// Search mocks base method
func (m *MockRead) Search(query string) ([]*entity.Ouvinte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*entity.Ouvinte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockReadMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRead)(nil).Search), query)
}

// List mocks base method
func (m *MockRead) List() ([]*entity.Ouvinte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*entity.Ouvinte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockReadMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRead)(nil).List))
}

// MockWrite is a mock of Write interface
type MockWrite struct {
	ctrl     *gomock.Controller
	recorder *MockWriteMockRecorder
}

// MockWriteMockRecorder is the mock recorder for MockWrite
type MockWriteMockRecorder struct {
	mock *MockWrite
}

// NewMockWrite creates a new mock instance
func NewMockWrite(ctrl *gomock.Controller) *MockWrite {
	mock := &MockWrite{ctrl: ctrl}
	mock.recorder = &MockWriteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWrite) EXPECT() *MockWriteMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockWrite) Create(e *entity.Ouvinte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockWriteMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWrite)(nil).Create), e)
}

// Update mocks base method
func (m *MockWrite) Update(e *entity.Ouvinte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockWriteMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWrite)(nil).Update), e)
}

// Delete mocks base method
func (m *MockWrite) Delete(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockWriteMockRecorder) Delete(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWrite)(nil).Delete), email)
}

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockRepository) Get(email string) (*entity.Ouvinte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", email)
	ret0, _ := ret[0].(*entity.Ouvinte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), email)
}

// Search mocks base method
func (m *MockRepository) Search(query string) ([]*entity.Ouvinte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*entity.Ouvinte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockRepositoryMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), query)
}

// List mocks base method
func (m *MockRepository) List() ([]*entity.Ouvinte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*entity.Ouvinte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List))
}

// Create mocks base method
func (m *MockRepository) Create(e *entity.Ouvinte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockRepositoryMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), e)
}

// Update mocks base method
func (m *MockRepository) Update(e *entity.Ouvinte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRepositoryMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), e)
}

// Delete mocks base method
func (m *MockRepository) Delete(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRepositoryMockRecorder) Delete(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), email)
}

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// GetOuvinte mocks base method
func (m *MockUseCase) GetOuvinte(email string) (*entity.Ouvinte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOuvinte", email)
	ret0, _ := ret[0].(*entity.Ouvinte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOuvinte indicates an expected call of GetOuvinte
func (mr *MockUseCaseMockRecorder) GetOuvinte(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOuvinte", reflect.TypeOf((*MockUseCase)(nil).GetOuvinte), email)
}

// SearchOuvintes mocks base method
func (m *MockUseCase) SearchOuvintes(query string) ([]*entity.Ouvinte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchOuvintes", query)
	ret0, _ := ret[0].([]*entity.Ouvinte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchOuvintes indicates an expected call of SearchOuvintes
func (mr *MockUseCaseMockRecorder) SearchOuvintes(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchOuvintes", reflect.TypeOf((*MockUseCase)(nil).SearchOuvintes), query)
}

// ListOuvintes mocks base method
func (m *MockUseCase) ListOuvintes() ([]*entity.Ouvinte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOuvintes")
	ret0, _ := ret[0].([]*entity.Ouvinte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOuvintes indicates an expected call of ListOuvintes
func (mr *MockUseCaseMockRecorder) ListOuvintes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOuvintes", reflect.TypeOf((*MockUseCase)(nil).ListOuvintes))
}

// CreateOuvinte mocks base method
func (m *MockUseCase) CreateOuvinte(email, password, birthday, primeironome, sobrenome string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOuvinte", email, password, birthday, primeironome, sobrenome)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOuvinte indicates an expected call of CreateOuvinte
func (mr *MockUseCaseMockRecorder) CreateOuvinte(email, password, birthday, primeironome, sobrenome interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOuvinte", reflect.TypeOf((*MockUseCase)(nil).CreateOuvinte), email, password, birthday, primeironome, sobrenome)
}

// UpdateOuvinte mocks base method
func (m *MockUseCase) UpdateOuvinte(e *entity.Ouvinte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOuvinte", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOuvinte indicates an expected call of UpdateOuvinte
func (mr *MockUseCaseMockRecorder) UpdateOuvinte(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOuvinte", reflect.TypeOf((*MockUseCase)(nil).UpdateOuvinte), e)
}

// DeleteOuvinte mocks base method
func (m *MockUseCase) DeleteOuvinte(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOuvinte", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOuvinte indicates an expected call of DeleteOuvinte
func (mr *MockUseCaseMockRecorder) DeleteOuvinte(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOuvinte", reflect.TypeOf((*MockUseCase)(nil).DeleteOuvinte), email)
}
