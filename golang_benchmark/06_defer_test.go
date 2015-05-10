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

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withDefer()
	}
}

func BenchmarkDeferWithout(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withoutDefer()
	}
}

func BenchmarkDeferTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withDoubleDefer()
	}
}

func BenchmarkDeferWithoutTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withoutDeferTwo()
	}
}
