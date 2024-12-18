// Code generated by MockGen. DO NOT EDIT.
// Source: microservices/playlist/repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddFavoritePlaylist mocks base method.
func (m *MockRepository) AddFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFavoritePlaylist", ctx, userID, playlistID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFavoritePlaylist indicates an expected call of AddFavoritePlaylist.
func (mr *MockRepositoryMockRecorder) AddFavoritePlaylist(ctx, userID, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFavoritePlaylist", reflect.TypeOf((*MockRepository)(nil).AddFavoritePlaylist), ctx, userID, playlistID)
}

// AddToPlaylist mocks base method.
func (m *MockRepository) AddToPlaylist(ctx context.Context, playlistID, trackOrder, trackID uint64) (*models.PlaylistTrack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToPlaylist", ctx, playlistID, trackOrder, trackID)
	ret0, _ := ret[0].(*models.PlaylistTrack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddToPlaylist indicates an expected call of AddToPlaylist.
func (mr *MockRepositoryMockRecorder) AddToPlaylist(ctx, playlistID, trackOrder, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToPlaylist", reflect.TypeOf((*MockRepository)(nil).AddToPlaylist), ctx, playlistID, trackOrder, trackID)
}

// CreatePlaylist mocks base method.
func (m *MockRepository) CreatePlaylist(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlaylist", ctx, playlist)
	ret0, _ := ret[0].(*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlaylist indicates an expected call of CreatePlaylist.
func (mr *MockRepositoryMockRecorder) CreatePlaylist(ctx, playlist interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlaylist", reflect.TypeOf((*MockRepository)(nil).CreatePlaylist), ctx, playlist)
}

// DeleteFavoritePlaylist mocks base method.
func (m *MockRepository) DeleteFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFavoritePlaylist", ctx, userID, playlistID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFavoritePlaylist indicates an expected call of DeleteFavoritePlaylist.
func (mr *MockRepositoryMockRecorder) DeleteFavoritePlaylist(ctx, userID, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFavoritePlaylist", reflect.TypeOf((*MockRepository)(nil).DeleteFavoritePlaylist), ctx, userID, playlistID)
}

// DeletePlaylist mocks base method.
func (m *MockRepository) DeletePlaylist(ctx context.Context, playlistID uint64) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlaylist", ctx, playlistID)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePlaylist indicates an expected call of DeletePlaylist.
func (mr *MockRepositoryMockRecorder) DeletePlaylist(ctx, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlaylist", reflect.TypeOf((*MockRepository)(nil).DeletePlaylist), ctx, playlistID)
}

// GetAllPlaylists mocks base method.
func (m *MockRepository) GetAllPlaylists(ctx context.Context) ([]*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPlaylists", ctx)
	ret0, _ := ret[0].([]*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPlaylists indicates an expected call of GetAllPlaylists.
func (mr *MockRepositoryMockRecorder) GetAllPlaylists(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPlaylists", reflect.TypeOf((*MockRepository)(nil).GetAllPlaylists), ctx)
}

// GetFavoritePlaylists mocks base method.
func (m *MockRepository) GetFavoritePlaylists(ctx context.Context, userID uuid.UUID) ([]*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoritePlaylists", ctx, userID)
	ret0, _ := ret[0].([]*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoritePlaylists indicates an expected call of GetFavoritePlaylists.
func (mr *MockRepositoryMockRecorder) GetFavoritePlaylists(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoritePlaylists", reflect.TypeOf((*MockRepository)(nil).GetFavoritePlaylists), ctx, userID)
}

// GetLengthPlaylist mocks base method.
func (m *MockRepository) GetLengthPlaylist(ctx context.Context, playlistID uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLengthPlaylist", ctx, playlistID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLengthPlaylist indicates an expected call of GetLengthPlaylist.
func (mr *MockRepositoryMockRecorder) GetLengthPlaylist(ctx, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLengthPlaylist", reflect.TypeOf((*MockRepository)(nil).GetLengthPlaylist), ctx, playlistID)
}

// GetPlaylist mocks base method.
func (m *MockRepository) GetPlaylist(ctx context.Context, playlistID uint64) (*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaylist", ctx, playlistID)
	ret0, _ := ret[0].(*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaylist indicates an expected call of GetPlaylist.
func (mr *MockRepositoryMockRecorder) GetPlaylist(ctx, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaylist", reflect.TypeOf((*MockRepository)(nil).GetPlaylist), ctx, playlistID)
}

// GetPopularPlaylists mocks base method.
func (m *MockRepository) GetPopularPlaylists(ctx context.Context) ([]*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPopularPlaylists", ctx)
	ret0, _ := ret[0].([]*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPopularPlaylists indicates an expected call of GetPopularPlaylists.
func (mr *MockRepositoryMockRecorder) GetPopularPlaylists(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPopularPlaylists", reflect.TypeOf((*MockRepository)(nil).GetPopularPlaylists), ctx)
}

// GetUserPlaylists mocks base method.
func (m *MockRepository) GetUserPlaylists(ctx context.Context, userID uuid.UUID) ([]*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPlaylists", ctx, userID)
	ret0, _ := ret[0].([]*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPlaylists indicates an expected call of GetUserPlaylists.
func (mr *MockRepositoryMockRecorder) GetUserPlaylists(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPlaylists", reflect.TypeOf((*MockRepository)(nil).GetUserPlaylists), ctx, userID)
}

// IsFavoritePlaylist mocks base method.
func (m *MockRepository) IsFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFavoritePlaylist", ctx, userID, playlistID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFavoritePlaylist indicates an expected call of IsFavoritePlaylist.
func (mr *MockRepositoryMockRecorder) IsFavoritePlaylist(ctx, userID, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFavoritePlaylist", reflect.TypeOf((*MockRepository)(nil).IsFavoritePlaylist), ctx, userID, playlistID)
}

// RemoveFromPlaylist mocks base method.
func (m *MockRepository) RemoveFromPlaylist(ctx context.Context, playlistID, trackID uint64) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromPlaylist", ctx, playlistID, trackID)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveFromPlaylist indicates an expected call of RemoveFromPlaylist.
func (mr *MockRepositoryMockRecorder) RemoveFromPlaylist(ctx, playlistID, trackID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromPlaylist", reflect.TypeOf((*MockRepository)(nil).RemoveFromPlaylist), ctx, playlistID, trackID)
}
