package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	listName := flag.String("name", "", "name of the list")
	listType := flag.String("type", "", "type of the list")
	packageName := flag.String("package", "", "package name of the list")
	addMutableFunctions := flag.Bool("mutable", false, "add mutable functions to the list")
	flag.Parse()

	if len(os.Args) <= 1 {
		flag.Usage()
		return
	}

	immutableListTemplate := template.Must(template.New("immutableListTemplate").Parse(immutableListTemplate))

	type value struct {
		PackageName string
		ListName    string
		ListType    string
		Primitive   bool
		Mutable     bool
	}
	v := value{
		PackageName: *packageName,
		ListName:    *listName,
		ListType:    *listType,
		Mutable:     *addMutableFunctions,
	}
	primitiveTypes := []string{"uint", "uint8,", "uint16", "uint32", "uint64", "int", "int8,", "int16", "int32", "int64", "float32", "float64", "{{.ListType}}", "bool", "string"}

	v.Primitive = false
	for _, pt := range primitiveTypes {
		if v.ListType == pt {
			v.Primitive = true
			break
		}
	}

	fileName := fmt.Sprintf("%s.go", flag.Lookup("name").Value.String())
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Unable to create file: %s - %s", fileName, err)
	}
	w := bufio.NewWriter(file)
	if err = immutableListTemplate.Execute(w, v); err != nil {
		log.Fatalf("Failed to create[%s]: %s", fileName, err)
	}
	if v.Mutable {
		mutableListTemplate := template.Must(template.New("mutableListTemplate").Parse(mutableListTemplate))
		if err = mutableListTemplate.Execute(w, v); err != nil {
			log.Fatalf("Failed to add mutable functions[%s]: %s", fileName, err)
		}
	}
	w.Flush()
	file.Close()

	if v.Primitive {
		fileName = fmt.Sprintf("%s_test.go", flag.Lookup("name").Value.String())
		file, err = os.Create(fileName)
		if err != nil {
			log.Fatalf("Unable to create file: %s - %s", fileName, err)
		}
		w = bufio.NewWriter(file)
		testCasesTemplate := template.Must(template.New("testCasesTemplate").Parse(testCases))
		if err = testCasesTemplate.Execute(w, v); err != nil {
			log.Fatalf("Failed to create[%s]: %s", fileName, err)
		}
		w.Flush()
		file.Close()
	}
}

