/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"testing"
)

func TestTestCaseTypeListCount(t *testing.T) {
	l := TestCaseTypeList_NewWithAll(sample)

	if l.Size() != 5 {
		t.Fail()
	}

	if l.IsEmpty() {
		t.Fail()
	}

	if l.NotEmpty() == false {
		t.Fail()
	}
}

func TestTestCaseTypeListSort(t *testing.T) {
	l := TestCaseTypeList_NewWithAll(sample)
	newList := l.Sorted(func(i, j *TestCaseType) bool {
		return i.S < j.S
	})

	var v *TestCaseType
	newList.Each(func(element *TestCaseType) {
		if v != nil && element.S < v.S {
			t.Fail()
		} else {
			v = element
		}
	})
}
