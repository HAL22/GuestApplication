// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GG_Backend_tech_challenge/src/repository (interfaces: GuestRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/GG_Backend_tech_challenge/src/model"
	gomock "github.com/golang/mock/gomock"
)

// MockGuestRepository is a mock of GuestRepository interface.
type MockGuestRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGuestRepositoryMockRecorder
}

// MockGuestRepositoryMockRecorder is the mock recorder for MockGuestRepository.
type MockGuestRepositoryMockRecorder struct {
	mock *MockGuestRepository
}

// NewMockGuestRepository creates a new mock instance.
func NewMockGuestRepository(ctrl *gomock.Controller) *MockGuestRepository {
	mock := &MockGuestRepository{ctrl: ctrl}
	mock.recorder = &MockGuestRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGuestRepository) EXPECT() *MockGuestRepositoryMockRecorder {
	return m.recorder
}

// AddGuest mocks base method.
func (m *MockGuestRepository) AddGuest(arg0 model.Guest) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddGuest", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AddGuest indicates an expected call of AddGuest.
func (mr *MockGuestRepositoryMockRecorder) AddGuest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGuest", reflect.TypeOf((*MockGuestRepository)(nil).AddGuest), arg0)
}

// GetGuestByName mocks base method.
func (m *MockGuestRepository) GetGuestByName(arg0 string) (bool, model.Guest) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGuestByName", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(model.Guest)
	return ret0, ret1
}

// GetGuestByName indicates an expected call of GetGuestByName.
func (mr *MockGuestRepositoryMockRecorder) GetGuestByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGuestByName", reflect.TypeOf((*MockGuestRepository)(nil).GetGuestByName), arg0)
}

// GetGuests mocks base method.
func (m *MockGuestRepository) GetGuests() (bool, []model.Guest) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGuests")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]model.Guest)
	return ret0, ret1
}

// GetGuests indicates an expected call of GetGuests.
func (mr *MockGuestRepositoryMockRecorder) GetGuests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGuests", reflect.TypeOf((*MockGuestRepository)(nil).GetGuests))
}
