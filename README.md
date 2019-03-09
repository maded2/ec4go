# ec4go

Attempt to add all the nice functions from Eclipse Collections.  Using code generation to get type safe functions.

Using normal 'go generate for immutable list'

    //go:generate ec4go -genOption list -typeName TestCaseImmutableTypeList -package example -valueType *TestCaseType
    type TestCaseImmutableTypeList []*TestCaseType
        
Add mutable flag to add additional mutable functions
    
    //go:generate ec4go -genOption list -typeName TestCaseTypeList -package example -valueType *TestCaseType -mutable
    type TestCaseTypeList []*TestCaseType
    
    
for immutable map functions

    //go:generate ec4go -genOption map -typeName TestCaseImmutableTypeMap -package example -keyType int -valueType *TestCaseType
    type TestCaseImmutableTypeMap map[int]*TestCaseType