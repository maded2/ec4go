/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"testing"
)

func TestIntListCount(t *testing.T) {
	l := (*IntList)(nil).NewWithAll([]int{5.0, 4.0, 3.0, 2.0, 1.0})

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

func TestIntListSort(t *testing.T) {
	l := (*IntList)(nil).NewWithAll([]int{5.0, 4.0, 3.0, 2.0, 1.0})
	newList := l.Sorted(func(i, j int) bool {
		return i < j
	})

	var v int
	newList.Each(func(element int) {
		if element < v {
			t.Fail()
		} else {
			v = element
		}
	})
}
