// Code generated by MockGen.
// Source: types/handler.go
// Chanes:
// + AnteHandler(...): calling `next` at the end of the function to run all anthe handler chain.

// Package mocks is a generated GoMock package.
package mocks

import (
	"reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/line/lbm-sdk/types"
)

// MockAnteDecorator is a mock of AnteDecorator interface.
type MockAnteDecorator struct {
	ctrl     *gomock.Controller
	recorder *MockAnteDecoratorMockRecorder
}

// MockAnteDecoratorMockRecorder is the mock recorder for MockAnteDecorator.
type MockAnteDecoratorMockRecorder struct {
	mock *MockAnteDecorator
}

// NewMockAnteDecorator creates a new mock instance.
func NewMockAnteDecorator(ctrl *gomock.Controller) *MockAnteDecorator {
	mock := &MockAnteDecorator{ctrl: ctrl}
	mock.recorder = &MockAnteDecoratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnteDecorator) EXPECT() *MockAnteDecoratorMockRecorder {
	return m.recorder
}

// AnteHandle mocks base method.
func (m *MockAnteDecorator) AnteHandle(ctx types.Context, tx types.Tx, simulate bool, next types.AnteHandler) (types.Context, error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AnteHandle", ctx, tx, simulate, next)
	// NOTE: we need to edit a generated code to call the "next handler"
	return next(ctx, tx, simulate)
}

// AnteHandle indicates an expected call of AnteHandle.
func (mr *MockAnteDecoratorMockRecorder) AnteHandle(ctx, tx, simulate, next interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnteHandle", reflect.TypeOf((*MockAnteDecorator)(nil).AnteHandle), ctx, tx, simulate, next)
}
