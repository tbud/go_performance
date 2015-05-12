package golang_benchmark

import (
	"testing"
)

var iarray = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var islice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func Benchmark03ArrayForIterator(b *testing.B) {
	sum := 0
	num := len(iarray)
	for i := 0; i < b.N; i++ {
		for j := 0; j < num; j++ { // HL
			sum += iarray[j] // HL
		} // HL
	}
	_ = sum
}

func Benchmark03SliceForIterator(b *testing.B) {
	sum := 0
	num := len(islice)
	for i := 0; i < b.N; i++ {
		for j := 0; j < num; j++ {
			sum += iarray[j]
		}
	}
	_ = sum
}

func Benchmark03ArrayRangeIterator(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for j := range iarray { // HL
			sum += iarray[j] // HL
		} // HL
	}
	_ = sum
}

func Benchmark03SliceRangeIterator(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for j := range islice {
			sum += iarray[j]
		}
	}
	_ = sum
}

func Benchmark03ArrayRangeValueIterator(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for _, v := range iarray { // HL
			sum += v // HL
		} // HL
	}
	_ = sum
}

func Benchmark03SliceRangeValueIterator(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for _, v := range islice {
			sum += v
		}
	}
	_ = sum
}
