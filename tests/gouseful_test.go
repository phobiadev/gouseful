package tests

import (
	"testing"
	. "github.com/phobiadev/gouseful"
)

func TestCompare(t *testing.T) {
	got := Compare([]int{1,2,3},[]int{1,2,3})
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCompareWeak(t *testing.T) {
	got := CompareWeak([]int{1,2,3},[]uint{1,2,3})
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCompareMultiple(t *testing.T) {
	got := CompareMultiple([]int{1,2,3},[]int{1,2,3},[]int{1,2,3})
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

