package tests

import (
	"testing"
	. "github.com/phobiadev/gouseful"
)

func TestNthPermutation(t *testing.T) {
	got := NthPermutation([]string{"a","b","c","d"},10)
	want := []string{"b","d","a","c"}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPermutations(t *testing.T) {
	got := Permutations([]int{1,2,3},2)
	want := [][]int{
		[]int{1,2},
		[]int{1,3},
		[]int{2,1},
		[]int{2,3},
		[]int{3,1},
		[]int{3,2},
	}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}