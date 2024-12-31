
// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package mocks

import (
    mock "github.com/stretchr/testify/mock"
)

 
// NewRequesterReturnElided creates a new instance of RequesterReturnElided. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequesterReturnElided (t interface {
	mock.TestingT
	Cleanup(func())
}) *RequesterReturnElided {
	mock := &RequesterReturnElided{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}


// RequesterReturnElided is an autogenerated mock type for the RequesterReturnElided type
type RequesterReturnElided struct {
	mock.Mock
}

type RequesterReturnElided_Expecter struct {
	mock *mock.Mock
}

func (_m *RequesterReturnElided) EXPECT() *RequesterReturnElided_Expecter {
	return &RequesterReturnElided_Expecter{mock: &_m.Mock}
}

 

// Get provides a mock function for the type RequesterReturnElided
func (_mock *RequesterReturnElided) Get(path string) (int, int, int, error) {  
	ret := _mock.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

		
	var r0 int
	var r1 int
	var r2 int
	var r3 error
	if returnFunc, ok := ret.Get(0).(func(string) (int, int, int, error)); ok {
		return returnFunc(path)
	} 
	if returnFunc, ok := ret.Get(0).(func(string) int); ok {
		r0 = returnFunc(path)
	} else {
		r0 = ret.Get(0).(int)
	} 
	if returnFunc, ok := ret.Get(1).(func(string) int); ok {
		r1 = returnFunc(path)
	} else {
		r1 = ret.Get(1).(int)
	} 
	if returnFunc, ok := ret.Get(2).(func(string) int); ok {
		r2 = returnFunc(path)
	} else {
		r2 = ret.Get(2).(int)
	} 
	if returnFunc, ok := ret.Get(3).(func(string) error); ok {
		r3 = returnFunc(path)
	} else {
		r3 = ret.Error(3)
	} 
	return r0, r1, r2, r3
}



// RequesterReturnElided_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type RequesterReturnElided_Get_Call struct {
	*mock.Call
}



// Get is a helper method to define mock.On call
//  - path
func (_e *RequesterReturnElided_Expecter) Get(path interface{}, ) *RequesterReturnElided_Get_Call {
	return &RequesterReturnElided_Get_Call{Call: _e.mock.On("Get",path, )}
}

func (_c *RequesterReturnElided_Get_Call) Run(run func(path string)) *RequesterReturnElided_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string),)
	})
	return _c
}

func (_c *RequesterReturnElided_Get_Call) Return(a int, b int, c int, err error) *RequesterReturnElided_Get_Call {
	_c.Call.Return(a, b, c, err)
	return _c
}

func (_c *RequesterReturnElided_Get_Call) RunAndReturn(run func(path string)(int, int, int, error)) *RequesterReturnElided_Get_Call {
	_c.Call.Return(run)
	return _c
}
 

// Put provides a mock function for the type RequesterReturnElided
func (_mock *RequesterReturnElided) Put(path string) (int, error) {  
	ret := _mock.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

		
	var r0 int
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(string) (int, error)); ok {
		return returnFunc(path)
	} 
	if returnFunc, ok := ret.Get(0).(func(string) int); ok {
		r0 = returnFunc(path)
	} else {
		r0 = ret.Get(0).(int)
	} 
	if returnFunc, ok := ret.Get(1).(func(string) error); ok {
		r1 = returnFunc(path)
	} else {
		r1 = ret.Error(1)
	} 
	return r0, r1
}



// RequesterReturnElided_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type RequesterReturnElided_Put_Call struct {
	*mock.Call
}



// Put is a helper method to define mock.On call
//  - path
func (_e *RequesterReturnElided_Expecter) Put(path interface{}, ) *RequesterReturnElided_Put_Call {
	return &RequesterReturnElided_Put_Call{Call: _e.mock.On("Put",path, )}
}

func (_c *RequesterReturnElided_Put_Call) Run(run func(path string)) *RequesterReturnElided_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string),)
	})
	return _c
}

func (_c *RequesterReturnElided_Put_Call) Return(n int, err error) *RequesterReturnElided_Put_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *RequesterReturnElided_Put_Call) RunAndReturn(run func(path string)(int, error)) *RequesterReturnElided_Put_Call {
	_c.Call.Return(run)
	return _c
}
  

