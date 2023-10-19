package tsmap

import (
	"sync"
	"sync/atomic"
)

type syncMap[K comparable, V any] struct {
	size *atomic.Int64
	mp   *sync.Map
}

func (s *syncMap[K, V]) Get(key K) (V, bool) {
	v, ok := s.mp.Load(key)
	return v.(V), ok
}

func (s *syncMap[K, V]) Set(key K, value V) {
	_, existed := s.mp.LoadOrStore(key, value)
	if !existed {
		s.increase()
	}
}

func (s *syncMap[K, V]) Delete(key K) {
	_, existed := s.mp.LoadAndDelete(key)
	if existed {
		s.decrease()
	}
}

func (s *syncMap[K, V]) Len() int {
	return int(s.size.Load())
}

func (s *syncMap[K, V]) Range(fn Operation[K, V], exception ...Exception[K, V]) {
	except := defaultException[K, V]
	if len(exception) > 0 {
		except = exception[0]
	}
	s.mp.Range(func(key, value any) bool {
		if !except(key.(K), value.(V)) {
			fn(key.(K), value.(V))
		}
		return true
	})
}

func (s *syncMap[K, V]) decrease() {
	s.size.Add(-1)
}

func (s *syncMap[K, V]) increase() {
	s.size.Add(1)
}
