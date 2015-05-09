package golang_benchmark

import (
	"testing"
)

func withDefer() {
	defer func() {

	}()
}

func withoutDefer() {

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
