package gouseful

import (
	"fmt"
)

type Num interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64	
}

func Compare[E any](a, b E) bool {
	return fmt.Sprintf("%#v",a) == fmt.Sprint("%#v",b)
}