package golang_benchmark

import (
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

func coverWithPanic() {
	defer func() {
		if r := recover(); r != nil {

		}
	}()

	withPanic()
}

func withPanic() {
	panic(1)
}

func Benchmark06Panic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coverWithPanic()
	}
}

func Benchmark06WithoutPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coverWithoutPanic()
	}
}
