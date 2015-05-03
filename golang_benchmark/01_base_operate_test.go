package golang_benchmark

import (
	"testing"
	"time"
)

func BenchmarkBaseOperate(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		sum += i // HL
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

func BenchmarkAddFuncCall(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		sum = add(i, sum) // HL
	}
	_ = sum
}

func BenchmarkAddIntStructCall(b *testing.B) {
	s := AddInt(0)
	for i := 0; i < b.N; i++ {
		s.add(i) // HL
	}
	_ = s
}

func BenchmarkAddStructCall(b *testing.B) {
	s := &Add{}
	for i := 0; i < b.N; i++ {
		s.add(i)
	}
	_ = s
}

func BenchmarkAddTwoStructCall(b *testing.B) {
	sum := 0
	s := &Add{}
	for i := 0; i < b.N; i++ {
		sum = s.addTwo(i, sum)
	}
	_ = s
	_ = sum
}
