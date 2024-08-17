// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jairogloz/go-l/pkg/ports (interfaces: LeagueRepository)
//
// Generated by this command:
//
//	mockgen -destination=mocks/mock_league_repository.go -package=mocks github.com/jairogloz/go-l/pkg/ports LeagueRepository
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/jairogloz/go-l/pkg/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockLeagueRepository is a mock of LeagueRepository interface.
type MockLeagueRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLeagueRepositoryMockRecorder
}

// MockLeagueRepositoryMockRecorder is the mock recorder for MockLeagueRepository.
type MockLeagueRepositoryMockRecorder struct {
	mock *MockLeagueRepository
}

// NewMockLeagueRepository creates a new mock instance.
func NewMockLeagueRepository(ctrl *gomock.Controller) *MockLeagueRepository {
	mock := &MockLeagueRepository{ctrl: ctrl}
	mock.recorder = &MockLeagueRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeagueRepository) EXPECT() *MockLeagueRepositoryMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockLeagueRepository) Insert(arg0 context.Context, arg1 *domain.League) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockLeagueRepositoryMockRecorder) Insert(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockLeagueRepository)(nil).Insert), arg0, arg1)
}
