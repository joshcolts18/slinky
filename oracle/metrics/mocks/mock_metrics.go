// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Metrics is an autogenerated mock type for the Metrics type
type Metrics struct {
	mock.Mock
}

// AddTick provides a mock function with given fields:
func (_m *Metrics) AddTick() {
	_m.Called()
}

// UpdateAggregatePrice provides a mock function with given fields: pairID, decimals, price
func (_m *Metrics) UpdateAggregatePrice(pairID string, decimals uint64, price float64) {
	_m.Called(pairID, decimals, price)
}

// UpdatePrice provides a mock function with given fields: name, handlerType, pairID, decimals, price
func (_m *Metrics) UpdatePrice(name string, handlerType string, pairID string, decimals uint64, price float64) {
	_m.Called(name, handlerType, pairID, decimals, price)
}

// NewMetrics creates a new instance of Metrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMetrics(t interface {
	mock.TestingT
	Cleanup(func())
}) *Metrics {
	mock := &Metrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
