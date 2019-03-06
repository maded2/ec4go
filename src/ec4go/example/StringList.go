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

func StringList_NewWith(element string) (newList StringList) {
	newList = append(StringList{}, element)
	return
}

func StringList_NewWithAll(elements []string) (newList StringList) {
	newList = append(StringList{}, elements...)
	return
}

func (l StringList) NewWithout(element string) (newList StringList) {
	newList = StringList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l StringList) NewWithoutAll(elements []string) (newList StringList) {
	newList = StringList{}
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

func (l StringList) Remove(element string) (newList StringList) {
	newList = StringList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return newList
}

func (l StringList) RemoveAll(elements []string) (newList StringList) {
	newList = StringList{}
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

func (l StringList) Size() (size int) {
	size = len(l)
	return
}

func (l StringList) IsEmpty() bool {
	return len(l) == 0
}

func (l StringList) NotEmpty() bool {
	return len(l) > 0
}

func (l StringList) GetAny() (result string) {
	if len(l) > 0 {
		result = l[rand.Intn(len(l)-1)]
	}
	return
}

func (l StringList) Contains(element string) bool {
	if l != nil {
		for _, e1 := range l {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l StringList) ContainsAll(elements []string) bool {
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

func (l StringList) Each(procedure func(element string)) StringList {
	if l != nil {
		for _, e := range l {
			procedure(e)
		}
	}
	return l
}

func (l StringList) Select(predicate func(element string) bool) (newList StringList) {
	newList = StringList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l StringList) Reject(predicate func(element string) bool) (newList StringList) {
	newList = StringList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) == false {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l StringList) Partition(predicate func(element string) bool) (accepted, rejected StringList) {
	accepted, rejected = StringList{}, StringList{}
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

func (l StringList) Detect(predicate func(element string) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l StringList) Count(predicate func(element string) bool) (count int) {
	for _, e := range l {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l StringList) AnySatisfy(predicate func(element string) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l StringList) AllSatisfy(predicate func(element string) bool) bool {
	for _, e := range l {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l StringList) NoneSatisfy(predicate func(element string) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l StringList) Sorted(compare func(i, j string) bool) (newList StringList) {
	newList = StringList{}
	if l != nil {
		newList = append(StringList{}, l...)
		sort.Slice(newList,
			func(i, j int) bool {
				return compare(newList[i], newList[j])
			})
	}
	return
}

// Mutable functions
func StringList_NewEmpty() (newList StringList) {
	newList = StringList{}
	return
}

func (l StringList) With(element string) StringList {
	l = append(l, element)
	return l
}

func (l StringList) Without(element string) StringList {
	for n, e := range l {
		if e == element {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

func (l StringList) WithAll(elements []string) StringList {
	l = append(l, elements...)
	return l
}

func (l StringList) WithoutAll(elements []string) StringList {
	for _, element := range elements {
		for n, e := range l {
			if e == element {
				l = append(l[:n], l[n+1:]...)
			}
		}
	}
	return l
}

func (l StringList) RemoveIf(predicate func(element string) bool) StringList {
	for n, e := range l {
		if predicate(e) {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

func (l StringList) AddAll(elements []string) StringList {
	return l.WithAll(elements)
}

func (l StringList) RetainAll(elements []string) StringList {
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
