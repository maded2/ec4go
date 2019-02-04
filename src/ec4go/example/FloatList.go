/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"math/rand"
	"sort"
)

type FloatList struct {
	list []float64
}

// Immutable functions

func (l *FloatList) NewWith(element float64) (newList *FloatList) {
	newList = &FloatList{list: append(l.list, element)}
	return
}

func (l *FloatList) NewWithout(element float64) (newList *FloatList) {
	newList = &FloatList{}
	for _, e := range l.list {
		if e != element {
			newList.list = append(newList.list, e)
		}
	}
	return newList
}

func (l *FloatList) NewWithAll(elements []float64) (newList *FloatList) {
	newList = &FloatList{}
	for _, e1 := range l.list {
		for _, e2 := range elements {
			if e1 != e2 {
				newList.list = append(newList.list, e1)
				break
			}
		}
	}
	return newList
}

func (l *FloatList) NewWithoutAll(elements []float64) (newList *FloatList) {
	newList = &FloatList{}
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
	return newList
}

func (l *FloatList) Size() int {
	return len(l.list)
}

func (l *FloatList) IsEmpty() bool {
	return len(l.list) == 0
}

func (l *FloatList) NotEmpty() bool {
	return len(l.list) > 0
}

func (l *FloatList) GetAny() (result float64) {
	if len(l.list) > 0 {
		result = l.list[rand.Intn(len(l.list)-1)]
	}
	return
}

func (l *FloatList) Contains(element float64) bool {
	for _, e1 := range l.list {
		if e1 == element {
			return true
		}
	}
	return false
}

func (l *FloatList) ContainsAll(elements []float64) bool {
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
	return false
}

func (l *FloatList) Each(procedure func(element float64)) *FloatList {
	for _, e := range l.list {
		procedure(e)
	}
	return l
}

func (l *FloatList) Select(predicate func(element float64) bool) (newList *FloatList) {
	newList = &FloatList{}
	for _, e := range l.list {
		if predicate(e) {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *FloatList) Reject(predicate func(element float64) bool) (newList *FloatList) {
	newList = &FloatList{}
	for _, e := range l.list {
		if predicate(e) == false {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *FloatList) Partition(predicate func(element float64) bool) (accepted, rejected *FloatList) {
	accepted, rejected = &FloatList{}, &FloatList{}
	for _, e := range l.list {
		if predicate(e) {
			accepted.list = append(accepted.list, e)
		} else {
			rejected.list = append(rejected.list, e)
		}
	}
	return
}

func (l *FloatList) Detect(predicate func(element float64) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *FloatList) Count(predicate func(element float64) bool) (count int) {
	for _, e := range l.list {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l *FloatList) AnySatisfy(predicate func(element float64) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *FloatList) AllSatisfy(predicate func(element float64) bool) bool {
	for _, e := range l.list {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l *FloatList) NoneSatisfy(predicate func(element float64) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l *FloatList) Sorted(compare func(i, j float64) bool) (newList *FloatList) {
	newList = &FloatList{list: append([]float64{}, l.list...)}
	sort.Slice(newList.list,
		func(i, j int) bool {
			return compare(newList.list[i], newList.list[j])
		})
	return
}

// Mutable functions
func (l *FloatList) NewEmpty() (newList *FloatList) {
	newList = &FloatList{}
	return l
}

func (l *FloatList) With(element float64) *FloatList {
	l.list = append(l.list, element)
	return l
}

func (l *FloatList) Without(element float64) *FloatList {
	for n, e := range l.list {
		if e == element {
			l.list = append(l.list[:n], l.list[n+1:]...)
		}
	}
	return l
}

func (l *FloatList) WithAll(elements []float64) *FloatList {
	l.list = append(l.list, elements...)
	return l
}

func (l *FloatList) WithoutAll(elements []float64) *FloatList {
	for _, element := range elements {
		for n, e := range l.list {
			if e == element {
				l.list = append(l.list[:n], l.list[n+1:]...)
			}
		}
	}
	return l
}

func (l *FloatList) RemoveIf(predicate func(element float64) bool) *FloatList {
	for n, e := range l.list {
		if predicate(e) {
			l.list = append(l.list[:n], l.list[n+1:]...)
		}
	}
	return l
}

func (l *FloatList) AddAll(elements []float64) *FloatList {
	return l.WithAll(elements)
}

func (l *FloatList) RemoveAll(elements []float64) *FloatList {
	return l.WithoutAll(elements)
}

func (l *FloatList) RetainAll(elements []float64) *FloatList {
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
