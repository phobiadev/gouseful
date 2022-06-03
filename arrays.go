package gouseful

import (
	"fmt"
)

func MakeCopy[T any](array []T) []T {
	newArray := make([]T, len(array))
	copy(newArray,array)
	return newArray
}

func Filter[T any](array []T, predicate func(T) bool) []T {
	newArray := []T{}
	for _, s := range array {
		if predicate(s) {
			newArray = append(newArray,s)
		}
	}
	return newArray
}

func FilterFalse[T any](array []T, predicate func(T) bool) []T {
	newArray := []T{}
	for _, s := range array {
		if !predicate(s) {
			newArray = append(newArray,s)
		}
	}
	return newArray
}

func MapOver[T any, R any](array []T, changer func(T) R) []R {
	newArray := []R{}
	for _, s := range array {
		newArray = append(newArray,changer(s))
	}
	return newArray
}

func Insert[T any](array []T, index int, item T) []T {
	newArray := MakeCopy(array)
	return append(append(newArray[:index],item),array[index:]...)
}

func Index[E comparable](array []E, item E) int {
	for i, s := range array {
		if s == item {
			return i
		}
	}
	return -1
}

func Remove[E comparable](array []E, item E) []E {
	index := Index(array, item)
	if index == -1 { 
		return array
	}
	return Pop(array,index)
}

func Chain[T any](arrays ...[]T) []T {
	newArray := []T{}
	for _, array := range arrays {
		newArray = append(newArray,array...)
	}
	return newArray
}

func Count[E comparable](array []E, item E) int {
	count := 0
	for _, s := range array {
		if s == item {
			count++
		}
	}
	return count
}

func Reverse[E any](array []E) []E {
	newArray := MakeCopy(array)
	for i, j := 0, len(newArray)-1; i < j; i, j = i+1, j-1 {
        newArray[i], newArray[j] = newArray[j], newArray[i]
    }
	return newArray
}

func Pop[T any](array []T, index int) []T {
	return append(array[:index],array[index+1:]...)
}

func PairWise[T any](array []T) [][]T {
	pairs := [][]T{}
	for i := 0; i < len(array)-1; i++ {
		pairs = append(pairs,array[i:i+2])
	}
	return pairs
}

func Compress[T any](data []T, selectors []int) []T {
	compressed := []T{}
	for i, s := range data {
		if selectors[i] == 1 {
			compressed = append(compressed, s)
		}
	}
	return compressed
}

func DropWhile[T any](array []T, predicate func(T) bool) []T {
	for i, s := range array {
		if !predicate(s) {
			return array[i:]
		}
	}
	return []T{}
}

func AccumulateAdd[N Num](array []N, initial N) []N {
	total := initial
	a := []N{}
	for _, s := range array {
		total += s
		a = append(a,total)
	}
	return a
}

func Accumulate[T any](array []T, function func(T,T) T, initial T) []T {
	 total := initial
	 a := []T{}
	 for _, s := range array {
		 total = function(total,s)
		 a = append(a,total)
	 }
	 return a
}

func Unshift[T any](array []T, item T) []T {
	return append([]T{item},array...)
}

func Shift[T any](array []T) (T, []T) {
	return array[0], array[1:]
}

func Includes[E comparable](array []E, item E) bool {
	return Index(array, item) != -1
}

func Every[T any](array []T, predicate func(T) bool) bool {
	for _, s := range array {
		if !predicate(s) {
			return false
		}
	}
	return true
}

func Some[T any](array []T, predicate func(T) bool) bool {
	for _, s := range array {
		if predicate(s) {
			return true
		}
	}
	return false
}

func FindIndex[T any](array []T, predicate func(T) bool) int {
	for i, s := range array {
		if predicate(s) {
			return i
		}
	}
	return -1
}

func Fill[T any](array []T, item T) []T {
	newArray := []T{}
	for i := 0; i < len(array); i++ {
		newArray = append(newArray,item)
	}
	return newArray
}

func Join[T any](array []T, connection string) string {
	result := fmt.Sprint(array[0])
	for _, s := range MapOver(array[1:],func(a T) string { return fmt.Sprint(a) }) {
		result += (connection + s)
	}
	return result
}

func Reduce[T any](array []T, function func(T, T) T, initial T ) T {
	total := initial
	for _, s := range array {
		total = function(total,s)
	}
	return total
}

func Sum[N Num](array []N) N {
	return Reduce(array, func(tot,curr N) N { return tot + curr }, 0)
}