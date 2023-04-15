package cgorithm

import (
	"strings"

	"golang.org/x/exp/constraints"
)

// Tests if every element of `arr` matches a predicate that accept
// index and value as parameters, and
// returns true, if element match, and false otherwise. The
// function returns true if every element matches, false
// otherwise.
func All[T any](arr []T, predicate func(int, T) bool) bool {
	for index, element := range arr {
		if !predicate(index, element) {
			return false
		}
	}
	return true
}

// Tests if any element of `arr` matches a predicate that accept
// index and value as parameters, and
// returns true, if element match, and false otherwise. The
// function returns true if any element matches, false
// otherwise.
func Any[T any](arr []T, predicate func(T) bool) bool {
	for _, element := range arr {
		if predicate(element) {
			return true
		}
	}
	return false
}

// Test if all pairs of elements in two slices satisfy a predicate
// that accept Index-Value pair of two slices, and returns true,
// if element match, and false otherwise. The function returns
// true, if all pairs of elements matches, false otherwise.
func AllSatisfy[T any,
	U any,
](arr1 []T, arr2 []U, predicate func(int, int, T, U) bool) bool {
	for index1, element1 := range arr1 {
		for index2, element2 := range arr2 {
			if !predicate(index1, index2, element1, element2) {
				return false
			}
		}
	}
	return true
}

// Test if any pair of elements in two slices satisfy a predicate
// that accept Index-Value pair of two slices, and returns true,
// if element match, and false otherwise. The function returns
// true, if any pair of elements matches, false otherwise.
func AnySatisfy[T any,
	U any,
](arr1 []T, arr2 []U, predicate func(int, int, T, U) bool) bool {
	for index1, element1 := range arr1 {
		for index2, element2 := range arr2 {
			if predicate(index1, index2, element1, element2) {
				return true
			}
		}
	}
	return false
}

// Concatenates all arguments into one string.
func Concatenate(elements ...string) string {
	var builder strings.Builder
	for _, element := range elements {
		builder.WriteString(element)
	}
	return builder.String()
}

// Concatenates strings from the slice into one.
func ConcatenateSlice(elements []string) string {
	var builder strings.Builder
	for _, element := range elements {
		builder.WriteString(element)
	}
	return builder.String()
}

// Counts element in slice.
func Count[T comparable](arr []T, element T) int {
	result := 0
	for _, x := range arr {
		if x == element {
			result++
		}
	}
	return result
}

// Counts elements in slice that match predicate. The predicate
// accept Index-Value pair, and returns true, if element need
// to be counted, false otherwise. The function returns count
// of counted values.
func CountIf[T any](arr []T, predicate func(int, T) bool) int {
	result := 0
	for index, element := range arr {
		if predicate(index, element) {
			result++
		}
	}
	return result
}

// Filters a slice from elements that not match predicate. The
// predicate accept Index-Value pair, and returns true, if
// element need to be skipped from resulting slice. The
// function returns filtered slice.
func Filter[T any](arr []T, predicate func(int, T) bool) []T {
	result := []T{}
	for index, element := range arr {
		if predicate(index, element) {
			result = append(result, element)
		}
	}
	return result
}

// Find element in the slice. Returns -1 if not found, index of element otherwise if found.
func Find[T comparable](arr []T, element T) int {
	for index, x := range arr {
		if x == element {
			return index
		}
	}
	return -1
}

// Finds element in the slice. Returns first element that
// matches predicate, that accept Index-Value pair, and
// returns true, if element was finded, false otherwise.
// The function returns -1 if none element match predicate.
func FindIf[T any](arr []T, predicate func(int, T) bool) int {
	for index, element := range arr {
		if predicate(index, element) {
			return index
		}
	}
	return -1
}

// ForEach loop action.
type ForeachAction int

const (
	// Do nothing.
	ForeachNoOp ForeachAction = 0
	// Break from loop.
	ForeachBreak ForeachAction = 1
	// Continues loop.
	ForeachContinue ForeachAction = 2
)

