package example

//go:generate ec4go -name IntImmutableList -package example -type int
//go:generate ec4go -name IntList -package example -type int -mutable
type IntType int

var sample_int = []int{5, 4, 3, 2, 1}
