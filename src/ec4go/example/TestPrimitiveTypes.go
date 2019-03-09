package example

// List examples

//go:generate ec4go -genOption list -typeName BoolImmutableList -package example -valueType bool
type BoolImmutableList []bool

//go:generate ec4go -genOption list -typeName BoolList -package example -valueType bool -mutable
type BoolList []bool

//go:generate ec4go -genOption list -typeName FloatImmutableList -package example -valueType float64
type FloatImmutableList []float64

//go:generate ec4go -genOption list -typeName FloatList -package example -valueType float64 -mutable
type FloatList []float64

//go:generate ec4go -genOption list -typeName IntImmutableList -package example -valueType int
type IntImmutableList []int

//go:generate ec4go -genOption list -typeName IntList -package example -valueType int -mutable
type IntList []int

//go:generate ec4go -genOption list -typeName StringImmutableList -package example -valueType string
type StringImmutableList []string

//go:generate ec4go -genOption list -typeName StringList -package example -valueType string -mutable
type StringList []string

// map examples

//go:generate ec4go -genOption map -typeName BoolImmutableMap -package example -keyType bool -valueType bool
type BoolImmutableMap map[bool]bool

//go:generate ec4go -genOption map -typeName IntImmutableMap -package example -keyType int -valueType int
type IntImmutableMap map[int]int

//go:generate ec4go -genOption map -typeName FloatImmutableMap -package example -keyType float64 -valueType float64
type FloatImmutableMap map[float64]float64

//go:generate ec4go -genOption map -typeName StringImmutableMap -package example -keyType string -valueType string
type StringImmutableMap map[string]string

// sample data for unit tests
var sample_bool = []bool{false, true, true, false, true}
var sample_float64 = []float64{5.0, 4.0, 3.0, 2.0, 1.0}
var sample_int = []int{5, 4, 3, 2, 1}
var sample_string = []string{"nop", "klm", "ghi", "def", "abc"}
