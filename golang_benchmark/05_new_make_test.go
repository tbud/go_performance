package golang_benchmark

import (
	"testing"
)

type hello struct {
	message string
}

func Benchmark05UseInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hello{"world!"} // HL
	}
}

func Benchmark05UseInitPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &hello{"world!"} // HL
	}
}

func Benchmark05UseNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := new(hello)      // HL
		h.message = "world!" // HL
		_ = h                // HL
	}
}

func Benchmark05AppendInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hl := []hello{}
		for i := 0; i < 10; i++ {
			hl = append(hl, hello{"world!"}) // HL
		}
	}
}

func Benchmark05AppendInitPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hl := []*hello{}
		for i := 0; i < 10; i++ {
			hl = append(hl, &hello{"world!"})
		}
	}
}

func Benchmark05InitFixedArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hl := [10]hello{}
		for i := 0; i < 10; i++ {
			hl[i].message = "world!"
		}
	}
}

func Benchmark05MakeFixedSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hl := make([]hello, 10)
		for i := 0; i < 10; i++ {
			hl[i].message = "world!"
		}
	}
}

func Benchmark05MakeFixedSliceAndAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hl := make([]*hello, 10)
		for i := 0; i < 10; i++ {
			hl = append(hl, &hello{"world!"})
		}
	}
}
