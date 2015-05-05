package golang_benchmark

import (
	"testing"
)

var intUnpackMap = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

var intUnpackInterfaceMap = []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

var intPackMap = make([]int, 10)

var intPackInterfaceMap = make([]interface{}, 10)

func BenchmarkInterfaceIntPackBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			intPackMap[i] = i
		}
	}
}

func BenchmarkInterfacePack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			intPackInterfaceMap[i] = i
		}
	}
}

func BenchmarkInterfaceIntUnpackBase(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			sum += intUnpackMap[i]
		}
	}
	_ = sum
}

func BenchmarkInterfaceUnPack(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			sum += intUnpackInterfaceMap[i].(int)
		}
	}
	_ = sum
}
