package golang_benchmark

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
