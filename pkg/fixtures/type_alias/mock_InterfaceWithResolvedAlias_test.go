// Code generated by mockery. DO NOT EDIT.

package type_alias_test

import mock "github.com/stretchr/testify/mock"

// InterfaceWithResolvedAlias is an autogenerated mock type for the Interface1 type
type InterfaceWithResolvedAlias struct {
	mock.Mock
}

type InterfaceWithResolvedAlias_Expecter struct {
	mock *mock.Mock
}

func (_m *InterfaceWithResolvedAlias) EXPECT() *InterfaceWithResolvedAlias_Expecter {
	return &InterfaceWithResolvedAlias_Expecter{mock: &_m.Mock}
}

// Foo provides a mock function with no fields
func (_m *InterfaceWithResolvedAlias) Foo() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Foo")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// InterfaceWithResolvedAlias_Foo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Foo'
type InterfaceWithResolvedAlias_Foo_Call struct {
	*mock.Call
}

// Foo is a helper method to define mock.On call
func (_e *InterfaceWithResolvedAlias_Expecter) Foo() *InterfaceWithResolvedAlias_Foo_Call {
	return &InterfaceWithResolvedAlias_Foo_Call{Call: _e.mock.On("Foo")}
}

func (_c *InterfaceWithResolvedAlias_Foo_Call) Run(run func()) *InterfaceWithResolvedAlias_Foo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *InterfaceWithResolvedAlias_Foo_Call) Return(_a0 int) *InterfaceWithResolvedAlias_Foo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *InterfaceWithResolvedAlias_Foo_Call) RunAndReturn(run func() int) *InterfaceWithResolvedAlias_Foo_Call {
	_c.Call.Return(run)
	return _c
}

// NewInterfaceWithResolvedAlias creates a new instance of InterfaceWithResolvedAlias. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInterfaceWithResolvedAlias(t interface {
	mock.TestingT
	Cleanup(func())
}) *InterfaceWithResolvedAlias {
	mock := &InterfaceWithResolvedAlias{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
