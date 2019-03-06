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

func BoolList_NewWith(element bool) (newList BoolList) {
	newList = append(BoolList{}, element)
	return
}

func BoolList_NewWithAll(elements []bool) (newList BoolList) {
	newList = append(BoolList{}, elements...)
	return
}

func (l BoolList) NewWithout(element bool) (newList BoolList) {
	newList = BoolList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l BoolList) NewWithoutAll(elements []bool) (newList BoolList) {
	newList = BoolList{}
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

func (l BoolList) Remove(element bool) (newList BoolList) {
	newList = BoolList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return newList
}

func (l BoolList) RemoveAll(elements []bool) (newList BoolList) {
	newList = BoolList{}
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

func (l BoolList) Size() (size int) {
	size = len(l)
	return
}

func (l BoolList) IsEmpty() bool {
	return len(l) == 0
}

func (l BoolList) NotEmpty() bool {
	return len(l) > 0
}

func (l BoolList) GetAny() (result bool) {
	if len(l) > 0 {
		result = l[rand.Intn(len(l)-1)]
	}
	return
}

func (l BoolList) Contains(element bool) bool {
	if l != nil {
		for _, e1 := range l {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l BoolList) ContainsAll(elements []bool) bool {
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

func (l BoolList) Each(procedure func(element bool)) BoolList {
	if l != nil {
		for _, e := range l {
			procedure(e)
		}
	}
	return l
}

func (l BoolList) Select(predicate func(element bool) bool) (newList BoolList) {
	newList = BoolList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l BoolList) Reject(predicate func(element bool) bool) (newList BoolList) {
	newList = BoolList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) == false {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l BoolList) Partition(predicate func(element bool) bool) (accepted, rejected BoolList) {
	accepted, rejected = BoolList{}, BoolList{}
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

func (l BoolList) Detect(predicate func(element bool) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l BoolList) Count(predicate func(element bool) bool) (count int) {
	for _, e := range l {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l BoolList) AnySatisfy(predicate func(element bool) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l BoolList) AllSatisfy(predicate func(element bool) bool) bool {
	for _, e := range l {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l BoolList) NoneSatisfy(predicate func(element bool) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l BoolList) Sorted(compare func(i, j bool) bool) (newList BoolList) {
	newList = BoolList{}
	if l != nil {
		newList = append(BoolList{}, l...)
		sort.Slice(newList,
			func(i, j int) bool {
				return compare(newList[i], newList[j])
			})
	}
	return
}

// Mutable functions
func BoolList_NewEmpty() (newList BoolList) {
	newList = BoolList{}
	return
}

func (l BoolList) With(element bool) BoolList {
	l = append(l, element)
	return l
}

func (l BoolList) Without(element bool) BoolList {
	for n, e := range l {
		if e == element {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

func (l BoolList) WithAll(elements []bool) BoolList {
	l = append(l, elements...)
	return l
}

func (l BoolList) WithoutAll(elements []bool) BoolList {
	for _, element := range elements {
		for n, e := range l {
			if e == element {
				l = append(l[:n], l[n+1:]...)
			}
		}
	}
	return l
}

func (l BoolList) RemoveIf(predicate func(element bool) bool) BoolList {
	for n, e := range l {
		if predicate(e) {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

func (l BoolList) AddAll(elements []bool) BoolList {
	return l.WithAll(elements)
}

func (l BoolList) RetainAll(elements []bool) BoolList {
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
