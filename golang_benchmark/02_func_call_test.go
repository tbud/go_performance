package golang_benchmark

import "testing"

func add(a int, b int) int {
	return a + b
}

type AddInt int

func (a *AddInt) add(b int) {
	*a += AddInt(b)
}

type Add struct {
	sum int
}

func (a *Add) add(b int) {
	a.sum += b
}

func (a *Add) addTwo(b int, c int) int {
	return b + c
}

// type Add end OMIT

func Benchmark02NormalAdd(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		sum += i // HL
	}
	_ = sum
}

func Benchmark02CallAddFunc(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		sum = add(i, sum) // HL
	}
	_ = sum
}

func Benchmark02CallPackedIntAdd(b *testing.B) {
	s := AddInt(0)
	for i := 0; i < b.N; i++ {
		s.add(i) // HL
	}
	_ = s
}

func Benchmark02CallIntStructAdd(b *testing.B) {
	s := &Add{}
	for i := 0; i < b.N; i++ {
		s.add(i)
	}
	_ = s
}

func Benchmark02CallStructAddFunc(b *testing.B) {
	sum := 0
	s := &Add{}
	for i := 0; i < b.N; i++ {
		sum = s.addTwo(i, sum)
	}
	_ = s
	_ = sum
}
