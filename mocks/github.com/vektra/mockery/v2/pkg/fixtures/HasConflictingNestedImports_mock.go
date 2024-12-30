// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package mocks

import (
	"net/http"

	mock "github.com/stretchr/testify/mock"
	my_http "github.com/vektra/mockery/v2/pkg/fixtures/http"
)

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

// Get provides a mock function for the type HasConflictingNestedImports
func (_mock *HasConflictingNestedImports) Get(path string) (http.Response, error) {
	ret := _mock.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 http.Response
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(string) (http.Response, error)); ok {
		return returnFunc(path)
	}
	if returnFunc, ok := ret.Get(0).(func(string) http.Response); ok {
		r0 = returnFunc(path)
	} else {
		r0 = ret.Get(0).(http.Response)
	}
	if returnFunc, ok := ret.Get(1).(func(string) error); ok {
		r1 = returnFunc(path)
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
//   - path
func (_e *HasConflictingNestedImports_Expecter) Get(path interface{}) *HasConflictingNestedImports_Get_Call {
	return &HasConflictingNestedImports_Get_Call{Call: _e.mock.On("Get", path)}
}

func (_c *HasConflictingNestedImports_Get_Call) Run(run func(path string)) *HasConflictingNestedImports_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(path)
	})
	return _c
}

func (_c *HasConflictingNestedImports_Get_Call) Return(responseOut http.Response, errOut error) *HasConflictingNestedImports_Get_Call {
	_c.Call.Return(responseOut, errOut)
	return _c
}

func (_c *HasConflictingNestedImports_Get_Call) RunAndReturn(run func(path string) (http.Response, error)) *HasConflictingNestedImports_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Z provides a mock function for the type HasConflictingNestedImports
func (_mock *HasConflictingNestedImports) Z() my_http.MyStruct {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for Z")
	}

	var r0 my_http.MyStruct
	if returnFunc, ok := ret.Get(0).(func() my_http.MyStruct); ok {
		r0 = returnFunc()
	} else {
		r0 = ret.Get(0).(my_http.MyStruct)
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

func (_c *HasConflictingNestedImports_Z_Call) Return(myStructOut my_http.MyStruct) *HasConflictingNestedImports_Z_Call {
	_c.Call.Return(myStructOut)
	return _c
}

func (_c *HasConflictingNestedImports_Z_Call) RunAndReturn(run func() my_http.MyStruct) *HasConflictingNestedImports_Z_Call {
	_c.Call.Return(run)
	return _c
}
