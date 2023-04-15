package cgorithm

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	arr1 := []int{2, 4, 6, 8, 10}
	arr2 := []int{2, 3, 5, 7, 9}
	arr3 := []int{2, 4, -5, 2, 1}
	isEven := func(_, n int) bool {
		return (n & 1) == 0
	}
	isNegative := func(_, n int) bool {
		return n < 0
	}
	assert.True(t, All(arr1, isEven), "all numbers should be even")
	assert.False(t, All(arr2, isEven), "not all numbers should be even")
	assert.False(t, All(arr3, isNegative), "all numbers should be positive except one")
}

func TestAny(t *testing.T) {
	arr1 := []int{1, 3, 4, 7, 9}
	arr2 := []int{2, 4, 5, 8, 10}
	arr3 := []int{2, 4, -5, 2, 1}
	isEven := func(n int) bool {
		return (n & 1) == 0
	}
	isOdd := func(n int) bool {
		return (n & 1) == 0
	}
	isNegative := func(n int) bool {
		return n < 0
	}
	isPositive := func(n int) bool {
		return n > 0
	}
	assert.True(t, Any(arr1, isEven), "one number should be odd")
	assert.True(t, Any(arr2, isOdd), "one number should be even")
	assert.True(t, Any(arr3, isNegative), "one number should be negative")
	assert.True(t, Any(arr3, isPositive), "one number should be positive")
}

func TestCount(t *testing.T) {
	arr1 := []bool{
		true, true, false, true, false,
		true, false, true, true, false,
		true, false, false, false, false}
	arr2 := []bool{
		false, false, true, false, true,
		true, true, false, false, false,
		false, false, false, false, true,
	}
	arr3 := []bool{
		false, true, false, false, false,
		false, true, false, true, true,
		true, true, false, false, false,
	}
	assert.Equal(t, 7, Count(arr1, true), "arr1 should have 7 truly values")
	assert.Equal(t, 8, Count(arr1, false), "arr1 should have 8 falsy values")
	assert.Equal(t, 5, Count(arr2, true), "arr2 should have 5 truly values")
	assert.Equal(t, 10, Count(arr2, false), "arr2 should have 10 falsy values")
	assert.Equal(t, 6, Count(arr3, true), "arr3 should have 7 truly values")
	assert.Equal(t, 9, Count(arr3, false), "arr3 should have 8 falsy values")
}

func TestCountIf(t *testing.T) {
	arr1 := []int{1, -2, 3, -4, -5, 6, 7, 8, 9, 10}
	arr2 := []int{-2, -5, -3, 4, 5, -6, -7, 8, 10, 11}
	isNegative := func(_, n int) bool {
		return n < 0
	}
	assert.Equal(t,
		3,
		CountIf(arr1, isNegative))
	assert.Equal(t,
		5,
		CountIf(arr2, isNegative))
}

func TestForeach(t *testing.T) {
	arr1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	result1 := make([]int, 0)
	assert.True(t, Foreach(arr1, func(_, x int) ForeachAction {
		result1 = append(result1, x*x)
		return ForeachNoOp
	}))
	assert.Equal(t, arr2, result1, "result should be squares of arr1")
}

func TestGenerate(t *testing.T) {
	assert.Equal(t, []int{
		0,
		1,
		4,
		9,
		16,
		25,
	}, Generate(6, func(index int) int {
		return index * index
	}))
	assert.Equal(t, []int{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
	}, Generate(10, func(index int) int {
		return index
	}))
}

func TestQsort(t *testing.T) {
	assert.Equal(t, []int{
		4,
		7,
		26,
		37,
		53,
		62,
		66,
		66,
		74,
		79,
	}, Qsort(
		[]int{
			37,
			53,
			74,
			7,
			66,
			62,
			79,
			66,
			26,
			4,
		},
		func(_, _, x, y int) int {
			return x - y
		}),
	)
	assert.Equal(t, []int{
		79,
		74,
		66,
		66,
		62,
		53,
		37,
		26,
		7,
		4,
	}, Qsort(
		[]int{
			37,
			53,
			74,
			7,
			66,
			62,
			79,
			66,
			26,
			4,
		},
		func(_, _, x, y int) int {
			return y - x
		}),
	)
}

