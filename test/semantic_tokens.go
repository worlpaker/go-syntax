// SYNTAX TEST "source.go" "Semantic Tokens Test"

package main

import (
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/build"
	"go/token"
	"io"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// functions

func Foo(foo string) FooX

func Fo0o(foo string) *[]FooX

func (m *FooX) Foo(foo string) FooX

func (m *FooX) Bar(bar any) func()

func Foo0(foo string) func(http.Handler) http.Handler

func Foo1[t FooX](foo t) {}

func Foo2(string) <-chan context.Context

func Foo02(a <-chan context.Context) <-chan context.Context { return a }

func Foo3(any) FooX

func Foo5(foo map[FooX]interface{}, bar any) {}

func Bar0(func(context.Context))

func Bar01(http.ResponseWriter, *http.Request) {}

func Bar02(w http.ResponseWriter, r *http.Request) {}

func Bar(context.Context, func(), context.Context) {
}

func Zoo(foo ...func(Context) context.Context)

func Chain(middlewares ...func(http.Handler) http.Handler)

func Z00(m ...func(http.Handler))

func Z000(m func(http.Handler))

func chain(middlewares []func(http.Handler) http.Handler, endpoint http.Handler)

func Bar2(context.Context, func(string, context.Context), io.Writer) (foo context.Context, bar string)

func Bar3(foox context.Context, barx func(foo context.Context, bar bool), foobarx io.Writer) (foo context.Context, bar string)

func Bar4(context.Context, func(foo context.Context, bar string)) (foo context.Context, bar string)

func Bar5() *Foo11 {
	return &Foo11{name: "Foo", age: 99}
}

func Bar55() *Foo11 {
	return &Foo11{
		name: "Foo",
		age:  100,
	}
}

func Bar6(w http.ResponseWriter, r *http.Request) context.Context {
	w.Write([]byte("Foo"))
	var a context.Context
	return a
}

func Bar7[T *FooX | string | *context.Context | int](foo T, bar T) {

}

func Bar8[T FooX](foo []*T, bar []T) {

}

func Bar9[
	T FooX, F *context.Context, X Foo15](foo, bar, foobar T) {

}

func Bar10[T FooX](foo []*T, bar T) {

}

func Bar11[
	a,
	b,
	c int, d, e, f FooX](foo a, bar f) {

}
func Bar12[a, b, c, t, z int, aa, bb, cc string](foo t, bar z) {

}

func Bar13(context.Context) {

}

func Bar14(string) {

}

func (*Foo11) Done() map[string]interface{} {
	return nil
}

func Bar15(
	a,
	b,
	c int) {
}

func Bar16(
	context.Context,
	string,
	context.Context) {
}

func Bar17(*context.Context, *context.Context, *context.Context) {
}

func Bar18(context.Context, FooX, context.Context) {
}

func Bar19(
	a, b,
	f context.Context, c, d context.Context) (context.Context, context.Context) {
	return a, b
}

func Bar20(
	a map[string]interface{},
	b, f context.Context, c, d context.Context) (foo map[string]interface{}, bar context.Context) {
	return a, b
}

func Bar21(a string, b context.Context) {}
func Bar22(ch <-chan context.Context) {

}
func Bar23(foo map[string]interface{}) {
}

func Bar24(
	a context.Context, b <-chan string,
	c context.Context) {
}

func Bar25(
	a chan<- context.Context, b <-chan struct{}, c context.Context) {
}

func Bar26(
	a, b, c <-chan context.Context) {
}

func Bar27(chan string, chan string, chan context.Context) {
}

func Bar28(
	a,
	b,
	c context.Context) {
}

func Bar29() {}

func Bar30(context.Context, map[Foo17]interface{},
	context.Context, func(), io.ReadCloser, io.ReadCloser,
	string, func(context.Context), *io.Writer, *io.Reader,
	error, map[Foo17]interface{}) {
}

func (m *FooX) Bar31(foo string) {

}

func Bar32(
	context.Context, context.Context,
	context.Context, io.ReadCloser,
	error, map[Foo17]interface{}) {

}

func Bar33(
	context.Context, func(), context.Context, io.ReadCloser, error, map[Foo17]interface{}) {

}

func Bar33_2(context.Context, func(), context.Context, io.ReadCloser, error, map[Foo17]interface{}) {

}

func Bar34(
	a, b, c context.Context) (foo context.Context, bar error) {

	return a, nil
}

func Bar35(

	foo context.Context, bar context.Context) {
}

func Bar36(context.Context, context.Context) {
}

func Bar37(bar string) {}

func (s *Foo15) Bar38(a chan<- int) {

}
func Bar39(context.Context, context.Context, string)

func (c *cancelCtx) Value(key any) any {
	if key == &cancelCtxKey {
		return c
	}
	return value(c.Context, key)
}
func Bar40(
	a map[string]interface{}, b <-chan *context.Context, c context.Context) context.Context {
	return c
}

func Bar41(
	a *context.Context,
	b string,
	c context.Context) (context.Context, string, context.Context) {
	return *a, b, c
}

func value(c Context, key any) any {
	for {
		switch ctx := c.(type) {
		case *cancelCtx:
			fmt.Println(ctx)
		default:
			return c.Value(key)
		}
	}
}

func WithValue(parent Context, key, val any) Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	if key == nil {
		panic("nil key")
	}
	return parent
}

func parentCancelCtx(parent Context) (*cancelCtx, bool) {
	done := parent.Done()
	if done == closedchan || done == nil {
		return nil, false
	}
	p, ok := parent.Value(&cancelCtxKey).(*cancelCtx)
	if !ok {
		return nil, false
	}
	pdone, _ := p.done.Load().(chan struct{})
	if pdone != done {
		return nil, false
	}
	return p, true
}

// var and const

const (
	CMDSet Command = "FOO"
	CMD            = "FOO"
)

var (
	CMDSet1 Command = "FOO"
	CMD1            = "FOO"
)

const (
	Fooo = "FOO"
	Barr = 911
)

var (
	background = new(Foo18)
	todo       = new(Foo18)
)

const CMDSet3 Command = "FOO"

const ntCatchAll nodeTyp = 3

var closedchan = make(chan struct{})

var foo, bar Foo17
var foo2 *sync.Mutex
var foo3 = "bar"
var foo4, bar4 FooX
var bar2 Command
var bar3 []*Command
var cancelCtxKey int
var FO FooX
var BO = 15

// types, structs and interfaces

type Command string

type Context context.Context

type ZooX <-chan context.Context

type Middlewares []func(http.Handler) http.Handler

type Middleware []func(http.Handler) <-chan http.Handler

type Foo4 *sync.Mutex

type Foo6 interface{ FooX }

type Foo06 interface{ *[]FooX }

type Foo07 struct {
	FooX
}

type Foo7 interface {
	[]FooX
	Context
}

type Foo8 interface{ FooX | *context.Context }

type Foo9 interface {
	FooX | Foo11
}

type Foo10 interface {
	FooX | <-chan *context.Context
}

type Foo12 interface {
	chan FooX | chan Foo11
}

type Foo13 interface {
	Bar(string, func(context.Context))
}

type Foo14 interface {
	Bar(w http.ResponseWriter, r *http.Request) (foo context.Context, bar string)
	Bar2(http.ResponseWriter, *http.Request) (context.Context, string)
	Bar3(string, func(bar context.Context)) (foo context.Context, bar string)
	Bar4(foo *context.Context, bar func(context.Context)) (foox context.Context, barx string)
	Bar5(context.Context, func(foo chan<- context.Context, bar string), *io.Reader) (foo context.Context, bar string)
	Bar6(context.Context, func(*context.Context, string)) (*context.Context, string)
	Bar7(foo *context.Context) (context.Context, context.Context)
	Bar8(foo string) context.Context
	Bar9(foo func()) <-chan context.Context
	Bar10() (context.Context, error)
	Bar11(foox any) (foo context.Context, bar context.Context)
	Bar12(foo map[FooX]interface{}) (context.Context, context.Context)
	Bar13() *map[FooX]interface{}
	Bar14() []FooX
}

type Foo18 interface {
	Foo(a string) func(a Foo11)
	Bar(a bool) (err error)
	FooBar(a context.Context) error
}

type Foo11 struct {
	name string
	age  int
}

type Foo19 *context.Context

type Foo20 struct {
	Foo15
	Foo17
	Foo       Foo16
	FFoo, Bar string
}

type Foo21 struct{ foo, bar Foo16 }

type Foo22 context.Context

type Foo23 map[string]interface{}

type Foo24 struct{ foo, bar FooX }

type Foo15 struct {
	foo, bar string
}

type Foo16 struct{}

type Foo17 struct {
	name string
	age  int
}

type nodeTyp uint8

type Foo25 struct {
	Foo15
	*Foo17
	Foo    func() (context.Context, func(), io.ReadCloser, error, map[Foo17]interface{})
	Bar    func() (foo context.Context, bar io.ReadCloser, err error)
	Fooo   Foo17
	Barr   string
	Foo2   map[Foo17]interface{}
	Bar2   func(foo FooX) (foox, barx, foobar io.Writer)
	FooBar *Foo17
	Foo3   func() (foo context.Context, bar io.ReadCloser, foobar *map[Foo17]interface{})
	Bar3   func() map[Foo17]interface{}
	Foo4   chan<- Foo17
	Bar4   <-chan Foo17
	Foo5   chan *Foo17
	Bar5   chan Foo17
	Foo6   []func(http.Handler) http.Handler
	Bar6   func() context.Context
	Foo7   []func(http.Handler) <-chan http.Handler
	Bar7   []nodeTyp
	Foo8   [ntCatchAll + 1]nodeTyp
	Bar8   []*nodeTyp
}

type Foo26 struct{ Foo15 }

type Foo27 struct{ foo <-chan Foo15 }

type Account struct {
	context.Context
	*Foo17
	Email      context.Context `json:"Email"`
	Password   string          `json:"Password"`
	First_name string          `json:"First_name"`
	Last_name  string          `json:"Last_name"`
	Birth_date *int            `json:"Birth_date"`
	Gender     *string         `json:"Gender"`
}

type Foo28 string

type Foo29 interface {
	Foo28 | sync.Mutex | int
}

type Foo30 interface {
	string | int64 | int32 | []byte
}

type Foo31 interface {
	*string | int64 | int32 | []*byte
}

type Foo32 struct {
	cancelCtx // foo
	timerCtx  /* bar */
}

type Foo33 interface {
	cancelCtx // foo
	timerCtx  /* bar */
}

type FooY int

type FooX struct{}

type timerCtx struct {
	cancelCtx
	timer *time.Timer // Under cancelCtx.mu.

	deadline time.Time
}

// A cancelCtx can be canceled. When canceled, it also cancels any children
// that implement canceler.
type cancelCtx struct {
	Context

	mu       sync.Mutex        // protects following fields
	done     atomic.Value      // of chan struct{}, created lazily, closed by first cancel call
	children map[FooX]struct{} // set to nil by the first cancel call
	err      error             // set to non-nil by the first cancel call
}

// function body
func main() {
	var foox map[FooX]string
	var fooy map[FooX]interface{}
	bar := 15
	fmt.Printf("foo %d", bar)
	foo := "bar"
	fmt.Println(foo, foox, fooy)
	time.Now().Weekday()
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case *cancelCtx:
			fmt.Println("I'm a cancelCtx")
		case int, string:
			fmt.Println("I'm an int")
		case FooX:
			fmt.Println("I'm a FooX")
		case bool:
			fmt.Println("I'm a bool")
		case Foo17, Foo16:
			fmt.Println("I'm a Foo17, Foo16")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	fmt.Println("Hello World!")
	data := []string{"foo", "bar"}

	f := &Foo17{
		name: "foo",
		age:  28}
	f2 := Foo17{name: "foo", age: 28}

	whatAmI(true)
	whatAmI(3222)
	whatAmI(foo)
	whatAmI(bar)
	whatAmI("hello world!")
	whatAmI(f2)

	foo2 := func(a string) {
		fmt.Println(a)
	}
	foo2(foo)

	fmt.Println(f)
	fmt.Println(f2)

	for _, i := range data {
		log.Printf("your data is %s", i)
	}
	for i := 0; i <= 10; i++ {
	}
	var i context.Context
	defer func(i context.Context) {
		if recover() == nil {
			fmt.Println("expected panic()")
		}
	}(i)

	mw := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
	var m http.Handler
	mw(m)

	type foo34 context.Context //v0.1.1

	//v0.2.1
	if foo == "foo" {
		fmt.Println("foo")
	} else if foo == "bar" {
		fmt.Println("bar")
	} else {
		fmt.Println("foobar")
	}

	//v0.2.2
	foo022 := func() {
		for {
			switch {
			}
			select {}
		}
	}
	foo022()

	//v0.2.4
	type Spec interface {
		[]string | string
	}

	type TestModel[TGen Spec] struct {
		a int
		b TGen
	}

	type TestModel2[foo Spec, bar FooX] struct {
		a int
		b foo
		c *bar
	}

	type Foo30 []func(http.Handler) <-chan http.Handler

	type Pointer[foo any] struct {
		_ [0]*foo
	}

	type Pointers struct {
		foo int
	}

	var _ = &Pointer[Pointers]{}

	type response context.Context

	type Pointerx struct {
		foo    int
		bar    *Pointer[*response]
		foobar Pointer[response]
		foo1   Pointer[int]
		bar1   map[response]interface{}
	}

}

// v0.2.5
type TestStorageTypes struct{}

func (t TestStorageTypes) TestingStorageTypes() error {
	t.normalName()
	t.bool()
	t.byte()
	t.error()
	t.complex64()
	t.complex128()
	t.float32()
	t.float64()
	t.int()
	t.int8()
	t.int16()
	t.int32()
	t.int64()
	t.uint()
	t.uint8()
	t.uint16()
	t.uint32()
	t.uint64()
	t.rune()
	t.string()
	t.uintptr()

	return nil
}

func (t TestStorageTypes) normalName() {}

func (t TestStorageTypes) bool() {}

func (t TestStorageTypes) byte() {}

func (t TestStorageTypes) error() {}

func (t TestStorageTypes) complex64() {}

func (t TestStorageTypes) complex128() {}

func (t TestStorageTypes) float32() {}

func (t TestStorageTypes) float64() {}

func (t TestStorageTypes) int() {}

func (t TestStorageTypes) int8() {}

func (t TestStorageTypes) int16() {}

func (t TestStorageTypes) int32() {}

func (t TestStorageTypes) int64() {}

func (t TestStorageTypes) uint() {}

func (t TestStorageTypes) uint8() {}

func (t TestStorageTypes) uint16() {}

func (t TestStorageTypes) uint32() {}

func (t TestStorageTypes) uint64() {}

func (t TestStorageTypes) rune() {}

func (t TestStorageTypes) string() {}

func (t TestStorageTypes) uintptr() {}

func (t TestStorageTypes) any() {}

// v0.2.6
func Bar42(w http.ResponseWriter, r *http.Request) context.Context {
	w.Write([]byte("Foo"))
	var a context.Context
	return a
}

func Bar43(w http.ResponseWriter, HTTPStatusCode int) error {
	w.WriteHeader(HTTPStatusCode)
	return json.NewEncoder(w).Encode(
		map[string]interface{}{
			"status": HTTPStatusCode,
			"error":  http.StatusText(HTTPStatusCode),
		})
}

