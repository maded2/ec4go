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

func IntImmutableList_NewWith(element int) (newList IntImmutableList) {
	newList = append(IntImmutableList{}, element)
	return
}

func IntImmutableList_NewWithAll(elements []int) (newList IntImmutableList) {
	newList = append(IntImmutableList{}, elements...)
	return
}

func (l IntImmutableList) NewWithout(element int) (newList IntImmutableList) {
	newList = IntImmutableList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l IntImmutableList) NewWithoutAll(elements []int) (newList IntImmutableList) {
	newList = IntImmutableList{}
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

func (l IntImmutableList) Remove(element int) (newList IntImmutableList) {
	newList = IntImmutableList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return newList
}

func (l IntImmutableList) RemoveAll(elements []int) (newList IntImmutableList) {
	newList = IntImmutableList{}
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

func (l IntImmutableList) Size() (size int) {
	size = len(l)
	return
}

func (l IntImmutableList) IsEmpty() bool {
	return len(l) == 0
}

func (l IntImmutableList) NotEmpty() bool {
	return len(l) > 0
}

func (l IntImmutableList) GetAny() (result int) {
	if len(l) > 0 {
		result = l[rand.Intn(len(l)-1)]
	}
	return
}

func (l IntImmutableList) Contains(element int) bool {
	if l != nil {
		for _, e1 := range l {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l IntImmutableList) ContainsAll(elements []int) bool {
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

func (l IntImmutableList) Each(procedure func(element int)) IntImmutableList {
	if l != nil {
		for _, e := range l {
			procedure(e)
		}
	}
	return l
}

func (l IntImmutableList) Select(predicate func(element int) bool) (newList IntImmutableList) {
	newList = IntImmutableList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l IntImmutableList) Reject(predicate func(element int) bool) (newList IntImmutableList) {
	newList = IntImmutableList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) == false {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l IntImmutableList) Partition(predicate func(element int) bool) (accepted, rejected IntImmutableList) {
	accepted, rejected = IntImmutableList{}, IntImmutableList{}
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

func (l IntImmutableList) Detect(predicate func(element int) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l IntImmutableList) Count(predicate func(element int) bool) (count int) {
	for _, e := range l {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l IntImmutableList) AnySatisfy(predicate func(element int) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l IntImmutableList) AllSatisfy(predicate func(element int) bool) bool {
	for _, e := range l {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l IntImmutableList) NoneSatisfy(predicate func(element int) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l IntImmutableList) Sorted(compare func(i, j int) bool) (newList IntImmutableList) {
	newList = IntImmutableList{}
	if l != nil {
		newList = append(IntImmutableList{}, l...)
		sort.Slice(newList,
			func(i, j int) bool {
				return compare(newList[i], newList[j])
			})
	}
	return
}
