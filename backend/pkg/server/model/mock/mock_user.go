// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/server/model/user.go
//
// Generated by this command:
//
//	mockgen -source=./pkg/server/model/user.go -destination ./pkg/server/model/mock/mock_user.go
//

// Package mock_model is a generated GoMock package.
package mock_model

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepositoryInterface is a mock of UserRepositoryInterface interface.
type MockUserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryInterfaceMockRecorder
}

// MockUserRepositoryInterfaceMockRecorder is the mock recorder for MockUserRepositoryInterface.
type MockUserRepositoryInterfaceMockRecorder struct {
	mock *MockUserRepositoryInterface
}

// NewMockUserRepositoryInterface creates a new mock instance.
func NewMockUserRepositoryInterface(ctrl *gomock.Controller) *MockUserRepositoryInterface {
	mock := &MockUserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryInterface) EXPECT() *MockUserRepositoryInterfaceMockRecorder {
	return m.recorder
}

// InsertUser mocks base method.
func (m *MockUserRepositoryInterface) InsertUser(uid string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", uid)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) InsertUser(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).InsertUser), uid)
}

// SelectUserIDByUID mocks base method.
func (m *MockUserRepositoryInterface) SelectUserIDByUID(uid string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserIDByUID", uid)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserIDByUID indicates an expected call of SelectUserIDByUID.
func (mr *MockUserRepositoryInterfaceMockRecorder) SelectUserIDByUID(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserIDByUID", reflect.TypeOf((*MockUserRepositoryInterface)(nil).SelectUserIDByUID), uid)
}

// SelectUserIDByUIDWithError mocks base method.
func (m *MockUserRepositoryInterface) SelectUserIDByUIDWithError(uid string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserIDByUIDWithError", uid)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserIDByUIDWithError indicates an expected call of SelectUserIDByUIDWithError.
func (mr *MockUserRepositoryInterfaceMockRecorder) SelectUserIDByUIDWithError(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserIDByUIDWithError", reflect.TypeOf((*MockUserRepositoryInterface)(nil).SelectUserIDByUIDWithError), uid)
}
