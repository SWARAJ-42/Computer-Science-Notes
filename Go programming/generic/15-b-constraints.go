package main

import "fmt"

// interface type lists
// Ordered is a type constraint that matches any ordered type.
// An ordered type is one that supports the <, <=, >, and >= operators.
type number interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
        ~float32 | ~float64 |
        ~string
}


// Below stringer acts as a constraint, which is nothing but an interface
func concat[T number](vals []T) T {
    var result T
    for _, val := range vals {
        // this is where the operation on number
        // is done. That's why we need a more specific
        // constraint instead of the any constraint
        result += val
    }
    return result
}

func main() {
    var nums []uint = []uint{1, 2, 3, 4,}
    sum := concat[uint](nums)
    fmt.Println(sum)

}
