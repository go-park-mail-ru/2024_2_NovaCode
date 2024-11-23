// Code generated by MockGen. DO NOT EDIT.
// Source: internal/track/repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/track/repository.go -destination=internal/track/mock/repository_mock.go -package=mock
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

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
	isgomock struct{}
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

// AddFavoriteTrack mocks base method.
func (m *MockRepo) AddFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFavoriteTrack", ctx, userID, trackID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFavoriteTrack indicates an expected call of AddFavoriteTrack.
func (mr *MockRepoMockRecorder) AddFavoriteTrack(ctx, userID, trackID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFavoriteTrack", reflect.TypeOf((*MockRepo)(nil).AddFavoriteTrack), ctx, userID, trackID)
}

// Create mocks base method.
func (m *MockRepo) Create(ctx context.Context, track *models.Track) (*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, track)
	ret0, _ := ret[0].(*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepoMockRecorder) Create(ctx, track any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepo)(nil).Create), ctx, track)
}

// DeleteFavoriteTrack mocks base method.
func (m *MockRepo) DeleteFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFavoriteTrack", ctx, userID, trackID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFavoriteTrack indicates an expected call of DeleteFavoriteTrack.
func (mr *MockRepoMockRecorder) DeleteFavoriteTrack(ctx, userID, trackID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFavoriteTrack", reflect.TypeOf((*MockRepo)(nil).DeleteFavoriteTrack), ctx, userID, trackID)
}

// FindById mocks base method.
func (m *MockRepo) FindById(ctx context.Context, trackID uint64) (*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, trackID)
	ret0, _ := ret[0].(*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockRepoMockRecorder) FindById(ctx, trackID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockRepo)(nil).FindById), ctx, trackID)
}

// FindByQuery mocks base method.
func (m *MockRepo) FindByQuery(ctx context.Context, query string) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByQuery", ctx, query)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByQuery indicates an expected call of FindByQuery.
func (mr *MockRepoMockRecorder) FindByQuery(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByQuery", reflect.TypeOf((*MockRepo)(nil).FindByQuery), ctx, query)
}

// GetAll mocks base method.
func (m *MockRepo) GetAll(ctx context.Context) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepoMockRecorder) GetAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepo)(nil).GetAll), ctx)
}

// GetAllByAlbumID mocks base method.
func (m *MockRepo) GetAllByAlbumID(ctx context.Context, albumID uint64) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByAlbumID", ctx, albumID)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByAlbumID indicates an expected call of GetAllByAlbumID.
func (mr *MockRepoMockRecorder) GetAllByAlbumID(ctx, albumID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByAlbumID", reflect.TypeOf((*MockRepo)(nil).GetAllByAlbumID), ctx, albumID)
}

// GetAllByArtistID mocks base method.
func (m *MockRepo) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByArtistID", ctx, artistID)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByArtistID indicates an expected call of GetAllByArtistID.
func (mr *MockRepoMockRecorder) GetAllByArtistID(ctx, artistID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByArtistID", reflect.TypeOf((*MockRepo)(nil).GetAllByArtistID), ctx, artistID)
}

// GetFavoriteTracks mocks base method.
func (m *MockRepo) GetFavoriteTracks(ctx context.Context, userID uuid.UUID) ([]*models.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoriteTracks", ctx, userID)
	ret0, _ := ret[0].([]*models.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoriteTracks indicates an expected call of GetFavoriteTracks.
func (mr *MockRepoMockRecorder) GetFavoriteTracks(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoriteTracks", reflect.TypeOf((*MockRepo)(nil).GetFavoriteTracks), ctx, userID)
}

// IsFavoriteTrack mocks base method.
func (m *MockRepo) IsFavoriteTrack(ctx context.Context, userID uuid.UUID, trackID uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFavoriteTrack", ctx, userID, trackID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFavoriteTrack indicates an expected call of IsFavoriteTrack.
func (mr *MockRepoMockRecorder) IsFavoriteTrack(ctx, userID, trackID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFavoriteTrack", reflect.TypeOf((*MockRepo)(nil).IsFavoriteTrack), ctx, userID, trackID)
}
