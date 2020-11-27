package pkg

import (
	"bufio"
	"context"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/vektra/mockery/v2/pkg/config"
)

const pkg = "test"

type GeneratorSuite struct {
	suite.Suite
	parser *Parser
	ctx    context.Context
}

func (s *GeneratorSuite) SetupTest() {
	s.parser = NewParser(nil)
	s.ctx = context.Background()
}

func (s *GeneratorSuite) getInterfaceFromFile(interfacePath, interfaceName string) *Interface {
	if !strings.Contains(interfacePath, fixturePath) {
		interfacePath = filepath.Join(fixturePath, interfacePath)
	}
	s.NoError(
		s.parser.Parse(s.ctx, interfacePath), "The parser is able to parse the given file.",
	)

	s.NoError(
		s.parser.Load(), "The parser is able to load the config.",
	)

	iface, err := s.parser.Find(interfaceName)
	s.Require().NoError(err)
	s.Require().NotNil(iface)
	return iface
}

func (s *GeneratorSuite) getGenerator(
	filepath, interfaceName string, inPackage bool, structName string,
) *Generator {
	return NewGenerator(
		s.ctx, config.Config{
			StructName: structName,
			InPackage:  inPackage,
		}, s.getInterfaceFromFile(filepath, interfaceName), pkg,
	)
}

func (s *GeneratorSuite) checkGeneration(
	filepath, interfaceName string, inPackage bool, structName string, expected string,
) *Generator {
	generator := s.getGenerator(filepath, interfaceName, inPackage, structName)
	s.NoError(generator.Generate(s.ctx), "The generator ran without errors.")

	// Mirror the formatting done by normally done by golang.org/x/tools/imports in Generator.Write.
	//
	// While we could possibly reuse Generator.Write here in addition to Generator.Generate,
	// it would require changing Write's signature to accept custom options, specifically to
	// allow the fragments in preexisting cases. It's assumed that this approximation,
	// just formatting the source, is sufficient for the needs of the current test styles.
	var actual []byte
	actual, fmtErr := format.Source(generator.buf.Bytes())
	s.NoError(fmtErr, "The formatter ran without errors.")

	// Compare lines for easier debugging via testify's slice diff output
	expectedLines := strings.Split(expected, "\n")
	actualLines := strings.Split(string(actual), "\n")

	s.Equal(
		expectedLines, actualLines,
		"The generator produced unexpected output.",
	)
	return generator
}

func (s *GeneratorSuite) checkPrologueGeneration(
	generator *Generator, expected string,
) {
	generator.GeneratePrologue(ctx, "mocks")
	s.Equal(
		expected, generator.buf.String(),
		"The generator produced an unexpected prologue.",
	)
}

func (s *GeneratorSuite) TestCalculateImport() {
	gp := []string{"a/src", "b/src"}

	s.Equal("c", calculateImport(ctx, gp, "a/src/c"))
	s.Equal("c", calculateImport(ctx, gp, "b/src/c"))
	s.Equal("d/src/c", calculateImport(ctx, gp, "d/src/c"))
}

func (s *GeneratorSuite) TestGenerator() {
	expected := `// Requester is an autogenerated mock type for the Requester type
type Requester struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *Requester) Get(path string) (string, error) {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
`
	s.checkGeneration(testFile, "Requester", false, "", expected)
}

func (s *GeneratorSuite) TestGeneratorFunction() {
	expected := `// SendFunc is an autogenerated mock type for the SendFunc type
type SendFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, data
func (_m *SendFunc) Execute(ctx context.Context, data string) (int, error) {
	ret := _m.Called(ctx, data)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "function.go"), "SendFunc", false, "",
		expected,
	)
}

func (s *GeneratorSuite) TestGeneratorSingleReturn() {
	expected := `// Requester2 is an autogenerated mock type for the Requester2 type
type Requester2 struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *Requester2) Get(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
`
	s.checkGeneration(testFile2, "Requester2", false, "", expected)
}

func (s *GeneratorSuite) TestGeneratorNoArguments() {
	expected := `// Requester3 is an autogenerated mock type for the Requester3 type
type Requester3 struct {
	mock.Mock
}

// Get provides a mock function with given fields:
func (_m *Requester3) Get() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "requester3.go"), "Requester3", false, "",
		expected,
	)
}

