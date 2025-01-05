// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package type_alias

import (
	mock "github.com/stretchr/testify/mock"
	"github.com/vektra/mockery/v3/internal/fixtures/type_alias/subpkg"
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

// Foo provides a mock function for the type MockInterface1
func (_mock *MockInterface1) Foo() Type {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for Foo")
	}

	var r0 Type
	if returnFunc, ok := ret.Get(0).(func() Type); ok {
		r0 = returnFunc()
	} else {
		r0 = ret.Get(0).(Type)
	}
	return r0
}

// MockInterface1_Foo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Foo'
type MockInterface1_Foo_Call struct {
	*mock.Call
}

// Foo is a helper method to define mock.On call
func (_e *MockInterface1_Expecter) Foo() *MockInterface1_Foo_Call {
	return &MockInterface1_Foo_Call{Call: _e.mock.On("Foo")}
}

func (_c *MockInterface1_Foo_Call) Run(run func()) *MockInterface1_Foo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterface1_Foo_Call) Return(v Type) *MockInterface1_Foo_Call {
	_c.Call.Return(v)
	return _c
}

func (_c *MockInterface1_Foo_Call) RunAndReturn(run func() Type) *MockInterface1_Foo_Call {
	_c.Call.Return(run)
	return _c
}

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

// MockInterface1 is an autogenerated mock type for the Interface2 type
type MockInterface1 struct {
	mock.Mock
}

type MockInterface1_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInterface1) EXPECT() *MockInterface1_Expecter {
	return &MockInterface1_Expecter{mock: &_m.Mock}
}

// F provides a mock function for the type MockInterface1
func (_mock *MockInterface1) F(v Type, v1 S, s subpkg.S) {
	_mock.Called(v, v1, s)
	return
}

// MockInterface1_F_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'F'
type MockInterface1_F_Call struct {
	*mock.Call
}

// F is a helper method to define mock.On call
//   - v
//   - v1
//   - s
func (_e *MockInterface1_Expecter) F(v interface{}, v1 interface{}, s interface{}) *MockInterface1_F_Call {
	return &MockInterface1_F_Call{Call: _e.mock.On("F", v, v1, s)}
}

func (_c *MockInterface1_F_Call) Run(run func(v Type, v1 S, s subpkg.S)) *MockInterface1_F_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Type), args[1].(S), args[2].(subpkg.S))
	})
	return _c
}

func (_c *MockInterface1_F_Call) Return() *MockInterface1_F_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockInterface1_F_Call) RunAndReturn(run func(v Type, v1 S, s subpkg.S)) *MockInterface1_F_Call {
	_c.Run(run)
	return _c
}
