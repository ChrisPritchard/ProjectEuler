package main

import (
	"cmp"
	"sort"
)

// returns a new slice containing only those items that the filter function returns true for
func filter[T any](items []T, filter func(item T) bool) []T {
	res := []T{}
	for i := range items {
		if filter(items[i]) {
			res = append(res, items[i])
		}
	}
	return res
}

// applies a mapping function to each item in a slice, returning a new slice
func transform[T any, U any](items []T, selector func(item T) U) []U {
	res := []U{}
	for i := range items {
		res = append(res, selector(items[i]))
	}
	return res
}

// reverses a slice and returns the result, without modifying the initial slice
func reverse[T comparable](slice []T) []T {
	n := make([]T, len(slice))
	for i := range slice {
		n[len(n)-1-i] = slice[i]
	}
	return n
}

// returns a map of slices with keys being the values returned by the selector for each item
func group_by[T any, U comparable](items []T, selector func(item T) U) map[U][]T {
	res := make(map[U][]T)
	for i := range items {
		key := selector(items[i])
		if v, exists := res[key]; exists {
			res[key] = append(v, items[i])
		} else {
			res[key] = []T{items[i]}
		}
	}
	return res
}

func sort_by[T any, K cmp.Ordered](items []T, selector func(T) K) []T {
	sorted := append([]T(nil), items...)
	sort.Slice(sorted, func(i, j int) bool {
		return selector(sorted[i]) < selector(sorted[j])
	})
	return sorted
}

type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}

func to_slice[K comparable, V any](m map[K]V) []KeyValue[K, V] {
	values := []KeyValue[K, V]{}
	for k, v := range m {
		values = append(values, KeyValue[K, V]{Key: k, Value: v})
	}
	return values
}
