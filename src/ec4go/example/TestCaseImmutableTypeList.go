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

func TestCaseImmutableTypeList_NewWith(element *TestCaseType) (newList TestCaseImmutableTypeList) {
	newList = append(TestCaseImmutableTypeList{}, element)
	return
}

func TestCaseImmutableTypeList_NewWithAll(elements []*TestCaseType) (newList TestCaseImmutableTypeList) {
	newList = append(TestCaseImmutableTypeList{}, elements...)
	return
}

func (l TestCaseImmutableTypeList) NewWithout(element *TestCaseType) (newList TestCaseImmutableTypeList) {
	newList = TestCaseImmutableTypeList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l TestCaseImmutableTypeList) NewWithoutAll(elements []*TestCaseType) (newList TestCaseImmutableTypeList) {
	newList = TestCaseImmutableTypeList{}
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

func (l TestCaseImmutableTypeList) Remove(element *TestCaseType) (newList TestCaseImmutableTypeList) {
	newList = TestCaseImmutableTypeList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return newList
}

func (l TestCaseImmutableTypeList) RemoveAll(elements []*TestCaseType) (newList TestCaseImmutableTypeList) {
	newList = TestCaseImmutableTypeList{}
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

func (l TestCaseImmutableTypeList) Size() (size int) {
	size = len(l)
	return
}

func (l TestCaseImmutableTypeList) IsEmpty() bool {
	return len(l) == 0
}

func (l TestCaseImmutableTypeList) NotEmpty() bool {
	return len(l) > 0
}

func (l TestCaseImmutableTypeList) GetAny() (result *TestCaseType) {
	if len(l) > 0 {
		result = l[rand.Intn(len(l)-1)]
	}
	return
}

func (l TestCaseImmutableTypeList) Contains(element *TestCaseType) bool {
	if l != nil {
		for _, e1 := range l {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l TestCaseImmutableTypeList) ContainsAll(elements []*TestCaseType) bool {
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

func (l TestCaseImmutableTypeList) Each(procedure func(element *TestCaseType)) TestCaseImmutableTypeList {
	if l != nil {
		for _, e := range l {
			procedure(e)
		}
	}
	return l
}

func (l TestCaseImmutableTypeList) Select(predicate func(element *TestCaseType) bool) (newList TestCaseImmutableTypeList) {
	newList = TestCaseImmutableTypeList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l TestCaseImmutableTypeList) Reject(predicate func(element *TestCaseType) bool) (newList TestCaseImmutableTypeList) {
	newList = TestCaseImmutableTypeList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) == false {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l TestCaseImmutableTypeList) Partition(predicate func(element *TestCaseType) bool) (accepted, rejected TestCaseImmutableTypeList) {
	accepted, rejected = TestCaseImmutableTypeList{}, TestCaseImmutableTypeList{}
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

func (l TestCaseImmutableTypeList) Detect(predicate func(element *TestCaseType) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l TestCaseImmutableTypeList) Count(predicate func(element *TestCaseType) bool) (count int) {
	for _, e := range l {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l TestCaseImmutableTypeList) AnySatisfy(predicate func(element *TestCaseType) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l TestCaseImmutableTypeList) AllSatisfy(predicate func(element *TestCaseType) bool) bool {
	for _, e := range l {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l TestCaseImmutableTypeList) NoneSatisfy(predicate func(element *TestCaseType) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l TestCaseImmutableTypeList) Sorted(compare func(i, j *TestCaseType) bool) (newList TestCaseImmutableTypeList) {
	newList = TestCaseImmutableTypeList{}
	if l != nil {
		newList = append(TestCaseImmutableTypeList{}, l...)
		sort.Slice(newList,
			func(i, j int) bool {
				return compare(newList[i], newList[j])
			})
	}
	return
}
