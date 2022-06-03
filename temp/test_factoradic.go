package main

import (
	"fmt"

	"github.com/phobiadev/gouseful"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(gouseful.Permutations(a, len(a))[20])
	fmt.Println(gouseful.NthPermutation(a, 20))
}