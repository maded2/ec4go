/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"math/rand"
	"sort"
)

type FloatImmutableList struct {
	list []float64
}

// Immutable functions

func (l *FloatImmutableList) NewWith(element float64) (newList *FloatImmutableList) {
	newList = &FloatImmutableList{list: append(l.list, element)}
	return
}

func (l *FloatImmutableList) NewWithout(element float64) (newList *FloatImmutableList) {
	newList = &FloatImmutableList{}
	for _, e := range l.list {
		if e != element {
			newList.list = append(newList.list, e)
		}
	}
	return newList
}

func (l *FloatImmutableList) NewWithAll(elements []float64) (newList *FloatImmutableList) {
	newList = &FloatImmutableList{}
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

func (l *FloatImmutableList) NewWithoutAll(elements []float64) (newList *FloatImmutableList) {
	newList = &FloatImmutableList{}
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

func (l *FloatImmutableList) Size() int {
	return len(l.list)
}

func (l *FloatImmutableList) IsEmpty() bool {
	return len(l.list) == 0
}

func (l *FloatImmutableList) NotEmpty() bool {
	return len(l.list) > 0
}

func (l *FloatImmutableList) GetAny() (result float64) {
	if len(l.list) > 0 {
		result = l.list[rand.Intn(len(l.list)-1)]
	}
	return
}

func (l *FloatImmutableList) Contains(element float64) bool {
	for _, e1 := range l.list {
		if e1 == element {
			return true
		}
	}
	return false
}

func (l *FloatImmutableList) ContainsAll(elements []float64) bool {
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

func (l *FloatImmutableList) Each(procedure func(element float64)) *FloatImmutableList {
	for _, e := range l.list {
		procedure(e)
	}
	return l
}

func (l *FloatImmutableList) Select(predicate func(element float64) bool) (newList *FloatImmutableList) {
	newList = &FloatImmutableList{}
	for _, e := range l.list {
		if predicate(e) {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *FloatImmutableList) Reject(predicate func(element float64) bool) (newList *FloatImmutableList) {
	newList = &FloatImmutableList{}
	for _, e := range l.list {
		if predicate(e) == false {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *FloatImmutableList) Partition(predicate func(element float64) bool) (accepted, rejected *FloatImmutableList) {
	accepted, rejected = &FloatImmutableList{}, &FloatImmutableList{}
	for _, e := range l.list {
		if predicate(e) {
			accepted.list = append(accepted.list, e)
		} else {
			rejected.list = append(rejected.list, e)
		}
	}
	return
}

func (l *FloatImmutableList) Detect(predicate func(element float64) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *FloatImmutableList) Count(predicate func(element float64) bool) (count int) {
	for _, e := range l.list {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l *FloatImmutableList) AnySatisfy(predicate func(element float64) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *FloatImmutableList) AllSatisfy(predicate func(element float64) bool) bool {
	for _, e := range l.list {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l *FloatImmutableList) NoneSatisfy(predicate func(element float64) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l *FloatImmutableList) Sorted(compare func(i, j float64) bool) (newList *FloatImmutableList) {
	newList = &FloatImmutableList{list: append([]float64{}, l.list...)}
	sort.Slice(newList.list,
		func(i, j int) bool {
			return compare(newList.list[i], newList.list[j])
		})
	return
}
