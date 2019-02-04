/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"math/rand"
	"sort"
)

type StringImmutableList struct {
	list []string
}

// Immutable functions

func (l *StringImmutableList) NewWith(element string) (newList *StringImmutableList) {
	newList = &StringImmutableList{list: append(l.list, element)}
	return
}

func (l *StringImmutableList) NewWithout(element string) (newList *StringImmutableList) {
	newList = &StringImmutableList{}
	for _, e := range l.list {
		if e != element {
			newList.list = append(newList.list, e)
		}
	}
	return newList
}

func (l *StringImmutableList) NewWithAll(elements []string) (newList *StringImmutableList) {
	newList = &StringImmutableList{}
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

func (l *StringImmutableList) NewWithoutAll(elements []string) (newList *StringImmutableList) {
	newList = &StringImmutableList{}
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

func (l *StringImmutableList) Size() int {
	return len(l.list)
}

func (l *StringImmutableList) IsEmpty() bool {
	return len(l.list) == 0
}

func (l *StringImmutableList) NotEmpty() bool {
	return len(l.list) > 0
}

func (l *StringImmutableList) GetAny() (result string) {
	if len(l.list) > 0 {
		result = l.list[rand.Intn(len(l.list)-1)]
	}
	return
}

func (l *StringImmutableList) Contains(element string) bool {
	for _, e1 := range l.list {
		if e1 == element {
			return true
		}
	}
	return false
}

func (l *StringImmutableList) ContainsAll(elements []string) bool {
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

func (l *StringImmutableList) Each(procedure func(element string)) *StringImmutableList {
	for _, e := range l.list {
		procedure(e)
	}
	return l
}

func (l *StringImmutableList) Select(predicate func(element string) bool) (newList *StringImmutableList) {
	newList = &StringImmutableList{}
	for _, e := range l.list {
		if predicate(e) {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *StringImmutableList) Reject(predicate func(element string) bool) (newList *StringImmutableList) {
	newList = &StringImmutableList{}
	for _, e := range l.list {
		if predicate(e) == false {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *StringImmutableList) Partition(predicate func(element string) bool) (accepted, rejected *StringImmutableList) {
	accepted, rejected = &StringImmutableList{}, &StringImmutableList{}
	for _, e := range l.list {
		if predicate(e) {
			accepted.list = append(accepted.list, e)
		} else {
			rejected.list = append(rejected.list, e)
		}
	}
	return
}

func (l *StringImmutableList) Detect(predicate func(element string) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *StringImmutableList) Count(predicate func(element string) bool) (count int) {
	for _, e := range l.list {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l *StringImmutableList) AnySatisfy(predicate func(element string) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *StringImmutableList) AllSatisfy(predicate func(element string) bool) bool {
	for _, e := range l.list {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l *StringImmutableList) NoneSatisfy(predicate func(element string) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l *StringImmutableList) Sorted(compare func(i, j string) bool) (newList *StringImmutableList) {
	newList = &StringImmutableList{list: append([]string{}, l.list...)}
	sort.Slice(newList.list,
		func(i, j int) bool {
			return compare(newList.list[i], newList.list[j])
		})
	return
}