func DecrementInLoop() {
	j := 10
	for i := 0; i < 10; i++ {
	}
	for i := 0; i < 10; i-- {
	}
	j++
	j--
}

// v0.2.7
func TypesSupportFunction() {
	var t1 bool
	var t2 byte
	var t3 error
	var t4 complex64
	var t5 complex128
	var t6 float32
	var t7 float64
	var t8 int
	var t9 int8
	var t10 int16
	var t11 int32
	var t12 int64
	var t13 uint
	var t14 uint8
	var t15 uint16
	var t16 uint32
	var t17 uint64
	var t18 rune
	var t19 string
	var t20 uintptr
	a1 := bool(t1)
	a2 := byte(t2)
	a3 := error(t3)
	a4 := complex64(t4)
	a5 := complex128(t5)
	a6 := float32(t6)
	a7 := float64(t7)
	a8 := int(t8)
	a9 := int8(t9)
	a10 := int16(t10)
	a11 := int32(t11)
	a12 := int64(t12)
	a13 := uint(t13)
	a14 := uint8(t14)
	a15 := uint16(t15)
	a16 := uint32(t16)
	a17 := uint64(t17)
	a18 := rune(t18)
	a19 := string(t19)
	a20 := uintptr(t20)
	fmt.Sprintln(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20)
}

// v0.2.8
func Bar44(context.Context, context.Context)

func Bar45(foo, bar,
	foobar context.Context)

func Bar46(foo, bar,
	foobar chan<- context.Context)

func Bar47(
	context.Context, context.Context,
	context.Context)

func Bar48(chan context.Context, context.Context,
	string, context.Context)

func Bar49(<-chan context.Context, context.Context,
	string, chan<- context.Context)

func Bar50(context.Context, <-chan context.Context,
	string, chan<- context.Context)

func Bar51(context.Context, <-chan func() context.Context,
	string, chan<- context.Context)

func Bar52(func() <-chan context.Context, context.Context,
	string, chan<- context.Context)

func Bar53(chan<- func() context.Context, func() context.Context,
	string, chan<- context.Context)

func Bar54(context.Context, func(), map[Foo17]interface{},
	context.Context, func(), io.ReadCloser, io.ReadCloser, context.Context,
	string, func(context.Context), *io.Writer, *io.Reader,
	context.Context, context.Context, context.Context,
	string, chan func(context.Context), bool, io.Reader,
	string, <-chan func(context.Context), bool, chan<- io.Reader,
	error, map[Foo17]interface{}, context.Context) {
}

func Bar56(context.Context, func(), map[Foo17]interface{}, context.Context, func(), io.ReadCloser, io.ReadCloser, context.Context, string, func(context.Context), *io.Writer, *io.Reader, context.Context, context.Context, context.Context, string, chan func(context.Context), bool, io.Reader, string, <-chan func(context.Context), bool, chan<- io.Reader, error, map[Foo17]interface{}, context.Context) {
}
func Bar57(context.Context, func(), <-chan context.Context)

func Bar58(<-chan context.Context, chan func(), <-chan context.Context)

func Bar59(<-chan context.Context, func(), chan<- func(context.Context), <-chan context.Context)

func Bar60(context.CancelCauseFunc, string)

func Bar61(foo context.Context, bar context.Context) {}

func Bar62(foo, bar, foobar context.Context)

func Bar63(
	foo, bar, foobar context.Context) {
}

func Bar64(
	context.CancelCauseFunc, context.CancelFunc, context.Context) {
}

func Bar65(
	context.CancelCauseFunc,
	context.CancelFunc,
	context.Context) {
}

func Bar66(context.Context, bool, byte, error, error, complex64, complex128, float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, rune, string, uintptr, context.Context, context.Context, bool, byte, complex64, complex128, float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, rune, string, uintptr) {
}

func Bar67() (context.Context, bool, byte, error, complex64, complex128, float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, rune, string, uintptr, context.Context) {
	var t0 context.Context
	var t1 bool
	var t2 byte
	var t3 error
	var t4 complex64
	var t5 complex128
	var t6 float32
	var t7 float64
	var t8 int
	var t9 int8
	var t10 int16
	var t11 int32
	var t12 int64
	var t13 uint
	var t14 uint8
	var t15 uint16
	var t16 uint32
	var t17 uint64
	var t18 rune
	var t19 string
	var t20 uintptr
	var t21 context.Context
	return t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21
}

func Bar68(t0 context.Context, t1 bool, t2 byte, t3 error, t4 complex64, t5 complex128, t6 float32, t7 float64, t8 int, t9 int8, t10 int16, t11 int32, t12 int64, t13 uint, t14 uint8, t15 uint16, t16 uint32, t17 uint64, t18 rune, t19 string, t20 uintptr, t21 context.Context) (bool, byte, error, complex64, complex128, float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, rune, string, uintptr) {
	a1 := bool(t1)
	a2 := byte(t2)
	a3 := error(t3)
	a4 := complex64(t4)
	a5 := complex128(t5)
	a6 := float32(t6)
	a7 := float64(t7)
	a8 := int(t8)
	a9 := int8(t9)
	a10 := int16(t10)
	a11 := int32(t11)
	a12 := int64(t12)
	a13 := uint(t13)
	a14 := uint8(t14)
	a15 := uint16(t15)
	a16 := uint32(t16)
	a17 := uint64(t17)
	a18 := rune(t18)
	a19 := string(t19)
	a20 := uintptr(t20)
	return a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20
}

func Bar69() (bool, byte, error, complex64, complex128, float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, rune, string, uintptr) {
	var t1 bool
	var t2 byte
	var t3 error
	var t4 complex64
	var t5 complex128
	var t6 float32
	var t7 float64
	var t8 int
	var t9 int8
	var t10 int16
	var t11 int32
	var t12 int64
	var t13 uint
	var t14 uint8
	var t15 uint16
	var t16 uint32
	var t17 uint64
	var t18 rune
	var t19 string
	var t20 uintptr
	return t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20
}

// v0.2.9
type cell struct{}

type board struct {
	matrixCells        [][]cell
	matrixCellsPointer *[][]cell
	rowCells           []cell
	rowCellsPointer    *[]cell
	matrixInts         [][]int
	matrixIntsPointer  *[][]int
	m0                 map[cell]interface{}
	m1                 []<-chan cell
	m2                 *[]chan<- cell
	m3                 [][]map[cell]interface{}
	m4                 *[][]map[cell]interface{}
	m5                 [][]<-chan func()
	m6                 *[][]chan<- func()
	m7                 [][]context.Context
	m8                 *[][]context.Context
	m9                 [][][][][][][][][]context.Context
	m10                *[][][][][][][][][]context.Context
	m11                [][][][][][][][][]cell
	m12                *[][][][][][][][][]cell
	m13                [][][][][][][][][]cell
	m14                *[][][][][][][][][]cell
	m15                [][][][][][][][][]<-chan cell
	m16                *[][][][][][][][][]chan<- cell
	m17                [][][][][][][][][]<-chan map[cell]interface{}
	m18                *[][][][][][][][][]chan<- map[cell]interface{}
	m19                [][][][][][][][][]func(http.Handler) <-chan http.Handler
	m20                *[][][][][][][][][]func(foo cell) (foox, barx, foobar io.Writer)
	m21                [][][][][][][][][]func() <-chan cell
	m22                *[][][][][][][][][]func() chan<- cell
	m23, m24, m25      [][][][][][][][][]cell
	m26, m27, m28      *[][][][][][][][][]cell
}
type board2 struct{ foo [][]cell }

type board3 struct{ foo *[][]cell }

// v0.2.10
type Foo35 struct {
	foo struct {
		foo context.Context
		bar struct {
			foo context.Context
			bar struct {
				foo context.Context
			}
			foobar context.Context
		}
	}
	bar context.Context
}

type Foo36 struct {
	foo, bar context.Context
	foo0     func()
	bar0     struct {
		foo context.Context
		bar struct {
			foo  func()
			foo0 func(context.Context)
			foo1 struct{}
			foo2 chan context.Context
			foo3 func()
		}
	}
	foobar context.Context
}

type Foo37 struct {
	Foo struct {
		matrixCells        [][]cell
		matrixCellsPointer *[][]cell
		rowCells           []cell
		rowCellsPointer    *[]cell
		matrixInts         [][]int
		matrixIntsPointer  *[][]int
		m0                 map[cell]interface{}
		m1                 []<-chan cell
		m2                 *[]chan<- cell
		m3                 [][]map[cell]interface{}
		m4                 *[][]map[cell]interface{}
		m5                 [][]<-chan func()
		m6                 *[][]chan<- func()
		m7                 [][]context.Context
		m8                 *[][]context.Context
		m9                 [][][][][][][][][]context.Context
		m10                *[][][][][][][][][]context.Context
		m11                [][][][][][][][][]cell
		m12                *[][][][][][][][][]cell
		m13                [][][][][][][][][]cell
		m14                *[][][][][][][][][]cell
		m15                [][][][][][][][][]<-chan cell
		m16                *[][][][][][][][][]chan<- cell
		m17                [][][][][][][][][]<-chan map[cell]interface{}
		m18                *[][][][][][][][][]chan<- map[cell]interface{}
		m19                [][][][][][][][][]func(http.Handler) <-chan http.Handler
		m20                *[][][][][][][][][]func(foo cell) (foox, barx, foobar io.Writer)
		m21                [][][][][][][][][]func() <-chan cell
		m22                *[][][][][][][][][]func() chan<- cell
		m23, m24, m25      [][][][][][][][][]cell
		m26, m27, m28      *[][][][][][][][][]cell
	}
	timerCtx struct {
		cancelCtx
		timer *time.Timer // Under cancelCtx.mu.

		deadline time.Time
	}
	Foo25 struct {
		Foo15
		*Foo17
		Foo    func() (context.Context, func(), io.ReadCloser, error, map[Foo17]interface{})
		Bar    func() (foo context.Context, bar io.ReadCloser, err error)
		Fooo   Foo17
		Barr   string
		Foo2   map[Foo17]interface{}
		Bar2   func(foo FooX) (foox, barx, foobar io.Writer)
		FooBar *Foo17
		Foo3   func() (foo context.Context, bar io.ReadCloser, foobar *map[Foo17]interface{})
		Bar3   func() map[Foo17]interface{}
		Foo4   chan<- Foo17
		Bar4   <-chan Foo17
		Foo5   chan *Foo17
		Bar5   chan Foo17
		Foo6   []func(http.Handler) http.Handler
		Bar6   func() context.Context
		Foo7   []func(http.Handler) <-chan http.Handler
		Bar7   []nodeTyp
		Foo8   [ntCatchAll + 1]nodeTyp
		Bar8   []*nodeTyp
	}
	bar string
}
type pageInfo struct {
	EndCursor string `json:"end_cursor"`
	NextPage  bool   `json:"has_next_page"`
}
type PageData struct {
	Rhxgis    string
	EntryData struct {
		ProfilePage struct {
			Graphql struct {
				User struct {
					Id    string
					Media struct {
						PageInfo pageInfo
					}
				}
			}
		}
	}
}

