package example

//go:generate ec4go -name BoolImmutableList -package example -type bool
type BoolImmutableList []bool

//go:generate ec4go -name BoolList -package example -type bool -mutable
type BoolList []bool

var sample_bool = []bool{false, true, true, false, true}
