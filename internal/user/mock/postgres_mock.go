// Code generated by MockGen. DO NOT EDIT.
// Source: internal/user/repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/user/repository.go -destination=internal/user/mock/postgres_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockPostgresRepo is a mock of PostgresRepo interface.
type MockPostgresRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPostgresRepoMockRecorder
	isgomock struct{}
}

// MockPostgresRepoMockRecorder is the mock recorder for MockPostgresRepo.
type MockPostgresRepoMockRecorder struct {
	mock *MockPostgresRepo
}

// NewMockPostgresRepo creates a new mock instance.
func NewMockPostgresRepo(ctrl *gomock.Controller) *MockPostgresRepo {
	mock := &MockPostgresRepo{ctrl: ctrl}
	mock.recorder = &MockPostgresRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostgresRepo) EXPECT() *MockPostgresRepoMockRecorder {
	return m.recorder
}

// FindByEmail mocks base method.
func (m *MockPostgresRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", ctx, email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockPostgresRepoMockRecorder) FindByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockPostgresRepo)(nil).FindByEmail), ctx, email)
}

// FindByID mocks base method.
func (m *MockPostgresRepo) FindByID(ctx context.Context, uuid uuid.UUID) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, uuid)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockPostgresRepoMockRecorder) FindByID(ctx, uuid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockPostgresRepo)(nil).FindByID), ctx, uuid)
}

// FindByUsername mocks base method.
func (m *MockPostgresRepo) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUsername", ctx, username)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUsername indicates an expected call of FindByUsername.
func (mr *MockPostgresRepoMockRecorder) FindByUsername(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUsername", reflect.TypeOf((*MockPostgresRepo)(nil).FindByUsername), ctx, username)
}

// Insert mocks base method.
func (m *MockPostgresRepo) Insert(ctx context.Context, user *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockPostgresRepoMockRecorder) Insert(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPostgresRepo)(nil).Insert), ctx, user)
}

// Update mocks base method.
func (m *MockPostgresRepo) Update(ctx context.Context, user *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPostgresRepoMockRecorder) Update(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostgresRepo)(nil).Update), ctx, user)
}
