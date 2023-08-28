// SYNTAX TEST "source.go" "Semantic Tokens Test"

package main

import (
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"go/ast"
	"io"
	"log"
	"mime/multipart"
	"net/http"
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
	_ = a.(FooX)
	_ = a.(*FooX)
	_ = a.(*[]FooX)
	_ = a.([]FooX)
	_ = a.([]*FooX)
	_ = a.(map[FooX]interface{})
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

func startIdx[E any](haystack, needle []E) int { return 0 }

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