type mainPageData struct {
	Rhxgis    string `json:"rhx_gis"`
	EntryData struct {
		ProfilePage []struct {
			Graphql struct {
				User struct {
					Id    string `json:"id"`
					Media struct {
						Edges []struct {
							Node struct {
								ImageURL     string `json:"display_url"`
								ThumbnailURL string `json:"thumbnail_src"`
								IsVideo      bool   `json:"is_video"`
								Date         int    `json:"date"`
								Dimensions   struct {
									Width  int `json:"width"`
									Height int `json:"height"`
								} `json:"dimensions"`
							} `json:"Node"`
						}
						PageInfo pageInfo `json:"page_info"`
					} `json:"edge_owner_to_timeline_media"`
				} `json:"user"`
			} `json:"graphql"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`
}

type nextPageData struct {
	Data struct {
		User struct {
			Container struct {
				PageInfo pageInfo `json:"page_info"`
				Edges    []struct {
					Node struct {
						ImageURL     string `json:"display_url"`
						ThumbnailURL string `json:"thumbnail_src"`
						IsVideo      bool   `json:"is_video"`
						Date         int    `json:"taken_at_timestamp"`
						Dimensions   struct {
							Width  int `json:"width"`
							Height int `json:"height"`
						}
					}
				} `json:"edges"`
			} `json:"edge_owner_to_timeline_media"`
		}
	} `json:"data"`
}

// v0.2.11
func Foo38() {
	type foo string
	type bar interface{}
	_ = make(map[string]Context)
	_ = make(map[string]*Context)
	_ = make(map[Context]map[Context]*map[Context]*Context)
	_ = make(map[Context]interface{})
	_ = make(map[Context]map[Context]*map[Context]interface{})
	_ = map[foo]bar{"foo": "bar"}
}

// v0.2.12
type Foo39 struct{}

type Foo40 interface{ *[]Foo39 | Bar70 }

type Foo41 interface{ <-chan func() []*Foo39 }

type Foo42 interface {
	Foo39 //bar
}

type Bar70 interface {
	Foo39 | []FooX
}

type Bar71 interface {
	Foo39 | *[]FooX
}
type Foo43 interface {
	FooX | <-chan *[]context.Context | chan *[]context.Context | chan<- *[]context.Context //foo
}

type Foo44 interface {
	chan *[]context.Context | FooX | func() <-chan *[]context.Context //bar
}

type hello interface {
	bar(bar *FooX) <-chan func() *[]context.Context //foo
}

type Generic interface {
	*[]Foo39 | FooX | map[FooX]string
}

func Foo45[T Generic](a string) string {
	return "foo"
}

func Foo46() {
	_ = Foo45[map[FooX]string]("bar")
	_ = Foo45[FooX]("bar")
	_ = Foo45[*[]Foo39]("bar")
}

func Foo47() {
	var a interface{}
	_, _ = a.(FooX)
	_, _ = a.(*FooX)
	_, _ = a.(*[]FooX)
	_, _ = a.([]FooX)
	_, _ = a.([]*FooX)
	_, _ = a.(map[FooX]interface{})
}

// v0.2.14
func Foo48() {
	_ = new(FooX)
	_ = new(*FooX)
	_ = new(*[]FooX)
	_ = new([]FooX)
	_ = new([]*FooX)
	_ = new(map[FooX]interface{})
	_ = new(context.Context)
}

// v0.2.15
func Foo49(a struct {
	b    Foo11
	c    FooX
	d    map[FooX]chan<- func(f FooX)
	e, f Foo07
}) {
	go func(a struct {
		b    Foo11
		c    FooX
		d    map[FooX]chan<- func(f FooX)
		e, f Foo07
	}) {
	}(a)
}

// v0.2.17
type newstruct1 struct{}
type newstruct2 struct{}

type newGeneric interface {
	newstruct1 | newstruct2
}

type newstructG[T newGeneric] struct {
	name string
	age  T
}

func (myvar *newstructG[newstruct1]) Hi() {

}

func (myvar *newstruct1) Hello() {

}

// v0.2.18
func Foo50(newstructG[newstruct1], newstructG[newstruct2]) {}

func Foo51(*newstructG[newstruct1], *newstructG[newstruct2]) {}

func Foo52(foo, bar,
	foobar newstructG[newstruct1])

func Foo53(foo, bar,
	foobar chan<- *newstructG[newstruct2])

func Foo54(
	newstructG[newstruct1], newstructG[newstruct2],
	newstructG[newstruct1])

func Foo55(chan newstructG[newstruct1], newstructG[newstruct2],
	string, newstructG[newstruct2], int64)

func Foo56(<-chan *newstructG[newstruct1], newstructG[newstruct1],
	string, chan<- newstructG[newstruct2])

func Foo57(newstructG[newstruct1], <-chan newstructG[newstruct2],
	string, chan<- newstructG[newstruct1], chan int16)

func Foo58(newstructG[newstruct1], <-chan func() newstructG[newstruct2],
	string, chan<- newstructG[newstruct2])

func Foo59(func() <-chan newstructG[newstruct1], newstructG[newstruct2],
	string, chan<- newstructG[newstruct2])

func Foo60(chan<- func() newstructG[newstruct1], func() newstructG[newstruct2],
	string, chan<- newstructG[newstruct2])

func Foo61(newstructG[newstruct1], func(), map[Foo17]interface{},
	newstructG[newstruct2], func(), io.ReadCloser, io.ReadCloser, context.Context,
	string, func(newstructG[newstruct1]), *io.Writer, *io.Reader,
	context.Context, context.Context, context.Context,
	string, chan func(*newstructG[newstruct2]), bool, io.Reader,
	string, <-chan func(context.Context), bool, chan<- io.Reader,
	error, map[Foo17]interface{}, newstructG[newstruct1]) {
}

func Foo62(newstructG[newstruct1], func(), map[Foo17]interface{}, newstructG[newstruct2], func(), io.ReadCloser, io.ReadCloser, context.Context, string, func(newstructG[newstruct1]), *io.Writer, *io.Reader, context.Context, newstructG[newstruct2], *newstructG[newstruct2], string, chan func(context.Context), bool, io.Reader, string, <-chan func(newstructG[newstruct1]), bool, chan<- io.Reader, error, map[Foo17]interface{}, newstructG[newstruct2]) {
}
func Foo63(newstructG[newstruct1], func(), <-chan newstructG[newstruct2])

func Foo64(<-chan context.Context, <-chan newstructG[newstruct1], chan func(), <-chan context.Context)

func Foo65(<-chan context.Context, func(), chan<- func(newstructG[newstruct1]), <-chan newstructG[newstruct2])

func Foo66(newstructG[newstruct1], string)

func Foo67(foo newstructG[newstruct1], bar newstructG[newstruct2]) {}

func Foo68(foo, bar, foobar newstructG[newstruct1])

func Foo69(
	foo, bar, foobar newstructG[newstruct2]) {
}

func Foo70(
	context.CancelCauseFunc, context.CancelFunc, context.Context, newstructG[newstruct1]) {
}

func Foo71(
	context.CancelCauseFunc,
	context.CancelFunc,
	context.Context,
	newstructG[newstruct1],
	newstructG[newstruct2]) {
}

func Foo72(
	context.CancelCauseFunc,
	context.CancelFunc,
	context.Context,
	*newstructG[newstruct1],
	*newstructG[newstruct2],
) {
}

func Foo73(bar newstructG[newstruct1]) {}

func Foo74(foo newstructG[newstruct1], bar *newstructG[newstruct2]) {}

func Foo75(
	foo *newstructG[newstruct1],
	bar newstructG[newstruct2],
) {

}

func Foo76() {
	type foo newstructG[newstruct1]
	type bar newstructG[newstruct2]
	type foox *newstructG[newstruct1]
	type barx *newstructG[newstruct2]
	type foobar map[newstructG[newstruct1]]interface{}

	var xfoo newstructG[newstruct1]
	var xbar newstructG[newstruct2]
	var xfoox *newstructG[newstruct1]
	var xbarx *newstructG[newstruct2]
	var foobarx map[newstructG[newstruct1]]interface{}

	type foo0 = newstructG[newstruct1]
	type bar0 = newstructG[newstruct2]
	type foox0 = *newstructG[newstruct1]
	type barx0 = *newstructG[newstruct2]
	type foobar0 = map[newstructG[newstruct1]]interface{}

	_, _, _, _, _ = xfoo, xbar, xfoox, xbarx, foobarx
}

// v0.2.19
type newstruct3 struct{}

func Foo77[foo newstruct1](bar foo) {

}

func Foo78[
	foo newstruct1](bar context.Context) {
}

func Foo79[foo newstruct1 | newstruct2 | newstruct3]() {

}

func Foo80[foo newstruct1 | newstruct2 | string | newstruct3]() {

}

func Foo81[
	foo newstruct1 |
		newstruct2 | string | newstruct3]() {
}

func Foo82[foo newstruct1 |
	newstruct2 | string | newstruct3]() {
}

func Foo83[
	foo,
	bar,
	foobar newstruct1,
]() {

}

func Foo84[
	foo,
	bar,
	foobar newstruct1]() {
}

func Foo85[
	foo,
	bar,
	foobar newstruct1]() {
}

func Foo86[
	foo,
	bar,
	foobar newstruct1,
	foox,
	barx,
	foobarx newstruct2,
]() {
}

func Foo87[
	foo, bar, foobar newstruct1,
	foox, barx, foobarx newstruct2]() {
}

func Foo88[
	foo, bar, foobar newstruct1,
	foox, barx, foobarx newstruct2,
	fooy, bary, foobary newstruct3,
]() {
}

func Foo89[
	foo, bar, foobar newstruct1,
	foox, barx, foobarx newstruct2,
	fooy, bary, foobary newstruct3]() {
}

func Foo90[
	foo newstruct1,
	bar newstruct2,
	foobar newstruct3,
]() {

}

func Foo91(struct {
	b    Foo11
	c    FooX
	d    map[FooX]chan<- func(f FooX)
	e, f Foo07
}) {
}

func Foo92() {
	_ = Foo11{}

	_ = &Foo11{}
	_ = Foo11{name: "foo", age: 100}
	_ = &Foo11{name: "foo", age: 100}
	_ = Foo11{
		name: "foo",
		age:  100,
	}
	_ = &Foo11{
		name: "foo",
		age:  100,
	}
	variable := 10
	value1, value2, value3 := 10, 20, 30
	switch variable {
	case value1, value2:
		// Code to execute
	case value3:
		// Code to execute
	}
	switch {
	case variable == value1:
		// Code to execute
	case variable > value2:
		// Code to execute
	}
	switch variable {
	case value1:
		// Code to execute
	case value2:
		// Code to execute
	case value3:
		// Code to execute
	}
}

func Foo93() {
outerLoop:
	for i := 0; i < 3; i++ {
	innerLoop:
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outerLoop
			} else {
				break innerLoop
			}
		}
	}
}

// v0.2.21
func Foo94[foo newstructG[newstruct1]](bar string) {}

func Foo95[foo newstructG[newstruct1], bar newstructG[newstruct2], foobar newstructG[newstruct1]]() {}

func Foo96[foo newstructG[newstruct1] | newstructG[newstruct2] | newstruct3]() {}

func Foo97() {
	foo := make([]int, 1)
	foo[0] = 10
}

// v0.2.22
func Foo98[T FooX](foo T) {}

func Foo99() {
	var (
		xfoo    newstructG[newstruct1]
		xbar    newstructG[newstruct2]
		xfoox   *newstructG[newstruct1]
		xbarx   *newstructG[newstruct2]
		foobarx map[newstructG[newstruct1]]interface{}
	)
	_, _, _, _, _ = xfoo, xbar, xbarx, xfoox, foobarx

	bar := "foo"
	foo := map[string]int{bar: 1}
	_ = foo[bar]
	fmt.Sprintln(foo[bar])

	_ = newstructG[newstruct1]{}
	foox := FooX{}
	Foo98[FooX](FooX{})
	Foo98[FooX](foox)
}

// v0.2.23
func Foo100() {
outerLoop:
	for i := 0; i < 3; i++ {
	innerLoop:
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				goto outerLoop
			} else {
				break innerLoop
			}
		}
	}
}

// v0.2.24
func Foo101() {
outerLoop:
	for i := 0; i < 3; i++ {
	innerLoop:
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				goto outerLoop /* Jump to the outer loop */
			} else if i == 2 && j == 2 {
				break outerLoop // Break out of the outer loop
			} else if i == 0 && j == 0 {
				continue innerLoop //Continue to the next iteration of the inner loop
			}
			fmt.Println(i, j)
		}
	}
}

// v0.2.25
func Foo102(expr ast.Expr) interface{} {
	switch x := expr.(type) { // first
	case *ast.BasicLit: // comment
		return x.Value
	case *ast.CompositeLit: /* comment */
		switch x.Type.(type) { /* second */
		case *ast.MapType: // comment
			m := make(map[interface{}]interface{})
			for _, el := range x.Elts {
				_ = el.(*ast.KeyValueExpr)
			}
			return m
		}
	}
	return nil
}

// v0.2.26
type newGeneric2 interface {
	newstructG[newstruct1] | newstructG[newstruct2] // comment
}

// v0.2.27
type newGeneric3 interface {
	newstructG[newstruct1] | newstructG[newstruct2] | // comment
		// comment
		int | float64
}

// v0.2.28
func Foo103(
	foo []string,
	bar string,
	foobar []struct {
		Path       string
		LineNumber int
		Line       string
	}) {
}

func Foo104(struct {
	Path       context.Context
	LineNumber int
	Line       string
}) {
}

func Foo105([]struct {
	Path       context.Context
	LineNumber int
	Line       string
}) {
}

func Foo106(foo []string, bar string, foobar []struct {
	Path       string
	LineNumber int
	Line       string
}) {
}

func Foo107() {
	foo := struct {
		Path       string
		LineNumber int
		Line       string
	}{
		Path:       "foo",
		LineNumber: 100,
		Line:       "bar",
	}

	bar := []struct {
		Path       string
		LineNumber int
		Line       string
	}{
		{
			Path:       "foo",
			LineNumber: 100,
			Line:       "bar",
		},
	}

	var foobar []struct {
		Path       context.Context
		LineNumber int
		Line       string
	}

	_, _, _ = foo, bar, foobar

}

// v0.2.29
func (in *FooX) Foo108([]newstruct1, func(n ast.Node, push bool)) {}

func (in *FooX) Foo109(types []newstruct1, f func(n ast.Node, push bool)) {}

func (in *FooX) Foo110(types []newstruct1, f func(n ast.Node, push bool) (proceed context.Context)) {}

func (in *FooX) Foo111(types []newstruct2,
	f func(n ast.Node, push bool) (proceed context.Context)) {
}

func Foo112() {
	const P = 3
	const C = 4
	type T [P * C]context.Context
}

func Foo113() {
	slc := []bool{true, false, true}
	_ = min(10, 20)
	_ = max(10, 20)
	clear(slc)
}

type newGen[a, b newstruct1, c newstruct2] struct{}

type Pair[T1, T2 any] struct {
	One T1
	Two T2
}

type GenFoo114[a, b newstruct1, c newstruct2] struct{}

func Foo114[T GenFoo114[newstruct1, newstruct1, newstruct2]](foo GenFoo114[newstruct1, newstruct1, newstruct2]) {
}

func Foo115[T newGen[newstruct1, newstruct1, newstruct2]]() {}

func Foo116() {
	cnt := 10
	arr := []string{"1"}
	_ = make(chan context.Context)
	_ = make(chan<- context.Context)
	_ = make(<-chan context.Context)
	_ = make(chan context.Context, 10)
	_ = make(map[context.Context]context.Context)
	_ = make([]context.Context, 0, 10)
	_ = make([]context.Context, 10)
	_ = make([]int, 0, 10)
	_ = make([]string, 0)
	_ = make(<-chan interface{}, 10)
	_ = make(map[string]interface{}, 10)
	_ = make([]context.Context, 0, cnt)
	_ = make([]context.Context, 0, len(arr))
}

type Iter[T any] func(yield func(T) bool) bool
type Iter2[T1, T2 any] func(yield func(T1, T2) bool) bool

func Enum[T any](iter Iter[T]) Pair[int, T] {
	return Pair[int, T]{}
}

func Equal[S ~[]E, E comparable](s1, s2 S) bool {
	return true
}

func EqualFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, eq func(E1, E2) bool) bool {
	return true
}

func Compare[S ~[]E, E cmp.Ordered](s1, s2 S) int {
	return 0
}

func CompareFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, cmp func(E1, E2) int) int {
	return 0
}

func Index[S ~[]E, E comparable](s S, v E) int {
	return -1
}

func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	return -1
}

func Contains[S ~[]E, E comparable](s S, v E) bool {
	return Index(s, v) >= 0
}

func ContainsFunc[S ~[]E, E any](s S, f func(E) bool) bool {
	return IndexFunc(s, f) >= 0
}

func Insert[S ~[]E, E any](s S, i int, v ...E) S {
	return s
}

func Delete[S ~[]E, E any](s S, i, j int) S {
	return s
}

func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S {
	return s
}

func Replace[S ~[]E, E any](s S, i, j int, v ...E) S {
	return s
}

func Clone[S ~[]E, E any](s S) S {
	return s
}

func Compact[S ~[]E, E comparable](s S) S {
	return s
}

func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S {
	return s
}

func Grow[S ~[]E, E any](s S, n int) S {
	return s
}

func Clip[S ~[]E, E any](s S) S {
	return s[:len(s):len(s)]
}

func rotateLeft[E any](s []E, r int) {

}

func rotateRight[E any](s []E, r int) {
}

func swap[E any](x, y []E) {}

func overlaps[E any](a, b []E) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	elemSize := unsafe.Sizeof(a[0])
	if elemSize == 0 {
		return false
	}
	return uintptr(unsafe.Pointer(&a[0])) <= uintptr(unsafe.Pointer(&b[len(b)-1]))+(elemSize-1) &&
		uintptr(unsafe.Pointer(&b[0])) <= uintptr(unsafe.Pointer(&a[len(a)-1]))+(elemSize-1)
}

func startIdx[E any](haystack, needle []E) int { _, _ = haystack, needle; return 0 }

func Reverse[S ~[]E, E any](s S) {}

// v0.2.30
func Foo117() {
	_ = new(int)
}

// v0.2.31
func (a *FooX) Foo118(foo string, bar map[string]*multipart.File, foobar map[Context]map[Context]*map[Context]*Context) {
}

func Foo119(
	foo string,
	bar map[string]*multipart.File,
	foobar map[Context]map[Context]*map[Context]*Context) {
}

func Foo120(map[string]*multipart.File)     {}
func Foo121(foo map[string]*multipart.File) {}
func Foo122(
	bar map[string]*multipart.File) {
}
func Foo123[foobar string](foo map[string]*multipart.File) {}

// v0.2.32
func Foo124() {
	_ = make([]Context, 0, 10)
}

// v0.2.33
func Foo125(chan context.Context) {}
func Foo126(
	chan context.Context) {
}
func Foo127(<-chan context.Context) {}
func Foo128(chan<- context.Context) {}

// v0.2.34
func Foo129(func() context.Context) {}
func Foo130(
	func() context.Context) {
}
func Foo131(func(i, j context.Context) context.Context) {}

func Foo132(
	func(i, j context.Context) context.Context) {
}

