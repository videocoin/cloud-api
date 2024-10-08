// Code generated by MockGen. DO NOT EDIT.
// Source: /go_workspace/src/github.com/videocoin/cloud-api/profiles/v1/profiles_service.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	types "github.com/gogo/protobuf/types"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/videocoin/cloud-api/profiles/v1"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockProfilesServiceClient is a mock of ProfilesServiceClient interface
type MockProfilesServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockProfilesServiceClientMockRecorder
}

// MockProfilesServiceClientMockRecorder is the mock recorder for MockProfilesServiceClient
type MockProfilesServiceClientMockRecorder struct {
	mock *MockProfilesServiceClient
}

// NewMockProfilesServiceClient creates a new mock instance
func NewMockProfilesServiceClient(ctrl *gomock.Controller) *MockProfilesServiceClient {
	mock := &MockProfilesServiceClient{ctrl: ctrl}
	mock.recorder = &MockProfilesServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProfilesServiceClient) EXPECT() *MockProfilesServiceClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockProfilesServiceClient) Get(ctx context.Context, in *v1.ProfileRequest, opts ...grpc.CallOption) (*v1.GetProfileResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*v1.GetProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockProfilesServiceClientMockRecorder) Get(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProfilesServiceClient)(nil).Get), varargs...)
}

// List mocks base method
func (m *MockProfilesServiceClient) List(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*v1.ProfileListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v1.ProfileListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockProfilesServiceClientMockRecorder) List(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProfilesServiceClient)(nil).List), varargs...)
}

// Render mocks base method
func (m *MockProfilesServiceClient) Render(ctx context.Context, in *v1.RenderRequest, opts ...grpc.CallOption) (*v1.RenderResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Render", varargs...)
	ret0, _ := ret[0].(*v1.RenderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Render indicates an expected call of Render
func (mr *MockProfilesServiceClientMockRecorder) Render(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockProfilesServiceClient)(nil).Render), varargs...)
}

// MockProfilesServiceServer is a mock of ProfilesServiceServer interface
type MockProfilesServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockProfilesServiceServerMockRecorder
}

// MockProfilesServiceServerMockRecorder is the mock recorder for MockProfilesServiceServer
type MockProfilesServiceServerMockRecorder struct {
	mock *MockProfilesServiceServer
}

// NewMockProfilesServiceServer creates a new mock instance
func NewMockProfilesServiceServer(ctrl *gomock.Controller) *MockProfilesServiceServer {
	mock := &MockProfilesServiceServer{ctrl: ctrl}
	mock.recorder = &MockProfilesServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProfilesServiceServer) EXPECT() *MockProfilesServiceServerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockProfilesServiceServer) Get(arg0 context.Context, arg1 *v1.ProfileRequest) (*v1.GetProfileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1.GetProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockProfilesServiceServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProfilesServiceServer)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockProfilesServiceServer) List(arg0 context.Context, arg1 *types.Empty) (*v1.ProfileListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ProfileListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockProfilesServiceServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProfilesServiceServer)(nil).List), arg0, arg1)
}

// Render mocks base method
func (m *MockProfilesServiceServer) Render(arg0 context.Context, arg1 *v1.RenderRequest) (*v1.RenderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Render", arg0, arg1)
	ret0, _ := ret[0].(*v1.RenderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Render indicates an expected call of Render
func (mr *MockProfilesServiceServerMockRecorder) Render(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockProfilesServiceServer)(nil).Render), arg0, arg1)
}