func (s *GeneratorSuite) TestGeneratorNoNothing() {
	expected := `// Requester4 is an autogenerated mock type for the Requester4 type
type Requester4 struct {
	mock.Mock
}

// Get provides a mock function with given fields:
func (_m *Requester4) Get() {
	_m.Called()
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "requester4.go"), "Requester4", false, "",
		expected,
	)
}

func (s *GeneratorSuite) TestGeneratorUnexported() {
	expected := `// mockRequester_unexported is an autogenerated mock type for the requester_unexported type
type mockRequester_unexported struct {
	mock.Mock
}

// Get provides a mock function with given fields:
func (_m *mockRequester_unexported) Get() {
	_m.Called()
}
`
	s.checkGeneration(
		"requester_unexported.go", "requester_unexported", true, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorPrologue() {
	generator := s.getGenerator(testFile, "Requester", false, "")
	expected := `package mocks

import mock "github.com/stretchr/testify/mock"
import test "github.com/vektra/mockery/v2/pkg/fixtures"

`
	s.checkPrologueGeneration(generator, expected)
}

func (s *GeneratorSuite) TestGeneratorPrologueWithImports() {
	generator := s.getGenerator("requester_ns.go", "RequesterNS", false, "")
	expected := `package mocks

import http "net/http"
import mock "github.com/stretchr/testify/mock"
import test "github.com/vektra/mockery/v2/pkg/fixtures"

`
	s.checkPrologueGeneration(generator, expected)
}

func (s *GeneratorSuite) TestGeneratorPrologueWithMultipleImportsSameName() {
	generator := s.getGenerator("same_name_imports.go", "Example", false, "")

	expected := `package mocks

import fixtureshttp "github.com/vektra/mockery/v2/pkg/fixtures/http"
import http "net/http"
import mock "github.com/stretchr/testify/mock"
import test "github.com/vektra/mockery/v2/pkg/fixtures"

`
	s.checkPrologueGeneration(generator, expected)
}

func (s *GeneratorSuite) TestGeneratorPrologueNote() {
	generator := s.getGenerator(testFile, "Requester", false, "")
	generator.GeneratePrologueNote("A\\nB")

	expected := `// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

// A
// B

`

	s.Equal(expected, generator.buf.String())
}

func (s *GeneratorSuite) TestGeneratorPrologueNoteBlockComment() {
	generator := s.getGenerator(testFile, "Requester", false, "")
	generator.GeneratePrologueNote("/*\n    BOILERPLATE\n*/")

	expected := `/*
    BOILERPLATE
*/

// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

`

	s.Equal(expected, generator.buf.String())
}

func (s *GeneratorSuite) TestGeneratorPrologueNoteNoVersionString() {
	generator := s.getGenerator(testFile, "Requester", false, "")
	generator.Config.DisableVersionString = true
	generator.GeneratePrologueNote("A\\nB")

	expected := `// Code generated by mockery. DO NOT EDIT.

// A
// B

`

	s.Equal(expected, generator.buf.String())
}

func (s *GeneratorSuite) TestVersionOnCorrectLine() {
	gen := s.getGenerator(testFile, "Requester", false, "")

	//Run everything that is ran by the GeneratorVisitor
	gen.GeneratePrologueNote("A\\nB")
	gen.GeneratePrologue(s.ctx, pkg)
	err := gen.Generate(s.ctx)

	require.NoError(s.T(), err)
	scan := bufio.NewScanner(&gen.buf)
	s.Contains("Code generated by", scan.Text())
}

func (s *GeneratorSuite) TestGeneratorChecksInterfacesForNilable() {
	expected := `// RequesterIface is an autogenerated mock type for the RequesterIface type
type RequesterIface struct {
	mock.Mock
}

// Get provides a mock function with given fields:
func (_m *RequesterIface) Get() io.Reader {
	ret := _m.Called()

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func() io.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	return r0
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "requester_iface.go"), "RequesterIface",
		false, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorPointers() {
	expected := `// RequesterPtr is an autogenerated mock type for the RequesterPtr type
type RequesterPtr struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *RequesterPtr) Get(path string) (*string, error) {
	ret := _m.Called(path)

	var r0 *string
	if rf, ok := ret.Get(0).(func(string) *string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "requester_ptr.go"), "RequesterPtr", false, "",
		expected,
	)
}

func (s *GeneratorSuite) TestGeneratorSlice() {
	expected := `// RequesterSlice is an autogenerated mock type for the RequesterSlice type
type RequesterSlice struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *RequesterSlice) Get(path string) ([]string, error) {
	ret := _m.Called(path)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "requester_slice.go"), "RequesterSlice",
		false, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorArrayLiteralLen() {
	expected := `// RequesterArray is an autogenerated mock type for the RequesterArray type
type RequesterArray struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *RequesterArray) Get(path string) ([2]string, error) {
	ret := _m.Called(path)

	var r0 [2]string
	if rf, ok := ret.Get(0).(func(string) [2]string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([2]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "requester_array.go"), "RequesterArray",
		false, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorNamespacedTypes() {
	expected := `// RequesterNS is an autogenerated mock type for the RequesterNS type
type RequesterNS struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *RequesterNS) Get(path string) (http.Response, error) {
	ret := _m.Called(path)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(string) http.Response); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(http.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "requester_ns.go"), "RequesterNS", false, "",
		expected,
	)
}

