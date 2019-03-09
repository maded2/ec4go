package example

type TestCaseType struct {
	S string
	I int
}

// list examples
//go:generate ec4go -genOption list -typeName TestCaseImmutableTypeList -package example -valueType *TestCaseType
type TestCaseImmutableTypeList []*TestCaseType

//go:generate ec4go -genOption list -typeName TestCaseTypeList -package example -valueType *TestCaseType -mutable
type TestCaseTypeList []*TestCaseType

// map examples
//go:generate ec4go -genOption map -typeName TestCaseImmutableTypeMap -package example -keyType int -valueType *TestCaseType
type TestCaseImmutableTypeMap map[int]*TestCaseType

var sample = []*TestCaseType{
	&TestCaseType{S: "Cde", I: 3},
	&TestCaseType{S: "Bcd", I: 2},
	&TestCaseType{S: "Abc", I: 1},
}
