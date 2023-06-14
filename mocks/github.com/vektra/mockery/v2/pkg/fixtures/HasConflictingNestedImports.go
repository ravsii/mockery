// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	http "net/http"

	fixtureshttp "github.com/vektra/mockery/v2/pkg/fixtures/http"

	mock "github.com/stretchr/testify/mock"
)

// HasConflictingNestedImports is an autogenerated mock type for the HasConflictingNestedImports type
type HasConflictingNestedImports struct {
	mock.Mock
}

type HasConflictingNestedImports_Expecter struct {
	mock *mock.Mock
}

func (_m *HasConflictingNestedImports) EXPECT() *HasConflictingNestedImports_Expecter {
	return &HasConflictingNestedImports_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: path
func (_m *HasConflictingNestedImports) Get(path string) (http.Response, error) {
	ret := _m.Called(path)

	var r0 http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (http.Response, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) http.Response); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(http.Response)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasConflictingNestedImports_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type HasConflictingNestedImports_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - path string
func (_e *HasConflictingNestedImports_Expecter) Get(path interface{}) *HasConflictingNestedImports_Get_Call {
	return &HasConflictingNestedImports_Get_Call{Call: _e.mock.On("Get", path)}
}

func (_c *HasConflictingNestedImports_Get_Call) Run(run func(path string)) *HasConflictingNestedImports_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *HasConflictingNestedImports_Get_Call) Return(_a0 http.Response, _a1 error) *HasConflictingNestedImports_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *HasConflictingNestedImports_Get_Call) RunAndReturn(run func(string) (http.Response, error)) *HasConflictingNestedImports_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Z provides a mock function with given fields:
func (_m *HasConflictingNestedImports) Z() fixtureshttp.MyStruct {
	ret := _m.Called()

	var r0 fixtureshttp.MyStruct
	if rf, ok := ret.Get(0).(func() fixtureshttp.MyStruct); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(fixtureshttp.MyStruct)
	}

	return r0
}

// HasConflictingNestedImports_Z_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Z'
type HasConflictingNestedImports_Z_Call struct {
	*mock.Call
}

// Z is a helper method to define mock.On call
func (_e *HasConflictingNestedImports_Expecter) Z() *HasConflictingNestedImports_Z_Call {
	return &HasConflictingNestedImports_Z_Call{Call: _e.mock.On("Z")}
}

func (_c *HasConflictingNestedImports_Z_Call) Run(run func()) *HasConflictingNestedImports_Z_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HasConflictingNestedImports_Z_Call) Return(_a0 fixtureshttp.MyStruct) *HasConflictingNestedImports_Z_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HasConflictingNestedImports_Z_Call) RunAndReturn(run func() fixtureshttp.MyStruct) *HasConflictingNestedImports_Z_Call {
	_c.Call.Return(run)
	return _c
}

// NewHasConflictingNestedImports creates a new instance of HasConflictingNestedImports. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHasConflictingNestedImports(t interface {
	mock.TestingT
	Cleanup(func())
}) *HasConflictingNestedImports {
	mock := &HasConflictingNestedImports{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
