package main

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

func {{.TypeName}}_NewWith(element {{.ValueType}}) (newList {{.TypeName}}) {
	newList = append({{.TypeName}}{}, element)
	return
}

func {{.TypeName}}_NewWithAll(elements []{{.ValueType}}) (newList {{.TypeName}}) {
	newList = append({{.TypeName}}{}, elements...)
	return
}

func (l {{.TypeName}}) NewWithout(element {{.ValueType}}) (newList {{.TypeName}}) {
	newList = {{.TypeName}}{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l {{.TypeName}}) NewWithoutAll(elements []{{.ValueType}}) (newList {{.TypeName}}) {
	newList = {{.TypeName}}{}
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

func (l {{.TypeName}}) Remove(element {{.ValueType}}) (newList {{.TypeName}}) {
	newList = {{.TypeName}}{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return newList
}

func (l {{.TypeName}}) RemoveAll(elements []{{.ValueType}}) (newList {{.TypeName}}) {
	newList = {{.TypeName}}{}
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

func (l {{.TypeName}}) Size() (size int) {
	size = len(l)
	return
}

func (l {{.TypeName}}) IsEmpty() bool {
	return len(l) == 0
}

func (l {{.TypeName}}) NotEmpty() bool {
	return len(l) > 0
}

func (l {{.TypeName}}) GetAny() (result {{.ValueType}}) {
	if len(l) > 0 {
		result = l[rand.Intn(len(l)-1)]
	}
	return
}

func (l {{.TypeName}}) Contains(element {{.ValueType}}) bool {
	if l != nil {
		for _, e1 := range l {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l {{.TypeName}}) ContainsAll(elements []{{.ValueType}}) bool {
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

func (l {{.TypeName}}) Each(procedure func(element {{.ValueType}})) {{.TypeName}} {
	if l != nil {
		for _, e := range l {
			procedure(e)
		}
	}
	return l
}

func (l {{.TypeName}}) Select(predicate func(element {{.ValueType}}) bool) (newList {{.TypeName}}) {
	newList = {{.TypeName}}{}
	if l != nil {
		for _, e := range l {
			if predicate(e) {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l {{.TypeName}}) Reject(predicate func(element {{.ValueType}}) bool) (newList {{.TypeName}}) {
	newList = {{.TypeName}}{}
	if l != nil {
		for _, e := range l {
			if predicate(e) == false {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l {{.TypeName}}) Partition(predicate func(element {{.ValueType}}) bool) (accepted, rejected {{.TypeName}}) {
	accepted, rejected = {{.TypeName}}{}, {{.TypeName}}{}
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

func (l {{.TypeName}}) Detect(predicate func(element {{.ValueType}}) bool) bool {
		for _, e := range l {
			if predicate(e) {
				return true
			}
		}
	return false
}

func (l {{.TypeName}}) Count(predicate func(element {{.ValueType}}) bool) (count int) {
		for _, e := range l {
			if predicate(e) {
				count++
			}
		}
	return
}

func (l {{.TypeName}}) AnySatisfy(predicate func(element {{.ValueType}}) bool) bool {
		for _, e := range l {
			if predicate(e) {
				return true
			}
		}
	return false
}

func (l {{.TypeName}}) AllSatisfy(predicate func(element {{.ValueType}}) bool) bool {
		for _, e := range l {
			if predicate(e) == false {
				return false
			}
		}
	return true
}

func (l {{.TypeName}}) NoneSatisfy(predicate func(element {{.ValueType}}) bool) bool {
		for _, e := range l {
			if predicate(e) {
				return false
			}
		}
	return true
}

func (l {{.TypeName}}) Sorted(compare func(i, j {{.ValueType}}) bool) (newList {{.TypeName}}) {
	newList = {{.TypeName}}{}
	if l != nil {
		newList = append({{.TypeName}}{}, l...)
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
func {{.TypeName}}_NewEmpty() (newList {{.TypeName}}) {
	newList = {{.TypeName}}{}
	return
}

func (l {{.TypeName}}) With(element {{.ValueType}}) {{.TypeName}} {
	l = append(l, element)
	return l
}

func (l {{.TypeName}}) Without(element {{.ValueType}}) {{.TypeName}} {
	for n, e := range l {
		if e == element {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

func (l {{.TypeName}}) WithAll(elements []{{.ValueType}}) {{.TypeName}} {
	l = append(l, elements...)
	return l
}

func (l {{.TypeName}}) WithoutAll(elements []{{.ValueType}}) {{.TypeName}} {
	for _, element := range elements {
		for n, e := range l {
			if e == element {
				l = append(l[:n], l[n+1:]...)
			}
		}
	}
	return l
}

func (l {{.TypeName}}) RemoveIf(predicate func(element {{.ValueType}}) bool) {{.TypeName}} {
	for n, e := range l {
		if predicate(e) {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

func (l {{.TypeName}}) AddAll(elements []{{.ValueType}}) {{.TypeName}} {
	return l.WithAll(elements)
}

func (l {{.TypeName}}) RetainAll(elements []{{.ValueType}}) {{.TypeName}} {
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
