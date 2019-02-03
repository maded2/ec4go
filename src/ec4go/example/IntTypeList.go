/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/
package example

import "math/rand"

type IntTypeList struct {
	list []int
}

// Immutable functions

func (l *IntTypeList) NewWith(element int) (newList *IntTypeList) {
	newList = &IntTypeList{list: append(l.list, element)}
	return
}

func (l *IntTypeList) NewWithout(element int) (newList *IntTypeList) {
	newList = &IntTypeList{}
	for _, e := range l.list {
		if e != element {
			newList.list = append(newList.list, e)
		}
	}
	return newList
}

func (l *IntTypeList) NewWithAll(elements []int) (newList *IntTypeList) {
	newList = &IntTypeList{}
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

func (l *IntTypeList) NewWithoutAll(elements []int) (newList *IntTypeList) {
	newList = &IntTypeList{}
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

func (l *IntTypeList) Size() int {
	return len(l.list)
}

func (l *IntTypeList) IsEmpty() bool {
	return len(l.list) == 0
}

func (l *IntTypeList) NotEmpty() bool {
	return len(l.list) > 0
}

func (l *IntTypeList) GetAny() (result int) {
	if len(l.list) > 0 {
		result = l.list[rand.Intn(len(l.list)-1)]
	}
	return
}

func (l *IntTypeList) Contains(element int) bool {
	for _, e1 := range l.list {
		if e1 == element {
			return true
		}
	}
	return false
}

func (l *IntTypeList) ContainsAll(elements []int) bool {
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

func (l *IntTypeList) Each(procedure func(element int)) *IntTypeList {
	for _, e := range l.list {
		procedure(e)
	}
	return l
}

func (l *IntTypeList) Select(predicate func(element int) bool) (newList *IntTypeList) {
	newList = &IntTypeList{}
	for _, e := range l.list {
		if predicate(e) {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *IntTypeList) Reject(predicate func(element int) bool) (newList *IntTypeList) {
	newList = &IntTypeList{}
	for _, e := range l.list {
		if predicate(e) == false {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *IntTypeList) Partition(predicate func(element int) bool) (accepted, rejected *IntTypeList) {
	accepted, rejected = &IntTypeList{}, &IntTypeList{}
	for _, e := range l.list {
		if predicate(e) {
			accepted.list = append(accepted.list, e)
		} else {
			rejected.list = append(rejected.list, e)
		}
	}
	return
}

func (l *IntTypeList) Detect(predicate func(element int) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *IntTypeList) Count(predicate func(element int) bool) (count int) {
	for _, e := range l.list {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l *IntTypeList) AnySatisfy(predicate func(element int) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *IntTypeList) AllSatisfy(predicate func(element int) bool) bool {
	for _, e := range l.list {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l *IntTypeList) NoneSatisfy(predicate func(element int) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return false
		}
	}
	return true
}

