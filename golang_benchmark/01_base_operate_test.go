package golang_benchmark

import (
	"testing"
	"time"
)

func Benchmark01NormalAdd(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		sum += i // HL
	}
	_ = sum
}

func Benchmark01Sleep1Microsecond(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		sum += i
	}
	time.Sleep(time.Microsecond) // HL
	_ = sum
}

func Benchmark01Sleep1Second(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		sum += i
	}
	time.Sleep(time.Second)
	_ = sum
}