func (s *GeneratorSuite) TestGeneratorWhereArgumentNameConflictsWithImport() {
	expected := `// RequesterArgSameAsImport is an autogenerated mock type for the RequesterArgSameAsImport type
type RequesterArgSameAsImport struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *RequesterArgSameAsImport) Get(_a0 string) *json.RawMessage {
	ret := _m.Called(_a0)

	var r0 *json.RawMessage
	if rf, ok := ret.Get(0).(func(string) *json.RawMessage); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*json.RawMessage)
		}
	}

	return r0
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "requester_arg_same_as_import.go"),
		"RequesterArgSameAsImport", false, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorWhereArgumentNameConflictsWithNamedImport() {
	expected := `// RequesterArgSameAsNamedImport is an autogenerated mock type for the RequesterArgSameAsNamedImport type
type RequesterArgSameAsNamedImport struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *RequesterArgSameAsNamedImport) Get(_a0 string) *json.RawMessage {
	ret := _m.Called(_a0)

	var r0 *json.RawMessage
	if rf, ok := ret.Get(0).(func(string) *json.RawMessage); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*json.RawMessage)
		}
	}

	return r0
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "requester_arg_same_as_named_import.go"),
		"RequesterArgSameAsNamedImport", false, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorWhereArgumentNameConflictsWithPackage() {
	expected := `// RequesterArgSameAsPkg is an autogenerated mock type for the RequesterArgSameAsPkg type
type RequesterArgSameAsPkg struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *RequesterArgSameAsPkg) Get(_a0 string) {
	_m.Called(_a0)
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "requester_arg_same_as_pkg.go"),
		"RequesterArgSameAsPkg", false, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorHavingNoNamesOnArguments() {
	expected := `// KeyManager is an autogenerated mock type for the KeyManager type
type KeyManager struct {
	mock.Mock
}

// GetKey provides a mock function with given fields: _a0, _a1
func (_m *KeyManager) GetKey(_a0 string, _a1 uint16) ([]byte, *test.Err) {
	ret := _m.Called(_a0, _a1)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, uint16) []byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 *test.Err
	if rf, ok := ret.Get(1).(func(string, uint16) *test.Err); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*test.Err)
		}
	}

	return r0, r1
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "custom_error.go"), "KeyManager", false, "",
		expected,
	)
}

