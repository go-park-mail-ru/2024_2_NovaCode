// Code generated by MockGen. DO NOT EDIT.
// Source: internal/genre/repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepo) Create(ctx context.Context, genre *models.Genre) (*models.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, genre)
	ret0, _ := ret[0].(*models.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepoMockRecorder) Create(ctx, genre interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepo)(nil).Create), ctx, genre)
}

// GetAll mocks base method.
func (m *MockRepo) GetAll(ctx context.Context) ([]*models.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*models.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepoMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepo)(nil).GetAll), ctx)
}

// GetAllByArtistID mocks base method.
func (m *MockRepo) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByArtistID", ctx, artistID)
	ret0, _ := ret[0].([]*models.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByArtistID indicates an expected call of GetAllByArtistID.
func (mr *MockRepoMockRecorder) GetAllByArtistID(ctx, artistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByArtistID", reflect.TypeOf((*MockRepo)(nil).GetAllByArtistID), ctx, artistID)
}

// GetAllByTrackID mocks base method.
func (m *MockRepo) GetAllByTrackID(ctx context.Context, trackID uint64) ([]*models.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByTrackID", ctx, trackID)
	ret0, _ := ret[0].([]*models.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByTrackID indicates an expected call of GetAllByTrackID.
func (mr *MockRepoMockRecorder) GetAllByTrackID(ctx, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByTrackID", reflect.TypeOf((*MockRepo)(nil).GetAllByTrackID), ctx, trackID)
}
