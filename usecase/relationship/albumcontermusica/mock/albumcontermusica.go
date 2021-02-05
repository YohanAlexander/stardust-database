// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/relationship/albumcontermusica/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/yohanalexander/deezefy-music/entity"
	reflect "reflect"
)

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

// Conter mocks base method
func (m_2 *MockUseCase) Conter(a *entity.Album, m *entity.Musica) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Conter", a, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Conter indicates an expected call of Conter
func (mr *MockUseCaseMockRecorder) Conter(a, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Conter", reflect.TypeOf((*MockUseCase)(nil).Conter), a, m)
}

// Desconter mocks base method
func (m_2 *MockUseCase) Desconter(a *entity.Album, m *entity.Musica) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Desconter", a, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Desconter indicates an expected call of Desconter
func (mr *MockUseCaseMockRecorder) Desconter(a, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Desconter", reflect.TypeOf((*MockUseCase)(nil).Desconter), a, m)
}
