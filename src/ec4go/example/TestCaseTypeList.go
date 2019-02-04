/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"math/rand"
	"sort"
)

type TestCaseTypeList struct {
	list []*TestCaseType
}

// Immutable functions

func (l *TestCaseTypeList) NewWith(element *TestCaseType) (newList *TestCaseTypeList) {
	newList = &TestCaseTypeList{}
	if l != nil {
		newList.list = append(l.list, element)
	} else {
		newList.list = append([]*TestCaseType{}, element)
	}
	return
}

func (l *TestCaseTypeList) NewWithAll(elements []*TestCaseType) (newList *TestCaseTypeList) {
	newList = &TestCaseTypeList{}
	if l == nil {
		newList.list = append(l.list, elements...)
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

func (l *TestCaseTypeList) NewWithout(element *TestCaseType) (newList *TestCaseTypeList) {
	newList = &TestCaseTypeList{}
	if l != nil {
		for _, e := range l.list {
			if e != element {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *TestCaseTypeList) NewWithoutAll(elements []*TestCaseType) (newList *TestCaseTypeList) {
	newList = &TestCaseTypeList{}
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

func (l *TestCaseTypeList) Remove(element *TestCaseType) (newList *TestCaseTypeList) {
	newList = &TestCaseTypeList{}
	if l != nil {
		for _, e := range l.list {
			if e != element {
				newList.list = append(newList.list, e)
			}
		}
	}
	return newList
}

func (l *TestCaseTypeList) RemoveAll(elements []*TestCaseType) (newList *TestCaseTypeList) {
	newList = &TestCaseTypeList{}
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

func (l *TestCaseTypeList) Size() (size int) {
	if l != nil {
		size = len(l.list)
	}
	return
}

func (l *TestCaseTypeList) IsEmpty() bool {
	if l != nil {
		return len(l.list) == 0
	} else {
		return true
	}
}

func (l *TestCaseTypeList) NotEmpty() bool {
	if l != nil {
		return len(l.list) > 0
	} else {
		return false
	}
}

func (l *TestCaseTypeList) GetAny() (result *TestCaseType) {
	if l != nil && len(l.list) > 0 {
		result = l.list[rand.Intn(len(l.list)-1)]
	}
	return
}

func (l *TestCaseTypeList) Contains(element *TestCaseType) bool {
	if l != nil {
		for _, e1 := range l.list {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l *TestCaseTypeList) ContainsAll(elements []*TestCaseType) bool {
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

func (l *TestCaseTypeList) Each(procedure func(element *TestCaseType)) *TestCaseTypeList {
	if l != nil {
		for _, e := range l.list {
			procedure(e)
		}
	}
	return l
}

func (l *TestCaseTypeList) Select(predicate func(element *TestCaseType) bool) (newList *TestCaseTypeList) {
	newList = &TestCaseTypeList{}
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *TestCaseTypeList) Reject(predicate func(element *TestCaseType) bool) (newList *TestCaseTypeList) {
	newList = &TestCaseTypeList{}
	if l != nil {
		for _, e := range l.list {
			if predicate(e) == false {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *TestCaseTypeList) Partition(predicate func(element *TestCaseType) bool) (accepted, rejected *TestCaseTypeList) {
	accepted, rejected = &TestCaseTypeList{}, &TestCaseTypeList{}
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

func (l *TestCaseTypeList) Detect(predicate func(element *TestCaseType) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return true
			}
		}
	}
	return false
}

func (l *TestCaseTypeList) Count(predicate func(element *TestCaseType) bool) (count int) {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				count++
			}
		}
	}
	return
}

func (l *TestCaseTypeList) AnySatisfy(predicate func(element *TestCaseType) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return true
			}
		}
	}
	return false
}

func (l *TestCaseTypeList) AllSatisfy(predicate func(element *TestCaseType) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) == false {
				return false
			}
		}
	}
	return true
}

func (l *TestCaseTypeList) NoneSatisfy(predicate func(element *TestCaseType) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return false
			}
		}
	}
	return true
}

func (l *TestCaseTypeList) Sorted(compare func(i, j *TestCaseType) bool) (newList *TestCaseTypeList) {
	newList = &TestCaseTypeList{}
	if l != nil {
		newList.list = append([]*TestCaseType{}, l.list...)
		sort.Slice(newList.list,
			func(i, j int) bool {
				return compare(newList.list[i], newList.list[j])
			})
	}
	return
}

// Mutable functions
func (l *TestCaseTypeList) NewEmpty() (newList *TestCaseTypeList) {
	newList = &TestCaseTypeList{}
	return l
}

func (l *TestCaseTypeList) With(element *TestCaseType) *TestCaseTypeList {
	l.list = append(l.list, element)
	return l
}

func (l *TestCaseTypeList) Without(element *TestCaseType) *TestCaseTypeList {
	for n, e := range l.list {
		if e == element {
			l.list = append(l.list[:n], l.list[n+1:]...)
		}
	}
	return l
}

func (l *TestCaseTypeList) WithAll(elements []*TestCaseType) *TestCaseTypeList {
	l.list = append(l.list, elements...)
	return l
}

func (l *TestCaseTypeList) WithoutAll(elements []*TestCaseType) *TestCaseTypeList {
	for _, element := range elements {
		for n, e := range l.list {
			if e == element {
				l.list = append(l.list[:n], l.list[n+1:]...)
			}
		}
	}
	return l
}

func (l *TestCaseTypeList) RemoveIf(predicate func(element *TestCaseType) bool) *TestCaseTypeList {
	for n, e := range l.list {
		if predicate(e) {
			l.list = append(l.list[:n], l.list[n+1:]...)
		}
	}
	return l
}

func (l *TestCaseTypeList) AddAll(elements []*TestCaseType) *TestCaseTypeList {
	return l.WithAll(elements)
}

func (l *TestCaseTypeList) RemoveAll(elements []*TestCaseType) *TestCaseTypeList {
	return l.WithoutAll(elements)
}

func (l *TestCaseTypeList) RetainAll(elements []*TestCaseType) *TestCaseTypeList {
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