func Foo133(chan func(i, j context.Context) context.Context)   {}
func Foo134(<-chan func(i, j context.Context) context.Context) {}
func Foo135(chan<- func(i, j context.Context) context.Context) {}
func Foo136(
	chan func(i, j context.Context) context.Context) {
}
func Foo137(
	<-chan func(i, j context.Context) context.Context) {
}
func Foo138(
	chan<- func(i, j context.Context) context.Context) {
}
func Foo139(func(i, j context.Context) chan context.Context)   {}
func Foo140(func(i, j context.Context) chan<- context.Context) {}
func Foo141(func(i, j context.Context) <-chan context.Context) {}
func Foo142(
	func(i, j context.Context) chan context.Context) {
}
func Foo143(
	func(i, j string) chan<- string) {
}
func Foo144(
	func(i, j context.Context) <-chan context.Context) {
}
func Foo145(foo func() context.Context) {}
func Foo146(
	foo func() context.Context) {
}
func Foo147(foo func(i, j context.Context) context.Context) {}

func Foo148(
	foo func(i, j context.Context) context.Context) {
}

func Foo149(foo chan func(i, j context.Context) context.Context)   {}
func Foo150(foo <-chan func(i, j context.Context) context.Context) {}
func Foo151(foo chan<- func(i, j context.Context) context.Context) {}
func Foo152(
	foo chan func(i, j context.Context) context.Context) {
}
func Foo153(
	foo <-chan func(i, j context.Context) context.Context) {
}
func Foo154(
	foo chan<- func(i, j context.Context) context.Context) {
}
func Foo155(foo func(i, j context.Context) chan context.Context, bar func(i, j context.Context) chan<- context.Context) {
}
func Foo156(
	foo func(i, j context.Context) <-chan context.Context, bar func(i, j context.Context) chan context.Context) {
}

func Foo157(
	foo func(i, j context.Context) chan context.Context,
	bar func(i, j context.Context) chan context.Context) {
}

func Foo158(foo func(i, j context.Context) chan<- context.Context) {}
func Foo159(foo func(i, j context.Context) <-chan context.Context) {}
func Foo160(
	foo func(i, j context.Context) chan context.Context) {
}
func Foo161(
	foo func(i, j string) chan<- context.Context) {
}
func Foo162(
	foo func(i, j context.Context) <-chan context.Context) {
}

// v0.2.35
func Foo163() {
	_ = func() *GenFoo114[newstruct1, newstruct1, newstruct2] {
		return nil
	}
}

// v0.2.36
func Foo164() {
	_ = func() (foo context.Context, bar context.Context) {
		return
	}
	_ = func() (context.Context, context.Context) {
		return nil, nil
	}
	_ = func() (foo, bar, foobar chan<- context.Context) {
		return nil, nil, nil
	}
}

// v0.2.37
func Foo165() {
	var foo func(con context.Context)
	var bar func(con <-chan context.Context)
	var foobar func(con <-chan context.Context, a, b, c context.Context)
	var fooxbar func(a context.Context, b context.Context)

	var (
		foox     func(con context.Context)
		barx     func(con context.Context)
		foobarx  func(con <-chan context.Context, a, b, c context.Context)
		fooxbarx func(a context.Context, b context.Context)
	)

	_, _, _, _ = foo, bar, foobar, fooxbar
	_, _, _, _ = foox, barx, foobarx, fooxbarx
}

// v0.2.38
func Foo166() {
	arr := []int{}
	_ = make([]GenFoo114[newstruct1, newstruct1, newstruct2], 10)
	_ = make([]newstructG[newstruct1], 10)
	_ = make([]GenFoo114[newstruct1, newstruct1, newstruct2], len(arr))
	_ = make([]newstructG[newstruct1], len(arr))
	_ = make([]GenFoo114[newstruct1, newstruct1, newstruct2], 0, len(arr))
	_ = make([]newstructG[newstruct1], 0, len(arr))
}

// v0.2.39
func Foo167() {
	arr := []int{}
	_ = make(map[GenFoo114[newstruct1, newstruct1, newstruct2]]int, 10)
	_ = make(map[newstructG[newstruct1]]context.Context, 10)
	_ = make(map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context, len(arr))
	_ = make(map[newstructG[newstruct1]]context.Context, len(arr))
	_ = make([]map[GenFoo114[newstruct1, newstruct1, newstruct2]]string, 0, len(arr))
	_ = make([]map[newstructG[newstruct1]]int, 0, len(arr))
	_ = make(map[GenFoo114[newstruct1, newstruct1, newstruct2]]map[GenFoo114[newstruct1, newstruct1, newstruct2]]*map[GenFoo114[newstruct1, newstruct1, newstruct2]]interface{})
	_ = make([]map[string]context.Context, 0, len(arr))
	_ = make(map[newstructG[newstruct1]]map[newstructG[newstruct1]]context.Context, len(arr))
	_ = make(map[GenFoo114[newstruct1, newstruct1, newstruct2]]map[GenFoo114[newstruct1, newstruct1, newstruct2]]*map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context)
	_ = make(map[GenFoo114[newstruct1, newstruct1, newstruct2]]map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context)
	_ = make(map[GenFoo114[newstruct1, newstruct1, newstruct2]]map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context, len(arr))
	_ = make([]map[GenFoo114[newstruct1, newstruct1, newstruct2]]map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context, 0, len(arr))
}

// v0.2.40
func Foo168() {
	var a1 map[GenFoo114[newstruct1, newstruct1, newstruct2]]newstruct1
	var a2 map[newstructG[newstruct1]]context.Context
	var a3 map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context
	var a4 map[newstructG[newstruct1]]context.Context
	var a5 map[newstructG[newstruct1]]context.Context
	var a6 []map[GenFoo114[newstruct1, newstruct1, newstruct2]]string
	var a7 []map[newstructG[newstruct1]]int
	var a8 map[GenFoo114[newstruct1, newstruct1, newstruct2]]map[GenFoo114[newstruct1, newstruct1, newstruct2]]*map[GenFoo114[newstruct1, newstruct1, newstruct2]]interface{}
	var a9 []map[string]context.Context
	var a10 map[newstructG[newstruct1]]map[newstructG[newstruct1]]context.Context
	var a11 map[GenFoo114[newstruct1, newstruct1, newstruct2]]map[GenFoo114[newstruct1, newstruct1, newstruct2]]*map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context
	var a12 map[GenFoo114[newstruct1, newstruct1, newstruct2]]map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context
	var a13 map[GenFoo114[newstruct1, newstruct1, newstruct2]]map[GenFoo114[newstruct1, newstruct1, newstruct2]]string
	var a14 []map[GenFoo114[newstruct1, newstruct1, newstruct2]]*map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context
	var a15 *[]map[GenFoo114[newstruct1, newstruct1, newstruct2]]*map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context

	var (
		b1  map[GenFoo114[newstruct1, newstruct1, newstruct2]]newstruct1
		b2  map[newstructG[newstruct1]]context.Context
		b3  map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context
		b4  map[newstructG[newstruct1]]context.Context
		b5  map[newstructG[newstruct1]]context.Context
		b6  []map[GenFoo114[newstruct1, newstruct1, newstruct2]]string
		b7  []map[newstructG[newstruct1]]int
		b8  map[GenFoo114[newstruct1, newstruct1, newstruct2]]map[GenFoo114[newstruct1, newstruct1, newstruct2]]*map[GenFoo114[newstruct1, newstruct1, newstruct2]]interface{}
		b9  []map[string]context.Context
		b10 map[newstructG[newstruct1]]map[newstructG[newstruct1]]context.Context
		b11 map[GenFoo114[newstruct1, newstruct1, newstruct2]]map[GenFoo114[newstruct1, newstruct1, newstruct2]]*map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context
		b12 map[GenFoo114[newstruct1, newstruct1, newstruct2]]map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context
		b13 map[GenFoo114[newstruct1, newstruct1, newstruct2]]map[GenFoo114[newstruct1, newstruct1, newstruct2]]string
		b14 []map[GenFoo114[newstruct1, newstruct1, newstruct2]]*map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context
		b15 *[]map[GenFoo114[newstruct1, newstruct1, newstruct2]]*map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context
	)

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b1, b2, b3, b4, b5, b6, b7, b8, b9, b10, b11, b12, b13, b14, b15

}

// v0.3.1
func Foo169() {
	_ = func(a *[]map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context) *[]map[GenFoo114[newstruct1, newstruct1, newstruct2]]context.Context {
		return a
	}
}

// v0.3.2
func Foo170() {
	type LongFooooooooooooooooooooooo context.Context
	var foo []LongFooooooooooooooooooooooo
	var bar []*LongFooooooooooooooooooooooo
	var (
		foox []LongFooooooooooooooooooooooo
		barx []*LongFooooooooooooooooooooooo
	)
	_, _, _, _ = foo, bar, foox, barx
}

// v0.3.3
func Foo171() {
	// fix preview (cursor)
	_ = &build.Default
	var (
		_ = &build.Default
		_ = build.Default
	)
}

// v0.3.4
func Foo172() {
	// better pre-highlighting variables after control keywords -before formatting with gofmt

	a := []string{"a"}
	b := make(chan bool)
	c := make(chan bool)

	for i, r := range a[0] {
		_, _ = i, r
	}
	if <-b {

	} else if <-c {
	}
}

// v0.3.5
func Foo173() {
	// better pre-highlighting variables after control keywords -before formatting with gofmt
	a := "foo"
	b := []*string{&a}
	for i, r := range *b[0] {
		_, _ = i, r
	}
}

// v0.3.6
func Foo174() {
	type foox string

	type foo[E, V context.Context] []E

	type bar[E, V any] []E

	type foobar[E, V any] context.Context

	type foobarx[E, V any] func(a context.Context, b context.Context) <-chan context.Context

	type foobar1[a string, b context.Context] *context.Context
	type foobar2[a string, b context.Context] []*context.Context
	type foobar3[a string, b context.Context] *[]context.Context
	type foobar4[a string, b context.Context] chan context.Context
	type foobar5[a string, b context.Context] <-chan func() context.Context
	type foobar6[a string, b context.Context] chan<- func() context.Context
	type foobar7[a string, b context.Context] chan *context.Context
	type foobar8[a string, b context.Context] <-chan []*context.Context
	type foobar9[a string, b context.Context] chan<- *[]context.Context
	type foobar10[a, b context.Context, d, e foox] context.Context
	type foobar11[a, b bar[context.Context, string], d, e bar[int, context.Context]] *context.Context
	type foobar12[a, b bar[context.Context, string], d, e bar[int, context.Context]] chan *context.Context
	type foobar13[a bar[context.Context, string], d bar[int, context.Context]] chan *context.Context
	type foobar14[a bar[context.Context, foo[context.Context, context.Context]], b bar[string, foo[context.Context, context.Context]]] chan *context.Context

	type foo1[a, b bar[context.Context, string], d, e bar[int, context.Context]] struct{}
	type foo2[a, b bar[context.Context, string], d, e bar[int, context.Context]] interface{}

	type foo3[a, b bar[context.Context, context.Context], c context.Context] struct {
	}

	type foo4[a, b bar[context.Context, context.Context], c, d context.Context] struct {
	}

	type foo5[a, b bar[context.Context, context.Context], c context.Context] interface {
	}

	type foo6[a, b bar[context.Context, context.Context], c, d context.Context] interface {
	}

	type foo7[a bar[context.Context, foo[context.Context, context.Context]], b bar[string, foo[context.Context, context.Context]]] struct{}
	type foo8[a bar[context.Context, foo[context.Context, context.Context]], b bar[string, foo[context.Context, context.Context]]] interface{}
	type foo9[a []bar[context.Context, []foo[context.Context, context.Context]], b *[]bar[string, []foo[context.Context, context.Context]]] struct{}
	type foo10[a *[]bar[context.Context, *[]foo[context.Context, context.Context]], b *[]bar[string, *[]foo[context.Context, context.Context]]] interface{}
	type foo11[a []*bar[context.Context, []*foo[context.Context, context.Context]], b []*bar[string, []*foo[context.Context, context.Context]]] interface{}

}

type Bar72[E, V any] []E

func (s *Bar72[E, V]) Foo175(func(E) V) {
}

// v0.3.7
func Foo176() {
	type foo struct{ context.Context }
}

// v0.3.8
func Foo177() {
	type foo func(index string, o ...func(*FooX)) (*FooX, error)
	type bar func(index string, o ...func(*FooX)) (a *FooX, b Context)
	type foobar func(index context.Context, o ...func(*FooX)) (a *FooX, b []*context.Context)
}

func Foo178() FooX // bar

func Foo179() FooX /* bar */

func Foo180(a, b, c context.Context) // bar

func Foo181(a, b, c context.Context) FooX // bar

func Foo182(a, b, c context.Context) <-chan FooX /* bar */

func Foo183(a, b, c context.Context) <-chan func(FooX) FooX /* bar */

func Foo184(string) FooX //bar

func Foo185(string) *FooX /* bar */

// v0.3.9
type VariantA struct{}

func (*VariantA) sealed() {}

type VariantB struct{}

func (*VariantB) sealed() {}

type MySumType interface {
	sealed()
}

func Foo186() {
	switch MySumType(nil).(type) { //foo
	case *VariantA: //bar
	}
}

// v0.4.0
func Foo187() {
	type foo map[string]func(args ...interface{}) context.Context
	_ = make(map[string]func(args context.Context) context.Context)
}

func f1(a int) string {
	return "I have " + strconv.Itoa(a) + " apples"
}

func f2(b string) string {
	return "there are " + b + " bananas"
}

type t1 struct {
	s1 map[string]func(args ...interface{}) context.Context
}

func Foo188() {
	type foo map[string]func(args ...interface{}) context.Context

	t := t1{
		s1: make(map[string]func(args ...interface{}) context.Context),
	}

	t.s1["test"] = func(args ...interface{}) context.Context {
		result1 := f1(5)
		result2 := f2("20")
		fmt.Println(result1)
		fmt.Println(result2)
		return context.TODO()
	}

	t.s1["test"]("test")
}

func Foo189() {
	_ = make(map[string]func(args context.Context) context.Context)
	_ = new(map[string]func(args context.Context) context.Context)
	type foo1 struct{ context.Context }             //foo
	type foo2 struct{ a, b <-chan context.Context } //foo
	type foo3 struct {
		context.Context //foo
	}

	type foo4 struct {
		a, b context.Context
	}
}

func Bar73() {
	// one line with semicolon without formatting gofmt

	// type foo1 struct { a  string; b  context.Context; c int;}
	type foo1 struct {
		a string
		b context.Context
		c int
	}

	// type foo2 struct { context.Context; a context.Context; b int;}
	type foo2 struct {
		context.Context
		a context.Context
		b int
	}

	// type foo3 struct { context.Context; a context.Context; b int}
	type foo3 struct {
		context.Context
		a context.Context
		b int
	}

	// type foo4 struct { string; context.Context; int}
	type foo4 struct {
		string
		context.Context
		int
	}
}

func Bar74(map[string]func(args context.Context) context.Context) {}

func Bar75(a, b, c map[string]func(args context.Context) context.Context) {}
func Bar76(a map[string]func(args context.Context) context.Context, b map[string]func(args context.Context) (context.Context, string)) {
}

func Bar77(
	a map[string]func(args context.Context) context.Context,
	b map[string]func(args context.Context) (context.Context, string)) {
}

func Bar78(a map[string]func(args context.Context) context.Context, b map[string]func(args context.Context) (context.Context, string),
	c map[string]func(args context.Context) (context.Context, string)) {
}

