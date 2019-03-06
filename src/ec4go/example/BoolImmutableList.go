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

func BoolImmutableList_NewWith(element bool) (newList BoolImmutableList) {
	newList = append(BoolImmutableList{}, element)
	return
}

func BoolImmutableList_NewWithAll(elements []bool) (newList BoolImmutableList) {
	newList = append(BoolImmutableList{}, elements...)
	return
}

func (l BoolImmutableList) NewWithout(element bool) (newList BoolImmutableList) {
	newList = BoolImmutableList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l BoolImmutableList) NewWithoutAll(elements []bool) (newList BoolImmutableList) {
	newList = BoolImmutableList{}
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

func (l BoolImmutableList) Remove(element bool) (newList BoolImmutableList) {
	newList = BoolImmutableList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return newList
}

func (l BoolImmutableList) RemoveAll(elements []bool) (newList BoolImmutableList) {
	newList = BoolImmutableList{}
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

func (l BoolImmutableList) Size() (size int) {
	size = len(l)
	return
}

func (l BoolImmutableList) IsEmpty() bool {
	return len(l) == 0
}

func (l BoolImmutableList) NotEmpty() bool {
	return len(l) > 0
}

func (l BoolImmutableList) GetAny() (result bool) {
	if len(l) > 0 {
		result = l[rand.Intn(len(l)-1)]
	}
	return
}

func (l BoolImmutableList) Contains(element bool) bool {
	if l != nil {
		for _, e1 := range l {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l BoolImmutableList) ContainsAll(elements []bool) bool {
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

func (l BoolImmutableList) Each(procedure func(element bool)) BoolImmutableList {
	if l != nil {
		for _, e := range l {
			procedure(e)
		}
	}
	return l
}

func (l BoolImmutableList) Select(predicate func(element bool) bool) (newList BoolImmutableList) {
	newList = BoolImmutableList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l BoolImmutableList) Reject(predicate func(element bool) bool) (newList BoolImmutableList) {
	newList = BoolImmutableList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) == false {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l BoolImmutableList) Partition(predicate func(element bool) bool) (accepted, rejected BoolImmutableList) {
	accepted, rejected = BoolImmutableList{}, BoolImmutableList{}
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

func (l BoolImmutableList) Detect(predicate func(element bool) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l BoolImmutableList) Count(predicate func(element bool) bool) (count int) {
	for _, e := range l {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l BoolImmutableList) AnySatisfy(predicate func(element bool) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l BoolImmutableList) AllSatisfy(predicate func(element bool) bool) bool {
	for _, e := range l {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l BoolImmutableList) NoneSatisfy(predicate func(element bool) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l BoolImmutableList) Sorted(compare func(i, j bool) bool) (newList BoolImmutableList) {
	newList = BoolImmutableList{}
	if l != nil {
		newList = append(BoolImmutableList{}, l...)
		sort.Slice(newList,
			func(i, j int) bool {
				return compare(newList[i], newList[j])
			})
	}
	return
}