func TestReduce(t *testing.T) {
	add := func(_ int, x int, y int) int {
		return x + y
	}
	mul := func(_ int, x int, y int) int {
		return x * y
	}
	arr1 := []int{1, 2, 3, 4, 5}
	result1 := Reduce(arr1, 0, add)
	result2 := Reduce(arr1, 1, mul)
	arr3 := []int{1, 2, 3, 4, 5, 6}
	result3 := Reduce(arr3, 1, mul)
	arr4 := []int{1, 2, 3, 4, 5, 6, 7}
	result4 := Reduce(arr4, 1, mul)
	assert.Equal(t, result1, 15, "result1 should be 15")
	assert.Equal(t, result2, 120, "result2 should be 120")
	assert.Equal(t, result3, 720, "result2 should be 720")
	assert.Equal(t, result4, 5040, "result2 should be 5040")
	myJoin := func(arr []string, sep string) string {
		return Reduce(arr, "", func(index int, x, y string) string {
			if index > 0 {
				return Concatenate(x, ", ", y)
			}
			return y
		})
	}
	assert.Equal(t, "012, 345, 678, 9ab, cde, fgh, ijk, lmn, opq, rst, uvw, xyz",
		myJoin([]string{
			"012",
			"345",
			"678",
			"9ab",
			"cde",
			"fgh",
			"ijk",
			"lmn",
			"opq",
			"rst",
			"uvw",
			"xyz",
		}, ", "),
	)
}

func TestRepeatArray(t *testing.T) {
	assert.Equal(t,
		[]int{
			1,
			2,
			3,
			1,
			2,
			3,
			1,
			2,
			3,
			1,
			2,
			3,
		},
		RepeatArray(4, []int{
			1,
			2,
			3,
		}))
}

func TestRepeatElement(t *testing.T) {
	assert.Equal(t, []int{
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
	}, RepeatElement(10, 1))
}

func TestSum(t *testing.T) {
	assert.Equal(t, 15, Sum([]int{
		1,
		2,
		3,
		4,
		5,
	}, 0))
	assert.Equal(t, 1.5, Sum([]float64{
		0.1,
		0.2,
		0.3,
		0.4,
		0.5,
	}, 0.0))
	assert.Equal(t, "abcdefghi", Sum([]string{
		"abc",
		"def",
		"ghi",
	}, ""))
	assert.Equal(t, "0123456789", Sum([]string{
		"56",
		"7",
		"89",
	}, "01234"))
}

func TestTransform(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t,
		[]int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100},
		Transform(arr1, func(_, x int) int {
			return x * x
		}))
}

func TestTransformReduce(t *testing.T) {
	arr1 := []string{"123", "456", "789"}
	assert.Equal(t, 1368, TransformReduce(arr1, 0, func(_, x, y int) int {
		return x + y
	}, func(_ int, i string) int {
		result, _ := strconv.Atoi(i)
		return result
	}))
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, 385, TransformReduce(arr2, 0, func(_, x, y int) int {
		return x + y
	}, func(_ int, x int) int {
		return x * x
	}))
}

func TestMAll(t *testing.T) {
	m1 := map[int]int{
		0:  0,
		1:  1,
		2:  4,
		3:  9,
		4:  16,
		5:  25,
		6:  36,
		7:  49,
		8:  64,
		9:  81,
		10: 100,
		11: 121,
		12: 144,
		13: 169,
		14: 196,
		15: 225,
		16: 256,
	}
	m2 := map[string]int{
		"1":   1,
		"2":   2,
		"3":   3,
		"11":  11,
		"12":  12,
		"123": 123,
		"456": 456,
		"789": 789,
	}
	assert.True(t, MAll(m1, func(x, y int) bool {
		return x*x == y
	}))
	assert.True(t, MAll(m2, func(x string, y int) bool {
		result, _ := strconv.Atoi(x)
		return result == y
	}))
}

func TestMAny(t *testing.T) {
	m1 := map[string][]string{
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ": {
			"abcdefghi",
			"jklmnopqr",
			"stuvwxyz",
			"<=>;:",
		},
		"hello world": {
			"!",
		},
		"abcdef": {
			"AB",
			"c",
			"DEF",
		},
	}
	assert.True(t, MAny(m1, func(x string, y []string) bool {
		return Any(y, func(z string) bool {
			return strings.Contains(x, z)
		})
	}))
}

func TestZip(t *testing.T) {
	arr1 := []int{0, 1, 2, 3, 4, 5}
	arr2 := []int{1, 1, 4, 9, 16, 25}
	result := make([]int, 0)
	assert.True(t, Zip(arr1, arr2, func(_, x, y int) ForeachAction {
		if (x * x) == y {
			result = append(result, y)
		}
		return ForeachNoOp
	}))
	assert.Equal(t, []int{1, 4, 9, 16, 25}, result)
}
