// SYNTAX TEST "source.go" "Semantic Tokens Test"

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
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

func Bar37(ah string) {}

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
