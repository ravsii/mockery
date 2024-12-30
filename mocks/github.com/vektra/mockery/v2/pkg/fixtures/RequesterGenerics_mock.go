// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package mocks

import (
	"io"

	mock "github.com/stretchr/testify/mock"
	test "github.com/vektra/mockery/v2/pkg/fixtures"
	"github.com/vektra/mockery/v2/pkg/fixtures/constraints"
)

// NewRequesterGenerics creates a new instance of RequesterGenerics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequesterGenerics[TAny any, TComparable comparable, TSigned constraints.Signed, TIntf test.GetInt, TExternalIntf io.Writer, TGenIntf test.GetGeneric[TSigned], TInlineType interface{ ~int | ~uint }, TInlineTypeGeneric interface {
	~int | test.GenericType[int, test.GetInt]
	comparable
}](t interface {
	mock.TestingT
	Cleanup(func())
}) *RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	mock := &RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// RequesterGenerics is an autogenerated mock type for the RequesterGenerics type
type RequesterGenerics[TAny any, TComparable comparable, TSigned constraints.Signed, TIntf test.GetInt, TExternalIntf io.Writer, TGenIntf test.GetGeneric[TSigned], TInlineType interface{ ~int | ~uint }, TInlineTypeGeneric interface {
	~int | test.GenericType[int, test.GetInt]
	comparable
}] struct {
	mock.Mock
}

type RequesterGenerics_Expecter[TAny any, TComparable comparable, TSigned constraints.Signed, TIntf test.GetInt, TExternalIntf io.Writer, TGenIntf test.GetGeneric[TSigned], TInlineType interface{ ~int | ~uint }, TInlineTypeGeneric interface {
	~int | test.GenericType[int, test.GetInt]
	comparable
}] struct {
	mock *mock.Mock
}

func (_m *RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) EXPECT() *RequesterGenerics_Expecter[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	return &RequesterGenerics_Expecter[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]{mock: &_m.Mock}
}

// GenericAnonymousStructs provides a mock function for the type RequesterGenerics
func (_mock *RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericAnonymousStructs(val struct{ Type1 TExternalIntf }) struct {
	Type2 test.GenericType[string, test.EmbeddedGet[int]]
} {
	ret := _mock.Called(val)

	if len(ret) == 0 {
		panic("no return value specified for GenericAnonymousStructs")
	}

	var r0 struct {
		Type2 test.GenericType[string, test.EmbeddedGet[int]]
	}
	if returnFunc, ok := ret.Get(0).(func(struct{ Type1 TExternalIntf }) struct {
		Type2 test.GenericType[string, test.EmbeddedGet[int]]
	}); ok {
		r0 = returnFunc(val)
	} else {
		r0 = ret.Get(0).(struct {
			Type2 test.GenericType[string, test.EmbeddedGet[int]]
		})
	}
	return r0
}

// RequesterGenerics_GenericAnonymousStructs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenericAnonymousStructs'
type RequesterGenerics_GenericAnonymousStructs_Call[TAny any, TComparable comparable, TSigned constraints.Signed, TIntf test.GetInt, TExternalIntf io.Writer, TGenIntf test.GetGeneric[TSigned], TInlineType interface{ ~int | ~uint }, TInlineTypeGeneric interface {
	~int | test.GenericType[int, test.GetInt]
	comparable
}] struct {
	*mock.Call
}

// GenericAnonymousStructs is a helper method to define mock.On call
//   - val
func (_e *RequesterGenerics_Expecter[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericAnonymousStructs(val interface{}) *RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	return &RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]{Call: _e.mock.On("GenericAnonymousStructs", val)}
}

func (_c *RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) Run(run func(val struct{ Type1 TExternalIntf })) *RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Run(func(args mock.Arguments) {
		run(val)
	})
	return _c
}

func (_c *RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) Return(valOut struct {
	Type2 test.GenericType[string, test.EmbeddedGet[int]]
}) *RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Return(valOut)
	return _c
}

func (_c *RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) RunAndReturn(run func(val struct{ Type1 TExternalIntf }) struct {
	Type2 test.GenericType[string, test.EmbeddedGet[int]]
}) *RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Return(run)
	return _c
}

