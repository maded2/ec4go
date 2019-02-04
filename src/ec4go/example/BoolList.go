/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"math/rand"
	"sort"
)

type BoolList struct {
	list []bool
}

// Immutable functions

func (l *BoolList) NewWith(element bool) (newList *BoolList) {
	newList = &BoolList{list: append(l.list, element)}
	return
}

func (l *BoolList) NewWithout(element bool) (newList *BoolList) {
	newList = &BoolList{}
	for _, e := range l.list {
		if e != element {
			newList.list = append(newList.list, e)
		}
	}
	return newList
}

func (l *BoolList) NewWithAll(elements []bool) (newList *BoolList) {
	newList = &BoolList{}
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

func (l *BoolList) NewWithoutAll(elements []bool) (newList *BoolList) {
	newList = &BoolList{}
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

func (l *BoolList) Size() int {
	return len(l.list)
}

func (l *BoolList) IsEmpty() bool {
	return len(l.list) == 0
}

func (l *BoolList) NotEmpty() bool {
	return len(l.list) > 0
}

func (l *BoolList) GetAny() (result bool) {
	if len(l.list) > 0 {
		result = l.list[rand.Intn(len(l.list)-1)]
	}
	return
}

func (l *BoolList) Contains(element bool) bool {
	for _, e1 := range l.list {
		if e1 == element {
			return true
		}
	}
	return false
}

func (l *BoolList) ContainsAll(elements []bool) bool {
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

func (l *BoolList) Each(procedure func(element bool)) *BoolList {
	for _, e := range l.list {
		procedure(e)
	}
	return l
}

func (l *BoolList) Select(predicate func(element bool) bool) (newList *BoolList) {
	newList = &BoolList{}
	for _, e := range l.list {
		if predicate(e) {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *BoolList) Reject(predicate func(element bool) bool) (newList *BoolList) {
	newList = &BoolList{}
	for _, e := range l.list {
		if predicate(e) == false {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *BoolList) Partition(predicate func(element bool) bool) (accepted, rejected *BoolList) {
	accepted, rejected = &BoolList{}, &BoolList{}
	for _, e := range l.list {
		if predicate(e) {
			accepted.list = append(accepted.list, e)
		} else {
			rejected.list = append(rejected.list, e)
		}
	}
	return
}

func (l *BoolList) Detect(predicate func(element bool) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *BoolList) Count(predicate func(element bool) bool) (count int) {
	for _, e := range l.list {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l *BoolList) AnySatisfy(predicate func(element bool) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *BoolList) AllSatisfy(predicate func(element bool) bool) bool {
	for _, e := range l.list {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l *BoolList) NoneSatisfy(predicate func(element bool) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l *BoolList) Sorted(compare func(i, j bool) bool) (newList *BoolList) {
	newList = &BoolList{list: append([]bool{}, l.list...)}
	sort.Slice(newList.list,
		func(i, j int) bool {
			return compare(newList.list[i], newList.list[j])
		})
	return
}

// Mutable functions
func (l *BoolList) NewEmpty() (newList *BoolList) {
	newList = &BoolList{}
	return l
}

func (l *BoolList) With(element bool) *BoolList {
	l.list = append(l.list, element)
	return l
}

func (l *BoolList) Without(element bool) *BoolList {
	for n, e := range l.list {
		if e == element {
			l.list = append(l.list[:n], l.list[n+1:]...)
		}
	}
	return l
}

func (l *BoolList) WithAll(elements []bool) *BoolList {
	l.list = append(l.list, elements...)
	return l
}

func (l *BoolList) WithoutAll(elements []bool) *BoolList {
	for _, element := range elements {
		for n, e := range l.list {
			if e == element {
				l.list = append(l.list[:n], l.list[n+1:]...)
			}
		}
	}
	return l
}

func (l *BoolList) RemoveIf(predicate func(element bool) bool) *BoolList {
	for n, e := range l.list {
		if predicate(e) {
			l.list = append(l.list[:n], l.list[n+1:]...)
		}
	}
	return l
}

func (l *BoolList) AddAll(elements []bool) *BoolList {
	return l.WithAll(elements)
}

func (l *BoolList) RemoveAll(elements []bool) *BoolList {
	return l.WithoutAll(elements)
}

func (l *BoolList) RetainAll(elements []bool) *BoolList {
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
