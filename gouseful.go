package gouseful

import (
	"fmt"
)

type Num interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64	
}

func Compare[E any](a, b E) bool {
	return fmt.Sprintf("%#v",a) == fmt.Sprintf("%#v",b)
}

func CompareWeak[A, B any](a A, b B) bool {
	return fmt.Sprint(a) == fmt.Sprint(b)
}

func CompareMultiple[E any](items ...E) bool {
	return Every(items, func(a E) bool { return Compare(a,items[0]) })
}