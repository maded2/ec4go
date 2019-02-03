package example

//go:generate ec4go -name TestCaseTypeList -package example -type  TestCaseType
type TestCaseType struct {
	S string
	I int
}
