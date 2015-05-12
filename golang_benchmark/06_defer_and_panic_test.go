package golang_benchmark

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

func withDefer() {
	defer func() {

	}()
}

func withDoubleDefer() {
	defer func() {

	}()

	withDefer()
}

func normalCall() {

}

func normalDoubleCall() {
	normalCall()
}

func Benchmark06NormalCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		normalCall()
	}
}

func Benchmark06NormalDoubleCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		normalDoubleCall()
	}
}

func Benchmark06Defer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withDefer()
	}
}

func Benchmark06DoubleDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withDoubleDefer()
	}
}

func withNumberPanic() {
	panic(1)
}

func withStringPanic() {
	panic("hello world")
}

var err = fmt.Errorf("error")

func withConstErrorPanic() {
	panic(err)
}

func withNewErrorPanic() {
	panic(errors.New("error"))
}

func withFmtErrorPanic() {
	panic(fmt.Errorf("error"))
}

func coverPanic(f func()) {
	defer func() {
		if r := recover(); r != nil {

		}
	}()

	f()
}

func Benchmark06CoverNoPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coverPanic(normalCall)
	}
}

func Benchmark06ConstErrorPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coverPanic(withConstErrorPanic)
	}
}

func Benchmark06NewErrorPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coverPanic(withNewErrorPanic)
	}
}

func Benchmark06FmtErrorPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coverPanic(withFmtErrorPanic)
	}
}

func Benchmark06NumberPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coverPanic(withNumberPanic)
	}
}

func Benchmark06StringPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coverPanic(withStringPanic)
	}
}

type LockTest struct {
	sync.Mutex
	a int
}

func (l *LockTest) get() (ret int) {
	l.Lock()
	ret = l.a
	l.Unlock()
	return
}

func (l *LockTest) getWithDefer() (ret int) {
	defer l.Unlock()
	l.Lock()
	ret = l.a
	return
}

func Benchmark06Lock(b *testing.B) {
	l := &LockTest{}
	for i := 0; i < b.N; i++ {
		l.get()
	}
}

func Benchmark06LockWithDefer(b *testing.B) {
	l := &LockTest{}
	for i := 0; i < b.N; i++ {
		l.getWithDefer()
	}
}
