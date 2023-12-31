// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package events is a generated GoMock package.
package events

import (
	db_actions "go_alert_bot/internal/db_actions"
	entities "go_alert_bot/internal/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEventRepo is a mock of EventRepo interface.
type MockEventRepo struct {
	ctrl     *gomock.Controller
	recorder *MockEventRepoMockRecorder
}

// MockEventRepoMockRecorder is the mock recorder for MockEventRepo.
type MockEventRepoMockRecorder struct {
	mock *MockEventRepo
}

// NewMockEventRepo creates a new mock instance.
func NewMockEventRepo(ctrl *gomock.Controller) *MockEventRepo {
	mock := &MockEventRepo{ctrl: ctrl}
	mock.recorder = &MockEventRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventRepo) EXPECT() *MockEventRepoMockRecorder {
	return m.recorder
}

// GetChannelFromChannelLink mocks base method.
func (m *MockEventRepo) GetChannelFromChannelLink(link entities.ChannelLink) *db_actions.ChannelDb {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannelFromChannelLink", link)
	ret0, _ := ret[0].(*db_actions.ChannelDb)
	return ret0
}

// GetChannelFromChannelLink indicates an expected call of GetChannelFromChannelLink.
func (mr *MockEventRepoMockRecorder) GetChannelFromChannelLink(link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannelFromChannelLink", reflect.TypeOf((*MockEventRepo)(nil).GetChannelFromChannelLink), link)
}

// GetStdoutChannelByChannelLink mocks base method.
func (m *MockEventRepo) GetStdoutChannelByChannelLink(channel *db_actions.ChannelDb) (*db_actions.ChannelDb, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStdoutChannelByChannelLink", channel)
	ret0, _ := ret[0].(*db_actions.ChannelDb)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStdoutChannelByChannelLink indicates an expected call of GetStdoutChannelByChannelLink.
func (mr *MockEventRepoMockRecorder) GetStdoutChannelByChannelLink(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStdoutChannelByChannelLink", reflect.TypeOf((*MockEventRepo)(nil).GetStdoutChannelByChannelLink), channel)
}

// GetTelegramChannelByChannelLink mocks base method.
func (m *MockEventRepo) GetTelegramChannelByChannelLink(channel *db_actions.ChannelDb) (*db_actions.ChannelDb, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTelegramChannelByChannelLink", channel)
	ret0, _ := ret[0].(*db_actions.ChannelDb)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTelegramChannelByChannelLink indicates an expected call of GetTelegramChannelByChannelLink.
func (mr *MockEventRepoMockRecorder) GetTelegramChannelByChannelLink(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTelegramChannelByChannelLink", reflect.TypeOf((*MockEventRepo)(nil).GetTelegramChannelByChannelLink), channel)
}

// IsExistChannelByChannelLink mocks base method.
func (m *MockEventRepo) IsExistChannelByChannelLink(link db_actions.ChannelLink) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExistChannelByChannelLink", link)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsExistChannelByChannelLink indicates an expected call of IsExistChannelByChannelLink.
func (mr *MockEventRepoMockRecorder) IsExistChannelByChannelLink(link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExistChannelByChannelLink", reflect.TypeOf((*MockEventRepo)(nil).IsExistChannelByChannelLink), link)
}

// MockSendEventRepo is a mock of SendEventRepo interface.
type MockSendEventRepo struct {
	ctrl     *gomock.Controller
	recorder *MockSendEventRepoMockRecorder
}

// MockSendEventRepoMockRecorder is the mock recorder for MockSendEventRepo.
type MockSendEventRepoMockRecorder struct {
	mock *MockSendEventRepo
}

// NewMockSendEventRepo creates a new mock instance.
func NewMockSendEventRepo(ctrl *gomock.Controller) *MockSendEventRepo {
	mock := &MockSendEventRepo{ctrl: ctrl}
	mock.recorder = &MockSendEventRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSendEventRepo) EXPECT() *MockSendEventRepoMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockSendEventRepo) Send(event Event, channel *db_actions.ChannelDb, counter int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", event, channel, counter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockSendEventRepoMockRecorder) Send(event, channel, counter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSendEventRepo)(nil).Send), event, channel, counter)
}
