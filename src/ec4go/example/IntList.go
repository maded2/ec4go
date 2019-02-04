/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"math/rand"
	"sort"
)

type IntList struct {
	list []int
}

// Immutable functions

func (l *IntList) NewWith(element int) (newList *IntList) {
	newList = &IntList{}
	if l != nil {
		newList.list = append(l.list, element)
	} else {
		newList.list = append([]int{}, element)
	}
	return
}

func (l *IntList) NewWithAll(elements []int) (newList *IntList) {
	newList = &IntList{}
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

func (l *IntList) NewWithout(element int) (newList *IntList) {
	newList = &IntList{}
	if l != nil {
		for _, e := range l.list {
			if e != element {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *IntList) NewWithoutAll(elements []int) (newList *IntList) {
	newList = &IntList{}
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

func (l *IntList) Remove(element int) (newList *IntList) {
	newList = &IntList{}
	if l != nil {
		for _, e := range l.list {
			if e != element {
				newList.list = append(newList.list, e)
			}
		}
	}
	return newList
}

func (l *IntList) RemoveAll(elements []int) (newList *IntList) {
	newList = &IntList{}
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

func (l *IntList) Size() (size int) {
	if l != nil {
		size = len(l.list)
	}
	return
}

func (l *IntList) IsEmpty() bool {
	if l != nil {
		return len(l.list) == 0
	} else {
		return true
	}
}

func (l *IntList) NotEmpty() bool {
	if l != nil {
		return len(l.list) > 0
	} else {
		return false
	}
}

func (l *IntList) GetAny() (result int) {
	if l != nil && len(l.list) > 0 {
		result = l.list[rand.Intn(len(l.list)-1)]
	}
	return
}

func (l *IntList) Contains(element int) bool {
	if l != nil {
		for _, e1 := range l.list {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l *IntList) ContainsAll(elements []int) bool {
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

func (l *IntList) Each(procedure func(element int)) *IntList {
	if l != nil {
		for _, e := range l.list {
			procedure(e)
		}
	}
	return l
}

func (l *IntList) Select(predicate func(element int) bool) (newList *IntList) {
	newList = &IntList{}
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *IntList) Reject(predicate func(element int) bool) (newList *IntList) {
	newList = &IntList{}
	if l != nil {
		for _, e := range l.list {
			if predicate(e) == false {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *IntList) Partition(predicate func(element int) bool) (accepted, rejected *IntList) {
	accepted, rejected = &IntList{}, &IntList{}
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

func (l *IntList) Detect(predicate func(element int) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return true
			}
		}
	}
	return false
}

func (l *IntList) Count(predicate func(element int) bool) (count int) {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				count++
			}
		}
	}
	return
}

func (l *IntList) AnySatisfy(predicate func(element int) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return true
			}
		}
	}
	return false
}

func (l *IntList) AllSatisfy(predicate func(element int) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) == false {
				return false
			}
		}
	}
	return true
}

func (l *IntList) NoneSatisfy(predicate func(element int) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return false
			}
		}
	}
	return true
}

func (l *IntList) Sorted(compare func(i, j int) bool) (newList *IntList) {
	newList = &IntList{}
	if l != nil {
		newList.list = append([]int{}, l.list...)
		sort.Slice(newList.list,
			func(i, j int) bool {
				return compare(newList.list[i], newList.list[j])
			})
	}
	return
}

// Mutable functions
func (l *IntList) NewEmpty() (newList *IntList) {
	newList = &IntList{}
	return l
}

func (l *IntList) With(element int) *IntList {
	l.list = append(l.list, element)
	return l
}

func (l *IntList) Without(element int) *IntList {
	for n, e := range l.list {
		if e == element {
			l.list = append(l.list[:n], l.list[n+1:]...)
		}
	}
	return l
}

func (l *IntList) WithAll(elements []int) *IntList {
	l.list = append(l.list, elements...)
	return l
}

func (l *IntList) WithoutAll(elements []int) *IntList {
	for _, element := range elements {
		for n, e := range l.list {
			if e == element {
				l.list = append(l.list[:n], l.list[n+1:]...)
			}
		}
	}
	return l
}

func (l *IntList) RemoveIf(predicate func(element int) bool) *IntList {
	for n, e := range l.list {
		if predicate(e) {
			l.list = append(l.list[:n], l.list[n+1:]...)
		}
	}
	return l
}

func (l *IntList) AddAll(elements []int) *IntList {
	return l.WithAll(elements)
}

func (l *IntList) RemoveAll(elements []int) *IntList {
	return l.WithoutAll(elements)
}

func (l *IntList) RetainAll(elements []int) *IntList {
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