func Bar79(a <-chan map[string]func(args context.Context) context.Context, b chan map[string]func(args context.Context) (context.Context, string),
	c chan<- map[string]func(args context.Context) (context.Context, string)) {
}

// v0.4.1
func Bar80[M ~map[K]V, K comparable, V any, C any](m M, del func(K, V) bool) {
	for k, v := range m {
		if del(k, v) {
			delete(m, k)
		}
	}
}

func Bar81[M ~map[K]func() context.Context, K comparable, V any, C any](m M, del func(K, V) bool) {
}

// v0.4.2
func Bar82() {
	type (
		foo      int
		bar      context.Context
		foox     *context.Context
		barx     *[]context.Context
		foobar   chan<- context.Context
		foobarx  func() chan context.Context
		fooxbarx func(a context.Context, b string) context.Context
		xfoox    struct {
			context.Context
			a, b, c context.Context
			d       func() chan context.Context
		}
	)
}

func Bar83() {
	type (
		foo      = int
		bar      = context.Context
		foox     = *context.Context
		barx     = *[]context.Context
		foobar   = chan<- context.Context
		foobarx  = func() chan context.Context
		fooxbarx = func(a context.Context, b string) context.Context
		xfoox    = struct {
			context.Context
			a, b, c context.Context
			d       func() chan context.Context
		}
	)
}

// v0.4.3
func Bar84() {
	// works with no space -before formatting with gofmt
	type Foo struct {
		context.Context
		foo, bar context.Context
	}
}

// v0.4.4
func Bar85(ctx context.Context, inputsStream <-chan <-chan context.Context) {}

func Bar86(
	ctx context.Context,
	inputsStream <-chan <-chan context.Context,
	a string) {
}

func Bar87(inputsStream <-chan <-chan context.Context) {}

func Bar88(a, b, c <-chan <-chan context.Context) {}

func Bar89(<-chan <-chan context.Context) {}

func Bar90(func() func() context.Context) {}

func Bar91(ctx context.Context, inputsStream func() func() context.Context) {}

func Bar92(
	ctx context.Context,
	inputsStream func() func() context.Context,
	a string) {
}

func Bar93(inputsStream func(a context.Context) <-chan func(b chan chan context.Context) context.Context) {
}

func Bar94[Input, Output func() func() context.Context](ctx context.Context, input <-chan Input, f func(ctx context.Context, input Input) Output) {
}

func Bar95[Input, Output <-chan <-chan context.Context](ctx context.Context, input <-chan Input, f func(ctx context.Context, input Input) Output) {
}

func Bar96[Input, Output func(a chan<- context.Context) func(b chan<- context.Context) context.Context](ctx context.Context, input <-chan Input, f func(ctx context.Context, input Input) Output) {

}

func Bar97[Input string,
	Output func(a chan<- context.Context) func(b chan<- context.Context) context.Context](ctx context.Context, input <-chan Input, f func(ctx context.Context, input Input) Output) {

}

func Bar98(foo func(a context.Context) func(b <-chan <-chan string) context.Context) {}

func Bar99() {
	// fixes 3rd party property variables
	// _ = &thirdparty.Config{
	// 	Queues: map[string]thirdparty.QueueConfig{
	// 		thirdparty.DefaultQueue: {MaxWorkers: 100},
	// 	},
	// }
	_ = &sync.Cond{
		L: nil,
	}
}

// v0.4.5
func Bar100() {
	type LongFooooooooooooooooooooooo struct{}
	_ = map[LongFooooooooooooooooooooooo]LongFooooooooooooooooooooooo{}
}

// v0.4.6
func Bar101() {
	type opt struct{}
	foo := func(a *opt, s string) {}
	bar := ""
	foo(new(opt), bar)
	_ = new(struct {
		a context.Context
		b string
	})
}

// v0.4.7
func Bar102() {
	a := 10
	b := 5
	// better pre-highlighting variables after control keywords -before formatting with gofmt
	// if a <b{
	// }
	// if a >b{
	// }
	// if a <=b{
	// }
	// if a >=b{
	// }
	// if a ==b{
	// }
	// if a !=b{
	// }
	// if a < b{
	// }
	// if a > b{
	// }
	// if a <= b{
	// }
	// if a >= b{
	// }
	// if a == b{
	// }
	// if a != b{
	// }

	if a > b {
		// implement
	}
}

// v0.4.8
func Bar103() {
	_ = make(chan chan context.Context)
}

func Bar104() {
	_ = func() chan chan context.Context {
		return make(chan chan context.Context)
	}
}

func Bar105() {
	_ = func(func(context.Context)) context.Context {
		return context.TODO()
	}

	_ = func(func(context.Context), context.Context) (context.Context, context.Context) {
		return context.TODO(), context.Background()
	}

	_ = func(a func(context.Context), b context.Context) (foo context.Context, bar context.Context) {
		_, _ = a, b
		return context.TODO(), context.Background()
	}

	_ = func(func(context.Context)) func(context.Context, func(string, context.Context), io.Writer) {
		a := func(context.Context, func(string, context.Context), io.Writer) {}
		return a
	}

	_ = func(context.Context, func(string, context.Context), io.Writer) chan chan context.Context {
		return make(chan chan context.Context)
	}

	_ = func(context.Context, func(string, context.Context), io.Writer) func(foo context.Context, bar context.Context) {
		a := func(context.Context, context.Context) {}
		return a
	}

	_ = func(a context.Context, b func(string, context.Context), c io.Writer) func(foo context.Context, bar context.Context) {
		_, _, _ = a, b, c
		d := func(context.Context, context.Context) {}
		return d
	}

	_ = func(context.Context, func(string, context.Context), io.Writer) func(context.Context, context.Context) chan chan context.Context {
		a := func(context.Context, context.Context) chan chan context.Context {
			return make(chan chan context.Context)
		}
		return a
	}

	_ = func(a context.Context, b func(string, context.Context), c io.Writer) func(c context.Context, d context.Context) chan chan context.Context {
		_, _, _ = a, b, c
		e := func(f context.Context, g context.Context) chan chan context.Context {
			return make(chan chan context.Context)
		}
		return e
	}
	_ = func(foo func(context.Context), bar []struct{ context.Context }, foobar context.Context) context.Context {
		_, _, _ = foo, bar, foobar
		return context.TODO()
	}
	_ = func(foo func(context.Context),
		bar []struct {
			a chan context.Context
			b context.Context
			c <-chan func() context.Context
		}) context.Context {
		_, _ = foo, bar
		return context.TODO()
	}
}

func Bar106() {
	var foo chan chan context.Context
	var bar chan<- chan<- *[]context.Context

	var (
		foo1 chan chan context.Context
		bar1 chan<- chan<- *[]context.Context
	)

	_, _, _, _ = foo, bar, foo1, bar1
}

func Bar107() {
	var foo []chan context.Context
	var bar *[]*<-chan context.Context

	var (
		foo1 []chan context.Context
		bar1 *[]*<-chan context.Context
	)

	_, _, _, _ = foo, bar, foo1, bar1

	type foox struct{ a *[]chan context.Context }
	type fooy struct {
		e *[]chan context.Context
		b []<-chan chan context.Context
	}

	_ = new([]chan context.Context)
	_ = new(*[]chan context.Context)
	_ = new(*[]<-chan context.Context)
	_ = new(*[]<-chan <-chan context.Context)

	_ = make([]chan context.Context, 1)
	_ = make([]<-chan context.Context, 1)
	_ = make([]chan<- context.Context, 1)
	_ = make([]chan chan context.Context, 1)
	_ = make([]<-chan *[]chan context.Context, 1)
	_ = make([]chan<- []chan<- context.Context, 1)

	a := []int{1, 2, 3}
	no := 10
	_ = make([]chan context.Context, 3, 5)
	_ = make([]<-chan context.Context, len(a))
	_ = make([]chan<- context.Context, len(a), 5)
	_ = make([]chan chan context.Context, no)
	_ = make([]<-chan *[]chan context.Context, len(a), no)
	_ = make([]chan<- []chan<- context.Context, 0, len(a))

}

func Bar108(*[]chan<- context.Context)

func Bar109(a *[]chan context.Context, b *[]<-chan context.Context)

func Bar110(a *[]chan context.Context, b *[]<-chan context.Context) {}

func Bar111(
	a []chan context.Context,
	b *[]<-chan context.Context,
	c *[]chan<- context.Context,
	d []chan<- chan context.Context,
) {
}

func Bar112(
	[]chan context.Context,
	*[]chan context.Context,
	*[]<-chan chan context.Context,
) {

}

func Bar113() {
	_ = new(newGen[newstruct1, newstruct1, newstruct2])
}

func Bar114() newstructG[newstruct1]

func Bar115() newstructG[newstruct1] {
	return newstructG[newstruct1]{}
}

func Bar116() map[newstructG[newstruct1]]context.Context {
	return make(map[newstructG[newstruct1]]context.Context)
}

func Bar117() chan chan context.Context

func Bar118() []chan context.Context

func Bar119() func(context.Context, func(string, context.Context), io.Writer) chan chan context.Context

func Bar120(chan chan context.Context) chan chan context.Context {
	return make(chan chan context.Context)
}

func Bar121(chan<- chan *context.Context) chan<- chan *context.Context {
	return make(chan<- chan *context.Context)
}

func Bar122(<-chan <-chan *[]context.Context) <-chan <-chan *[]context.Context {
	return make(<-chan <-chan *[]context.Context)
}

func Bar123(foo func() chan chan context.Context) func() chan chan context.Context {
	a := func() chan chan context.Context {
		return make(chan chan context.Context)
	}
	return a
}

func Bar124(a *[]chan context.Context) *[]chan context.Context

func Bar125(a *[]chan context.Context) *[]chan context.Context {
	return &[]chan context.Context{}
}

func Bar126(a *[]chan context.Context) (*[]chan context.Context, []<-chan context.Context)

func Bar127(a *[]chan context.Context) (foo *[]chan context.Context, bar *[]<-chan context.Context) {
	return &[]chan context.Context{}, nil
}

// v0.5.0
func Bar128() {
	var _ <-chan <-chan chan<- <-chan <-chan context.Context
}

func Bar129() {
	type foo0123456789 context.Context
	var foo **context.Context // all operators should be keyword.operator.address.go
	var bar **foo0123456789   // all operators should be keyword.operator.address.go
	fmt.Print(1*2 + 3)        // should be keyword.operator.arithmetic.go
	fmt.Print(4 * 5)          // should be keyword.operator.arithmetic.go
	_, _ = foo, bar
}

// # This is a heading
//
//   - A bullet.
//
//     Another paragraph of that first bullet.
//
//   - A second bullet.
//
// Package json implements encoding and decoding of JSON as defined in
// [RFC 7159]. The mapping between JSON and Go values is described
// in the documentation for the Marshal and Unmarshal functions.
//
// For an introduction to this package, see the article
// “[JSON and Go].”
//
// [RFC 7159]: https://tools.ietf.org/html/rfc7159
// [JSON and Go]: https://golang.org/doc/articles/json_and_go.html
func Bar130() {
	// example for go/doc: headings, lists, and links in Go doc comments
	// https://github.com/golang/proposal/blob/master/design/51082-godocfmt.md
}

func Bar131() {
	const (
		val1 = iota
		val2
		val3
		val4
	)
}

func Bar132() {
	// import declarations with semicolon between -before formatting with gofmt

	// // example-1
	// package main
	// import (foo "foobar/a"; bar "foobar/b")

	// func main(){
	// 	foo.Hello()
	// 	bar.Hi()
	// }

	// // example-2
	// package main
	// import ("foo";"bar")

	// // example-3
	// package main
	// import "foo"; import "bar"

	// // example-4
	// package main
	// import foo "foobar/a"; import bar "foobar/b"
}

func Bar133() {
	type foo struct {
		Context `json:"example"`
		x       context.Context
	}
	type bar struct {
		Context `json:"example"` // comment
		x       context.Context
	}
	type foobar struct {
		Context `json:"example"` /* comment */
		x       context.Context
	}

	// better property variables highlighting -before formatting with gofmt
	type foobarx struct {
		A       Context `json:"example"`  // comment
		B       Context `json:"example2"` // comment
		c, d, e Context /* comment */
	}
}

func Bar134(foo []struct{ context.Context }) {
}

func Bar135(foo *[]struct{ context.Context }) {
}

func Bar136(foo chan *[]struct{ context.Context }) {
}

func Bar137(foo <-chan *[]struct{ context.Context }) {
}

func Bar138(foo []chan<- chan *[]struct{ context.Context }) {
}

func Bar139(foo []context.Context, bar Context, foobar []struct {
	A       context.Context `json:"example"`  // comment
	B       context.Context `json:"example2"` // comment
	c, d, e context.Context
}) {
}

func Bar140(foo []context.Context, bar Context, foobar *[]struct {
	A       context.Context `json:"example"`  // comment
	B       context.Context `json:"example2"` // comment
	c, d, e context.Context
}) {
}

func Bar141(foo []context.Context, bar Context, foobar chan<- chan *[]struct {
	A       context.Context `json:"example"`  // comment
	B       context.Context `json:"example2"` // comment
	c, d, e context.Context
}) {
}

func Bar142(foo []context.Context, bar Context, foobar *[]chan<- chan *[]struct {
	A       context.Context `json:"example"`  // comment
	B       context.Context `json:"example2"` // comment
	c, d, e context.Context
}) {
}

func Bar143(foo chan<- *[]context.Context, bar func(a, b context.Context) context.Context, foobar []struct {
	Foo15
	Foo17
	Foo       Foo16
	FFoo, Bar string
}) {

}

func Bar144(foo chan<- *[]context.Context, bar func(a, b context.Context) context.Context, foobar chan<- *[]struct {
	Foo15
	Foo17
	Foo       Foo16
	FFoo, Bar string
}) {

}

func Bar145(
	foo chan<- *[]context.Context,
	bar func(a, b context.Context) context.Context,
	foobar chan<- chan *[]struct {
		Foo15
		Foo17
		Foo       Foo16
		FFoo, Bar string
	}) {

}

func Bar146() {
	type foo1 *[]context.Context
	type foo2 chan *[]context.Context
	type foo3 *chan context.Context
	type foo4 chan<- *[]context.Context
	type foo5 chan<- *chan<- *[]context.Context

	var _ *[]*<-chan context.Context
	var _ *<-chan context.Context
	var _ chan<- *context.Context

}

func Bar147(foo *[]context.Context) {}

func Bar148(foo <-chan *[]context.Context) {}

func Bar149(foo *[]<-chan *[]context.Context) {}

func Bar150(foo *[]map[*[]FooX]interface{}) {}

func Bar151(foo *[]map[string]string) {}

func Bar152() *[]chan context.Context {
	return &[]chan context.Context{}
}

func Bar153() *[]<-chan context.Context {
	return &[]<-chan context.Context{}
}

func Bar154() *[]chan<- chan context.Context {
	return &[]chan<- chan context.Context{}
}

func Bar155() []*context.Context {
	return []*context.Context{}
}

func Bar156() *[]context.Context {
	return &[]context.Context{}
}

