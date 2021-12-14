package benching

import (
	"sync"
	"testing"
)

type Counter struct {
	A int
	B int
}

func NewIncrementer() *Counter {

	counter := Counter{
		A: 0,
		B: 0,
	}
	return &counter
}

func Incrementer(s *Counter) {
	// a := NewIncrementer()
	s.A = s.A + 1
	s.B = s.B + 1
}

var pool = sync.Pool{
	New: func() interface{} { return NewIncrementer() },
}

func (c *Counter) Reset() {
	*c = Counter{}
}

func putPool(c *Counter) {
	pool.Put(c)
}

func BenchmarkWithoutPool(b *testing.B) {
	var s *Counter
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			// s = pool.Get().(*Counter)
			s = NewIncrementer()
			b.StopTimer()
			Incrementer(s)
			b.StartTimer()
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var s *Counter
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			s = pool.Get().(*Counter)
			b.StopTimer()
			Incrementer(s)
			b.StartTimer()
			s.Reset()
			pool.Put(s)
		}
	}
}
