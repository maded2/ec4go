/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"testing"
)

func TestIntImmutableListCount(t *testing.T) {
	l := (*IntImmutableList)(nil).NewWithAll(sample_int)

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

func TestIntImmutableListSort(t *testing.T) {
	l := (*IntImmutableList)(nil).NewWithAll(sample_int)
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
