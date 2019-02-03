/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import "math/rand"

type StringList struct {
	list []*string
}

// Immutable functions

func (l *StringList) NewWith(element *string) (newList *StringList) {
	newList = &StringList{list: append(l.list, element)}
	return
}

func (l *StringList) NewWithout(element *string) (newList *StringList) {
	newList = &StringList{}
	for _, e := range l.list {
		if e != element {
			newList.list = append(newList.list, e)
		}
	}
	return newList
}

func (l *StringList) NewWithAll(elements []*string) (newList *StringList) {
	newList = &StringList{}
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

func (l *StringList) NewWithoutAll(elements []*string) (newList *StringList) {
	newList = &StringList{}
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

func (l *StringList) Size() int {
	return len(l.list)
}

func (l *StringList) IsEmpty() bool {
	return len(l.list) == 0
}

func (l *StringList) NotEmpty() bool {
	return len(l.list) > 0
}

func (l *StringList) GetAny() (result *string) {
	if len(l.list) > 0 {
		result = l.list[rand.Intn(len(l.list)-1)]
	}
	return
}

func (l *StringList) Contains(element *string) bool {
	for _, e1 := range l.list {
		if e1 == element {
			return true
		}
	}
	return false
}

func (l *StringList) ContainsAll(elements []*string) bool {
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

func (l *StringList) Each(procedure func(element *string)) *StringList {
	for _, e := range l.list {
		procedure(e)
	}
	return l
}

func (l *StringList) Select(predicate func(element *string) bool) (newList *StringList) {
	newList = &StringList{}
	for _, e := range l.list {
		if predicate(e) {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *StringList) Reject(predicate func(element *string) bool) (newList *StringList) {
	newList = &StringList{}
	for _, e := range l.list {
		if predicate(e) == false {
			newList.list = append(newList.list, e)
		}
	}
	return
}

func (l *StringList) Partition(predicate func(element *string) bool) (accepted, rejected *StringList) {
	accepted, rejected = &StringList{}, &StringList{}
	for _, e := range l.list {
		if predicate(e) {
			accepted.list = append(accepted.list, e)
		} else {
			rejected.list = append(rejected.list, e)
		}
	}
	return
}

func (l *StringList) Detect(predicate func(element *string) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *StringList) Count(predicate func(element *string) bool) (count int) {
	for _, e := range l.list {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l *StringList) AnySatisfy(predicate func(element *string) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l *StringList) AllSatisfy(predicate func(element *string) bool) bool {
	for _, e := range l.list {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l *StringList) NoneSatisfy(predicate func(element *string) bool) bool {
	for _, e := range l.list {
		if predicate(e) {
			return false
		}
	}
	return true
}

// Mutable functions
func (l *StringList) NewEmpty() (newList *StringList) {
	newList = &StringList{}
	return l
}

func (l *StringList) With(element *string) *StringList {
	l.list = append(l.list, element)
	return l
}

func (l *StringList) Without(element *string) *StringList {
	for n, e := range l.list {
		if e == element {
			l.list = append(l.list[:n], l.list[n+1:]...)
		}
	}
	return l
}

func (l *StringList) WithAll(elements []*string) *StringList {
	l.list = append(l.list, elements...)
	return l
}

func (l *StringList) WithoutAll(elements []*string) *StringList {
	for __, element := range elements {
		for n, e := range l.list {
			if e == element {
				l.list = append(l.list[:n], l.list[n+1:]...)
			}
		}
	}
	return l
}

func (l *StringList) RemoveIf(predicate func(element *string) bool) (l *StringList) {
	for n, e := range l.list {
		if predicate(e) {
			l.list = append(l.list[:n], l.list[n+1:]...)
		}
	}
	return l
}

func (l *StringList) AddAll(elements []*string) *StringList {
	return l.WithAll(elements)
}

func (l *StringList) RemoveAll(elements []*string) *StringList {
	return l.WithoutAll(elements)
}

func (l *StringList) RetainAll(elements []*string) *StringList {
	for n, e := range l.list {
		retain := true
		for __, element := range elements {
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
