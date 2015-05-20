package golang_benchmark

import (
	"encoding/json"
	"testing"

	"github.com/pquerna/ffjson/ffjson"
)

type NormalHello struct {
	Message string
}

func init() {
	b, _ := json.Marshal(&Hello{Message: "world"})
	println(string(b))
}

func Benchmark09NormalMarshalJson(b *testing.B) {
	h := &NormalHello{Message: "world"}
	for i := 0; i < b.N; i++ {
		json.Marshal(h)
	}
}

func Benchmark09NormalMarshalWithffjsonObject(b *testing.B) {
	h := &Hello{Message: "world"}
	for i := 0; i < b.N; i++ {
		json.Marshal(h)
	}
}

func Benchmark09Marshalffjson(b *testing.B) {
	h := &Hello{Message: "world"}
	for i := 0; i < b.N; i++ {
		ffjson.Marshal(h)
	}
}

func Benchmark09MarshalffjsonObject(b *testing.B) {
	h := &Hello{Message: "world"}
	for i := 0; i < b.N; i++ {
		h.MarshalJSON()
	}
}

func Benchmark09MarshalffjsonObjectWithPool(b *testing.B) {
	h := &Hello{Message: "world"}
	for i := 0; i < b.N; i++ {
		buf, _ := h.MarshalJSON()
		ffjson.Pool(buf)
	}
}

func Benchmark09NormualUnmarshalJson(b *testing.B) {
	h := &NormalHello{}
	buf := []byte(`{"Message":"world"}`)
	for i := 0; i < b.N; i++ {
		json.Unmarshal(buf, h)
	}
}

func Benchmark09NormalUnMarshalWithffjsonObject(b *testing.B) {
	h := &Hello{}
	buf := []byte(`{"Message":"world"}`)
	for i := 0; i < b.N; i++ {
		json.Unmarshal(buf, h)
	}
}

func Benchmark09Unmarshalffjson(b *testing.B) {
	h := &Hello{}
	buf := []byte(`{"Message":"world"}`)
	for i := 0; i < b.N; i++ {
		ffjson.Unmarshal(buf, h)
	}
}

func Benchmark09UnmarshalffjsonObject(b *testing.B) {
	h := &Hello{}
	buf := []byte(`{"Message":"world"}`)
	for i := 0; i < b.N; i++ {
		h.UnmarshalJSON(buf)
	}
}
