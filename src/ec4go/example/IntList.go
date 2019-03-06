/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"math/rand"
	"sort"
)

// Immutable functions

func IntList_NewWith(element int) (newList IntList) {
	newList = append(IntList{}, element)
	return
}

func IntList_NewWithAll(elements []int) (newList IntList) {
	newList = append(IntList{}, elements...)
	return
}

func (l IntList) NewWithout(element int) (newList IntList) {
	newList = IntList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l IntList) NewWithoutAll(elements []int) (newList IntList) {
	newList = IntList{}
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

func (l IntList) Remove(element int) (newList IntList) {
	newList = IntList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return newList
}

func (l IntList) RemoveAll(elements []int) (newList IntList) {
	newList = IntList{}
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

func (l IntList) Size() (size int) {
	size = len(l)
	return
}

func (l IntList) IsEmpty() bool {
	return len(l) == 0
}

func (l IntList) NotEmpty() bool {
	return len(l) > 0
}

func (l IntList) GetAny() (result int) {
	if len(l) > 0 {
		result = l[rand.Intn(len(l)-1)]
	}
	return
}

func (l IntList) Contains(element int) bool {
	if l != nil {
		for _, e1 := range l {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l IntList) ContainsAll(elements []int) bool {
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

func (l IntList) Each(procedure func(element int)) IntList {
	if l != nil {
		for _, e := range l {
			procedure(e)
		}
	}
	return l
}

func (l IntList) Select(predicate func(element int) bool) (newList IntList) {
	newList = IntList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l IntList) Reject(predicate func(element int) bool) (newList IntList) {
	newList = IntList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) == false {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l IntList) Partition(predicate func(element int) bool) (accepted, rejected IntList) {
	accepted, rejected = IntList{}, IntList{}
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

func (l IntList) Detect(predicate func(element int) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l IntList) Count(predicate func(element int) bool) (count int) {
	for _, e := range l {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l IntList) AnySatisfy(predicate func(element int) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l IntList) AllSatisfy(predicate func(element int) bool) bool {
	for _, e := range l {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l IntList) NoneSatisfy(predicate func(element int) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l IntList) Sorted(compare func(i, j int) bool) (newList IntList) {
	newList = IntList{}
	if l != nil {
		newList = append(IntList{}, l...)
		sort.Slice(newList,
			func(i, j int) bool {
				return compare(newList[i], newList[j])
			})
	}
	return
}

// Mutable functions
func IntList_NewEmpty() (newList IntList) {
	newList = IntList{}
	return
}

func (l IntList) With(element int) IntList {
	l = append(l, element)
	return l
}

func (l IntList) Without(element int) IntList {
	for n, e := range l {
		if e == element {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

func (l IntList) WithAll(elements []int) IntList {
	l = append(l, elements...)
	return l
}

func (l IntList) WithoutAll(elements []int) IntList {
	for _, element := range elements {
		for n, e := range l {
			if e == element {
				l = append(l[:n], l[n+1:]...)
			}
		}
	}
	return l
}

func (l IntList) RemoveIf(predicate func(element int) bool) IntList {
	for n, e := range l {
		if predicate(e) {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

func (l IntList) AddAll(elements []int) IntList {
	return l.WithAll(elements)
}

func (l IntList) RetainAll(elements []int) IntList {
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