func Bar157() {
	type foo *[]struct {
		a string
		b context.Context
	}

	var c foo

	c = &[]struct {
		a string
		b context.Context
	}{
		{a: "foo", b: context.TODO()},
		{a: "bar", b: context.TODO()},
	}

	for _, item := range *c {
		_ = fmt.Sprintf("f: %s, bar: %v\n", item.a, item.b)
	}
}

func Bar158() {
	type foo []struct {
		a string
		b context.Context
	}

	c := foo{{a: "foo", b: context.TODO()}, {a: "bar", b: context.TODO()}}
	_ = fmt.Sprintln(c)
}

func Bar159() {
	type foo []interface {
		foox(a, b, c context.Context) context.Context
	}

	type bar []interface {
		context.Context
	}

	type foobar *[]interface {
		context.Context
	}
}

func Bar160() {

	var foo []struct {
		a, b, c context.Context
	}

	var bar *[]struct {
		a, b, c context.Context
	}

	var (
		foo1 []struct {
			a, b, c context.Context
		}

		bar1 *[]struct {
			a, b, c context.Context
		}
	)

	var foox []struct {
		a string
		b context.Context
	} = []struct {
		a string
		b context.Context
	}{{a: "foo", b: context.TODO()}}

	var (
		barx []struct {
			a string
			b context.Context
		} = []struct {
			a string
			b context.Context
		}{{a: "foo", b: context.TODO()}}

		bary []struct {
			a string
			b context.Context
		} = []struct {
			a string
			b context.Context
		}{{a: "foo", b: context.TODO()}}
	)

	_, _, _, _, _, _, _ = foo, bar, foo1, bar1, foox, barx, bary
}

func Bar161() {
	type MyType int

	var myInt int = 42

	myTypeVar := (MyType)(myInt)

	_ = myTypeVar
}

func Bar162() {
	str := "foo"
	str1, str2, str3 := "foo", "bar", "foobar"

	switch str {
	case str1, str2: // comment

	case str3: /* comment */

	}

	switch {
	case str == str1: // comment

	case str > str2: // comment
	}

	switch str {
	case str1:
		// Code to execute
	case str2:
		// Code to execute
	case str3:
		// Code to execute
	}
}

func Bar163(
	context.Context, string, context.Context, // comment
) {

}

func Bar164(
	context.Context, // comment
	string, // comment
	func() chan<- context.Context, // comment
) {
}

func Bar165(
	foo, bar context.Context, // comment
) {
}

func Bar166(
	foo, // comment
	bar, // comment
	foobar context.Context, // comment
) {

}

func Bar167() {
	_ = func() []context.Context {
		return nil
	}
	_ = func() *[]context.Context {
		return nil
	}
	_ = func() []*context.Context {
		return nil
	}
	_ = func() map[string]context.Context {
		return nil
	}
	_ = func() *map[string]context.Context {
		return nil
	}
	_ = func() *[]map[string]context.Context {
		return nil
	}
	_ = func() []*map[string]context.Context {
		return nil
	}
}

func Bar168() {
	var foo interface{}
	_, _ = foo.(FooX)
	_, _ = foo.(*FooX)
	_, _ = foo.([]FooX)
	_, _ = foo.(chan FooX)
	_, _ = foo.(chan *FooX)
	_, _ = foo.(chan<- []FooX)
	_, _ = foo.(*[]chan<- *[]FooX)
	_, _ = foo.(chan<- []FooX)
	_, _ = foo.(func() FooX)
	_, _ = foo.(func() *[]FooX)
	_, _ = foo.(func() (string, FooX))
}

func Bar169() {
	_ = map[string]struct{ bar context.Context }{"foo": {bar: context.TODO()}}
	_ = map[string]struct {
		bar    context.Context
		foobar chan<- *context.Context
	}{
		"a": {
			bar: context.TODO(),
		},
		"b": {
			foobar: make(chan<- *context.Context),
		},
	}
	_ = map[string][]struct{ bar context.Context }{"foo": {{bar: context.TODO()}}}
}

func Bar170(func() context.Context, func() context.Context, func() context.Context) {}

func Bar171(
	func() chan context.Context,
	func() context.Context, func() context.Context,
) {
}

func Bar172(
	foo func() context.Context,
	bar context.Context,
	foobar ...func(a, b, c chan context.Context) http.Handler,
) {
}

func Bar173(
	foo func(a context.Context, b context.Context) (context.Context, context.Context),
	bar context.Context,
	foobar ...func(a, b, c chan context.Context) (d, e, f chan context.Context),
) {
}

type Options string

func Bar174(fset *token.FileSet, filename string, src []byte, opt *Options) (*ast.File, func(orig, src []byte) []byte, error) {
	return nil, nil, nil
}

func Bar175(*ast.File, func(orig, src []byte) byte) {
}

func Bar176(context.Context, func(http.Handler) http.Handler, context.Context) {
}

func Bar177(context.Context, func(http.Handler) chan<- http.Handler, chan context.Context) {
}

func Bar178(
	context.Context, func(http.Handler) chan<- http.Handler, chan context.Context) {
}

func Bar179(
	context.Context, func(a, b, c chan context.Context) (context.Context, context.Context), func(http.Handler) chan<- http.Handler, chan context.Context) {
}

func Bar180(func(a, b, c chan context.Context) (context.Context, context.Context), func(http.Handler) chan<- http.Handler, chan context.Context) {
}

func Bar181(func(a, b, c chan context.Context) (context.Context, context.Context), func(a, b, c chan context.Context) (context.Context, context.Context), func(a context.Context, b context.Context) (context.Context, context.Context)) {
}

func Bar182(
	func(a, b, c chan context.Context) (context.Context, context.Context), func(http.Handler) chan<- http.Handler, chan context.Context) {
}

func Bar183(
	func(a, b, c chan context.Context) (context.Context, context.Context), func(a, b, c chan context.Context) (context.Context, context.Context), func(a context.Context, b context.Context) (context.Context, context.Context)) {
}

type LongFooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo context.Context

func Bar184(LongFooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo, chan<- *[]LongFooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo) {
}

func Bar185(a LongFooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo, b chan<- *[]LongFooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo) {
}

func Bar186(
	a LongFooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo,
	b chan<- *[]LongFooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo) {
}

func Bar187(
	func(a, b, c chan context.Context) (context.Context, context.Context),
	func(http.Handler) chan<- http.Handler, chan context.Context,
) {
}

func Bar188(
	context.Context,
	func(http.Handler) chan<- http.Handler, chan context.Context) {
}

func Bar189(
	context.Context,
	func(http.Handler) chan<- http.Handler,
	chan context.Context,
) {
}

func Bar190(context.Context, context.Context, ...func(http.Handler) (context.Context, error)) {
}

func Bar191(context.Context, context.Context, func(http.Handler) (context.Context, error), func(http.Handler) (context.Context, error)) {
}

func Bar192(context.Context, context.Context, context.Context) {
}

func Bar193(context.Context, func(), context.Context, func(http.Handler) (context.Context, error)) {
}

func Bar194(a, b, c func() func(a, b, c context.Context) context.Context) {}

func Bar195(
	a context.Context,
	b func() func(context.Context, string) (context.Context, <-chan context.Context),
) {
}

func Bar196(context.Context, func() func(context.Context, string)) {}

func Bar197(func() func(context.Context, string)) {}

func Bar198() func() func(context.Context, string)

func Bar199(context.Context, func() func(context.Context, string) context.Context) {}

func Bar200(context.Context, func() func(context.Context, string) (context.Context, <-chan context.Context)) {
}

func Bar201(
	context.Context,
	func() func(context.Context, string) (context.Context, <-chan context.Context),
) {
}

func Bar202(
	context.Context,
	func() func(context.Context, string) (context.Context, <-chan context.Context),
	func(a, b, c context.Context) func(context.Context, string) (context.Context, <-chan context.Context)) {
}

type Bar203[E, V context.Context] []E

type Bar204[E, V any] []E

func Bar205[a Bar204[context.Context, Bar204[context.Context, context.Context]], b Bar204[Bar204[context.Context, context.Context], Bar204[string, Bar204[context.Context, context.Context]]]]() {
}
func Bar206[
	a Bar204[context.Context, Bar204[context.Context, context.Context]],
	b Bar204[Bar204[context.Context, context.Context], Bar204[string, Bar204[context.Context, context.Context]]]]() {
}

type Bar207[a Bar204[context.Context, Bar204[context.Context, context.Context]], b Bar204[Bar204[context.Context, context.Context], Bar204[string, Bar204[context.Context, context.Context]]]] interface{}

type Bar208[a Bar204[context.Context, Bar203[context.Context, context.Context]], b Bar204[string, Bar203[context.Context, context.Context]]] chan *context.Context

func Bar209() {
	type (
		a Bar204[context.Context, Bar203[context.Context, context.Context]]
		b Bar204[string, Bar203[context.Context, context.Context]]
		c map[string]interface{}
		d func(a context.Context, b string, c context.Context) (d context.Context, e context.Context)
		e func(a chan context.Context) context.Context
	)

	type (
		f = Bar204[context.Context, Bar203[context.Context, context.Context]]
		g = Bar204[string, Bar203[context.Context, context.Context]]
		h = map[string]interface{}
		j = func(a context.Context, b string, c context.Context) (d context.Context, e context.Context)
		k = func(a chan context.Context) context.Context
	)

}

func Bar210(
	context.Context,
	func() func(context.Context, string) (context.Context, <-chan context.Context),
) (
	context.Context,
	func() func(context.Context, string) (context.Context, <-chan context.Context),
) {
	a := func() func(context.Context, string) (context.Context, <-chan context.Context) {
		b := func(context.Context, string) (context.Context, <-chan context.Context) {
			return context.TODO(), make(<-chan context.Context)
		}
		return b
	}
	return context.TODO(), a
}

func Bar211(
	context.Context,
	func() func(context.Context, string) (context.Context, <-chan context.Context),
) (context.Context, func() func(context.Context, string) (context.Context, <-chan context.Context)) {
	a := func() func(context.Context, string) (context.Context, <-chan context.Context) {
		b := func(context.Context, string) (context.Context, <-chan context.Context) {
			return context.TODO(), make(<-chan context.Context)
		}
		return b
	}
	return context.TODO(), a
}

func Bar212() {
	_ = func() func(context.Context, string) (context.Context, <-chan context.Context) {
		foo := func(context.Context, string) (context.Context, <-chan context.Context) {
			return context.TODO(), make(<-chan context.Context)
		}
		return foo
	}
	_ = func(a, b, c func() func(a, b, c context.Context) context.Context) {

	}
}

func Bar213() {
	_ = func(a, b, c func() context.Context, e context.Context) {
	}
	_ = func(foo context.Context, a, b, c func() context.Context) {
	}
}

func Bar214() {

	type foo map[string]func(args ...interface{}) context.Context

	_ = make(map[string]func(args context.Context) context.Context)
	_ = make(map[string]func() (context.Context, context.Context))
	_ = make(map[string]func() context.Context)

	_ = func() func(context.Context, string) (context.Context, <-chan context.Context) {
		a := func(context.Context, string) (context.Context, <-chan context.Context) {
			return context.TODO(), make(<-chan context.Context)
		}
		return a
	}

	_ = func(a, b, c func(), e context.Context) {
	}

	_ = func(a, b, c func() (context.Context, context.Context), foo context.Context) {
	}

}

func Bar215() {
	_ = func() (context.Context, func(), error, context.Context) {
		a := func() {}
		return context.TODO(), a, nil, context.TODO()
	}

	_ = func() (context.Context, func() (context.Context, context.Context), error, context.Context) {
		a := func() (context.Context, context.Context) {
			return context.TODO(), context.TODO()
		}
		return context.TODO(), a, nil, context.TODO()
	}

	_ = func() (context.Context, func(a, b, c context.Context) (context.Context, context.Context), error, context.Context) {
		a := func(a, b, c context.Context) (context.Context, context.Context) {
			return context.TODO(), context.TODO()
		}
		return context.TODO(), a, nil, context.TODO()
	}

	_ = func() (a context.Context, b func(), c io.ReadCloser, d error, e context.Context) {
		return
	}

	_ = func() (a func(), b context.Context, c io.ReadCloser, d error, e context.Context) {
		return
	}

	_ = func() (func() (context.Context, context.Context), context.Context, error, context.Context) {
		a := func() (context.Context, context.Context) {
			return context.TODO(), context.TODO()
		}
		return a, context.TODO(), nil, context.TODO()
	}

	_ = func() (func(a, b, c context.Context) (context.Context, context.Context), context.Context, error, context.Context) {
		foo := func(a, b, c context.Context) (context.Context, context.Context) {
			return context.TODO(), context.TODO()
		}
		return foo, context.TODO(), nil, context.TODO()
	}

	_ = new(map[string]func(args context.Context) context.Context)
	_ = new(map[string]func() (context.Context, context.Context))
	_ = new(map[string]func() context.Context)
	_ = make(map[string]func() (context.Context, string))
}

func Bar216() Bar204[context.Context, Bar204[context.Context, context.Context]] {
	return nil
}

func Bar217() Bar204[Bar204[context.Context, context.Context], Bar204[string, Bar204[context.Context, context.Context]]] {
	return nil
}

func Bar218() (a Bar204[context.Context, Bar204[context.Context, context.Context]], b Bar204[Bar204[context.Context, context.Context], Bar204[string, Bar204[context.Context, context.Context]]]) {
	return
}

func Bar219() *[]chan<- *[]Bar204[Bar204[context.Context, context.Context], context.Context] {
	return nil
}

func Bar220() {
	var a, b, c func(context.Context, string) (context.Context, context.Context) // comment

	var foo func(context.Context, string) (context.Context, context.Context)
	var bar func() func(context.Context, string) (context.Context, <-chan context.Context)

	var (
		ax, bx, cx func(context.Context, string) (context.Context, context.Context)
		foox       func(context.Context, string) (context.Context, context.Context)
		barx       func() func(context.Context, string) (context.Context, <-chan context.Context)
	)

	_, _, _, _, _, _, _, _, _, _ = a, b, c, ax, bx, cx, foo, bar, foox, barx
}

// v0.5.1
func Bar221() {
	type Bar struct{}

	type AppConfig struct {
		Brand    string                 `json:"brand"`
		BundleID string                 `json:"bundleId"`
		AppName  string                 `json:"appName"`
		Skin     string                 `json:"skin"`
		Foo      *string                `json:"foo"`
		FooX     chan<- context.Context `json:"foox"`
		Ctx      context.Context        `json:"ctx"`
		Bar      Bar                    `json:"bar"`
		Barx     *Bar                   `json:"barx"`
		FooBar   *Bar                   `json:"foobar"`
		TimerFoo time.Duration          `json:"timerfoo"`
		TimerBar time.Duration          `json:"timerbar"`
	}

	config := AppConfig{
		Brand:    "foo",
		BundleID: "foo",
		AppName:  "foo",
		Skin:     "foo",
	}

	// Update(v0.7.3): "variable.other.property.field.go" is deprecated.
	appConfig := AppConfig{
		Brand:    config.Brand,    // comment
		BundleID: config.BundleID, /* comment */
		AppName:  config.AppName,
		Skin:     config.Skin,
		Foo:      &config.Brand,
		FooX:     make(chan<- context.Context),
		Ctx:      context.TODO(),
		Bar:      Bar{},
		Barx:     &Bar{},
		FooBar:   new(Bar),
		TimerFoo: time.Second,
		TimerBar: 5 * time.Second,
	}

	_ = appConfig
}

