/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"math/rand"
	"sort"
)

type IntImmutableList struct {
	list []int
}

// Immutable functions

func (l *IntImmutableList) NewWith(element int) (newList *IntImmutableList) {
	newList = &IntImmutableList{}
	if l != nil {
		newList.list = append(l.list, element)
	} else {
		newList.list = append([]int{}, element)
	}
	return
}

func (l *IntImmutableList) NewWithAll(elements []int) (newList *IntImmutableList) {
	newList = &IntImmutableList{}
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

func (l *IntImmutableList) NewWithout(element int) (newList *IntImmutableList) {
	newList = &IntImmutableList{}
	if l != nil {
		for _, e := range l.list {
			if e != element {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *IntImmutableList) NewWithoutAll(elements []int) (newList *IntImmutableList) {
	newList = &IntImmutableList{}
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

func (l *IntImmutableList) Remove(element int) (newList *IntImmutableList) {
	newList = &IntImmutableList{}
	if l != nil {
		for _, e := range l.list {
			if e != element {
				newList.list = append(newList.list, e)
			}
		}
	}
	return newList
}

func (l *IntImmutableList) RemoveAll(elements []int) (newList *IntImmutableList) {
	newList = &IntImmutableList{}
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

func (l *IntImmutableList) Size() (size int) {
	if l != nil {
		size = len(l.list)
	}
	return
}

func (l *IntImmutableList) IsEmpty() bool {
	if l != nil {
		return len(l.list) == 0
	} else {
		return true
	}
}

func (l *IntImmutableList) NotEmpty() bool {
	if l != nil {
		return len(l.list) > 0
	} else {
		return false
	}
}

func (l *IntImmutableList) GetAny() (result int) {
	if l != nil && len(l.list) > 0 {
		result = l.list[rand.Intn(len(l.list)-1)]
	}
	return
}

func (l *IntImmutableList) Contains(element int) bool {
	if l != nil {
		for _, e1 := range l.list {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l *IntImmutableList) ContainsAll(elements []int) bool {
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

func (l *IntImmutableList) Each(procedure func(element int)) *IntImmutableList {
	if l != nil {
		for _, e := range l.list {
			procedure(e)
		}
	}
	return l
}

func (l *IntImmutableList) Select(predicate func(element int) bool) (newList *IntImmutableList) {
	newList = &IntImmutableList{}
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *IntImmutableList) Reject(predicate func(element int) bool) (newList *IntImmutableList) {
	newList = &IntImmutableList{}
	if l != nil {
		for _, e := range l.list {
			if predicate(e) == false {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *IntImmutableList) Partition(predicate func(element int) bool) (accepted, rejected *IntImmutableList) {
	accepted, rejected = &IntImmutableList{}, &IntImmutableList{}
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

func (l *IntImmutableList) Detect(predicate func(element int) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return true
			}
		}
	}
	return false
}

func (l *IntImmutableList) Count(predicate func(element int) bool) (count int) {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				count++
			}
		}
	}
	return
}

func (l *IntImmutableList) AnySatisfy(predicate func(element int) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return true
			}
		}
	}
	return false
}

func (l *IntImmutableList) AllSatisfy(predicate func(element int) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) == false {
				return false
			}
		}
	}
	return true
}

func (l *IntImmutableList) NoneSatisfy(predicate func(element int) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return false
			}
		}
	}
	return true
}

func (l *IntImmutableList) Sorted(compare func(i, j int) bool) (newList *IntImmutableList) {
	newList = &IntImmutableList{}
	if l != nil {
		newList.list = append([]int{}, l.list...)
		sort.Slice(newList.list,
			func(i, j int) bool {
				return compare(newList.list[i], newList.list[j])
			})
	}
	return
}