// Calls predicate for each element in the slice. The
// predicate returns action. Returns true if loop was
// success (the predicates returned values ​​in the
// range 0-2), false otherwise.
func Foreach[T any](arr []T, predicate func(int, T) ForeachAction) bool {
	for index, element := range arr {
		switch predicate(index, element) {
		case ForeachNoOp:
		case ForeachBreak:
			break
		case ForeachContinue:
			continue
		default:
			return false
		}
	}
	return true
}

// Generates a array filled with `count` values
// returned by predicate, and return it.
func Generate[T any](count int, predicate func(int) T) []T {
	result := make([]T, count)
	for i := 0; i < count; i++ {
		result[i] = predicate(i)
	}
	return result
}

// Same as `All`, but only for maps, and
// the predicate accepts Key-Value pair.
func MAll[TK comparable, TV any](m map[TK]TV, predicate func(TK, TV) bool) bool {
	for key, value := range m {
		if !predicate(key, value) {
			return false
		}
	}
	return true
}

// Same as `Any`, but only for maps, and
// the predicate accepts Key-Value pair.
func MAny[TK comparable, TV any](m map[TK]TV, predicate func(TK, TV) bool) bool {
	for key, value := range m {
		if predicate(key, value) {
			return true
		}
	}
	return false
}

// Same as `AllSatisfy`, but only for maps, and
// the predicate accepts two Key-Value pairs of both maps.
func MAllSatisfy[TK comparable,
	TV any,
	UK comparable,
	UV any,
](m1 map[TK]TV, m2 map[UK]UV, predicate func(TK, UK, TV, UV) bool) bool {
	for key1, value1 := range m1 {
		for key2, value2 := range m2 {
			if !predicate(key1, key2, value1, value2) {
				return false
			}
		}
	}
	return true
}

// Same as `AnySatisfy`, but only for maps, and
// the predicate accepts two Key-Value pairs of both maps.
func MAnySatisfy[TK comparable,
	TV any,
	UK comparable,
	UV any,
](m1 map[TK]TV, m2 map[UK]UV, predicate func(TK, UK, TV, UV) bool) bool {
	for key1, value1 := range m1 {
		for key2, value2 := range m2 {
			if predicate(key1, key2, value1, value2) {
				return true
			}
		}
	}
	return false
}

// Same as `Count`, but only for maps, and the element
// argument is value of pair to be counted.
func MCount[TK comparable, TV comparable](m map[TK]TV, element TV) int {
	result := 0
	for _, v := range m {
		if v == element {
			result++
		}
	}
	return result
}

// Same as `CountIf`, but only for maps, and
// the predicate accept Key-Value pair.
func MCountIf[TK comparable, TV any](m map[TK]TV, predicate func(TK, TV) bool) int {
	result := 0
	for key, value := range m {
		if predicate(key, value) {
			result++
		}
	}
	return result
}

// Same as `Filter`, but only for maps, and
// the predicate accept Key-Value pair.
func MFilter[TK comparable, TV any](m map[TK]TV, predicate func(TK, TV) bool) map[TK]TV {
	result := map[TK]TV{}
	for key, value := range m {
		if predicate(key, value) {
			result[key] = value
		}
	}
	return result
}

// Same as `Find`, but only for maps, and
// the predicate accept Key-Value pair.
func MFindK[TK comparable, TV comparable](m map[TK]TV, element TV) []TK {
	for key, value := range m {
		if value == element {
			return []TK{key}
		}
	}
	return []TK{}
}

func MFindV[TK comparable, TV comparable](m map[TK]TV, element TV) []TK {
	result := []TK{}
	for key, value := range m {
		if value == element {
			result = append(result, key)
		}
	}
	return result
}

// Same as `FindIf`, but only for maps, and
// the predicate accept Key-Value pair. The
// function returns array of keys, values
// of that matches the predicate.
func MFindIf[TK comparable, TV any](m map[TK]TV, predicate func(TK, TV) bool) []TK {
	result := []TK{}
	for key, value := range m {
		if predicate(key, value) {
			result = append(result, key)
		}
	}
	return result
}

