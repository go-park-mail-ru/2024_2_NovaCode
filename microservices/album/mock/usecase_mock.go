// Code generated by MockGen. DO NOT EDIT.
// Source: microservices/album/usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	dto "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/dto"
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

// AddFavoriteAlbum mocks base method.
func (m *MockUsecase) AddFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFavoriteAlbum", ctx, userID, albumID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFavoriteAlbum indicates an expected call of AddFavoriteAlbum.
func (mr *MockUsecaseMockRecorder) AddFavoriteAlbum(ctx, userID, albumID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFavoriteAlbum", reflect.TypeOf((*MockUsecase)(nil).AddFavoriteAlbum), ctx, userID, albumID)
}

// DeleteFavoriteAlbum mocks base method.
func (m *MockUsecase) DeleteFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFavoriteAlbum", ctx, userID, albumID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFavoriteAlbum indicates an expected call of DeleteFavoriteAlbum.
func (mr *MockUsecaseMockRecorder) DeleteFavoriteAlbum(ctx, userID, albumID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFavoriteAlbum", reflect.TypeOf((*MockUsecase)(nil).DeleteFavoriteAlbum), ctx, userID, albumID)
}

// GetAll mocks base method.
func (m *MockUsecase) GetAll(ctx context.Context) ([]*dto.AlbumDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*dto.AlbumDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockUsecaseMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUsecase)(nil).GetAll), ctx)
}

// GetAllByArtistID mocks base method.
func (m *MockUsecase) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*dto.AlbumDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByArtistID", ctx, artistID)
	ret0, _ := ret[0].([]*dto.AlbumDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByArtistID indicates an expected call of GetAllByArtistID.
func (mr *MockUsecaseMockRecorder) GetAllByArtistID(ctx, artistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByArtistID", reflect.TypeOf((*MockUsecase)(nil).GetAllByArtistID), ctx, artistID)
}

// GetFavoriteAlbums mocks base method.
func (m *MockUsecase) GetFavoriteAlbums(ctx context.Context, userID uuid.UUID) ([]*dto.AlbumDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoriteAlbums", ctx, userID)
	ret0, _ := ret[0].([]*dto.AlbumDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoriteAlbums indicates an expected call of GetFavoriteAlbums.
func (mr *MockUsecaseMockRecorder) GetFavoriteAlbums(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoriteAlbums", reflect.TypeOf((*MockUsecase)(nil).GetFavoriteAlbums), ctx, userID)
}

// IsFavoriteAlbum mocks base method.
func (m *MockUsecase) IsFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFavoriteAlbum", ctx, userID, albumID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFavoriteAlbum indicates an expected call of IsFavoriteAlbum.
func (mr *MockUsecaseMockRecorder) IsFavoriteAlbum(ctx, userID, albumID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFavoriteAlbum", reflect.TypeOf((*MockUsecase)(nil).IsFavoriteAlbum), ctx, userID, albumID)
}

// Search mocks base method.
func (m *MockUsecase) Search(ctx context.Context, name string) ([]*dto.AlbumDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, name)
	ret0, _ := ret[0].([]*dto.AlbumDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockUsecaseMockRecorder) Search(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockUsecase)(nil).Search), ctx, name)
}

// View mocks base method.
func (m *MockUsecase) View(ctx context.Context, albumID uint64) (*dto.AlbumDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "View", ctx, albumID)
	ret0, _ := ret[0].(*dto.AlbumDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// View indicates an expected call of View.
func (mr *MockUsecaseMockRecorder) View(ctx, albumID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "View", reflect.TypeOf((*MockUsecase)(nil).View), ctx, albumID)
}
