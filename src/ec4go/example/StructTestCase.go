package example

//go:generate ec4go -name TestCaseImmutableTypeList -package example -type TestCaseType
//go:generate ec4go -name TestCaseTypeList -package example -type TestCaseType -mutable
type TestCaseType struct {
	S string
	I int
}
