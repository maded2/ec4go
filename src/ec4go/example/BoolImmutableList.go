/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"math/rand"
	"sort"
)

type BoolImmutableList struct {
	list []bool
}

// Immutable functions

func (l *BoolImmutableList) NewWith(element bool) (newList *BoolImmutableList) {
	newList = &BoolImmutableList{}
	if l != nil {
		newList.list = append(l.list, element)
	} else {
		newList.list = append([]bool{}, element)
	}
	return
}

func (l *BoolImmutableList) NewWithAll(elements []bool) (newList *BoolImmutableList) {
	newList = &BoolImmutableList{}
	if l == nil {
		newList.list = append([]bool{}, elements...)
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

func (l *BoolImmutableList) NewWithout(element bool) (newList *BoolImmutableList) {
	newList = &BoolImmutableList{}
	if l != nil {
		for _, e := range l.list {
			if e != element {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *BoolImmutableList) NewWithoutAll(elements []bool) (newList *BoolImmutableList) {
	newList = &BoolImmutableList{}
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

func (l *BoolImmutableList) Remove(element bool) (newList *BoolImmutableList) {
	newList = &BoolImmutableList{}
	if l != nil {
		for _, e := range l.list {
			if e != element {
				newList.list = append(newList.list, e)
			}
		}
	}
	return newList
}

func (l *BoolImmutableList) RemoveAll(elements []bool) (newList *BoolImmutableList) {
	newList = &BoolImmutableList{}
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

func (l *BoolImmutableList) Size() (size int) {
	if l != nil {
		size = len(l.list)
	}
	return
}

func (l *BoolImmutableList) IsEmpty() bool {
	if l != nil {
		return len(l.list) == 0
	} else {
		return true
	}
}

func (l *BoolImmutableList) NotEmpty() bool {
	if l != nil {
		return len(l.list) > 0
	} else {
		return false
	}
}

func (l *BoolImmutableList) GetAny() (result bool) {
	if l != nil && len(l.list) > 0 {
		result = l.list[rand.Intn(len(l.list)-1)]
	}
	return
}

func (l *BoolImmutableList) Contains(element bool) bool {
	if l != nil {
		for _, e1 := range l.list {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l *BoolImmutableList) ContainsAll(elements []bool) bool {
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

func (l *BoolImmutableList) Each(procedure func(element bool)) *BoolImmutableList {
	if l != nil {
		for _, e := range l.list {
			procedure(e)
		}
	}
	return l
}

func (l *BoolImmutableList) Select(predicate func(element bool) bool) (newList *BoolImmutableList) {
	newList = &BoolImmutableList{}
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *BoolImmutableList) Reject(predicate func(element bool) bool) (newList *BoolImmutableList) {
	newList = &BoolImmutableList{}
	if l != nil {
		for _, e := range l.list {
			if predicate(e) == false {
				newList.list = append(newList.list, e)
			}
		}
	}
	return
}

func (l *BoolImmutableList) Partition(predicate func(element bool) bool) (accepted, rejected *BoolImmutableList) {
	accepted, rejected = &BoolImmutableList{}, &BoolImmutableList{}
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

func (l *BoolImmutableList) Detect(predicate func(element bool) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return true
			}
		}
	}
	return false
}

func (l *BoolImmutableList) Count(predicate func(element bool) bool) (count int) {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				count++
			}
		}
	}
	return
}

func (l *BoolImmutableList) AnySatisfy(predicate func(element bool) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return true
			}
		}
	}
	return false
}

func (l *BoolImmutableList) AllSatisfy(predicate func(element bool) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) == false {
				return false
			}
		}
	}
	return true
}

func (l *BoolImmutableList) NoneSatisfy(predicate func(element bool) bool) bool {
	if l != nil {
		for _, e := range l.list {
			if predicate(e) {
				return false
			}
		}
	}
	return true
}

func (l *BoolImmutableList) Sorted(compare func(i, j bool) bool) (newList *BoolImmutableList) {
	newList = &BoolImmutableList{}
	if l != nil {
		newList.list = append([]bool{}, l.list...)
		sort.Slice(newList.list,
			func(i, j int) bool {
				return compare(newList.list[i], newList.list[j])
			})
	}
	return
}
