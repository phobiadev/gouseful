package tests

import (
	"testing"
	. "github.com/phobiadev/gouseful"
)

func TestDivMod(t *testing.T) {
	a, b := DivMod(10,3)
	wantA := 3
	wantB := 1
	if !(a == wantA && Compare(b,wantB)) {
		t.Errorf("got [%v %v], wanted [%v %v] ", a, b, wantA, wantB)
	}
}

func TestIFloor(t *testing.T) {
	got := IFloor(4.5)
	want := 4
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestICeil(t *testing.T) {
	got := ICeil(4.5)
	want := 5
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestISqrt(t *testing.T) {
	got := ISqrt(345)
	want := 18
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsPrime(t *testing.T) {
	got := IsPrime(31)
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsWithin(t *testing.T) {
	got := IsWithin(4.5,1.0,9.0)
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}