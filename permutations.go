package gouseful

func factoradic(n int) []int {
	i := 1
	var r int
	repr := []int{}
	for n > 0 {
		n, r = DivMod(n,i)
		repr = Unshift(repr,r)
		i++
	}
	return repr
}

func NthPermutation[T any](array []T, n int) []T {
	newArray := MakeCopy(array)
	factRepr := factoradic(n)
	permutation := []T{}
	for _, s :=  range factRepr {
		permutation = append(permutation,newArray[s])
		newArray = Pop(newArray,s)
	}
	return permutation
} 

func Permutations[E comparable](array []E, lim int) [][]E {
	permutations := [][]E{}
	var inner func([]E)
    inner = func(curr []E) {
		if len(curr) == lim {
			permutations = append(permutations,curr)
			return
		}
		temp := FilterFalse(array,func(a E) bool { return Includes(curr,a) })
		for _, s := range temp {
			inner(append(curr,s))
		}
    }
	inner([]E{})
	return permutations
}