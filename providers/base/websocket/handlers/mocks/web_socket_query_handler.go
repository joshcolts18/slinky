// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/skip-mev/slinky/providers/types"
)

// WebSocketQueryHandler is an autogenerated mock type for the WebSocketQueryHandler type
type WebSocketQueryHandler[K types.ResponseKey, V types.ResponseValue] struct {
	mock.Mock
}

// Start provides a mock function with given fields: ctx, ids, responseCh
func (_m *WebSocketQueryHandler[K, V]) Start(ctx context.Context, ids []K, responseCh chan<- types.GetResponse[K, V]) error {
	ret := _m.Called(ctx, ids, responseCh)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []K, chan<- types.GetResponse[K, V]) error); ok {
		r0 = rf(ctx, ids, responseCh)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewWebSocketQueryHandler creates a new instance of WebSocketQueryHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWebSocketQueryHandler[K types.ResponseKey, V types.ResponseValue](t interface {
	mock.TestingT
	Cleanup(func())
}) *WebSocketQueryHandler[K, V] {
	mock := &WebSocketQueryHandler[K, V]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
