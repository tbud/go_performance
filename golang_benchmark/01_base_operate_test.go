package golang_benchmark

import (
	"testing"
	"time"
)

func BenchmarkBaseOperate(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		sum += i
	}
	_ = sum
}

func BenchmarkBaseOperateSleep1Microsecond(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		sum += i
	}
	time.Sleep(time.Microsecond) // HL
	_ = sum
}

func BenchmarkBaseOperateSleep1Second(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		sum += i
	}
	time.Sleep(time.Second)
	_ = sum
}

func BenchmarkFuncCall(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		sum = add(i, sum)
	}
	_ = sum
}
