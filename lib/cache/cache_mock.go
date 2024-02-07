// Code generated by MockGen. DO NOT EDIT.
// Source: lib/cache/interface.go

// Package cache is a generated GoMock package.
package cache

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	codec "github.com/sdimon13/cache/lib/codec"
	store "github.com/sdimon13/cache/lib/store"
)

// MockCacheInterface is a mock of CacheInterface interface.
type MockCacheInterface[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockCacheInterfaceMockRecorder[T]
}

// MockCacheInterfaceMockRecorder is the mock recorder for MockCacheInterface.
type MockCacheInterfaceMockRecorder[T any] struct {
	mock *MockCacheInterface[T]
}

// NewMockCacheInterface creates a new mock instance.
func NewMockCacheInterface[T any](ctrl *gomock.Controller) *MockCacheInterface[T] {
	mock := &MockCacheInterface[T]{ctrl: ctrl}
	mock.recorder = &MockCacheInterfaceMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheInterface[T]) EXPECT() *MockCacheInterfaceMockRecorder[T] {
	return m.recorder
}

// Clear mocks base method.
func (m *MockCacheInterface[T]) Clear(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear.
func (mr *MockCacheInterfaceMockRecorder[T]) Clear(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockCacheInterface[T])(nil).Clear), ctx)
}

// Delete mocks base method.
func (m *MockCacheInterface[T]) Delete(ctx context.Context, key any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCacheInterfaceMockRecorder[T]) Delete(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCacheInterface[T])(nil).Delete), ctx, key)
}

// Get mocks base method.
func (m *MockCacheInterface[T]) Get(ctx context.Context, key any) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheInterfaceMockRecorder[T]) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacheInterface[T])(nil).Get), ctx, key)
}

// GetType mocks base method.
func (m *MockCacheInterface[T]) GetType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetType indicates an expected call of GetType.
func (mr *MockCacheInterfaceMockRecorder[T]) GetType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockCacheInterface[T])(nil).GetType))
}

// Invalidate mocks base method.
func (m *MockCacheInterface[T]) Invalidate(ctx context.Context, options ...store.InvalidateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Invalidate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Invalidate indicates an expected call of Invalidate.
func (mr *MockCacheInterfaceMockRecorder[T]) Invalidate(ctx interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidate", reflect.TypeOf((*MockCacheInterface[T])(nil).Invalidate), varargs...)
}

// Set mocks base method.
func (m *MockCacheInterface[T]) Set(ctx context.Context, key any, object T, options ...store.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key, object}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Set", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheInterfaceMockRecorder[T]) Set(ctx, key, object interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key, object}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCacheInterface[T])(nil).Set), varargs...)
}

// MockCacheKeyGenerator is a mock of CacheKeyGenerator interface.
type MockCacheKeyGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockCacheKeyGeneratorMockRecorder
}

// MockCacheKeyGeneratorMockRecorder is the mock recorder for MockCacheKeyGenerator.
type MockCacheKeyGeneratorMockRecorder struct {
	mock *MockCacheKeyGenerator
}

// NewMockCacheKeyGenerator creates a new mock instance.
func NewMockCacheKeyGenerator(ctrl *gomock.Controller) *MockCacheKeyGenerator {
	mock := &MockCacheKeyGenerator{ctrl: ctrl}
	mock.recorder = &MockCacheKeyGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheKeyGenerator) EXPECT() *MockCacheKeyGeneratorMockRecorder {
	return m.recorder
}

// GetCacheKey mocks base method.
func (m *MockCacheKeyGenerator) GetCacheKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCacheKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCacheKey indicates an expected call of GetCacheKey.
func (mr *MockCacheKeyGeneratorMockRecorder) GetCacheKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCacheKey", reflect.TypeOf((*MockCacheKeyGenerator)(nil).GetCacheKey))
}

// MockSetterCacheInterface is a mock of SetterCacheInterface interface.
type MockSetterCacheInterface[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockSetterCacheInterfaceMockRecorder[T]
}

// MockSetterCacheInterfaceMockRecorder is the mock recorder for MockSetterCacheInterface.
type MockSetterCacheInterfaceMockRecorder[T any] struct {
	mock *MockSetterCacheInterface[T]
}

// NewMockSetterCacheInterface creates a new mock instance.
func NewMockSetterCacheInterface[T any](ctrl *gomock.Controller) *MockSetterCacheInterface[T] {
	mock := &MockSetterCacheInterface[T]{ctrl: ctrl}
	mock.recorder = &MockSetterCacheInterfaceMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSetterCacheInterface[T]) EXPECT() *MockSetterCacheInterfaceMockRecorder[T] {
	return m.recorder
}

// Clear mocks base method.
func (m *MockSetterCacheInterface[T]) Clear(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear.
func (mr *MockSetterCacheInterfaceMockRecorder[T]) Clear(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockSetterCacheInterface[T])(nil).Clear), ctx)
}

// Delete mocks base method.
func (m *MockSetterCacheInterface[T]) Delete(ctx context.Context, key any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSetterCacheInterfaceMockRecorder[T]) Delete(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSetterCacheInterface[T])(nil).Delete), ctx, key)
}

// Get mocks base method.
func (m *MockSetterCacheInterface[T]) Get(ctx context.Context, key any) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSetterCacheInterfaceMockRecorder[T]) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSetterCacheInterface[T])(nil).Get), ctx, key)
}

// GetCodec mocks base method.
func (m *MockSetterCacheInterface[T]) GetCodec() codec.CodecInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCodec")
	ret0, _ := ret[0].(codec.CodecInterface)
	return ret0
}

// GetCodec indicates an expected call of GetCodec.
func (mr *MockSetterCacheInterfaceMockRecorder[T]) GetCodec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCodec", reflect.TypeOf((*MockSetterCacheInterface[T])(nil).GetCodec))
}

// GetType mocks base method.
func (m *MockSetterCacheInterface[T]) GetType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetType indicates an expected call of GetType.
func (mr *MockSetterCacheInterfaceMockRecorder[T]) GetType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockSetterCacheInterface[T])(nil).GetType))
}

// GetWithTTL mocks base method.
func (m *MockSetterCacheInterface[T]) GetWithTTL(ctx context.Context, key any) (T, time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithTTL", ctx, key)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(time.Duration)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetWithTTL indicates an expected call of GetWithTTL.
func (mr *MockSetterCacheInterfaceMockRecorder[T]) GetWithTTL(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithTTL", reflect.TypeOf((*MockSetterCacheInterface[T])(nil).GetWithTTL), ctx, key)
}

// Invalidate mocks base method.
func (m *MockSetterCacheInterface[T]) Invalidate(ctx context.Context, options ...store.InvalidateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Invalidate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Invalidate indicates an expected call of Invalidate.
func (mr *MockSetterCacheInterfaceMockRecorder[T]) Invalidate(ctx interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invalidate", reflect.TypeOf((*MockSetterCacheInterface[T])(nil).Invalidate), varargs...)
}

// Set mocks base method.
func (m *MockSetterCacheInterface[T]) Set(ctx context.Context, key any, object T, options ...store.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key, object}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Set", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockSetterCacheInterfaceMockRecorder[T]) Set(ctx, key, object interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key, object}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockSetterCacheInterface[T])(nil).Set), varargs...)
}