// GenericArguments provides a mock function for the type RequesterGenerics
func (_mock *RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericArguments(v1 TAny, v2 TComparable) (TSigned, TIntf) {
	ret := _mock.Called(v1, v2)

	if len(ret) == 0 {
		panic("no return value specified for GenericArguments")
	}

	var r0 TSigned
	var r1 TIntf
	if returnFunc, ok := ret.Get(0).(func(TAny, TComparable) (TSigned, TIntf)); ok {
		return returnFunc(v1, v2)
	}
	if returnFunc, ok := ret.Get(0).(func(TAny, TComparable) TSigned); ok {
		r0 = returnFunc(v1, v2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(TSigned)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(TAny, TComparable) TIntf); ok {
		r1 = returnFunc(v1, v2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(TIntf)
		}
	}
	return r0, r1
}

// RequesterGenerics_GenericArguments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenericArguments'
type RequesterGenerics_GenericArguments_Call[TAny any, TComparable comparable, TSigned constraints.Signed, TIntf test.GetInt, TExternalIntf io.Writer, TGenIntf test.GetGeneric[TSigned], TInlineType interface{ ~int | ~uint }, TInlineTypeGeneric interface {
	~int | test.GenericType[int, test.GetInt]
	comparable
}] struct {
	*mock.Call
}

// GenericArguments is a helper method to define mock.On call
//   - v1
//   - v2
func (_e *RequesterGenerics_Expecter[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericArguments(v1 interface{}, v2 interface{}) *RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	return &RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]{Call: _e.mock.On("GenericArguments", v1, v2)}
}

func (_c *RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) Run(run func(v1 TAny, v2 TComparable)) *RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Run(func(args mock.Arguments) {
		run(v1, v2)
	})
	return _c
}

func (_c *RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) Return(vOut1 TSigned, vOut2 TIntf) *RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Return(vOut1, vOut2)
	return _c
}

func (_c *RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) RunAndReturn(run func(v1 TAny, v2 TComparable) (TSigned, TIntf)) *RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Return(run)
	return _c
}

// GenericStructs provides a mock function for the type RequesterGenerics
func (_mock *RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericStructs(genericType test.GenericType[TAny, TIntf]) test.GenericType[TSigned, TIntf] {
	ret := _mock.Called(genericType)

	if len(ret) == 0 {
		panic("no return value specified for GenericStructs")
	}

	var r0 test.GenericType[TSigned, TIntf]
	if returnFunc, ok := ret.Get(0).(func(test.GenericType[TAny, TIntf]) test.GenericType[TSigned, TIntf]); ok {
		r0 = returnFunc(genericType)
	} else {
		r0 = ret.Get(0).(test.GenericType[TSigned, TIntf])
	}
	return r0
}

// RequesterGenerics_GenericStructs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenericStructs'
type RequesterGenerics_GenericStructs_Call[TAny any, TComparable comparable, TSigned constraints.Signed, TIntf test.GetInt, TExternalIntf io.Writer, TGenIntf test.GetGeneric[TSigned], TInlineType interface{ ~int | ~uint }, TInlineTypeGeneric interface {
	~int | test.GenericType[int, test.GetInt]
	comparable
}] struct {
	*mock.Call
}

// GenericStructs is a helper method to define mock.On call
//   - genericType
func (_e *RequesterGenerics_Expecter[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericStructs(genericType interface{}) *RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	return &RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]{Call: _e.mock.On("GenericStructs", genericType)}
}

func (_c *RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) Run(run func(genericType test.GenericType[TAny, TIntf])) *RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Run(func(args mock.Arguments) {
		run(genericType)
	})
	return _c
}

func (_c *RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) Return(genericTypeOut test.GenericType[TSigned, TIntf]) *RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Return(genericTypeOut)
	return _c
}

func (_c *RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) RunAndReturn(run func(genericType test.GenericType[TAny, TIntf]) test.GenericType[TSigned, TIntf]) *RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Return(run)
	return _c
}
