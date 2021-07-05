// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GG_Backend_tech_challenge/src/repository (interfaces: TableRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/GG_Backend_tech_challenge/src/model"
	gomock "github.com/golang/mock/gomock"
)

// MockTableRepository is a mock of TableRepository interface.
type MockTableRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTableRepositoryMockRecorder
}

// MockTableRepositoryMockRecorder is the mock recorder for MockTableRepository.
type MockTableRepositoryMockRecorder struct {
	mock *MockTableRepository
}

// NewMockTableRepository creates a new mock instance.
func NewMockTableRepository(ctrl *gomock.Controller) *MockTableRepository {
	mock := &MockTableRepository{ctrl: ctrl}
	mock.recorder = &MockTableRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTableRepository) EXPECT() *MockTableRepositoryMockRecorder {
	return m.recorder
}

// AddTable mocks base method.
func (m *MockTableRepository) AddTable(arg0, arg1 int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTable", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AddTable indicates an expected call of AddTable.
func (mr *MockTableRepositoryMockRecorder) AddTable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTable", reflect.TypeOf((*MockTableRepository)(nil).AddTable), arg0, arg1)
}

// AssignTableToGuest mocks base method.
func (m *MockTableRepository) AssignTableToGuest(arg0 model.Table, arg1 *model.Guest) (bool, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignTableToGuest", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// AssignTableToGuest indicates an expected call of AssignTableToGuest.
func (mr *MockTableRepositoryMockRecorder) AssignTableToGuest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignTableToGuest", reflect.TypeOf((*MockTableRepository)(nil).AssignTableToGuest), arg0, arg1)
}

// DecreaseGuestSeats mocks base method.
func (m *MockTableRepository) DecreaseGuestSeats(arg0 model.Table, arg1 int) (bool, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecreaseGuestSeats", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// DecreaseGuestSeats indicates an expected call of DecreaseGuestSeats.
func (mr *MockTableRepositoryMockRecorder) DecreaseGuestSeats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecreaseGuestSeats", reflect.TypeOf((*MockTableRepository)(nil).DecreaseGuestSeats), arg0, arg1)
}

// DoesTableExist mocks base method.
func (m *MockTableRepository) DoesTableExist(arg0 int) (bool, model.Table) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesTableExist", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(model.Table)
	return ret0, ret1
}

// DoesTableExist indicates an expected call of DoesTableExist.
func (mr *MockTableRepositoryMockRecorder) DoesTableExist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesTableExist", reflect.TypeOf((*MockTableRepository)(nil).DoesTableExist), arg0)
}

// GetEmptySeats mocks base method.
func (m *MockTableRepository) GetEmptySeats() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmptySeats")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetEmptySeats indicates an expected call of GetEmptySeats.
func (mr *MockTableRepositoryMockRecorder) GetEmptySeats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmptySeats", reflect.TypeOf((*MockTableRepository)(nil).GetEmptySeats))
}

// IncreaseGuestSeats mocks base method.
func (m *MockTableRepository) IncreaseGuestSeats(arg0 model.Table, arg1 int) (bool, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseGuestSeats", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// IncreaseGuestSeats indicates an expected call of IncreaseGuestSeats.
func (mr *MockTableRepositoryMockRecorder) IncreaseGuestSeats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseGuestSeats", reflect.TypeOf((*MockTableRepository)(nil).IncreaseGuestSeats), arg0, arg1)
}

// RemoveGuestFromTable mocks base method.
func (m *MockTableRepository) RemoveGuestFromTable(arg0 model.Guest, arg1 model.Table) (bool, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveGuestFromTable", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// RemoveGuestFromTable indicates an expected call of RemoveGuestFromTable.
func (mr *MockTableRepositoryMockRecorder) RemoveGuestFromTable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveGuestFromTable", reflect.TypeOf((*MockTableRepository)(nil).RemoveGuestFromTable), arg0, arg1)
}