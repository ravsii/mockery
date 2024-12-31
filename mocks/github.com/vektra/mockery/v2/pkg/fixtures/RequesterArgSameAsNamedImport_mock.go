
// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package mocks

import (
	"encoding/json"
    mock "github.com/stretchr/testify/mock"
)

 
// NewRequesterArgSameAsNamedImport creates a new instance of RequesterArgSameAsNamedImport. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequesterArgSameAsNamedImport (t interface {
	mock.TestingT
	Cleanup(func())
}) *RequesterArgSameAsNamedImport {
	mock := &RequesterArgSameAsNamedImport{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}


// RequesterArgSameAsNamedImport is an autogenerated mock type for the RequesterArgSameAsNamedImport type
type RequesterArgSameAsNamedImport struct {
	mock.Mock
}

type RequesterArgSameAsNamedImport_Expecter struct {
	mock *mock.Mock
}

func (_m *RequesterArgSameAsNamedImport) EXPECT() *RequesterArgSameAsNamedImport_Expecter {
	return &RequesterArgSameAsNamedImport_Expecter{mock: &_m.Mock}
}

 

// Get provides a mock function for the type RequesterArgSameAsNamedImport
func (_mock *RequesterArgSameAsNamedImport) Get(json1 string) *json.RawMessage {  
	ret := _mock.Called(json1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

		
	var r0 *json.RawMessage 
	if returnFunc, ok := ret.Get(0).(func(string) *json.RawMessage); ok {
		r0 = returnFunc(json1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*json.RawMessage)
		}
	} 
	return r0
}



// RequesterArgSameAsNamedImport_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type RequesterArgSameAsNamedImport_Get_Call struct {
	*mock.Call
}



// Get is a helper method to define mock.On call
//  - json1
func (_e *RequesterArgSameAsNamedImport_Expecter) Get(json1 interface{}, ) *RequesterArgSameAsNamedImport_Get_Call {
	return &RequesterArgSameAsNamedImport_Get_Call{Call: _e.mock.On("Get",json1, )}
}

func (_c *RequesterArgSameAsNamedImport_Get_Call) Run(run func(json1 string)) *RequesterArgSameAsNamedImport_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string),)
	})
	return _c
}

func (_c *RequesterArgSameAsNamedImport_Get_Call) Return(rawMessage *json.RawMessage) *RequesterArgSameAsNamedImport_Get_Call {
	_c.Call.Return(rawMessage)
	return _c
}

func (_c *RequesterArgSameAsNamedImport_Get_Call) RunAndReturn(run func(json1 string)*json.RawMessage) *RequesterArgSameAsNamedImport_Get_Call {
	_c.Call.Return(run)
	return _c
}
  

