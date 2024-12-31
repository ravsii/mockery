
// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package mocks

import (
    mock "github.com/stretchr/testify/mock"
)

 
// NewIssue766 creates a new instance of Issue766. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIssue766 (t interface {
	mock.TestingT
	Cleanup(func())
}) *Issue766 {
	mock := &Issue766{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}


// Issue766 is an autogenerated mock type for the Issue766 type
type Issue766 struct {
	mock.Mock
}

type Issue766_Expecter struct {
	mock *mock.Mock
}

func (_m *Issue766) EXPECT() *Issue766_Expecter {
	return &Issue766_Expecter{mock: &_m.Mock}
}

 

// FetchData provides a mock function for the type Issue766
func (_mock *Issue766) FetchData(fetchFunc func(x ...int) ([]int, error)) ([]int, error) {  
	ret := _mock.Called(fetchFunc)

	if len(ret) == 0 {
		panic("no return value specified for FetchData")
	}

		
	var r0 []int
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(func(x ...int) ([]int, error)) ([]int, error)); ok {
		return returnFunc(fetchFunc)
	} 
	if returnFunc, ok := ret.Get(0).(func(func(x ...int) ([]int, error)) []int); ok {
		r0 = returnFunc(fetchFunc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	} 
	if returnFunc, ok := ret.Get(1).(func(func(x ...int) ([]int, error)) error); ok {
		r1 = returnFunc(fetchFunc)
	} else {
		r1 = ret.Error(1)
	} 
	return r0, r1
}



// Issue766_FetchData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchData'
type Issue766_FetchData_Call struct {
	*mock.Call
}



// FetchData is a helper method to define mock.On call
//  - fetchFunc
func (_e *Issue766_Expecter) FetchData(fetchFunc interface{}, ) *Issue766_FetchData_Call {
	return &Issue766_FetchData_Call{Call: _e.mock.On("FetchData",fetchFunc, )}
}

func (_c *Issue766_FetchData_Call) Run(run func(fetchFunc func(x ...int) ([]int, error))) *Issue766_FetchData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(x ...int) ([]int, error)),)
	})
	return _c
}

func (_c *Issue766_FetchData_Call) Return(ints []int, err error) *Issue766_FetchData_Call {
	_c.Call.Return(ints, err)
	return _c
}

func (_c *Issue766_FetchData_Call) RunAndReturn(run func(fetchFunc func(x ...int) ([]int, error))([]int, error)) *Issue766_FetchData_Call {
	_c.Call.Return(run)
	return _c
}
  