// Same as `Foreach`, but only for maps, and
// the predicate accept Key-Value pair.
func MForeach[TK comparable, TV any](m map[TK]TV, predicate func(TK, TV) ForeachAction) bool {
	for key, value := range m {
		switch predicate(key, value) {
		case ForeachNoOp:
		case ForeachBreak:
			break
		case ForeachContinue:
			continue
		default:
			return false
		}
	}
	return true
}

// Same as `Reduce`, but only for maps.
func MReduce[TK comparable, TV any, U any](m map[TK]TV, init U, predicate func(TK, U, TV) U) U {
	result := init
	for key, value := range m {
		result = predicate(key,
			result,
			value)
	}
	return result
}

// Same as `Transform`, but only for maps, and
// the predicate accept Key-Value pair, and
// returns Key-Value pair, which will stored in
// map.
func MTransform[TK comparable, TV any, UK comparable, UV any](m map[TK]TV, predicate func(TK, TV) (UK, UV)) map[UK]UV {
	result := map[UK]UV{}
	for key, value := range m {
		rkey, rvalue := predicate(key, value)
		result[rkey] = rvalue
	}
	return result
}

// Same as `TransformReduce`, but only for maps.
func MTransformReduce[TK comparable, TV any, U any](arr map[TK]TV, init U, reduce func(TK, U, U) U, transform func(TK, TV) U) U {
	result := init
	for key, value := range arr {
		result = reduce(
			key,
			result,
			transform(key, value))
	}
	return result
}

// Finds max value of x and y values.
func Max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// Finds min value of x and y values.
func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Sorts slice. The function does same as `qsort` in C,
// but the predicate accepts indexes.
func Qsort[T any](arr []T, predicate func(int, int, T, T) int) []T {
	result := arr
	n := len(arr) - 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if predicate(i, j, result[j], result[j+1]) > 0 {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// Reduces the slice using the predicate. The predicate
// accept index, result, and element to be reduced, and
// returns the result. The function returns result of
// predicates.
func Reduce[T any, U any](arr []T, init U, predicate func(int, U, T) U) U {
	result := init
	for index, element := range arr {
		result = predicate(index,
			result,
			element)
	}
	return result
}

// Repeats elements in one slice N times.
func RepeatArray[T any](count int, elements []T) []T {
	length := len(elements)
	result := make([]T, length*count)
	for i := 0; i < length*count; i++ {
		result[i] = elements[i%length]
	}
	return result
}

// Repeats a one element N times and writes in the slice.
func RepeatElement[T any](count int, element T) []T {
	result := make([]T, count)
	for i := 0; i < count; i++ {
		result[i] = element
	}
	return result
}

// Apply predicate to every element in the slice. The
// predicate accepts Index-Value pair, and returns
// changed value with (un-)changed type. The function
// returns slice of elements with predicate applied.
func Transform[T any, U any](arr []T, predicate func(int, T) U) []U {
	result := []U{}
	for index, element := range arr {
		result = append(result, predicate(index, element))
	}
	return result
}

// Same as `Reduce` and `Transform`.
func TransformReduce[T any, U any](arr []T, init U, reduce func(int, U, U) U, transform func(int, T) U) U {
	result := init
	for index, element := range arr {
		result = reduce(
			index,
			result,
			transform(
				index,
				element,
			),
		)
	}
	return result
}

func Sort[T constraints.Ordered](arr []T) []T {
	result := arr
	n := len(arr) - 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// Sums the slice.
func Sum[T constraints.Ordered](arr []T, init T) T {
	result := init
	for _, element := range arr {
		result += element
	}
	return result
}

// Same as `Foreach`, but accepts a two slices and the predicate
// accepts index, and elements of two slices.
func Zip[T any, U any](arr1 []T, arr2 []U, predicate func(int, T, U) ForeachAction) bool {
	length := Min(len(arr1), len(arr2))
	for index := 0; index < length; index++ {
		switch predicate(index, arr1[index], arr2[index]) {
		case ForeachNoOp:
		case ForeachBreak:
			break
		case ForeachContinue:
			continue
		default:
			return false
		}
	}
	return true
}