func (s *GeneratorSuite) TestGeneratorElidedType() {
	expected := `// RequesterElided is an autogenerated mock type for the RequesterElided type
type RequesterElided struct {
	mock.Mock
}

// Get provides a mock function with given fields: path, url
func (_m *RequesterElided) Get(path string, url string) error {
	ret := _m.Called(path, url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(path, url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "requester_elided.go"), "RequesterElided",
		false, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorReturnElidedType() {
	expected := `// RequesterReturnElided is an autogenerated mock type for the RequesterReturnElided type
type RequesterReturnElided struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *RequesterReturnElided) Get(path string) (int, int, int, error) {
	ret := _m.Called(path)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(string) int); ok {
		r2 = rf(path)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(string) error); ok {
		r3 = rf(path)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "requester_ret_elided.go"),
		"RequesterReturnElided", false, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorVariadicArgs() {
	// Read the expected output from a "golden" file that we can also import in CompatSuite
	// to asserts its actual behavior.
	//
	// To regenerate the golden file: make fixture
	expectedBytes, err := ioutil.ReadFile(filepath.Join(fixturePath, "mocks", "requester_variadic.go"))
	s.NoError(err)
	expected := string(expectedBytes)
	expected = expected[strings.Index(expected, "// RequesterVariadic is"):]
	s.checkGeneration(
		filepath.Join(fixturePath, "requester_variadic.go"),
		"RequesterVariadic", false, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorArgumentIsFuncType() {
	expected := `// Fooer is an autogenerated mock type for the Fooer type
type Fooer struct {
	mock.Mock
}

// Bar provides a mock function with given fields: f
func (_m *Fooer) Bar(f func([]int)) {
	_m.Called(f)
}

// Baz provides a mock function with given fields: path
func (_m *Fooer) Baz(path string) func(string) string {
	ret := _m.Called(path)

	var r0 func(string) string
	if rf, ok := ret.Get(0).(func(string) func(string) string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(string) string)
		}
	}

	return r0
}

// Foo provides a mock function with given fields: f
func (_m *Fooer) Foo(f func(string) string) error {
	ret := _m.Called(f)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(string) string) error); ok {
		r0 = rf(f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "func_type.go"), "Fooer", false, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorChanType() {
	expected := `// AsyncProducer is an autogenerated mock type for the AsyncProducer type
type AsyncProducer struct {
	mock.Mock
}

// Input provides a mock function with given fields:
func (_m *AsyncProducer) Input() chan<- bool {
	ret := _m.Called()

	var r0 chan<- bool
	if rf, ok := ret.Get(0).(func() chan<- bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan<- bool)
		}
	}

	return r0
}

// Output provides a mock function with given fields:
func (_m *AsyncProducer) Output() <-chan bool {
	ret := _m.Called()

	var r0 <-chan bool
	if rf, ok := ret.Get(0).(func() <-chan bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan bool)
		}
	}

	return r0
}

// Whatever provides a mock function with given fields:
func (_m *AsyncProducer) Whatever() chan bool {
	ret := _m.Called()

	var r0 chan bool
	if rf, ok := ret.Get(0).(func() chan bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan bool)
		}
	}

	return r0
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "async.go"), "AsyncProducer", false, "",
		expected,
	)
}

