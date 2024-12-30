// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	test "github.com/vektra/mockery/v2/pkg/fixtures"
)

// NewUsesOtherPkgIface creates a new instance of UsesOtherPkgIface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsesOtherPkgIface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UsesOtherPkgIface {
	mock := &UsesOtherPkgIface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// UsesOtherPkgIface is an autogenerated mock type for the UsesOtherPkgIface type
type UsesOtherPkgIface struct {
	mock.Mock
}

type UsesOtherPkgIface_Expecter struct {
	mock *mock.Mock
}

func (_m *UsesOtherPkgIface) EXPECT() *UsesOtherPkgIface_Expecter {
	return &UsesOtherPkgIface_Expecter{mock: &_m.Mock}
}

// DoSomethingElse provides a mock function for the type UsesOtherPkgIface
func (_mock *UsesOtherPkgIface) DoSomethingElse(obj test.Sibling) {
	_mock.Called(obj)
	return
}

// UsesOtherPkgIface_DoSomethingElse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DoSomethingElse'
type UsesOtherPkgIface_DoSomethingElse_Call struct {
	*mock.Call
}

// DoSomethingElse is a helper method to define mock.On call
//   - obj
func (_e *UsesOtherPkgIface_Expecter) DoSomethingElse(obj interface{}) *UsesOtherPkgIface_DoSomethingElse_Call {
	return &UsesOtherPkgIface_DoSomethingElse_Call{Call: _e.mock.On("DoSomethingElse", obj)}
}

func (_c *UsesOtherPkgIface_DoSomethingElse_Call) Run(run func(obj test.Sibling)) *UsesOtherPkgIface_DoSomethingElse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(obj)
	})
	return _c
}

func (_c *UsesOtherPkgIface_DoSomethingElse_Call) Return() *UsesOtherPkgIface_DoSomethingElse_Call {
	_c.Call.Return()
	return _c
}

func (_c *UsesOtherPkgIface_DoSomethingElse_Call) RunAndReturn(run func(obj test.Sibling)) *UsesOtherPkgIface_DoSomethingElse_Call {
	_c.Run(run)
	return _c
}
