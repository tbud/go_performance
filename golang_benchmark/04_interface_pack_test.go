package golang_benchmark

import (
	"testing"
)

var constIntSlice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

var constPackedIntSlice = []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

var intSlice = make([]int, 10)

var packedIntSlice = make([]interface{}, 10)

func Benchmark04SetIntSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			intSlice[i] = i
		}
	}
}

func Benchmark04SetPackedIntSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			packedIntSlice[i] = i // HL
		}
	}
}

func Benchmark04GetIntSlice(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			sum += constIntSlice[i]
		}
	}
	_ = sum
}

func Benchmark04GetPackedIntSlice(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			sum += constPackedIntSlice[i].(int) // HL
		}
	}
	_ = sum
}
