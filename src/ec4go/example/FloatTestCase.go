package example

//go:generate ec4go -name FloatImmutableList -package example -type float64
type FloatImmutableList []float64

//go:generate ec4go -name FloatList -package example -type float64 -mutable
type FloatList []float64

var sample_float64 = []float64{5.0, 4.0, 3.0, 2.0, 1.0}