func (s *GeneratorSuite) TestGeneratorFromImport() {
	expected := `// MyReader is an autogenerated mock type for the MyReader type
type MyReader struct {
	mock.Mock
}

// Read provides a mock function with given fields: p
func (_m *MyReader) Read(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "io_import.go"), "MyReader", false, "",
		expected,
	)
}

func (s *GeneratorSuite) TestGeneratorComplexChanFromConsul() {
	expected := `// ConsulLock is an autogenerated mock type for the ConsulLock type
type ConsulLock struct {
	mock.Mock
}

// Lock provides a mock function with given fields: _a0
func (_m *ConsulLock) Lock(_a0 <-chan struct{}) (<-chan struct{}, error) {
	ret := _m.Called(_a0)

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func(<-chan struct{}) <-chan struct{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(<-chan struct{}) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unlock provides a mock function with given fields:
func (_m *ConsulLock) Unlock() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "consul.go"), "ConsulLock", false, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorForEmptyInterface() {
	expected := `// Blank is an autogenerated mock type for the Blank type
type Blank struct {
	mock.Mock
}

// Create provides a mock function with given fields: x
func (_m *Blank) Create(x interface{}) error {
	ret := _m.Called(x)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(x)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "empty_interface.go"), "Blank", false, "",
		expected,
	)
}

func (s *GeneratorSuite) TestGeneratorArgumentIsMapFunc() {
	expected := `// MapFunc is an autogenerated mock type for the MapFunc type
type MapFunc struct {
	mock.Mock
}

// Get provides a mock function with given fields: m
func (_m *MapFunc) Get(m map[string]func(string) string) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]func(string) string) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "map_func.go"), "MapFunc", false, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorForMethodUsingInterface() {
	expected := `// UsesOtherPkgIface is an autogenerated mock type for the UsesOtherPkgIface type
type UsesOtherPkgIface struct {
	mock.Mock
}

// DoSomethingElse provides a mock function with given fields: obj
func (_m *UsesOtherPkgIface) DoSomethingElse(obj test.Sibling) {
	_m.Called(obj)
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "mock_method_uses_pkg_iface.go"),
		"UsesOtherPkgIface", false, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorForMethodUsingInterfaceInPackage() {
	expected := `// MockUsesOtherPkgIface is an autogenerated mock type for the UsesOtherPkgIface type
type MockUsesOtherPkgIface struct {
	mock.Mock
}

// DoSomethingElse provides a mock function with given fields: obj
func (_m *MockUsesOtherPkgIface) DoSomethingElse(obj Sibling) {
	_m.Called(obj)
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "mock_method_uses_pkg_iface.go"),
		"UsesOtherPkgIface", true, "", expected,
	)
}

func (s *GeneratorSuite) TestGeneratorWithAliasing() {
	expected := `// Example is an autogenerated mock type for the Example type
type Example struct {
	mock.Mock
}

// A provides a mock function with given fields:
func (_m *Example) A() http.Flusher {
	ret := _m.Called()

	var r0 http.Flusher
	if rf, ok := ret.Get(0).(func() http.Flusher); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Flusher)
		}
	}

	return r0
}

// B provides a mock function with given fields: _a0
func (_m *Example) B(_a0 string) fixtureshttp.MyStruct {
	ret := _m.Called(_a0)

	var r0 fixtureshttp.MyStruct
	if rf, ok := ret.Get(0).(func(string) fixtureshttp.MyStruct); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(fixtureshttp.MyStruct)
	}

	return r0
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "same_name_imports.go"), "Example", false, "",
		expected,
	)
}

func (s *GeneratorSuite) TestGeneratorWithImportSameAsLocalPackageInpkgNoCycle() {
	iface := s.getInterfaceFromFile("imports_same_as_package.go", "ImportsSameAsPackage")
	pkg := iface.QualifiedName
	gen := NewGenerator(s.ctx, config.Config{
		InPackage: true,
	}, iface, pkg)
	gen.GeneratePrologue(s.ctx, pkg)
	s.NotContains(gen.buf.String(), `import test "github.com/vektra/mockery/v2/pkg/fixtures/test"`)
}

func (s *GeneratorSuite) TestMapToInterface() {
	expected := `// MapToInterface is an autogenerated mock type for the MapToInterface type
type MapToInterface struct {
	mock.Mock
}

// Foo provides a mock function with given fields: arg1
func (_m *MapToInterface) Foo(arg1 ...map[string]interface{}) {
	_va := make([]interface{}, len(arg1))
	for _i := range arg1 {
		_va[_i] = arg1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}
`
	s.checkGeneration(
		"MapToInterface.go", "MapToInterface", false, "",
		expected,
	)

}

func (s *GeneratorSuite) TestGeneratorWithImportSameAsLocalPackage() {
	expected := `// ImportsSameAsPackage is an autogenerated mock type for the ImportsSameAsPackage type
type ImportsSameAsPackage struct {
	mock.Mock
}

// A provides a mock function with given fields:
func (_m *ImportsSameAsPackage) A() test.B {
	ret := _m.Called()

	var r0 test.B
	if rf, ok := ret.Get(0).(func() test.B); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(test.B)
	}

	return r0
}

// B provides a mock function with given fields:
func (_m *ImportsSameAsPackage) B() fixtures.KeyManager {
	ret := _m.Called()

	var r0 fixtures.KeyManager
	if rf, ok := ret.Get(0).(func() fixtures.KeyManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fixtures.KeyManager)
		}
	}

	return r0
}

