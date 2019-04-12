// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/heptio/developer-dash/internal/objectstore (interfaces: ObjectStore)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	objectstoreutil "github.com/heptio/developer-dash/pkg/objectstoreutil"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	cache "k8s.io/client-go/tools/cache"
	reflect "reflect"
)

// MockObjectStore is a mock of ObjectStore interface
type MockObjectStore struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStoreMockRecorder
}

// MockObjectStoreMockRecorder is the mock recorder for MockObjectStore
type MockObjectStoreMockRecorder struct {
	mock *MockObjectStore
}

// NewMockObjectStore creates a new mock instance
func NewMockObjectStore(ctrl *gomock.Controller) *MockObjectStore {
	mock := &MockObjectStore{ctrl: ctrl}
	mock.recorder = &MockObjectStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockObjectStore) EXPECT() *MockObjectStoreMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockObjectStore) Get(arg0 context.Context, arg1 objectstoreutil.Key) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockObjectStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockObjectStore)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockObjectStore) List(arg0 context.Context, arg1 objectstoreutil.Key) ([]*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockObjectStoreMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockObjectStore)(nil).List), arg0, arg1)
}

// Watch mocks base method
func (m *MockObjectStore) Watch(arg0 objectstoreutil.Key, arg1 cache.ResourceEventHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Watch indicates an expected call of Watch
func (mr *MockObjectStoreMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockObjectStore)(nil).Watch), arg0, arg1)
}