package example

type TestCaseType struct {
	S string
	I int
}

//go:generate ec4go -name TestCaseImmutableTypeList -package example -type *TestCaseType
type TestCaseImmutableTypeList []*TestCaseType

//go:generate ec4go -name TestCaseTypeList -package example -type *TestCaseType -mutable
type TestCaseTypeList []*TestCaseType
