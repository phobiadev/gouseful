# Go Useful

gouseful is a package that provides some useful, generic functions offered by other languages, as well as some more complex ones, such as integer square root, and n'th permutation

live tests can be found at [go.dev/play/p/0-nLoNiqP_x](https://go.dev/play/p/0-nLoNiqP_x)
<br>
see detailed usage below

## General

<br>

`Num` - interface that contains all number types (float, int, uint) except complex64 and complex128 for use in generic functions
```go
func add[N gouseful.Num](a, b N) N {
    return a + b
}
fmt.Println(add(2.5,2.2),add(3,2))
```
output: `4.7 5`

<br>

`Compare` - compares two variables of any type (even non-comparables) by converting them both to strings using Sprintf with %#v
```go
fmt.Println(gouseful.Compare([]int{1,2,3},[]int{2,3,4}))
```
output: `false`

<br>

`CompareWeak` - weakly compares two variables of any (seperate) type by converting them to strings using Sprint rather than Sprintf
```go
fmt.Println(gouseful.CompareWeak([]int{1,2,3},[]uint{1,2,3}))
```
output: `true`

<br>

`CompareMultiple` - compare infinite items of the same type (any)
```go
fmt.Println(gouseful.CompareMultiple([]int{1,2,3},[]int{1,2,3},[]int{1,2,3}))
```
output: `true`

<br>

## Arrays 
see tests/arrays_test.go

<br>

`MakeCopy` - returns a copied version of the slice
```go
a := []int{1,2,3}
b := gouseful.MakeCopy(a)
a = gouseful.Pop(a,1)
fmt.Println(a,b)
```
output: `[1 3] [1 2 3]`

<br>

`Filter` - returns a slice of all array elements with which the predicate function evaluates to true
```go
a := []int{1,5,6,3,8,2,9}
b := gouseful.Filter(a,func(x int) bool { return x > 5 })
fmt.Println(b)
```
output: `[6 8 9]`

<br>

`FilterFalse` - does the same thing as Filter, but when the predicate evaluates to false
```go
a := []int{1,5,6,3,8,2,9}
b := gouseful.FilterFalse(a,func(x int) bool { return x > 5 })
fmt.Println(b)
```
output: `[1 5 3 2]`

<br>

`MapOver` - applies the function provided to each of the elements in the array, and returns the new array
```go
a := []int{1,2,3}
fmt.Println(gouseful.MapOver(a,func(a int) string { return fmt.Sprintf("-%v-",a+1) }))
```
output: `[-2- -3- -4-]`

<br>

`Insert` - inserts the given element at the given index, shifting all elements after this index one to the right, and returns the new array
```go
a := []string{"a","c","d"}
fmt.Println(gouseful.Insert(a,1,"b"))
```
output: `[a b c d]`

<br>

`Index` - returns the index of the item in the array (if not found, returns -1)
```go
a := []string{"a","b","c"}
fmt.Println(gouseful.Index(a,"c"))
```
output: `2`

<br>

`IndexNonComparable` - returns the index of the item in the array of non-comparables (if not found, returns -1)
```go
a := [][]int{[]int{1,2},[]int{2,3},[]int{3,4}}
fmt.Println(IndexNonComparable(a,[]int{3,4}))
```
output: `2`

<br>

`Remove` - removes an item from the array (first occurence, not all), and returns the new array
```go
a := []string{"a","b","c"}
fmt.Println(gouseful.Remove(a,"b"))
```
output: `[a c]`

<br>

`RemoveNonComparable` - removes the item from the array of non-comparables (first occurence, not all)
```go
a := [][]int{[]int{1,2},[]int{2,3},[]int{3,4}}
fmt.Println(RemoveNonComparable(a,[]int{3,4}))
```
output: `[[1 2] [3 4]]`

<br>

`Chain` - chains together multiple lists into one, and returns this big joined up list
```go
fmt.Println(Chain([]string{"a","b"}, []string{"c","d"}, []string{"e","f"}))
```
output: `[a b c d e f]`

<br>

`Count` - counts the number of occurences of an item in an array (returns as integer)
```go
a := []string{"a","b","b","c","d"}
fmt.Println(gouseful.Count(a,"b"))
```
output: `2`

<br>

`CountNonComparable` - counts the number of occurences of an item in an array of non-comparables (returns as integer)
```go
a := [][]int{[]int{1,2},[]int{2,3},[]int{3,4},[]int{3,4}}
fmt.Println(gouseful.CountNonComparable(a,[]int{3,4}))
```
output: `2`

<br>

`Reverse` - reverses a list and returns the reversed version
```go
a := []string{"a","b","c"}
fmt.Println(gouseful.Reverse(a))
```
output: `[c b a]`

<br>

`Pop` - removes the item at the given index from the array, and returns the new array
```go
a := []string{"a","b","c"}
fmt.Println(gouseful.Pop(a,1))
```
output: `[a c]`

<br>

`PairWise` - groups the array into overlapping pairs, eg. [a b], [b c], etc., returning these pairs
```go
a := []string{"a","b","c","d"}
fmt.Println(gouseful.PairWise(a))
```
output: `[a b] [b c] [c d]`

<br>

`DropWhile` - drops elements until one meets the provided predicate function, and returns the shortened array
```go
a := []int{1,2,6,7,1}
fmt.Println(gouseful.DropWhile(a,func(a int) bool { return a < 5 }))
```
output: `[6 7 1]`

<br>

`AccumulateAdd` - accumulates all of the elements in an array by adding each element to the sum of all the elements before it and the initial value provided
```go
a := []int{1,2,3,4}
fmt.Println(gouseful.AccumulateAdd(a,0))
```
output: `[6 7 1]`

<br>

`Accumulate` - accumulates all of the elements in an array using a custom function, that accepts two parameters (total and current), with total initially being set to the initial parameter provided
```go
a := []int{2,4,6,8}
fmt.Println(gouseful.Accumulate(a,func(a,b int) int { return a * b },1))
```
output: `[2 8 48 384]`

<br>

`Unshift` - appends an item to the beginning of an array
```go
a := []int{2,3,4}
fmt.Println(gouseful.Unshift(a,1))
```
output: `[1 2 3 4]`

<br>

`Shift` - removes the first item of the array, returning two values: the previous first item and the new array
```go
a, b := gouseful.Shift([]string{"a","b","c"})
fmt.Println(a,b)
```
output: `a [b c]`

<br>

`Includes` - checks whether or not an item exists in an array of non-comparables
```go
a := []string{"a","b","c"}
fmt.Println(gouseful.Includes(a,"b"))
```
output: `true`

<br>

`IncludesNonComparable` - checks whether or not an item exists in an array
```go
a := [][]int{[]int{1,2},[]int{2,3},[]int{3,4}}
fmt.Println(gouseful.Includes(a,,[]int{3,4}))
```
output: `true`

<br>

`Every` - checks if the provided predicate evaluates to true for every item in the array
```go
a := []int{1,6,7}
fmt.Println(gouseful.Every(a,func(x int) bool { return x > 5 }))
```
output: `false`

<br>

`Some` - checks if the provided predicate evaluates to true for any of the items in the array
```go
a := []int{1,6,7}
fmt.Println(gouseful.Some(a,func(x int) bool { return x > 5 }))
```
output: `true`

<br>

`FindIndex` - finds the index of the first item in the array that the provided predicate function evaluates to true with, returns -1 if it does not for any of the items
a := []int{1,6,7}

<br>

`Fill` - makes every item in the array equal to the item provided, and returns the new array
```go
fmt.Println(gouseful.Fill(make([]int, 3),1))
```
output: `[1 1 1]`

<br>

`Join` - joins together an array (of any type) into a string using a provided connector
```go
a := []string{"a","b","c"}
fmt.Println(gouseful.Join(a," - "))
```
output: `a - b - c`

<br>

`Reduce` - does a very similar thing to the Accumulate function, but instead of returning the accumulated array, only the final value (total) is returned
```go
fmt.Println(gouseful.Reduce([]int{2,4,6,8},func(a, b int) int { return a * b }, 1))
```
output: `384`

<br>

`Sum` - adds together all of the elements in the array, and returns the total
```go
a := []float64{.1,.2,.3,.4}
fmt.Println(gouseful.Sum(a))
```
output: `1.0`

<br>

`Min` - returns the smallest item in the array
```go
a := []int{3,5,4,6}
fmt.Println(gouseful.Min(a))
```
output: `3`

<br>

`Max` - returns the biggest item in the array
```go
a := []int{3,5,4,6}
fmt.Println(gouseful.Max(a))
```
output: `6`

<br>

## Math
see tests/math_test.go

<br>

`DivMod` - accepts two integers, and returns the result of their integer division, and the remainder (result of modulus)
```go
a, b := gouseful.DivMod(10,3)
fmt.Println(a,b)
```
output: `3 1`

<br>

`IFloor` - returns the rounded down version of any number as an integer (as opposed to as a float64, like the math.Floor function)
```go
fmt.Println(gouseful.IFloor(4.5))
```
output: `4`

<br>

`ICeil` - returns the rounded up version of any number as an integer (as opposed to as a float64, like the math.Ceil function)
```go
fmt.Println(gouseful.ICeil(4.5))
```
output: `5`

<br>

`ISqrt` - uses an algorithm (not mine) that uses unsigned integeres and bit shifting to quickly get the integer square root of a number (as opposed to the expensive traditional square root function)
```go
fmt.Println(gouseful.ISqrt(345))
```
output: `18`

<br>

`IsPrime` - checks whether or not provided number is prime (higher than 1 and divisible by only itself and 1)
```go
fmt.Println(gouseful.IsPrime(31))
```
output: `true`

<br>

`IsWithin` - checks if a number is within the provided min and max (inclusive)
```go
fmt.Println(gouseful.IsWithin(4.5,1.0,9.0))
```
output: `true`

<br>