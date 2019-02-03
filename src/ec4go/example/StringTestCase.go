package example

//go:generate ec4go -name StringImmutableList -package example -type string
//go:generate ec4go -name StringList -package example -type string -mutable
type StringType string