func Bar222() {
	type Foo[
		Input, Output context.Context,
	] func(ctx context.Context, input Input) (output Output, err error)

	type Foo2[
		Input, Output context.Context,
	] context.Context

	type Foo3[
		Input,
		Output context.Context,
	] *[]<-chan *[]context.Context

	type Foo4[
		Input, Output context.Context,
	] struct {
		context.Context
		a, b, c chan context.Context
		foo     context.Context
		bar     func(a context.Context) error
	}

	type Foo5[
		Input,
		Output context.Context,
	] *[]struct {
		context.Context
		a, b, c chan context.Context
		foo     context.Context
		bar     func(a context.Context) error
	}

	type Foo6[
		Input,
		Output context.Context] *[]chan<- *[]struct {
		context.Context
		a, b, c chan context.Context
		foo     context.Context
		bar     func(a context.Context) error
	}

	type Foo7[
		Input, Output context.Context,
	] *[]<-chan *map[context.Context]context.Context

	type Foo8[
		Input, Output context.Context,
	] *[]<-chan *[]func(ctx context.Context, input Input) (output Output, err error)

	type Foo9[
		Input, Output context.Context,
	] *[]<-chan *map[context.Context]map[context.Context]interface{}

	type Foo10[
		Input,
		Output context.Context] map[context.Context]interface{}

	type (
		Bar[
			Input, Output context.Context,
		] func(ctx context.Context, input Input) (output Output, err error)

		Bar2[
			Input, Output context.Context,
		] context.Context

		Bar3[
			Input,
			Output context.Context,
		] *[]<-chan *[]context.Context

		Bar4[
			Input, Output context.Context,
		] struct {
			context.Context
			a, b, c chan context.Context
			foo     context.Context
			bar     func(a context.Context) error
		}

		Bar5[
			Input,
			Output context.Context,
		] *[]struct {
			context.Context
			a, b, c chan context.Context
			foo     context.Context
			bar     func(a context.Context) error
		}

		Bar6[
			Input,
			Output context.Context] *[]chan<- *[]struct {
			context.Context
			a, b, c chan context.Context
			foo     context.Context
			bar     func(a context.Context) error
		}

		Bar7[
			Input, Output context.Context,
		] *[]<-chan *map[context.Context]context.Context

		Bar8[
			Input, Output context.Context,
		] *[]<-chan *[]func(ctx context.Context, input Input) (output Output, err error)

		Bar9[
			Input, Output context.Context,
		] *[]<-chan *map[context.Context]map[context.Context]interface{}

		Bar10[
			Input,
			Output context.Context] map[context.Context]interface{}
	)

}

func Bar223() {
	// before formatting with gofmt
	// type foo struct{context.Context; a chan context.Context; b chan context.Context; c chan context.Context; foo context.Context; bar func(a context.Context) error}
	type foo struct {
		context.Context
		a   chan context.Context
		b   chan context.Context
		c   chan context.Context
		foo context.Context
		bar func(a context.Context) error
	}
}

func Bar224(a []struct {
	context.Context
	a, b, c chan context.Context
	foo     context.Context
	bar     func(a context.Context) error
}, b context.Context) {

}

func Bar225(a, b, c context.Context) ([]struct {
	context.Context
	a, b, c chan context.Context
	foo     context.Context
	bar     func(a context.Context) error
}, error) {
	return nil, nil
}

func Bar226() []struct {
	context.Context
	a, b, c chan context.Context
	foo     context.Context
	bar     func(a context.Context) error
}

func Bar227(a, b, c context.Context) []struct {
	context.Context
	a, b, c chan context.Context
	foo     context.Context
	bar     func(a context.Context) error
} {
	return nil
}

func Bar228() *[]struct {
	context.Context
	a, b, c chan context.Context
	foo     context.Context
	bar     func(a context.Context) error
}

func Bar229(a, b, c context.Context) *[]struct {
	context.Context
	a, b, c chan context.Context
	foo     context.Context
	bar     func(a context.Context) error
} {
	return nil
}

func Bar230() *[]<-chan []struct {
	context.Context
	a, b, c chan context.Context
	foo     context.Context
	bar     func(a context.Context) error
}

func Bar231(a, b, c context.Context) *[]chan<- *[]struct {
	context.Context
	a, b, c chan context.Context
	foo     context.Context
	bar     func(a context.Context) error
} {
	return nil
}

func Bar232(a, b, c context.Context) struct {
	context.Context
	a, b, c chan context.Context
	foo     context.Context
	bar     func(a context.Context) error
} {
	return struct {
		context.Context
		a   chan context.Context
		b   chan context.Context
		c   chan context.Context
		foo context.Context
		bar func(a context.Context) error
	}{}
}

func Bar233() struct {
	context.Context
	a, b, c chan context.Context
	foo     context.Context
	bar     func(a context.Context) error
}

func Bar234() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
	}

}

func Bar235() {
	// better pre-highlighting variables after control keywords -before formatting with gofmt
	arr := []bool{true, false}
	num := 1
	a, b := 0, 0

	// if arr[num-1]{
	// }
	// if arr[num+1]{
	// }
	// if arr[num%1]{
	// }
	// if arr[num/1]{
	// }

	// if num > a+ b{
	// }
	// if num > a% b{
	// }
	// if num > a/ b{
	// }
	// if num > a- b{
	// }
	// if num > a* b{
	// }

	// for i:=0; i<10; a+= b{
	// }

	// for i:=0; i<10; a-= b{
	// }

	// for i:=0; i<10; a/= b{
	// }

	// for i:=0; i<10; a*= b{
	// }

	// for i:=0; i<10; a%= b{
	// }

	if arr[num+1] {
		// implement
	}

	if num > a*b {
		// implement
	}

	for i := 0; i < 10; a += b {
		// implement
	}

}

// v0.5.2
func Bar236() {
	// better pre-highlighting variables after control keywords -before formatting with gofmt
	nums := []int{1, 2, 3, 4, 5}
	ok, notOk := true, false
	// for i := range nums {
	// 	for j := range nums[i+1:]{
	// 		_ = j
	// 	}
	// }

	// if ok || notOk{
	// }
	// if ok && notOk{
	// }

	for i := range nums {
		for j := range nums[i+1:] {
			_ = j
		}
	}

	if ok || notOk {
		// implement
	}

}

// v0.5.4
func Bar237() {
	var a1 string
	var a2, a3, a4 string = "foo", "foo", "foo"
	var a5 = "foo"
	var a6 string = "foo"

	a1, a2, a3 = a3, a4, a5
	a3, a4 = a1, a2
	a5 = a6

	b1, b2, b3, b4, b5 := "foo", "foo", "foo", "foo", "foo"
	b6 := "foo"

	b1, b2, b3 = b3, b4, b5
	b3, b4 = b1, b2
	b5 = b6

	var (
		c1, c2, c3                 = "foo", "foo", "foo"
		c4, c5     string          = "foo", "foo"
		c6         context.Context = context.TODO()
	)

	c1, c2, c3 = c3, c4, c5
	c4, c5 = c1, c2
	_ = c6

	a, b, c := new(string), new(string), new(string)
	*a, *b, *c = "foo", "foo", "foo"
}

func Bar238() {
	arr := []int{1, 2, 3, 4, 5}
	start := 0
	capacity := 3
	end := 2
	_ = arr[start:]
	_ = arr[:end]
	_ = arr[start:end]
	_ = arr[start:end:capacity]

	type Foo struct {
		end int
	}

	e := Foo{end: 2}
	_ = arr[e.end:]
	_ = arr[:e.end]
	_ = arr[start:e.end]
	_ = arr[start:e.end:capacity]
}

// v0.5.5
func Bar239[T any]() {
	t := reflect.TypeOf((*T)(nil)).Elem()
	_ = fmt.Sprintf("%v: %v", reflect.TypeOf((*T)(nil)), t)
}

func Bar240() {
	type Foo struct {
		min, max int
		nums     []Foo
	}

	type Bar struct {
		a1, a2, a3 bool
		a4, a5, a6 int
	}

	a := Foo{min: 2, max: 4, nums: []Foo{{min: 1, max: 2}}}

	b := Bar{
		a1: a.max == a.min,
		a2: a.max == a.min || a.max > a.min,
		a3: a.min != a.max && a.min <= a.max,
		a4: a.max + a.min,
		a5: a.max / a.min,
		a6: a.nums[0].min - a.nums[0].max,
	}

	_ = b
}

func Bar241() {
	arr := []int{1, 2, 3, 4, 5}

	type Foo struct {
		min, max int
	}

	f := Foo{min: 1, max: 4}

	_ = arr[f.max+f.min]
	_ = arr[f.max-f.min:]
	_ = arr[f.max+f.min:]
	_ = arr[f.max/f.min:]
}

// v0.5.6
func Bar242() {
	type Foo struct {
		min, max int
	}

	f := Foo{min: 1, max: 4}
	_ = Foo{
		max: f.max | f.min,
		min: f.max & f.min,
	}
}

func Bar243() {
	// better pre-highlighting variables after control keywords -before formatting with gofmt
	ok, notOk := true, false
	// if ok&&notOk{
	// }

	if ok || notOk {
		// implement
	}

}

func Bar244() {
	arr := []int{1, 2, 3, 4, 5}

	type Foo struct {
		min, max int
	}

	f := Foo{min: 1, max: 4}

	_, _ = arr[f.max|f.min:], arr[f.max&f.min:]
}

func Bar245() {
	var foo1 chan<- struct {
		a context.Context
	}
	var foo2 <-chan func(a context.Context)
	var foo3 []<-chan *[]func(a, b, c context.Context) <-chan context.Context

	var (
		bar1 chan<- struct {
			a context.Context
		}
		bar2 <-chan func(a context.Context)
		bar3 []<-chan *[]func(a, b, c context.Context) <-chan context.Context
	)

	_, _, _, _, _, _ = foo1, foo2, foo3, bar1, bar2, bar3
}

// v0.5.7
func Bar246() {
	type Foo map[string]int // [bar.foo] foo.bar
}

func Bar247() {
	var a map[string]interface{}                // [bar.foo] foo.bar
	var b map[string][][][][][]*context.Context /* [bar.foo] foo.bar */

	var (
		c map[string]interface{}                // [bar.foo] foo.bar
		d map[string][][][][][]*context.Context /* [bar.foo] foo.bar */
	)

	_, _, _, _ = a, b, c, d
}

func Bar248() {
	type Foo[a any] string // [bar.foo] foo.bar

	type Bar Foo[context.Context] // [bar.foo] foo.bar
}

func Bar249() {
	var a map[context.Context]chan<- context.Context
	var b map[context.Context][]*chan<- context.Context

	var (
		c map[context.Context]chan<- context.Context
		d map[context.Context][]*chan<- context.Context
	)

	_, _, _, _ = a, b, c, d
}

func Bar250() {
	type Foo[a any] struct{}    // [bar.foo] foo.bar
	type Boo[a any] interface{} // [bar.foo] foo.bar
}

func Bar251() {
	type foo string
	_ = make(chan (foo))
}

func Bar252() {
	type foo string
	var a chan (foo)
	var (
		b chan (foo)
	)
	_, _ = a, b
}

func Bar253() {
	_ = func() (error, string) { // func() (error, string) { return nil, "" } foo.bar
		return nil, ""
	}
}

func Bar254[a any](b a) a {
	return b
}

func Bar255() {
	type foo struct {
		s1 map[string]func(args ...interface{}) context.Context
		s2 map[rune]func(args ...interface{}) context.Context
		s3 map[int]func(args ...interface{}) context.Context
	}
	bar := foo{}
	bar.s1["foo"]()
	bar.s2['f']()
	bar.s3[0]()

	type foox string
	var barx map[foox]interface{}
	_, _ = Bar254[map[foox]interface{}](barx), Bar254[string]("bar") // [foo.bar](foo) bar.foo
}

func Bar256() {
	type foo struct{}
	type bar[a context.Context, b foo] struct{}

	var a bar[context.Context, foo]

	var (
		b bar[context.Context, foo]
	)

	_, _ = Bar254[bar[context.Context, foo]](a), Bar254[bar[context.Context, foo]](b)

}

// v0.5.8
func Bar257() {
	var c Context

	type (
		context1 Context
		context2 Context
		context3 Context
		context4 Context
		context5 Context
		context6 cancelCtx
		context7 cancelCtx
		context8 cancelCtx
		context9 cancelCtx
	)

	switch ctx := c.(type) {
	case context1, context2, /* foo */
		context3, context4: // bar
		fmt.Println(ctx)
	case context5: // foo
		fmt.Println(ctx)
	case *context6, // foo
		*context7, // bar
		*context8, /* foo */
		*context9: /* bar */
		fmt.Println(ctx)
	}
}

// v0.5.9
func Bar258() {
	foo := func(i int) func(func(int, string) bool) {
		return func(f func(int, string) bool) {
			result := f(i, "foo")
			fmt.Println("result:", result)
		}
	}

	value := 10

	foo(value)(func(i int, s string) bool {
		return i == 10
	})
}

// v0.6.0
func Bar259() {
	ch := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for {
		select {
		case s, ok := <-ch:
			if ok {
				fmt.Println(s)
			}
		case <-ctx.Done():
			fmt.Println("foo")
		}
	}
}

func Bar260() {
	for _, s := range []string{"foo", "bar"} {
		fmt.Println(s)
	}

	type foo struct {
		num int
	}

	for _, i := range []foo{{1}, {2}, {3}} {
		fmt.Println(i.num)
	}

	type bar []string

	for _, s := range []bar{[]string{"foo", "bar"}} {
		fmt.Println(s)
	}

	type foobar map[string]interface{}

	a := foobar{
		"name": "Foo",
		"age":  100,
	}

	b := foobar{
		"name": "Bar",
		"age":  100,
	}

	for _, s := range []foobar{a, b} {
		fmt.Println(s)
	}

}

// v0.6.1

// better pre-highlighting types after function declaration -before formatting with gofmt
// func Bar()context.Context

func Bar261() {
	// better pre-highlighting variables after control keywords -before formatting with gofmt
	foo := new(bool)
	bar := true
	foo = &bar

	// if *foo{
	// }

	if *foo {
		fmt.Println("foo")
	}
}

func Bar262() {
	_ = make([]float64, 1024*10)
	_ = make([]context.Context, 10/2)
	_ = make([]context.Context, 10*2)
	_ = make([]context.Context, 10-2)
	_ = make([]context.Context, 10+2)
	_ = make([]context.Context, 10%2)
	_ = make([]context.Context, 10|2)
	_ = make([]context.Context, 10&2)
	_ = make([]context.Context, 10<<2)
	_ = make([]context.Context, 10>>2)
}

type Bar263 interface {
	map[context.Context]func(a context.Context) context.Context | map[string]func(a, b, c context.Context)
}

type Bar264 interface {
	Bar204[string, Bar203[context.Context, context.Context]]
}

type Bar265 interface {
	Bar204[Bar204[Bar204[context.Context, context.Context], Bar204[string, Bar204[context.Context, context.Context]]], Bar204[string, Bar204[context.Context, context.Context]]]
}

type Bar266 interface {
	Bar204[string, Bar203[context.Context, context.Context]] | Bar204[Bar204[context.Context, context.Context], Bar204[string, Bar204[context.Context, context.Context]]]
}

