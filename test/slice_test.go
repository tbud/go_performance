package test

import (
	"strconv"
	"testing"
)

var strSlice []string

func init() {
	for i := 0; i < 100; i++ {
		strSlice = append(strSlice, strconv.Itoa(i))
	}
}

func forSlice(str string) {
	s := str
	_ = s
}

func forSlicePoint(str *string) {
	s := *str
	_ = s
}

func BenchmarkSliceForRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, s := range strSlice {
			forSlice(s)
		}
	}
}

func BenchmarkSliceForRangePoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, s := range strSlice {
			forSlicePoint(&s)
		}
	}
}

func BenchmarkSliceForRangeIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range strSlice {
			forSlice(strSlice[j])
		}
	}
}

func BenchmarkSliceForRangeIndexPoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range strSlice {
			forSlicePoint(&strSlice[j])
		}
	}
}

func BenchmarkSliceForIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(strSlice); j++ {
			forSlice(strSlice[j])
		}
	}
}

func BenchmarkSliceForIndexPoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(strSlice); j++ {
			forSlicePoint(&strSlice[j])
		}
	}
}
