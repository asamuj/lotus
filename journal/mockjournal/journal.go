// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/filecoin-project/lotus/journal (interfaces: Journal)

// Package mockjournal is a generated GoMock package.
package mockjournal

import (
	reflect "reflect"

	journal "github.com/filecoin-project/lotus/journal"
	gomock "github.com/golang/mock/gomock"
)

// MockJournal is a mock of Journal interface.
type MockJournal struct {
	ctrl     *gomock.Controller
	recorder *MockJournalMockRecorder
}

// MockJournalMockRecorder is the mock recorder for MockJournal.
type MockJournalMockRecorder struct {
	mock *MockJournal
}

// NewMockJournal creates a new mock instance.
func NewMockJournal(ctrl *gomock.Controller) *MockJournal {
	mock := &MockJournal{ctrl: ctrl}
	mock.recorder = &MockJournalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJournal) EXPECT() *MockJournalMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockJournal) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockJournalMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockJournal)(nil).Close))
}

// RecordEvent mocks base method.
func (m *MockJournal) RecordEvent(arg0 journal.EventType, arg1 func() interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RecordEvent", arg0, arg1)
}

// RecordEvent indicates an expected call of RecordEvent.
func (mr *MockJournalMockRecorder) RecordEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordEvent", reflect.TypeOf((*MockJournal)(nil).RecordEvent), arg0, arg1)
}

// RegisterEventType mocks base method.
func (m *MockJournal) RegisterEventType(arg0, arg1 string) journal.EventType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterEventType", arg0, arg1)
	ret0, _ := ret[0].(journal.EventType)
	return ret0
}

// RegisterEventType indicates an expected call of RegisterEventType.
func (mr *MockJournalMockRecorder) RegisterEventType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterEventType", reflect.TypeOf((*MockJournal)(nil).RegisterEventType), arg0, arg1)
}