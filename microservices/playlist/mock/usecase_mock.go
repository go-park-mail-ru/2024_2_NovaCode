// Code generated by MockGen. DO NOT EDIT.
// Source: microservices/playlist/usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	dto "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/playlist/dto"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// AddFavoritePlaylist mocks base method.
func (m *MockUsecase) AddFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFavoritePlaylist", ctx, userID, playlistID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFavoritePlaylist indicates an expected call of AddFavoritePlaylist.
func (mr *MockUsecaseMockRecorder) AddFavoritePlaylist(ctx, userID, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFavoritePlaylist", reflect.TypeOf((*MockUsecase)(nil).AddFavoritePlaylist), ctx, userID, playlistID)
}

// AddToPlaylist mocks base method.
func (m *MockUsecase) AddToPlaylist(ctx context.Context, playlistTrackDTO *dto.PlaylistTrackDTO) (*models.PlaylistTrack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToPlaylist", ctx, playlistTrackDTO)
	ret0, _ := ret[0].(*models.PlaylistTrack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddToPlaylist indicates an expected call of AddToPlaylist.
func (mr *MockUsecaseMockRecorder) AddToPlaylist(ctx, playlistTrackDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToPlaylist", reflect.TypeOf((*MockUsecase)(nil).AddToPlaylist), ctx, playlistTrackDTO)
}

// CreatePlaylist mocks base method.
func (m *MockUsecase) CreatePlaylist(ctx context.Context, newPlaylistDTO *dto.PlaylistDTO) (*dto.PlaylistDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlaylist", ctx, newPlaylistDTO)
	ret0, _ := ret[0].(*dto.PlaylistDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlaylist indicates an expected call of CreatePlaylist.
func (mr *MockUsecaseMockRecorder) CreatePlaylist(ctx, newPlaylistDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlaylist", reflect.TypeOf((*MockUsecase)(nil).CreatePlaylist), ctx, newPlaylistDTO)
}

// DeleteFavoritePlaylist mocks base method.
func (m *MockUsecase) DeleteFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFavoritePlaylist", ctx, userID, playlistID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFavoritePlaylist indicates an expected call of DeleteFavoritePlaylist.
func (mr *MockUsecaseMockRecorder) DeleteFavoritePlaylist(ctx, userID, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFavoritePlaylist", reflect.TypeOf((*MockUsecase)(nil).DeleteFavoritePlaylist), ctx, userID, playlistID)
}

// DeletePlaylist mocks base method.
func (m *MockUsecase) DeletePlaylist(ctx context.Context, playlistID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlaylist", ctx, playlistID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePlaylist indicates an expected call of DeletePlaylist.
func (mr *MockUsecaseMockRecorder) DeletePlaylist(ctx, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlaylist", reflect.TypeOf((*MockUsecase)(nil).DeletePlaylist), ctx, playlistID)
}

// GetAllPlaylists mocks base method.
func (m *MockUsecase) GetAllPlaylists(ctx context.Context) ([]*dto.PlaylistDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPlaylists", ctx)
	ret0, _ := ret[0].([]*dto.PlaylistDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPlaylists indicates an expected call of GetAllPlaylists.
func (mr *MockUsecaseMockRecorder) GetAllPlaylists(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPlaylists", reflect.TypeOf((*MockUsecase)(nil).GetAllPlaylists), ctx)
}

// GetFavoritePlaylists mocks base method.
func (m *MockUsecase) GetFavoritePlaylists(ctx context.Context, userID uuid.UUID) ([]*dto.PlaylistDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoritePlaylists", ctx, userID)
	ret0, _ := ret[0].([]*dto.PlaylistDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoritePlaylists indicates an expected call of GetFavoritePlaylists.
func (mr *MockUsecaseMockRecorder) GetFavoritePlaylists(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoritePlaylists", reflect.TypeOf((*MockUsecase)(nil).GetFavoritePlaylists), ctx, userID)
}

// GetPlaylist mocks base method.
func (m *MockUsecase) GetPlaylist(ctx context.Context, playlistID uint64) (*dto.PlaylistDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaylist", ctx, playlistID)
	ret0, _ := ret[0].(*dto.PlaylistDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaylist indicates an expected call of GetPlaylist.
func (mr *MockUsecaseMockRecorder) GetPlaylist(ctx, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaylist", reflect.TypeOf((*MockUsecase)(nil).GetPlaylist), ctx, playlistID)
}

// GetPopularPlaylists mocks base method.
func (m *MockUsecase) GetPopularPlaylists(ctx context.Context) ([]*dto.PlaylistDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPopularPlaylists", ctx)
	ret0, _ := ret[0].([]*dto.PlaylistDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPopularPlaylists indicates an expected call of GetPopularPlaylists.
func (mr *MockUsecaseMockRecorder) GetPopularPlaylists(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPopularPlaylists", reflect.TypeOf((*MockUsecase)(nil).GetPopularPlaylists), ctx)
}

// GetUserPlaylists mocks base method.
func (m *MockUsecase) GetUserPlaylists(ctx context.Context, userID uuid.UUID) ([]*dto.PlaylistDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPlaylists", ctx, userID)
	ret0, _ := ret[0].([]*dto.PlaylistDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPlaylists indicates an expected call of GetUserPlaylists.
func (mr *MockUsecaseMockRecorder) GetUserPlaylists(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPlaylists", reflect.TypeOf((*MockUsecase)(nil).GetUserPlaylists), ctx, userID)
}

// IsFavoritePlaylist mocks base method.
func (m *MockUsecase) IsFavoritePlaylist(ctx context.Context, userID uuid.UUID, playlistID uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFavoritePlaylist", ctx, userID, playlistID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFavoritePlaylist indicates an expected call of IsFavoritePlaylist.
func (mr *MockUsecaseMockRecorder) IsFavoritePlaylist(ctx, userID, playlistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFavoritePlaylist", reflect.TypeOf((*MockUsecase)(nil).IsFavoritePlaylist), ctx, userID, playlistID)
}

// RemoveFromPlaylist mocks base method.
func (m *MockUsecase) RemoveFromPlaylist(ctx context.Context, playlistTrackDTO *dto.PlaylistTrackDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromPlaylist", ctx, playlistTrackDTO)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromPlaylist indicates an expected call of RemoveFromPlaylist.
func (mr *MockUsecaseMockRecorder) RemoveFromPlaylist(ctx, playlistTrackDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromPlaylist", reflect.TypeOf((*MockUsecase)(nil).RemoveFromPlaylist), ctx, playlistTrackDTO)
}
