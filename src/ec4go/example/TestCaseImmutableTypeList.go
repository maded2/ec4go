/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"math/rand"
	"sort"
)

type TestCaseImmutableTypeList struct {
	list []*TestCaseType
}

// Immutable functions

func (l *TestCaseImmutableTypeList) NewWith(element *TestCaseType) (newList *TestCaseImmutableTypeList) {
	newList = &TestCaseImmutableTypeList{list: append(l.list, element)}
	return
}

func (l *TestCaseImmutableTypeList) NewWithout(element *TestCaseType) (newList *TestCaseImmutableTypeList) {
	newList = &TestCaseImmutableTypeList{}
	for _, e := range l.list {
		if e != element {
			newList.list = append(newList.list, e)
		}
	}
	return newList
}

func (l *TestCaseImmutableTypeList) NewWithAll(elements []*TestCaseType) (newList *TestCaseImmutableTypeList) {
	newList = &TestCaseImmutableTypeList{}
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

func (l *TestCaseImmutableTypeList) NewWithoutAll(elements []*TestCaseType) (newList *TestCaseImmutableTypeList) {
	newList = &TestCaseImmutableTypeList{}
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

func (l *TestCaseImmutableTypeList) Size() int {
	return len(l.list)
}

func (l *TestCaseImmutableTypeList) IsEmpty() bool {
	return len(l.list) == 0
}

func (l *TestCaseImmutableTypeList) NotEmpty() bool {
	return len(l.list) > 0
}

func (l *TestCaseImmutableTypeList) GetAny() (result *TestCaseType) {
	if len(l.list) > 0 {
		result = l.list[rand.Intn(len(l.list)-1)]
	}
	return
}

func (l *TestCaseImmutableTypeList) Contains(element *TestCaseType) bool {
	for _, e1 := range l.list {
		if e1 == element {
			return true
		}
	}
	return false
}

func (l *TestCaseImmutableTypeList) ContainsAll(elements []*TestCaseType) bool {
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

func (l *TestCaseImmutableTypeList) Each(procedure func(element *TestCaseType)) *TestCaseImmutableTypeList {
	for _, e := range l.list {
		procedure(e)
	}
	return l
}

func (l *TestCaseImmutableTypeList) Select(predicate func(element *TestCaseType) bool) (newList *TestCaseImmutableTypeList) {
	newList = &TestCaseImmutableTypeList{}
	for _, e := range l.list {
		if predicate(e) {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *TestCaseImmutableTypeList) Reject(predicate func(element *TestCaseType) bool) (newList *TestCaseImmutableTypeList) {
	newList = &TestCaseImmutableTypeList{}
	for _, e := range l.list {
		if predicate(e) == false {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *TestCaseImmutableTypeList) Partition(predicate func(element *TestCaseType) bool) (accepted, rejected *TestCaseImmutableTypeList) {
	accepted, rejected = &TestCaseImmutableTypeList{}, &TestCaseImmutableTypeList{}
	for _, e := range l.list {
		if predicate(e) {
			accepted.list = append(accepted.list, e)
		} else {
			rejected.list = append(rejected.list, e)
		}
	}
	return
}

func (l *TestCaseImmutableTypeList) Detect(predicate func(element *TestCaseType) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *TestCaseImmutableTypeList) Count(predicate func(element *TestCaseType) bool) (count int) {
	for _, e := range l.list {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l *TestCaseImmutableTypeList) AnySatisfy(predicate func(element *TestCaseType) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *TestCaseImmutableTypeList) AllSatisfy(predicate func(element *TestCaseType) bool) bool {
	for _, e := range l.list {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l *TestCaseImmutableTypeList) NoneSatisfy(predicate func(element *TestCaseType) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l *TestCaseImmutableTypeList) Sorted(compare func(i, j *TestCaseType) bool) (newList *TestCaseImmutableTypeList) {
	newList = &TestCaseImmutableTypeList{list: append([]*TestCaseType{}, l.list...)}
	sort.Slice(newList.list,
		func(i, j int) bool {
			return compare(newList.list[i], newList.list[j])
		})
	return
}
