package golang_benchmark

import (
	"math/rand"
	"sync"
	"testing"
)

func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000)+5000000)
}

type bufferPool struct {
	b []byte
}

func Benchmark07NormalMake(b *testing.B) {
	pool := make([][]byte, 20)
	for i := 0; i < b.N; i++ {
		b := makeBuffer()

		i := rand.Intn(len(pool))
		pool[i] = b

		bytes := 0

		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}
	}
}

func Benchmark07ChanReuse(b *testing.B) {
	pool := make([][]byte, 20)

	buffer := make(chan []byte, 5)

	for i := 0; i < b.N; i++ {
		var b []byte
		select {
		case b = <-buffer:
		default:
			b = makeBuffer()
		}

		i := rand.Intn(len(pool))
		if pool[i] != nil {
			select {
			case buffer <- pool[i]:
				pool[i] = nil
			default:
			}
		}

		pool[i] = b

		bytes := 0

		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}
	}
}

func Benchmark07PoolReuse(b *testing.B) {
	pool := make([][]byte, 20)

	p := sync.Pool{
		New: func() interface{} {
			return make([]byte, rand.Intn(5000000)+5000000)
		},
	}

	for i := 0; i < b.N; i++ {
		b := p.Get().([]byte)

		i := rand.Intn(len(pool))

		p.Put(pool[i])

		pool[i] = b

		bytes := 0

		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}
	}
}
