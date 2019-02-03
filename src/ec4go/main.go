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

	fileName := fmt.Sprintf("%s.go", flag.Lookup("name").Value.String())
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Unable to create file: %s - %s", fileName, err)
	}
	w := bufio.NewWriter(file)
	immutableListTemplate := template.Must(template.New("immutableListTemplate").Parse(immutableListTemplate))

	type value struct {
		PackageName string
		ListName    string
		ListType    string
	}
	v := value{
		PackageName: *packageName,
		ListName:    *listName,
		ListType:    *listType,
	}
	primitiveTypes := []string{"int", "int8,", "int16", "int32", "int64", "float32", "float64", "bool", "string"}

	primitive := false
	for _, pt := range primitiveTypes {
		if v.ListType == pt {
			primitive = true
			break
		}
	}
	if !primitive {
		v.ListType = "*" + v.ListType
	}

	if err = immutableListTemplate.Execute(w, v); err != nil {
		log.Fatalf("Failed to create[%s]: %s", fileName, err)
	}
	if *addMutableFunctions {
		mutableListTemplate := template.Must(template.New("mutableListTemplate").Parse(mutableListTemplate))
		if err = mutableListTemplate.Execute(w, v); err != nil {
			log.Fatalf("Failed to add mutable functions[%s]: %s", fileName, err)
		}
	}
	w.Flush()
	file.Close()
}

var immutableListTemplate = `/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/
package {{.PackageName}}

import "math/rand"

type {{.ListName}} struct {
	list []{{.ListType}}
}

// Immutable functions

func (l *{{.ListName}}) NewWith(element {{.ListType}}) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{list: append(l.list, element)}
	return
}

func (l *{{.ListName}}) NewWithout(element {{.ListType}}) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
	for _, e := range l.list {
		if e != element {
			newList.list = append(newList.list, e)
		}
	}
	return newList
}

func (l *{{.ListName}}) NewWithAll(elements []{{.ListType}}) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
	for _, e1 := range l.list {
		for _, e2 := range elements {
			if e1 != e2 {
				newList.list = append(newList.list, e1)
				break
			}
		}
	}
	return newList
}

func (l *{{.ListName}}) NewWithoutAll(elements []{{.ListType}}) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
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
	return newList
}

func (l *{{.ListName}}) Size() int {
	return len(l.list)
}

func (l *{{.ListName}}) IsEmpty() bool {
	return len(l.list) == 0
}

func (l *{{.ListName}}) NotEmpty() bool {
	return len(l.list) > 0
}

func (l *{{.ListName}}) GetAny() (result {{.ListType}}) {
	if len(l.list) > 0 {
		result = l.list[rand.Intn(len(l.list)-1)]
	}
	return
}

func (l *{{.ListName}}) Contains(element {{.ListType}}) bool {
	for _, e1 := range l.list {
		if e1 == element {
			return true
		}
	}
	return false
}

func (l *{{.ListName}}) ContainsAll(elements []{{.ListType}}) bool {
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
	return false
}

func (l *{{.ListName}}) Each(procedure func(element {{.ListType}})) *{{.ListName}} {
	for _, e := range l.list {
		procedure(e)
	}
	return l
}

func (l *{{.ListName}}) Select(predicate func(element {{.ListType}}) bool) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
	for _, e := range l.list {
		if predicate(e) {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *{{.ListName}}) Reject(predicate func(element {{.ListType}}) bool) (newList *{{.ListName}}) {
	newList = &{{.ListName}}{}
	for _, e := range l.list {
		if predicate(e) == false {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *{{.ListName}}) Partition(predicate func(element {{.ListType}}) bool) (accepted, rejected *{{.ListName}}) {
	accepted, rejected = &{{.ListName}}{}, &{{.ListName}}{}
	for _, e := range l.list {
		if predicate(e) {
			accepted.list = append(accepted.list, e)
		} else {
			rejected.list = append(rejected.list, e)
		}
	}
	return
}

func (l *{{.ListName}}) Detect(predicate func(element {{.ListType}}) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *{{.ListName}}) Count(predicate func(element {{.ListType}}) bool) (count int) {
	for _, e := range l.list {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l *{{.ListName}}) AnySatisfy(predicate func(element {{.ListType}}) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *{{.ListName}}) AllSatisfy(predicate func(element {{.ListType}}) bool) bool {
	for _, e := range l.list {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l *{{.ListName}}) NoneSatisfy(predicate func(element {{.ListType}}) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return false
		}
	}
	return true
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
	for __, element := range elements {
		for n, e := range l.list {
			if e == element {
				l.list = append(l.list[:n], l.list[n+1:]...)
			}
		}
	}
	return l
}

func (l *{{.ListName}}) RemoveIf(predicate func(element {{.ListType}}) bool) (l *{{.ListName}}) {
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

func (l *{{.ListName}}) RemoveAll(elements []{{.ListType}}) *{{.ListName}} {
	return l.WithoutAll(elements)
}

func (l *{{.ListName}}) RetainAll(elements []{{.ListType}}) *{{.ListName}} {
	for n, e := range l.list {
		retain := true
		for __, element := range elements {
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
