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
	primitiveTypes := []string{"int", "int8,", "int16", "int32", "int64", "float32", "float64", "{{.ListType}}", "bool", "string"}

	v.Primitive = false
	for _, pt := range primitiveTypes {
		if v.ListType == pt {
			v.Primitive = true
			break
		}
	}
	if !v.Primitive {
		v.ListType = "*" + v.ListType
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

var immutableListTemplate = `/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/
package {{.PackageName}}

import (
	"math/rand"
	"sort"
)

type {{.ListName}} struct {
	list []{{.ListType}}
}

// Immutable functions

func (l *{{.ListName}}) NewWith(element {{.ListType}}) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
	if l != nil {
		newList.list = append(l.list, element)
	} else {
		newList.list = append([]{{.ListType}}{}, element)
	}
	return
}

func (l *{{.ListName}}) NewWithAll(elements []{{.ListType}}) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
	if l == nil {
		newList.list = append([]{{.ListType}}{}, elements...)
	} else {
		for _, e1 := range l.list {
			for _, e2 := range elements {
				if e1 != e2 {
					newList.list = append(newList.list, e1)
					break
				}
			}
		}
	}

	return newList
}

func (l *{{.ListName}}) NewWithout(element {{.ListType}}) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
	if l != nil {
		for _, e := range l.list {
			if e != element {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *{{.ListName}}) NewWithoutAll(elements []{{.ListType}}) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
	if l != nil {
		for _, e1 := range l.list {
			found := false
			for _, e2 := range elements {
				if e1 == e2 {
					found = true
					break
				}
			}
			if !found {
				newList.list = append(newList.list, e1)
			}
		}
	}
	return newList
}

func (l *{{.ListName}}) Remove(element {{.ListType}}) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
	if l != nil {
		for _, e := range l.list {
			if e != element {
				newList.list = append(newList.list, e)
			}
		}
	}
	return newList
}

func (l *{{.ListName}}) RemoveAll(elements []{{.ListType}}) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
	if l != nil {
		for _, e1 := range l.list {
			found := false
			for _, e2 := range elements {
				if e1 == e2 {
					found = true
					break
				}
			}
			if !found {
				newList.list = append(newList.list, e1)
			}
		}
	}
	return newList
}

func (l *{{.ListName}}) Size() (size int) {
	if l != nil {
		size = len(l.list)
	}
	return
}

func (l *{{.ListName}}) IsEmpty() bool {
	if l != nil {
		return len(l.list) == 0
	} else {
		return true
	}
}

func (l *{{.ListName}}) NotEmpty() bool {
	if l != nil {
		return len(l.list) > 0
	} else {
		return false
	}
}

func (l *{{.ListName}}) GetAny() (result {{.ListType}}) {
	if l != nil && len(l.list) > 0 {
		result = l.list[rand.Intn(len(l.list)-1)]
	}
	return
}

func (l *{{.ListName}}) Contains(element {{.ListType}}) bool {
	if l != nil {
		for _, e1 := range l.list {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l *{{.ListName}}) ContainsAll(elements []{{.ListType}}) bool {
	if l != nil {
		n := 0
		for _, e1 := range l.list {
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

func (l *{{.ListName}}) Each(procedure func(element {{.ListType}})) *{{.ListName}} {
	if l != nil {
		for _, e := range l.list {
			procedure(e)
		}
	}
	return l
}

func (l *{{.ListName}}) Select(predicate func(element {{.ListType}}) bool) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *{{.ListName}}) Reject(predicate func(element {{.ListType}}) bool) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
	if l != nil {
		for _, e := range l.list {
			if predicate(e) == false {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *{{.ListName}}) Partition(predicate func(element {{.ListType}}) bool) (accepted, rejected *{{.ListName}}) {
	accepted, rejected = &{{.ListName}}{}, &{{.ListName}}{}
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				accepted.list = append(accepted.list, e)
			} else {
				rejected.list = append(rejected.list, e)
			}
		}
	}
	return
}

func (l *{{.ListName}}) Detect(predicate func(element {{.ListType}}) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return true
			}
		}
	}
	return false
}

func (l *{{.ListName}}) Count(predicate func(element {{.ListType}}) bool) (count int) {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				count++
			}
		}
	}
	return
}

func (l *{{.ListName}}) AnySatisfy(predicate func(element {{.ListType}}) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return true
			}
		}
	}
	return false
}

func (l *{{.ListName}}) AllSatisfy(predicate func(element {{.ListType}}) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) == false {
				return false
			}
		}
	}
	return true
}

func (l *{{.ListName}}) NoneSatisfy(predicate func(element {{.ListType}}) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return false
			}
		}
	}
	return true
}

func (l *{{.ListName}}) Sorted(compare func(i, j {{.ListType}}) bool) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
	if l != nil {
		newList.list = append([]{{.ListType}}{}, l.list...)
		sort.Slice(newList.list,
			func(i, j int) bool {
				return compare(newList.list[i], newList.list[j])
			})
	}
	return
}


`

var mutableListTemplate = `

// Mutable functions
func (l *{{.ListName}}) NewEmpty() (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
	return l
}

func (l *{{.ListName}}) With(element {{.ListType}}) *{{.ListName}} {
	l.list = append(l.list, element)
	return l
}

func (l *{{.ListName}}) Without(element {{.ListType}}) *{{.ListName}} {
	for n, e := range l.list {
		if e == element {
			l.list = append(l.list[:n], l.list[n+1:]...)
		}
	}
	return l
}

func (l *{{.ListName}}) WithAll(elements []{{.ListType}}) *{{.ListName}} {
	l.list = append(l.list, elements...)
	return l
}

func (l *{{.ListName}}) WithoutAll(elements []{{.ListType}}) *{{.ListName}} {
	for _, element := range elements {
		for n, e := range l.list {
			if e == element {
				l.list = append(l.list[:n], l.list[n+1:]...)
			}
		}
	}
	return l
}

func (l *{{.ListName}}) RemoveIf(predicate func(element {{.ListType}}) bool) *{{.ListName}} {
	for n, e := range l.list {
		if predicate(e) {
			l.list = append(l.list[:n], l.list[n+1:]...)
		}
	}
	return l
}

func (l *{{.ListName}}) AddAll(elements []{{.ListType}}) *{{.ListName}} {
	return l.WithAll(elements)
}

func (l *{{.ListName}}) RetainAll(elements []{{.ListType}}) *{{.ListName}} {
	for n, e := range l.list {
		retain := true
		for _, element := range elements {
			if e != element {
				retain = false
				break
			}
		}
		if !retain {
			l.list = append(l.list[:n], l.list[n+1:]...)
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
	l := (*{{.ListName}})(nil).NewWithAll([]{{.ListType}}{5.0, 4.0, 3.0, 2.0, 1.0})

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


func Test{{.ListName}}Sort(t *testing.T) {
	l := (*{{.ListName}})(nil).NewWithAll([]{{.ListType}}{5.0, 4.0, 3.0, 2.0, 1.0})
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
`
