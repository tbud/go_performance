package golang_benchmark

//go:generate ffjson $GOFILE
type Hello struct {
	Message string
}
