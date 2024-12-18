// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/db/s3/repository.go
//
// Generated by this command:
//
//	mockgen -source=pkg/db/s3/repository.go -destination=internal/user/mock/s3_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	s3 "github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3"
	gomock "github.com/golang/mock/gomock"
	v7 "github.com/minio/minio-go/v7"
)

// MockS3Repo is a mock of S3Repo interface.
type MockS3Repo struct {
	ctrl     *gomock.Controller
	recorder *MockS3RepoMockRecorder
	isgomock struct{}
}

// MockS3RepoMockRecorder is the mock recorder for MockS3Repo.
type MockS3RepoMockRecorder struct {
	mock *MockS3Repo
}

// NewMockS3Repo creates a new mock instance.
func NewMockS3Repo(ctrl *gomock.Controller) *MockS3Repo {
	mock := &MockS3Repo{ctrl: ctrl}
	mock.recorder = &MockS3RepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3Repo) EXPECT() *MockS3RepoMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockS3Repo) Get(ctx context.Context, bucket, filename string) (*v7.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, bucket, filename)
	ret0, _ := ret[0].(*v7.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockS3RepoMockRecorder) Get(ctx, bucket, filename any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockS3Repo)(nil).Get), ctx, bucket, filename)
}

// Put mocks base method.
func (m *MockS3Repo) Put(ctx context.Context, upload s3.Upload) (*v7.UploadInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ctx, upload)
	ret0, _ := ret[0].(*v7.UploadInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockS3RepoMockRecorder) Put(ctx, upload any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockS3Repo)(nil).Put), ctx, upload)
}

// Remove mocks base method.
func (m *MockS3Repo) Remove(ctx context.Context, bucket, filename string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, bucket, filename)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockS3RepoMockRecorder) Remove(ctx, bucket, filename any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockS3Repo)(nil).Remove), ctx, bucket, filename)
}