// C provides a mock function with given fields: _a0
func (_m *ImportsSameAsPackage) C(_a0 fixtures.C) {
	_m.Called(_a0)
}
`
	s.checkGeneration(
		"imports_same_as_package.go", "ImportsSameAsPackage", false, "",
		expected,
	)
}

func (s *GeneratorSuite) TestPrologueWithImportSameAsLocalPackage() {
	generator := s.getGenerator(
		"imports_same_as_package.go", "ImportsSameAsPackage", false, "",
	)
	expected := `package mocks

import fixtures "` + generator.iface.QualifiedName + `"
import mock "github.com/stretchr/testify/mock"
import test "github.com/vektra/mockery/v2/pkg/fixtures/test"

`
	s.checkPrologueGeneration(generator, expected)
}

func (s *GeneratorSuite) TestPrologueWithImportFromNestedInterface() {
	generator := s.getGenerator(
		"imports_from_nested_interface.go", "HasConflictingNestedImports", false, "",
	)
	expected := `package mocks

import fixtureshttp "github.com/vektra/mockery/v2/pkg/fixtures/http"
import http "net/http"
import mock "github.com/stretchr/testify/mock"
import test "github.com/vektra/mockery/v2/pkg/fixtures"

`

	s.checkPrologueGeneration(generator, expected)
}

func (s *GeneratorSuite) TestGeneratorForStructValueReturn() {
	expected := `// A is an autogenerated mock type for the A type
type A struct {
	mock.Mock
}

// Call provides a mock function with given fields:
func (_m *A) Call() (test.B, error) {
	ret := _m.Called()

	var r0 test.B
	if rf, ok := ret.Get(0).(func() test.B); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(test.B)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
`
	s.checkGeneration(
		filepath.Join(fixturePath, "struct_value.go"), "A", false, "",
		expected,
	)
}

func (s *GeneratorSuite) TestStructNameOverride() {
	expected := `// Requester2OverrideName is an autogenerated mock type for the Requester2 type
type Requester2OverrideName struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *Requester2OverrideName) Get(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
`
	s.checkGeneration(testFile2, "Requester2", false, "Requester2OverrideName", expected)
}

func (s *GeneratorSuite) TestKeepTreeInPackageCombined() {
	type testData struct {
		path     string
		name     string
		expected string
	}

	tests := []testData{
		{path: "example_project/root.go", name: "Root", expected: `package example_project

import example_project "github.com/vektra/mockery/v2/pkg/fixtures/example_project"
import foo "github.com/vektra/mockery/v2/pkg/fixtures/example_project/foo"
import mock "github.com/stretchr/testify/mock"

`},
		{path: "example_project/foo/foo.go", name: "Foo", expected: `package foo

import foo "github.com/vektra/mockery/v2/pkg/fixtures/example_project/foo"
import mock "github.com/stretchr/testify/mock"

`},
	}

	for _, test := range tests {
		generator := NewGenerator(
			s.ctx,
			config.Config{InPackage: true, KeepTree: true},
			s.getInterfaceFromFile(test.path, test.name),
			pkg,
		)
		s.checkPrologueGeneration(generator, test.expected)
	}
}

func TestGeneratorSuite(t *testing.T) {
	generatorSuite := new(GeneratorSuite)
	suite.Run(t, generatorSuite)
}
