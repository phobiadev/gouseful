package gouseful

import (
	"math"
)

func DivMod(a,b int) (int, int) {
	return a / b, a % b
}

func IFloor[N Num](x N) int {
	return int(math.Floor(float64(x)))
}

func ICeil[N Num](x N) int {
	return int(math.Ceil(float64(x)))
}

func ISqrt[N Num](x N) int {
	// !! Not my algorithm !! - fast integer square root using bitwise
	// https://groups.google.com/g/golang-nuts/c/5qVyilXt4B0?pli=1
	var t, b, r uint
	t = uint(x)
	p := uint(1 << 30)

	for p > t {
		p >>= 2
	}

	for ; p != 0; p >>= 2 {
		b = r | p
		r >>= 1
		if t >= b {
			t -= b
			r |= p
		}
	}

	return int(r)
}

func IsPrime(x int) bool {
	lim := ISqrt(x)
	for i := 2; i <= lim; i++ {
		if x % i == 0 {
			return false
		}
	}
	return true
} 

func IsWithin[N Num](n, min, max N) bool {
	return n >= min && n <= max
}

