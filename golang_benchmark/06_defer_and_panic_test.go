package golang_benchmark

import (
	"fmt"
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

func withoutDefer() {

}

func withoutDeferTwo() {
	withoutDefer()
}

func Benchmark06Defer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withDefer()
	}
}

func Benchmark06WithoutDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withoutDefer()
	}
}

func Benchmark06TwoDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withDoubleDefer()
	}
}

func Benchmark06WithoutTwoDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withoutDeferTwo()
	}
}

func coverWithoutPanic() {
	withoutPanic()
}

func withoutPanic() {

}

func coverNoPanic() {
	defer func() {
		if r := recover(); r != nil {

		}
	}()

	withoutPanic()
}

func withNumberPanic() {
	panic(1)
}

func withStringPanic() {
	panic("hello world")
}

func withErrorPanic() {
	panic(fmt.Errorf("error"))
}

func coverWithPanic(f func()) {
	defer func() {
		if r := recover(); r != nil {

		}
	}()

	f()
}

func Benchmark06NumberPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coverWithPanic(withNumberPanic)
	}
}

func Benchmark06StringPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coverWithPanic(withStringPanic)
	}
}

func Benchmark06ErrorPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coverWithPanic(withErrorPanic)
	}
}

func Benchmark06WithoutPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coverWithoutPanic()
	}
}

func Benchmark06NoPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coverNoPanic()
	}
}
