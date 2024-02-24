package conv

import (
	"math/rand"
	"reflect"
	"strings"
)

// ToSlice converts any type slice or array to the specified type slice.
func ToSlice[T any](a any) []T {
	r, _ := ToSliceE[T](a)
	return r
}

// ToSliceE converts any type slice or array to the specified type slice.
// An error will be returned if an error occurred.
func ToSliceE[T any](a any) ([]T, error) {
	if a == nil {
		return nil, nil
	}
	switch v := a.(type) {
	case []T:
		return v, nil
	case string:
		return ToSliceE[T](strings.Fields(v))
	}

	kind := reflect.TypeOf(a).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		// If input is a slice or array.
		v := reflect.ValueOf(a)
		if kind == reflect.Slice && v.IsNil() {
			return nil, nil
		}
		s := make([]T, v.Len())
		for i := 0; i < v.Len(); i++ {
			val, err := ToAnyE[T](v.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			s[i] = val
		}
		return s, nil
	default:
		// If input is a single value.
		v, err := ToAnyE[T](a)
		if err != nil {
			return nil, err
		}
		return []T{v}, nil
	}
}

// SplitStrToSlice splits a string to a slice by the specified separator.
func SplitStrToSlice[T any](s, sep string) []T {
	v, _ := SplitStrToSliceE[T](s, sep)
	return v
}

// SplitStrToSliceE splits a string to a slice by the specified separator and returns an error if occurred.
// Note that this function is implemented through 1.18 generics, so the element type needs to
// be specified when calling it, e.g. SplitStrToSliceE[int]("1,2,3", ",").
func SplitStrToSliceE[T any](s, sep string) ([]T, error) {
	ss := strings.Split(s, sep)
	r := make([]T, len(ss))
	for i := range ss {
		v, err := ToAnyE[T](ss[i])
		if err != nil {
			return nil, err
		}
		r[i] = v
	}
	return r, nil
}

func InSlice[T comparable](slice []T, in T) bool {
	for _, v := range slice {
		if v == in {
			return true
		}
	}
	return false
}

// MapSlice manipulates a slice and transforms it to a slice of another type.
// Play: https://go.dev/play/p/OkPcYAhBo0D
func MapSlice[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item, i)
	}

	return result
}

// UniqueSlice returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array.
// Play: https://go.dev/play/p/DTzbeXZ6iEN
func UniqueSlice[T comparable](collection []T) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[T]struct{}, len(collection))

	for _, item := range collection {
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

// FilterSlice iterates over elements of collection, returning an array of all elements predicate returns truthy for.
// Play: https://go.dev/play/p/Apjg3WeSi7K
func FilterSlice[V any](collection []V, predicate func(item V, index int) bool) []V {
	result := make([]V, 0, len(collection))

	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}

	return result
}

// GroupBySlice returns an object composed of keys generated from the results of running each element of collection through iteratee.
// Play: https://go.dev/play/p/XnQBd_v6brd
func GroupBySlice[T any, U comparable](collection []T, iteratee func(item T) U) map[U][]T {
	result := map[U][]T{}

	for _, item := range collection {
		key := iteratee(item)

		result[key] = append(result[key], item)
	}

	return result
}

func UpdateSlice[T any](slice []T, update func(v T) T) []T {
	var results = make([]T, 0, len(slice))
	for _, v := range slice {
		results = append(results, update(v))
	}
	return results
}

func DeleteSlice[T comparable](slice []T, e T) []T {
	//idx := 0
	//for _, v := range slice {
	//	if !v != e {
	//		slice[idx] = v
	//		idx++
	//	}
	//}
	//return slice[:idx]

	idx := 0
	for _, v := range slice {
		if v == e {
			slice[idx] = slice[len(slice)-1]
			return slice[:len(slice)-1]
		}
		idx++
	}
	return slice
}

// Shuffle returns an array of shuffled values. Uses the Fisher-Yates shuffle algorithm.
// Play: https://go.dev/play/p/Qp73bnTDnc7
func ShuffleSlice[T any](collection []T) []T {
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})

	return collection
}

// Reverse reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.
// Play: https://go.dev/play/p/fhUMLvZ7vS6
func ReverseSlice[T any](collection []T) []T {
	length := len(collection)
	half := length / 2

	for i := 0; i < half; i = i + 1 {
		j := length - 1 - i
		collection[i], collection[j] = collection[j], collection[i]
	}

	return collection
}
