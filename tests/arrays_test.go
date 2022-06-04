package tests

import (
	"testing"
	"fmt"
	. "github.com/phobiadev/gouseful"
)

func TestMakeCopy(t *testing.T) {
	a := []int{1,2,3}
	b := MakeCopy(a)
	a = Pop(a,1)
	wantA := []int{1,3}
	wantB := []int{1,2,3}
	if Compare(a, wantA) && !Compare(b,wantB) {
		t.Errorf("got [%v %v], wanted [%v %v] ", a, b, wantA, wantB)
	}
}

func TestFilter(t *testing.T) {
	got := Filter([]int{1,5,6,3,8,2,9},func(a int) bool { return a > 5 })
	want := []int{6,8,9}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFilterFalse(t *testing.T) {
	got := FilterFalse([]int{1,5,6,3,8,2,9},func(a int) bool { return a > 5 })
	want := []int{1,5,3,2}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMapOver(t *testing.T) {
	got := MapOver([]int{1,2,3},func(a int) string { return fmt.Sprintf("-%v-",a+1) })
	want := []string{"-2-","-3-","-4-"}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestInsert(t *testing.T) {
	got := Insert([]string{"a","c","d"},1,"b")
	want := []string{"a","b","c","d"}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIndex(t *testing.T) {
	got := Index([]string{"a","b","c"},"c")
	want := 2
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIndexNonComparable(t *testing.T) {
	got := IndexNonComparable([][]int{[]int{1,2},[]int{2,3},[]int{3,4}},[]int{3,4})
	want := 2
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRemove(t *testing.T) {
	got := Remove([]string{"a","b","c"},"b")
	want := []string{"a","c"}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestRemoveNonComparable(t *testing.T) {
	got := RemoveNonComparable([][]int{[]int{1,2},[]int{2,3},[]int{3,4}},[]int{2,3})
	want := [][]int{[]int{1,2},[]int{3,4}}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestChain(t *testing.T) {
	got := Chain([]string{"a","b"},[]string{"c","d"},[]string{"e","f"})
	want := []string{"a","b","c","d","e","f"}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCount(t *testing.T) {
	got := Count([]string{"a","b","b","c","d"},"b")
	want := 2
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCountNonComparable(t *testing.T) {
	got := CountNonComparable([][]int{[]int{1,2},[]int{2,3},[]int{3,4},[]int{3,4}},[]int{3,4})
	want := 2
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestReverse(t *testing.T) {
	got := Reverse([]string{"a","b","c"})
	want := []string{"c","b","a"}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPop(t *testing.T) {
	got := Pop([]string{"a","b","c"},1)
	want := []string{"a","c"}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPairWise(t *testing.T) {
	got := PairWise([]string{"a","b","c","d"})
	want := [][]string{
		[]string{"a","b"},
		[]string{"b","c"},
		[]string{"c","d"},
	}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDropWhile(t *testing.T) {
	got := DropWhile([]int{1,2,6,7,1},func(a int) bool { return a < 5 })
	want := []int{6,7,1}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestAccumulateAdd(t *testing.T) {
	got := AccumulateAdd([]int{1,2,3,4},0)
	want := []int{1,3,6,10}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestAccumulate(t *testing.T) {
	got := Accumulate([]int{2,4,6,8},func(a,b int) int { return a * b }, 1)
	want := []int{2,8,48,384}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUnshift(t *testing.T) {
	got := Unshift([]int{2,3,4},1)
	want := []int{1,2,3,4}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestShift(t *testing.T) {
	a, b := Shift([]string{"a","b","c"})
	wantA := "a"
	wantB := []string{"b","c"}
	if !(a == wantA && Compare(b,wantB)) {
		t.Errorf("got [%v %v], wanted [%v %v] ", a, b, wantA, wantB)
	}
}

func TestIncludes(t *testing.T) {
	got := Includes([]string{"a","b","c"},"b")
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIncludesNonComparable(t *testing.T) {
	got := IncludesNonComparable([][]int{[]int{1,2},[]int{2,3},[]int{3,4}},[]int{3,4})
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestEvery(t *testing.T) {
	got := Every([]int{1,6,7}, func(a int) bool { return a > 5 })
	want := false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSome(t *testing.T) {
	got := Some([]int{1,6,7}, func(a int) bool { return a > 5 })
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFindIndex(t *testing.T) {
	got := FindIndex([]int{1,6,7}, func(a int) bool { return a > 5 })
	want := 1
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestFill(t *testing.T) {
	got := Fill(make([]int, 3),1)
	want := []int{1,1,1}
	if !Compare(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestJoin(t *testing.T) {
	got := Join([]string{"a","b","c"}," - ")
	want := "a - b - c"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestReduce(t *testing.T) {
	got := Reduce([]int{2,4,6,8},func(a, b int) int { return a * b }, 1)
	want := 384
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSum(t *testing.T) {
	got := Sum([]float64{.1,.2,.3,.4})
	want := 1.0
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	} 
}