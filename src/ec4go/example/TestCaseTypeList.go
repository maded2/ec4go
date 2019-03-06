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

func TestCaseTypeList_NewWith(element *TestCaseType) (newList TestCaseTypeList) {
	newList = append(TestCaseTypeList{}, element)
	return
}

func TestCaseTypeList_NewWithAll(elements []*TestCaseType) (newList TestCaseTypeList) {
	newList = append(TestCaseTypeList{}, elements...)
	return
}

func (l TestCaseTypeList) NewWithout(element *TestCaseType) (newList TestCaseTypeList) {
	newList = TestCaseTypeList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l TestCaseTypeList) NewWithoutAll(elements []*TestCaseType) (newList TestCaseTypeList) {
	newList = TestCaseTypeList{}
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

func (l TestCaseTypeList) Remove(element *TestCaseType) (newList TestCaseTypeList) {
	newList = TestCaseTypeList{}
	if l != nil {
		for _, e := range l {
			if e != element {
				newList = append(newList, e)
			}
		}
	}
	return newList
}

func (l TestCaseTypeList) RemoveAll(elements []*TestCaseType) (newList TestCaseTypeList) {
	newList = TestCaseTypeList{}
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

func (l TestCaseTypeList) Size() (size int) {
	size = len(l)
	return
}

func (l TestCaseTypeList) IsEmpty() bool {
	return len(l) == 0
}

func (l TestCaseTypeList) NotEmpty() bool {
	return len(l) > 0
}

func (l TestCaseTypeList) GetAny() (result *TestCaseType) {
	if len(l) > 0 {
		result = l[rand.Intn(len(l)-1)]
	}
	return
}

func (l TestCaseTypeList) Contains(element *TestCaseType) bool {
	if l != nil {
		for _, e1 := range l {
			if e1 == element {
				return true
			}
		}
	}
	return false
}

func (l TestCaseTypeList) ContainsAll(elements []*TestCaseType) bool {
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

func (l TestCaseTypeList) Each(procedure func(element *TestCaseType)) TestCaseTypeList {
	if l != nil {
		for _, e := range l {
			procedure(e)
		}
	}
	return l
}

func (l TestCaseTypeList) Select(predicate func(element *TestCaseType) bool) (newList TestCaseTypeList) {
	newList = TestCaseTypeList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l TestCaseTypeList) Reject(predicate func(element *TestCaseType) bool) (newList TestCaseTypeList) {
	newList = TestCaseTypeList{}
	if l != nil {
		for _, e := range l {
			if predicate(e) == false {
				newList = append(newList, e)
			}
		}
	}
	return
}

func (l TestCaseTypeList) Partition(predicate func(element *TestCaseType) bool) (accepted, rejected TestCaseTypeList) {
	accepted, rejected = TestCaseTypeList{}, TestCaseTypeList{}
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

func (l TestCaseTypeList) Detect(predicate func(element *TestCaseType) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l TestCaseTypeList) Count(predicate func(element *TestCaseType) bool) (count int) {
	for _, e := range l {
		if predicate(e) {
			count++
		}
	}
	return
}

func (l TestCaseTypeList) AnySatisfy(predicate func(element *TestCaseType) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return true
		}
	}
	return false
}

func (l TestCaseTypeList) AllSatisfy(predicate func(element *TestCaseType) bool) bool {
	for _, e := range l {
		if predicate(e) == false {
			return false
		}
	}
	return true
}

func (l TestCaseTypeList) NoneSatisfy(predicate func(element *TestCaseType) bool) bool {
	for _, e := range l {
		if predicate(e) {
			return false
		}
	}
	return true
}

func (l TestCaseTypeList) Sorted(compare func(i, j *TestCaseType) bool) (newList TestCaseTypeList) {
	newList = TestCaseTypeList{}
	if l != nil {
		newList = append(TestCaseTypeList{}, l...)
		sort.Slice(newList,
			func(i, j int) bool {
				return compare(newList[i], newList[j])
			})
	}
	return
}

// Mutable functions
func TestCaseTypeList_NewEmpty() (newList TestCaseTypeList) {
	newList = TestCaseTypeList{}
	return
}

func (l TestCaseTypeList) With(element *TestCaseType) TestCaseTypeList {
	l = append(l, element)
	return l
}

func (l TestCaseTypeList) Without(element *TestCaseType) TestCaseTypeList {
	for n, e := range l {
		if e == element {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

func (l TestCaseTypeList) WithAll(elements []*TestCaseType) TestCaseTypeList {
	l = append(l, elements...)
	return l
}

func (l TestCaseTypeList) WithoutAll(elements []*TestCaseType) TestCaseTypeList {
	for _, element := range elements {
		for n, e := range l {
			if e == element {
				l = append(l[:n], l[n+1:]...)
			}
		}
	}
	return l
}

func (l TestCaseTypeList) RemoveIf(predicate func(element *TestCaseType) bool) TestCaseTypeList {
	for n, e := range l {
		if predicate(e) {
			l = append(l[:n], l[n+1:]...)
		}
	}
	return l
}

func (l TestCaseTypeList) AddAll(elements []*TestCaseType) TestCaseTypeList {
	return l.WithAll(elements)
}

func (l TestCaseTypeList) RetainAll(elements []*TestCaseType) TestCaseTypeList {
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
