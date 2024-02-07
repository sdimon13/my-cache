// Code generated by MockGen. DO NOT EDIT.
// Source: lib/metrics/interface.go

// Package metrics is a generated GoMock package.
package metrics

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	codec "github.com/sdimon13/cache/lib/v4/codec"
)

// MockMetricsInterface is a mock of MetricsInterface interface.
type MockMetricsInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsInterfaceMockRecorder
}

// MockMetricsInterfaceMockRecorder is the mock recorder for MockMetricsInterface.
type MockMetricsInterfaceMockRecorder struct {
	mock *MockMetricsInterface
}

// NewMockMetricsInterface creates a new mock instance.
func NewMockMetricsInterface(ctrl *gomock.Controller) *MockMetricsInterface {
	mock := &MockMetricsInterface{ctrl: ctrl}
	mock.recorder = &MockMetricsInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetricsInterface) EXPECT() *MockMetricsInterfaceMockRecorder {
	return m.recorder
}

// RecordFromCodec mocks base method.
func (m *MockMetricsInterface) RecordFromCodec(codec codec.CodecInterface) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RecordFromCodec", codec)
}

// RecordFromCodec indicates an expected call of RecordFromCodec.
func (mr *MockMetricsInterfaceMockRecorder) RecordFromCodec(codec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordFromCodec", reflect.TypeOf((*MockMetricsInterface)(nil).RecordFromCodec), codec)
}
