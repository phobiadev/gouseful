package gouseful

import (
	"math/rand"
	"time"
)

func resetSeed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func init() {
	resetSeed()
}

func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandChoice[E any](array []E) E {
	return array[RandInt(0,len(array)-1)]
}

func Shuffle[E any](array []E) []E {
	newArray := MakeCopy(array)
	rand.Shuffle(len(array), func(i, j int) {
		newArray[i], newArray[j] = newArray[j], newArray[i]
	})
	return newArray
}


