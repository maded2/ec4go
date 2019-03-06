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

func StringImmutableList_NewWith(element string) (newList StringImmutableList) {
	newList = append(StringImmutableList{}, element)
	return
}

func StringImmutableList_NewWithAll(elements []string) (newList StringImmutableList) {
	newList = append(StringImmutableList{}, elements...)
	return
}

func (l StringImmutableList) NewWithout(element string) (newList StringImmutableList) {
	newList = StringImmutableList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l StringImmutableList) NewWithoutAll(elements []string) (newList StringImmutableList) {
	newList = StringImmutableList{}
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

func (l StringImmutableList) Remove(element string) (newList StringImmutableList) {
	newList = StringImmutableList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return newList
}

func (l StringImmutableList) RemoveAll(elements []string) (newList StringImmutableList) {
	newList = StringImmutableList{}
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

func (l StringImmutableList) Size() (size int) {
	size = len(l)
	return
}

func (l StringImmutableList) IsEmpty() bool {
	return len(l) == 0
}

func (l StringImmutableList) NotEmpty() bool {
	return len(l) > 0
}

func (l StringImmutableList) GetAny() (result string) {
	if len(l) > 0 {
		result = l[rand.Intn(len(l)-1)]
	}
	return
}

func (l StringImmutableList) Contains(element string) bool {
	if l != nil {
		for _, e1 := range l {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l StringImmutableList) ContainsAll(elements []string) bool {
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

func (l StringImmutableList) Each(procedure func(element string)) StringImmutableList {
	if l != nil {
		for _, e := range l {
			procedure(e)
		}
	}
	return l
}

func (l StringImmutableList) Select(predicate func(element string) bool) (newList StringImmutableList) {
	newList = StringImmutableList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l StringImmutableList) Reject(predicate func(element string) bool) (newList StringImmutableList) {
	newList = StringImmutableList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) == false {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l StringImmutableList) Partition(predicate func(element string) bool) (accepted, rejected StringImmutableList) {
	accepted, rejected = StringImmutableList{}, StringImmutableList{}
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

func (l StringImmutableList) Detect(predicate func(element string) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l StringImmutableList) Count(predicate func(element string) bool) (count int) {
	for _, e := range l {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l StringImmutableList) AnySatisfy(predicate func(element string) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l StringImmutableList) AllSatisfy(predicate func(element string) bool) bool {
	for _, e := range l {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l StringImmutableList) NoneSatisfy(predicate func(element string) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l StringImmutableList) Sorted(compare func(i, j string) bool) (newList StringImmutableList) {
	newList = StringImmutableList{}
	if l != nil {
		newList = append(StringImmutableList{}, l...)
		sort.Slice(newList,
			func(i, j int) bool {
				return compare(newList[i], newList[j])
			})
	}
	return
}
