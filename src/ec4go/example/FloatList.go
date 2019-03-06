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

func FloatList_NewWith(element float64) (newList FloatList) {
	newList = append(FloatList{}, element)
	return
}

func FloatList_NewWithAll(elements []float64) (newList FloatList) {
	newList = append(FloatList{}, elements...)
	return
}

func (l FloatList) NewWithout(element float64) (newList FloatList) {
	newList = FloatList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l FloatList) NewWithoutAll(elements []float64) (newList FloatList) {
	newList = FloatList{}
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

func (l FloatList) Remove(element float64) (newList FloatList) {
	newList = FloatList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return newList
}

func (l FloatList) RemoveAll(elements []float64) (newList FloatList) {
	newList = FloatList{}
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

func (l FloatList) Size() (size int) {
	size = len(l)
	return
}

func (l FloatList) IsEmpty() bool {
	return len(l) == 0
}

func (l FloatList) NotEmpty() bool {
	return len(l) > 0
}

func (l FloatList) GetAny() (result float64) {
	if len(l) > 0 {
		result = l[rand.Intn(len(l)-1)]
	}
	return
}

func (l FloatList) Contains(element float64) bool {
	if l != nil {
		for _, e1 := range l {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l FloatList) ContainsAll(elements []float64) bool {
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

func (l FloatList) Each(procedure func(element float64)) FloatList {
	if l != nil {
		for _, e := range l {
			procedure(e)
		}
	}
	return l
}

func (l FloatList) Select(predicate func(element float64) bool) (newList FloatList) {
	newList = FloatList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l FloatList) Reject(predicate func(element float64) bool) (newList FloatList) {
	newList = FloatList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) == false {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l FloatList) Partition(predicate func(element float64) bool) (accepted, rejected FloatList) {
	accepted, rejected = FloatList{}, FloatList{}
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

func (l FloatList) Detect(predicate func(element float64) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l FloatList) Count(predicate func(element float64) bool) (count int) {
	for _, e := range l {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l FloatList) AnySatisfy(predicate func(element float64) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l FloatList) AllSatisfy(predicate func(element float64) bool) bool {
	for _, e := range l {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l FloatList) NoneSatisfy(predicate func(element float64) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l FloatList) Sorted(compare func(i, j float64) bool) (newList FloatList) {
	newList = FloatList{}
	if l != nil {
		newList = append(FloatList{}, l...)
		sort.Slice(newList,
			func(i, j int) bool {
				return compare(newList[i], newList[j])
			})
	}
	return
}

// Mutable functions
func FloatList_NewEmpty() (newList FloatList) {
	newList = FloatList{}
	return
}

func (l FloatList) With(element float64) FloatList {
	l = append(l, element)
	return l
}

func (l FloatList) Without(element float64) FloatList {
	for n, e := range l {
		if e == element {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

func (l FloatList) WithAll(elements []float64) FloatList {
	l = append(l, elements...)
	return l
}

func (l FloatList) WithoutAll(elements []float64) FloatList {
	for _, element := range elements {
		for n, e := range l {
			if e == element {
				l = append(l[:n], l[n+1:]...)
			}
		}
	}
	return l
}

func (l FloatList) RemoveIf(predicate func(element float64) bool) FloatList {
	for n, e := range l {
		if predicate(e) {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

func (l FloatList) AddAll(elements []float64) FloatList {
	return l.WithAll(elements)
}

func (l FloatList) RetainAll(elements []float64) FloatList {
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
