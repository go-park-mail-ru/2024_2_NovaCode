// Code generated by MockGen. DO NOT EDIT.
// Source: microservices/artist/usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	dto "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/dto"
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

// AddFavoriteArtist mocks base method.
func (m *MockUsecase) AddFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFavoriteArtist", ctx, userID, artistID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFavoriteArtist indicates an expected call of AddFavoriteArtist.
func (mr *MockUsecaseMockRecorder) AddFavoriteArtist(ctx, userID, artistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFavoriteArtist", reflect.TypeOf((*MockUsecase)(nil).AddFavoriteArtist), ctx, userID, artistID)
}

// DeleteFavoriteArtist mocks base method.
func (m *MockUsecase) DeleteFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFavoriteArtist", ctx, userID, artistID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFavoriteArtist indicates an expected call of DeleteFavoriteArtist.
func (mr *MockUsecaseMockRecorder) DeleteFavoriteArtist(ctx, userID, artistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFavoriteArtist", reflect.TypeOf((*MockUsecase)(nil).DeleteFavoriteArtist), ctx, userID, artistID)
}

// GetAll mocks base method.
func (m *MockUsecase) GetAll(ctx context.Context) ([]*dto.ArtistDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*dto.ArtistDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockUsecaseMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUsecase)(nil).GetAll), ctx)
}

// GetFavoriteArtists mocks base method.
func (m *MockUsecase) GetFavoriteArtists(ctx context.Context, userID uuid.UUID) ([]*dto.ArtistDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoriteArtists", ctx, userID)
	ret0, _ := ret[0].([]*dto.ArtistDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoriteArtists indicates an expected call of GetFavoriteArtists.
func (mr *MockUsecaseMockRecorder) GetFavoriteArtists(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoriteArtists", reflect.TypeOf((*MockUsecase)(nil).GetFavoriteArtists), ctx, userID)
}

// GetPopular mocks base method.
func (m *MockUsecase) GetPopular(ctx context.Context) ([]*dto.ArtistDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPopular", ctx)
	ret0, _ := ret[0].([]*dto.ArtistDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPopular indicates an expected call of GetPopular.
func (mr *MockUsecaseMockRecorder) GetPopular(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPopular", reflect.TypeOf((*MockUsecase)(nil).GetPopular), ctx)
}

// IsFavoriteArtist mocks base method.
func (m *MockUsecase) IsFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFavoriteArtist", ctx, userID, artistID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFavoriteArtist indicates an expected call of IsFavoriteArtist.
func (mr *MockUsecaseMockRecorder) IsFavoriteArtist(ctx, userID, artistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFavoriteArtist", reflect.TypeOf((*MockUsecase)(nil).IsFavoriteArtist), ctx, userID, artistID)
}

// Search mocks base method.
func (m *MockUsecase) Search(ctx context.Context, query string) ([]*dto.ArtistDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, query)
	ret0, _ := ret[0].([]*dto.ArtistDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockUsecaseMockRecorder) Search(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockUsecase)(nil).Search), ctx, query)
}

// View mocks base method.
func (m *MockUsecase) View(ctx context.Context, artistID uint64) (*dto.ArtistDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "View", ctx, artistID)
	ret0, _ := ret[0].(*dto.ArtistDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// View indicates an expected call of View.
func (mr *MockUsecaseMockRecorder) View(ctx, artistID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "View", reflect.TypeOf((*MockUsecase)(nil).View), ctx, artistID)
}