// v0.6.2
func Bar267() {
	// functions inline with multi return types -before formatting with gofmt
	// _ = func(a, b, c string) (func(a string) (context.Context, error), error) { a = "foo"; b = "bar"; return func(a string) (context.Context, error) { fmt.Println(a, b, c); return context.Background(), nil }, nil } // foo
	_ = func(a, b, c string) (func(a string) (context.Context, error), error) {
		a = "foo"
		b = "bar"
		return func(a string) (context.Context, error) { fmt.Println(a, b, c); return context.Background(), nil }, nil
	}

	_ = func() (context.Context, func(), error, context.Context) {
		a := func() {}
		return context.TODO(), a, nil, context.TODO()
	}

	_ = func(func(a, b, c string) context.Context) (func() context.Context, context.Context) {
		f := func() context.Context {
			return context.TODO()
		}
		return f, context.TODO()
	}
}

// v0.6.3
func Bar268(foo func(a string, b interface {
	Bar() context.Context
})) {

}

func Bar269[T interface {
	Bar() context.Context
}](foo T) {
	foo.Bar()
}

type Bar270 struct {
	foo interface {
		A() context.Context
	}
	bar context.Context
}

type Bar271 struct {
	foo func(a string, b interface {
		A() context.Context
	})
	bar context.Context
}

// v0.6.4
func Bar272(foo func(a string, b chan<- interface {
	Bar() context.Context
})) {

}

func Bar273[T chan interface {
	Bar() context.Context
}](foo T) {

}

type Bar274 struct {
	foo chan interface {
		A() context.Context
	}
	bar context.Context
}

type Bar275 struct {
	foo chan func(a string, b chan interface {
		A() context.Context
	})
	bar context.Context
}

type Bar276 struct {
	foo chan struct {
		foo chan<- func(a string, b <-chan interface {
			A() context.Context
		})
		bar chan<- context.Context
	}
	bar context.Context
}

type Bar277 struct {
	foo *[]chan<- *[]struct {
		foo *[]chan<- *[]func(a string, b <-chan interface {
			A() context.Context
		})
		bar chan<- context.Context
	}
	bar context.Context
}

// v0.6.5
type Type string

type Type1 string

func Bar278[T ~[]Type | ~map[Type]Type1](t T) {}

func Bar279[T *[]Type | ~map[Type]string](t T) {}

func Bar280[T *[]Type | ~map[*[]Type]map[*[]Type]Type](t T) {}

func Bar281[T *[]map[*[]func(a context.Context) (b string, c context.Context)]interface{}](foo T) {

}

func Bar282(foo *[]map[*[]func(a context.Context) (b string, c context.Context)]interface{}) {

}

// v0.6.6

// better support for struct type one line with semicolon(;) -before formatting with gofmt
// type oneLineSemi struct{a1, a2, a3 context.Context; a4, a5 string; a6 context.Context; a7 context.Context; context.Context; a8, a9 context.Context}

type Bar283 = struct{ context.Context }

type Bar284 = struct{ foo context.Context }

type Bar285 = []struct{ foo <-chan context.Context }

type Bar286 = struct {
	foo    context.Context
	bar    string
	foobar *[]<-chan *[]context.Context
}

type Bar287 = []struct {
	foo    context.Context
	bar    string
	foobar *[]<-chan *[]context.Context
}

// v0.6.8
func Bar288() {
	// better highlighting struct properties and types when hovering with the mouse
	type Foo struct{ name context.Context }
	_ = Foo{name: context.TODO()}
}

// v0.6.9
func Bar289() {
	a1 := []int{10}
	a2 := 20
	foo := func([]byte) error {
		return nil
	}

	var (
		_ = foo(make([]byte, len(a1)*int(a2)))
	)

	if 5 > 10 {

	}

	fmt.Println("Hello World!")
}

// v0.7.0
// Despite the warnings, the following syntax still should be supported without causing any bugs.
// Warning: struct field tag "json:\"bar\n\t,omitempty\"" not compatible with reflect.StructTag.Get: bad syntax for struct tag value
type Bar290 struct {
	FooX `json:"foox
	,omitempty"` // foo
	Type `json:"type"
	,omitempty"`
	Bar string `json:"bar
	,omitempty"` // bar
	Foobar string `json"foobar
	,omitempty"`
}

type Bar291 struct {
	Foo string `json:"foo"
		foo bar 'string'
	` // foo bar
	Bar string // foo bar
}

// v0.7.2
func Bar292() {
	// struct field hover with generics and functions
	type bar1 struct{}
	type bar2 struct{}

	type bar[b1 bar1, b2 bar2] struct{}

	type foo struct {
		a func(foo, bar, foobar context.Context)
		b bar[bar1, bar2]
	}
	f := foo{}
	_ = f.a
	_ = f.b
}

func Bar293() {
	type bar[a, b any] struct{}

	type foo[a, b string] bar[context.Context, context.Context]

	type foobar[a, b string] chan<- bar[context.Context, context.Context]
}

// v0.7.3
func Bar294() {
	type Bar struct{}

	type Foo struct {
		bar map[context.Context]chan *Bar
	}

	_ = &Foo{
		bar: map[context.Context]chan *Bar{},
	}
}

// v0.7.4
func Bar295() {
	type (
		FooBar interface {
			Foo(ctx context.Context) context.Context
			Bar(ctx context.Context) <-chan context.Context
		}
	)
}

// v0.7.5
func Bar296() {
	ctx := []Context{context.TODO(), context.TODO()}

	type (
		context1 Context
		context2 cancelCtx
	)

	i := 0
	switch c := ctx[i+1].(type) {
	case context1,
		context2:
		fmt.Println(c)
	case Context:
		fmt.Println(c)
	default:
		fmt.Println(c)
	}
}

// v0.7.6
func Bar297() {
	type Foo string
	type Bar[T any] string
	type Baz string

	var foo map[Foo]*Bar[Baz]
	var bar map[Foo]newGen[newstruct1, newstruct1, newstruct2]

	_, _ = foo, bar
}

// v0.7.7
func Bar298() {
	type (
		bar1 string
		bar2 string
		bar3 string
		bar4 string
		bar5 string
		bar6 string
		bar7 string
		bar8 string
		bar9 string
	)

	type Foo struct {
		bar1 `"foo"`
		bar2 "`foo`" // `foo` bar
		bar3 `foo`   /* `foo` bar */
		bar4 `foo
		bar "foobar"
		foo bar baz` // `foo` "bar"
		bar5 // `"foo"` bar
		bar6 // "`foo`" bar
		bar7 // `foo` bar
		bar8 /* `foo bar` foo bar */
		bar9 /* foo
		`foo` `bar` baz
		"`foobar`" `"baz"` */
		foo1 string `"foo"`
		foo2 string "`foo`" // `foo` bar
		foo3 string `foo`   /* `bar` bar */
		foo4 string `foo
		bar "foobar"
		foo bar baz` // `foo` "bar"
		foo5 string // `"foo"` bar
		foo6 string // "`foo`" bar
		foo7 string // `foo` bar
		foo8 string /* `foo bar` foo bar */
		foo9 string /* foo
		`foo` `bar` baz
		"`foobar`" `"baz"` */
	}
}

// v0.7.8
func Bar299() {
	type anyX[a any] any
	type Constraint = int
	type G[P Constraint] = anyX[string]

	type Map[K comparable, V any] context.Context
	type Set[K comparable] = Map[K, bool]

	type integers interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}
	type IntSet[K integers] = Set[K]

	type Point3D[E any] = struct{ x, y, z E }
	p := Point3D[float64]{1.0, 2.71828, 3.14159}
	fmt.Println(p)

	type IntSet1[A, B integers, C, D string] = Set[A]
	type IntSet2[A, B integers, C, D string] = Set[B] // foo  "" `` '' [] bar
	type IntSet3[A, B integers, C, D string] = Set[C] /* foo  "" `` '' [] bar */
	type IntSet4[A, B integers, C, D string] = Set[D] /* foo
	"" `` '' [] bar */

	type Foo1[T comparable] = Map[T, string]

	type (
		Foo2[T comparable] = Map[T, string]

		IntSet5[A, B integers, C, D string] = Set[A]
		IntSet6[A, B integers, C, D string] = Set[B] // foo  "" `` '' [] bar
		IntSet7[A, B integers, C, D string] = Set[C] /* foo  "" `` '' [] bar */
		IntSet8[A, B integers, C, D string] = Set[D] /* foo
		"" `` '' [] bar */
	)

	fmt.Println("Hello, World!")
}

// v0.7.9
func Bar300() func() int {
	return func() int {
		return 5
	}
}

var _ = Bar301()(func() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
})

func Bar301() func(...interface{}) bool {
	return func(...interface{}) bool {
		return true
	}
}

func Bar302() func() { return nil }

var _ = (func() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
})

// v0.8.0
func Bar303() {
	const foo = "bar"

	const (
		BAR = 1
	)
}

// v0.8.1
type Bar304 comparable

type slot[K comparable, V any] struct {
	key   K
	value V
}

type group[K comparable, V any] struct {
	slots [10]slot[K, V]
}

func (g *group[K, V]) Bar305(i uint32) *slot[K, V] {
	return (*slot[K, V])(unsafe.Add(unsafe.Pointer(&g.slots[0]), uintptr(i)*unsafe.Sizeof(g.slots[0])))
}

func Bar306() {
	type foo[K, V, M any] struct{}
	type foox int
	type fooy string
	bar1 := foo[foox, fooy, string]{}
	bar2 := &foo[foox, fooy, string]{}
	bar3 := &[]foo[foox, fooy, string]{}
	bar4 := []foo[foox, fooy, string]{}
	bar5 := &foo[foox, fooy, map[foox]struct{}]{}
	_ = (foo[foox, fooy, string])(bar1)
	_ = (*foo[foox, fooy, string])(bar2)
	_ = (*[]foo[foox, fooy, string])(bar3)
	_ = (*[]foo[foox, fooy, string])(&bar4)
	_ = ([]foo[foox, fooy, string])(bar4)
	_ = (*foo[foox, fooy, map[foox]struct{}])(bar5)
}

func Bar307() {
	type foo[T any] struct{}
	type bar int
	type baz[T any] string
	type foobar[K, V any] bool
	type foobarbaz[K, V, M any] struct{}
	_ = foo[map[bar]struct{}]{}
	_ = []foo[map[string]struct{}]{}
	_ = &foo[map[string]struct{}]{}
	_ = &foo[map[bar]bar]{}
	_ = &foo[baz[baz[bar]]]{}
	_ = &foo[foobar[bar, baz[struct{}]]]{}
	_ = &[]foobarbaz[foo[bar], foo[bar], bool]{}
	_ = &foobarbaz[foo[bar], baz[bar], foobar[string, bar]]{}
	_ = &foobarbaz[foo[bar], baz[bar], foobar[string, bar]]{}
	_ = foo[foobar[foobar[foobar[foobar[foobar[foobar[foobar[bar, string], map[string]interface{}], map[string]struct{}], struct{}], bar], bar], bar]]{}
}

func Bar308() {
	type foo[T any] struct{}
	type baz[T any] struct{}
	type foobar[K, V, M any] struct{}
	type foox int
	type fooy string
	type fooz bool

	bar1 := &foobar[foo[string], baz[int], fooz]{}
	bar2 := &foobar[foo[foox], baz[fooy], foobar[foox, fooy, fooz]]{}
	bar3 := &[]foobar[foo[foox], baz[fooy], foobar[foox, fooy, fooz]]{}
	bar4 := foobar[foo[map[string]interface{}], map[string]struct{}, foobar[foox, fooy, map[string]struct{}]]{}
	bar5 := &foobar[foobar[foobar[foobar[foobar[foobar[foobar[foox, fooy, fooz], fooy, fooz], fooy, fooz], foo[foox], baz[fooy]], foo[foox], baz[fooy]], foo[foox], baz[fooy]], foo[foox], baz[fooy]]{}
	_ = (*foobar[foo[string], baz[int], foox])(bar1)
	_ = (foobar[foo[string], baz[int], fooy])(*bar1)
	_ = (*foobar[foo[foox], baz[fooy], foobar[foox, fooy, fooz]])(bar2)
	_ = (*foobar[foo[foox], baz[fooy], foobar[foox, fooy, fooz]])(bar2)
	_ = (*[]foobar[foo[foox], baz[fooy], foobar[foox, fooy, fooz]])(bar3)
	_ = ([]foobar[foo[foox], baz[fooy], foobar[foox, fooy, fooz]])(*bar3)
	_ = (foobar[foo[map[string]interface{}], map[string]struct{}, foobar[foox, fooy, map[string]struct{}]])(bar4)
	_ = (*foobar[foobar[foobar[foobar[foobar[foobar[foobar[foox, fooy, fooz], fooy, fooz], fooy, fooz], foo[foox], baz[fooy]], foo[foox], baz[fooy]], foo[foox], baz[fooy]], foo[foox], baz[fooy]])(bar5)
	_ = (foobar[foobar[foobar[foobar[foobar[foobar[foobar[foox, fooy, fooz], fooy, fooz], fooy, fooz], foo[foox], baz[fooy]], foo[foox], baz[fooy]], foo[foox], baz[fooy]], foo[foox], baz[fooy]])(*bar5)
}

// v0.8.2
func Bar309() {
	a := 100
	b := 50
	c := 30
	_ = 100 - 50&30
	_ = a - b&c

	type foo struct {
		x, y float64
	}

	f := &foo{x: 100, y: 20}
	_ = math.Sqrt(f.x*f.x + f.y*f.y)
	_ = math.Sqrt((f.x * f.x) + (f.y * f.y))
}

// v0.8.3
func Bar310() {

	type (
		foo[x, y, z any]    struct{}
		foobar[K, V, M any] struct{}
		baz[T any]          struct{}

		x int
		y int
		z int
	)

	var bar any

	_, _ = bar.(foo[x, y, z])
	_, _ = bar.(*foo[x, y, z])
	_, _ = bar.([]foo[x, y, z])
	_, _ = bar.([]*foo[x, y, z])
	_, _ = bar.(*[]foo[x, y, z])
	_, _ = bar.(foobar[foobar[foobar[foobar[foobar[foobar[foobar[x, y, z], x, y], x, y], foo[x, y, z], baz[x]], baz[y], baz[z]], foo[x, y, z], baz[y]], foo[x, y, z], baz[z]])
	_, _ = bar.(*foobar[foobar[foobar[foobar[foobar[foobar[foobar[x, y, z], x, y], x, y], foo[x, y, z], baz[x]], baz[y], baz[z]], foo[x, y, z], baz[y]], foo[x, y, z], baz[z]])
	_, _ = bar.([]*foobar[foobar[foobar[foobar[foobar[foobar[foobar[x, y, z], x, y], x, y], foo[x, y, z], baz[x]], baz[y], baz[z]], foo[x, y, z], baz[y]], foo[x, y, z], baz[z]])
	_, _ = bar.(*[]foobar[foobar[foobar[foobar[foobar[foobar[foobar[x, y, z], x, y], x, y], foo[x, y, z], baz[x]], baz[y], baz[z]], foo[x, y, z], baz[y]], foo[x, y, z], baz[z]])
}

// v0.8.4
func Bar311() {
	type foo struct { //nolint:bgolint // foo bar
	}
}
