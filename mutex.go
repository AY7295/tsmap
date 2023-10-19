package tsmap

import (
	"sync"
)

type mutexMap[K comparable, V any] struct {
	sync.Mutex
	mp map[K]V
}

func (s *mutexMap[K, V]) Get(key K) (V, bool) {
	s.Lock()
	defer s.Unlock()
	v, ok := s.mp[key]
	return v, ok
}

func (s *mutexMap[K, V]) Set(key K, value V) {
	s.Lock()
	defer s.Unlock()
	s.mp[key] = value
}

func (s *mutexMap[K, V]) Delete(key K) {
	s.Lock()
	defer s.Unlock()
	delete(s.mp, key)
}

func (s *mutexMap[K, V]) Len() int {
	s.Lock()
	defer s.Unlock()
	return len(s.mp)
}

func (s *mutexMap[K, V]) Range(fn Operation[K, V], exception ...Exception[K, V]) {
	s.Lock()
	defer s.Unlock()
	except := defaultException[K, V]
	if len(exception) > 0 {
		except = exception[0]
	}
	for key, value := range s.mp {
		if !except(key, value) {
			fn(key, value)
		}
	}
}
