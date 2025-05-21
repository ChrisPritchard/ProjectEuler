package main

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T])
	s.Add(items...)
	return s
}

func (s Set[T]) Add(items ...T) {
	for _, v := range items {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

func (s Set[T]) Remove(items ...T) {
	for _, v := range items {
		delete(s, v)
	}
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s))
	for v := range s {
		slice = append(slice, v)
	}
	return slice
}