var immutableListTemplate = `/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/
package {{.PackageName}}

import (
	"math/rand"
	"sort"
)

// Immutable functions

func {{.ListName}}_NewWith(element {{.ListType}}) (newList {{.ListName}}) {
	newList = append({{.ListName}}{}, element)
	return
}

func {{.ListName}}_NewWithAll(elements []{{.ListType}}) (newList {{.ListName}}) {
	newList = append({{.ListName}}{}, elements...)
	return
}

func (l {{.ListName}}) NewWithout(element {{.ListType}}) (newList {{.ListName}}) {
	newList = {{.ListName}}{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l {{.ListName}}) NewWithoutAll(elements []{{.ListType}}) (newList {{.ListName}}) {
	newList = {{.ListName}}{}
	if l != nil {
		for _, e1 := range l {
			found := false
			for _, e2 := range elements {
				if e1 == e2 {
					found = true
					break
				}
			}
			if !found {
				newList = append(newList, e1)
			}
		}
	}
	return
}

func (l {{.ListName}}) Remove(element {{.ListType}}) (newList {{.ListName}}) {
	newList = {{.ListName}}{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return newList
}

func (l {{.ListName}}) RemoveAll(elements []{{.ListType}}) (newList {{.ListName}}) {
	newList = {{.ListName}}{}
	if l != nil {
		for _, e1 := range l {
			found := false
			for _, e2 := range elements {
				if e1 == e2 {
					found = true
					break
				}
			}
			if !found {
				newList = append(newList, e1)
			}
		}
	}
	return
}

func (l {{.ListName}}) Size() (size int) {
	size = len(l)
	return
}

func (l {{.ListName}}) IsEmpty() bool {
	return len(l) == 0
}

func (l {{.ListName}}) NotEmpty() bool {
	return len(l) > 0
}

func (l {{.ListName}}) GetAny() (result {{.ListType}}) {
	if len(l) > 0 {
		result = l[rand.Intn(len(l)-1)]
	}
	return
}

func (l {{.ListName}}) Contains(element {{.ListType}}) bool {
	if l != nil {
		for _, e1 := range l {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l {{.ListName}}) ContainsAll(elements []{{.ListType}}) bool {
	if l != nil {
		n := 0
		for _, e1 := range l {
			for _, e2 := range elements {
				if e1 == e2 {
					n++
					break
				}
			}
			if n == len(elements) {
				return true
			}
		}
	}
	return false
}

func (l {{.ListName}}) Each(procedure func(element {{.ListType}})) {{.ListName}} {
	if l != nil {
		for _, e := range l {
			procedure(e)
		}
	}
	return l
}

func (l {{.ListName}}) Select(predicate func(element {{.ListType}}) bool) (newList {{.ListName}}) {
	newList = {{.ListName}}{}
	if l != nil {
		for _, e := range l {
			if predicate(e) {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l {{.ListName}}) Reject(predicate func(element {{.ListType}}) bool) (newList {{.ListName}}) {
	newList = {{.ListName}}{}
	if l != nil {
		for _, e := range l {
			if predicate(e) == false {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l {{.ListName}}) Partition(predicate func(element {{.ListType}}) bool) (accepted, rejected {{.ListName}}) {
	accepted, rejected = {{.ListName}}{}, {{.ListName}}{}
	if l != nil {
		for _, e := range l {
			if predicate(e) {
				accepted = append(accepted, e)
			} else {
				rejected = append(rejected, e)
			}
		}
	}
	return
}

func (l {{.ListName}}) Detect(predicate func(element {{.ListType}}) bool) bool {
		for _, e := range l {
			if predicate(e) {
				return true
			}
		}
	return false
}

func (l {{.ListName}}) Count(predicate func(element {{.ListType}}) bool) (count int) {
		for _, e := range l {
			if predicate(e) {
				count++
			}
		}
	return
}

func (l {{.ListName}}) AnySatisfy(predicate func(element {{.ListType}}) bool) bool {
		for _, e := range l {
			if predicate(e) {
				return true
			}
		}
	return false
}

func (l {{.ListName}}) AllSatisfy(predicate func(element {{.ListType}}) bool) bool {
		for _, e := range l {
			if predicate(e) == false {
				return false
			}
		}
	return true
}

func (l {{.ListName}}) NoneSatisfy(predicate func(element {{.ListType}}) bool) bool {
		for _, e := range l {
			if predicate(e) {
				return false
			}
		}
	return true
}

func (l {{.ListName}}) Sorted(compare func(i, j {{.ListType}}) bool) (newList {{.ListName}}) {
	newList = {{.ListName}}{}
	if l != nil {
		newList = append({{.ListName}}{}, l...)
		sort.Slice(newList,
			func(i, j int) bool {
				return compare(newList[i], newList[j])
			})
	}
	return
}


`

var mutableListTemplate = `

// Mutable functions
func {{.ListName}}_NewEmpty() (newList {{.ListName}}) {
	newList = {{.ListName}}{}
	return
}

func (l {{.ListName}}) With(element {{.ListType}}) {{.ListName}} {
	l = append(l, element)
	return l
}

func (l {{.ListName}}) Without(element {{.ListType}}) {{.ListName}} {
	for n, e := range l {
		if e == element {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

func (l {{.ListName}}) WithAll(elements []{{.ListType}}) {{.ListName}} {
	l = append(l, elements...)
	return l
}

func (l {{.ListName}}) WithoutAll(elements []{{.ListType}}) {{.ListName}} {
	for _, element := range elements {
		for n, e := range l {
			if e == element {
				l = append(l[:n], l[n+1:]...)
			}
		}
	}
	return l
}

func (l {{.ListName}}) RemoveIf(predicate func(element {{.ListType}}) bool) {{.ListName}} {
	for n, e := range l {
		if predicate(e) {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

func (l {{.ListName}}) AddAll(elements []{{.ListType}}) {{.ListName}} {
	return l.WithAll(elements)
}

func (l {{.ListName}}) RetainAll(elements []{{.ListType}}) {{.ListName}} {
	for n, e := range l {
		retain := true
		for _, element := range elements {
			if e != element {
				retain = false
				break
			}
		}
		if !retain {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

`

var testCases = `/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package {{.PackageName}}

import (
	"testing"
)

func Test{{.ListName}}Count(t *testing.T) {
	l := {{.ListName}}_NewWithAll(sample_{{.ListType}})

	if l.Size() != 5 {
		t.Fail()
	}

	if l.IsEmpty() {
		t.Fail()
	}

	if l.NotEmpty() == false {
		t.Fail()
	}
}


{{if ne .ListType "bool"}}
func Test{{.ListName}}Sort(t *testing.T) {
	l := {{.ListName}}_NewWithAll(sample_{{.ListType}})
	newList := l.Sorted(func(i, j {{.ListType}}) bool {
		return i < j
	})

	var v {{.ListType}}
	newList.Each(func(element {{.ListType}}) {
		if element < v {
			t.Fail()
		} else {
			v = element
		}
	})
}
{{end}}
`
