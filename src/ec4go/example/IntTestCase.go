package example

//go:generate ec4go -name IntImmutableList -package example -type int
type IntImmutableList []int

//go:generate ec4go -name IntList -package example -type int -mutable
type IntList []int

var sample_int = []int{5, 4, 3, 2, 1}
