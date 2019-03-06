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

func FloatImmutableList_NewWith(element float64) (newList FloatImmutableList) {
	newList = append(FloatImmutableList{}, element)
	return
}

func FloatImmutableList_NewWithAll(elements []float64) (newList FloatImmutableList) {
	newList = append(FloatImmutableList{}, elements...)
	return
}

func (l FloatImmutableList) NewWithout(element float64) (newList FloatImmutableList) {
	newList = FloatImmutableList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l FloatImmutableList) NewWithoutAll(elements []float64) (newList FloatImmutableList) {
	newList = FloatImmutableList{}
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

func (l FloatImmutableList) Remove(element float64) (newList FloatImmutableList) {
	newList = FloatImmutableList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return newList
}

func (l FloatImmutableList) RemoveAll(elements []float64) (newList FloatImmutableList) {
	newList = FloatImmutableList{}
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

func (l FloatImmutableList) Size() (size int) {
	size = len(l)
	return
}

func (l FloatImmutableList) IsEmpty() bool {
	return len(l) == 0
}

func (l FloatImmutableList) NotEmpty() bool {
	return len(l) > 0
}

func (l FloatImmutableList) GetAny() (result float64) {
	if len(l) > 0 {
		result = l[rand.Intn(len(l)-1)]
	}
	return
}

func (l FloatImmutableList) Contains(element float64) bool {
	if l != nil {
		for _, e1 := range l {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l FloatImmutableList) ContainsAll(elements []float64) bool {
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

func (l FloatImmutableList) Each(procedure func(element float64)) FloatImmutableList {
	if l != nil {
		for _, e := range l {
			procedure(e)
		}
	}
	return l
}

func (l FloatImmutableList) Select(predicate func(element float64) bool) (newList FloatImmutableList) {
	newList = FloatImmutableList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l FloatImmutableList) Reject(predicate func(element float64) bool) (newList FloatImmutableList) {
	newList = FloatImmutableList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) == false {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l FloatImmutableList) Partition(predicate func(element float64) bool) (accepted, rejected FloatImmutableList) {
	accepted, rejected = FloatImmutableList{}, FloatImmutableList{}
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

func (l FloatImmutableList) Detect(predicate func(element float64) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l FloatImmutableList) Count(predicate func(element float64) bool) (count int) {
	for _, e := range l {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l FloatImmutableList) AnySatisfy(predicate func(element float64) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l FloatImmutableList) AllSatisfy(predicate func(element float64) bool) bool {
	for _, e := range l {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l FloatImmutableList) NoneSatisfy(predicate func(element float64) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l FloatImmutableList) Sorted(compare func(i, j float64) bool) (newList FloatImmutableList) {
	newList = FloatImmutableList{}
	if l != nil {
		newList = append(FloatImmutableList{}, l...)
		sort.Slice(newList,
			func(i, j int) bool {
				return compare(newList[i], newList[j])
			})
	}
	return
}
