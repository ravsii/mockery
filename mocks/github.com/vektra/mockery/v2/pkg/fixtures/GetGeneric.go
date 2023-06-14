// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	constraints "github.com/vektra/mockery/v2/pkg/fixtures/constraints"
)

// GetGeneric is an autogenerated mock type for the GetGeneric type
type GetGeneric[T constraints.Integer] struct {
	mock.Mock
}

type GetGeneric_Expecter[T constraints.Integer] struct {
	mock *mock.Mock
}

func (_m *GetGeneric[T]) EXPECT() *GetGeneric_Expecter[T] {
	return &GetGeneric_Expecter[T]{mock: &_m.Mock}
}

// Get provides a mock function with given fields:
func (_m *GetGeneric[T]) Get() T {
	ret := _m.Called()

	var r0 T
	if rf, ok := ret.Get(0).(func() T); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(T)
	}

	return r0
}

// GetGeneric_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type GetGeneric_Get_Call[T constraints.Integer] struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *GetGeneric_Expecter[T]) Get() *GetGeneric_Get_Call[T] {
	return &GetGeneric_Get_Call[T]{Call: _e.mock.On("Get")}
}

func (_c *GetGeneric_Get_Call[T]) Run(run func()) *GetGeneric_Get_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GetGeneric_Get_Call[T]) Return(_a0 T) *GetGeneric_Get_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GetGeneric_Get_Call[T]) RunAndReturn(run func() T) *GetGeneric_Get_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewGetGeneric creates a new instance of GetGeneric. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetGeneric[T constraints.Integer](t interface {
	mock.TestingT
	Cleanup(func())
}) *GetGeneric[T] {
	mock := &GetGeneric[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
