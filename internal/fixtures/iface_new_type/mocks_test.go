// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package iface_new_type

import (
	mock "github.com/stretchr/testify/mock"
)

// NewMockInterface1 creates a new instance of MockInterface1. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInterface1(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInterface1 {
	mock := &MockInterface1{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockInterface1 is an autogenerated mock type for the Interface1 type
type MockInterface1 struct {
	mock.Mock
}

type MockInterface1_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInterface1) EXPECT() *MockInterface1_Expecter {
	return &MockInterface1_Expecter{mock: &_m.Mock}
}

// Method1 provides a mock function for the type MockInterface1
func (_mock *MockInterface1) Method1() {
	_mock.Called()
	return
}

// MockInterface1_Method1_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Method1'
type MockInterface1_Method1_Call struct {
	*mock.Call
}

// Method1 is a helper method to define mock.On call
func (_e *MockInterface1_Expecter) Method1() *MockInterface1_Method1_Call {
	return &MockInterface1_Method1_Call{Call: _e.mock.On("Method1")}
}

func (_c *MockInterface1_Method1_Call) Run(run func()) *MockInterface1_Method1_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterface1_Method1_Call) Return() *MockInterface1_Method1_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockInterface1_Method1_Call) RunAndReturn(run func()) *MockInterface1_Method1_Call {
	_c.Run(run)
	return _c
}
