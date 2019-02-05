/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"testing"
)

func TestStringListCount(t *testing.T) {
	l := (*StringList)(nil).NewWithAll(sample_string)

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

func TestStringListSort(t *testing.T) {
	l := (*StringList)(nil).NewWithAll(sample_string)
	newList := l.Sorted(func(i, j string) bool {
		return i < j
	})

	var v string
	newList.Each(func(element string) {
		if element < v {
			t.Fail()
		} else {
			v = element
		}
	})
}
