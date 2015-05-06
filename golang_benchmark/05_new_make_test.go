package golang_benchmark

import (
	"testing"
)

type hello struct {
	message string
}

func BenchmarkCreateInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &hello{"world!"}
	}
}

func BenchmarkCreateNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := new(hello)
		h.message = "world!"
		_ = h
	}
}

func BenchmarkCreateMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hl := []hello{}
		for i := 0; i < 10; i++ {
			hl = append(hl, hello{"world!"})
		}
	}
}

func BenchmarkCreateMakeFixSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hl := make([]hello, 10)
		for i := 0; i < 10; i++ {
			hl[i].message = "world!"
		}
	}
}
