# ec4go

Attempt to add all the nice functions from Eclipse Collections.  Using code generation to get type safe functions.

Using normal 'go generate to immutable list'

    //go:generate ec4go -name TestCaseImmutableTypeList -package example -type  TestCaseType
    
Add mutable flag to add the additional mutable functions
    
    //go:generate ec4go -name TestCaseTypeList -package example -type  TestCaseType -mutable