// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/artista/interface.go

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
func (m *MockRead) Get(email string) (*entity.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", email)
	ret0, _ := ret[0].(*entity.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockReadMockRecorder) Get(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRead)(nil).Get), email)
}

// Search mocks base method
func (m *MockRead) Search(query string) ([]*entity.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*entity.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockReadMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRead)(nil).Search), query)
}

// List mocks base method
func (m *MockRead) List() ([]*entity.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*entity.Artista)
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
func (m *MockWrite) Create(e *entity.Artista) (string, error) {
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
func (m *MockWrite) Update(e *entity.Artista) error {
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
func (m *MockRepository) Get(email string) (*entity.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", email)
	ret0, _ := ret[0].(*entity.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), email)
}

// Search mocks base method
func (m *MockRepository) Search(query string) ([]*entity.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*entity.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockRepositoryMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), query)
}

// List mocks base method
func (m *MockRepository) List() ([]*entity.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*entity.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List))
}

// Create mocks base method
func (m *MockRepository) Create(e *entity.Artista) (string, error) {
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
func (m *MockRepository) Update(e *entity.Artista) error {
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

// GetArtista mocks base method
func (m *MockUseCase) GetArtista(email string) (*entity.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArtista", email)
	ret0, _ := ret[0].(*entity.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArtista indicates an expected call of GetArtista
func (mr *MockUseCaseMockRecorder) GetArtista(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArtista", reflect.TypeOf((*MockUseCase)(nil).GetArtista), email)
}

// SearchArtistas mocks base method
func (m *MockUseCase) SearchArtistas(query string) ([]*entity.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchArtistas", query)
	ret0, _ := ret[0].([]*entity.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchArtistas indicates an expected call of SearchArtistas
func (mr *MockUseCaseMockRecorder) SearchArtistas(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchArtistas", reflect.TypeOf((*MockUseCase)(nil).SearchArtistas), query)
}

// ListArtistas mocks base method
func (m *MockUseCase) ListArtistas() ([]*entity.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListArtistas")
	ret0, _ := ret[0].([]*entity.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListArtistas indicates an expected call of ListArtistas
func (mr *MockUseCaseMockRecorder) ListArtistas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListArtistas", reflect.TypeOf((*MockUseCase)(nil).ListArtistas))
}

// CreateArtista mocks base method
func (m *MockUseCase) CreateArtista(email, password, birthday, nomeartistico, biografia string, anoformacao int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArtista", email, password, birthday, nomeartistico, biografia, anoformacao)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArtista indicates an expected call of CreateArtista
func (mr *MockUseCaseMockRecorder) CreateArtista(email, password, birthday, nomeartistico, biografia, anoformacao interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArtista", reflect.TypeOf((*MockUseCase)(nil).CreateArtista), email, password, birthday, nomeartistico, biografia, anoformacao)
}

// UpdateArtista mocks base method
func (m *MockUseCase) UpdateArtista(e *entity.Artista) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateArtista", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateArtista indicates an expected call of UpdateArtista
func (mr *MockUseCaseMockRecorder) UpdateArtista(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArtista", reflect.TypeOf((*MockUseCase)(nil).UpdateArtista), e)
}

// DeleteArtista mocks base method
func (m *MockUseCase) DeleteArtista(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArtista", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteArtista indicates an expected call of DeleteArtista
func (mr *MockUseCaseMockRecorder) DeleteArtista(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArtista", reflect.TypeOf((*MockUseCase)(nil).DeleteArtista), email)
}
