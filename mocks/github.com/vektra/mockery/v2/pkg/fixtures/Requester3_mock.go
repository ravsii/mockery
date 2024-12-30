// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package mocks

import (
	mock "github.com/stretchr/testify/mock"
)

// NewRequester3 creates a new instance of Requester3. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequester3(t interface {
	mock.TestingT
	Cleanup(func())
}) *Requester3 {
	mock := &Requester3{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Requester3 is an autogenerated mock type for the Requester3 type
type Requester3 struct {
	mock.Mock
}

type Requester3_Expecter struct {
	mock *mock.Mock
}

func (_m *Requester3) EXPECT() *Requester3_Expecter {
	return &Requester3_Expecter{mock: &_m.Mock}
}

// Get provides a mock function for the type Requester3
func (_mock *Requester3) Get() error {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func() error); ok {
		r0 = returnFunc()
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// Requester3_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Requester3_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *Requester3_Expecter) Get() *Requester3_Get_Call {
	return &Requester3_Get_Call{Call: _e.mock.On("Get")}
}

func (_c *Requester3_Get_Call) Run(run func()) *Requester3_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Requester3_Get_Call) Return(errOut error) *Requester3_Get_Call {
	_c.Call.Return(errOut)
	return _c
}

func (_c *Requester3_Get_Call) RunAndReturn(run func() error) *Requester3_Get_Call {
	_c.Call.Return(run)
	return _c
}
