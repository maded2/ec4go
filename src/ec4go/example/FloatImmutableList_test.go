/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"testing"
)

func TestFloatImmutableListCount(t *testing.T) {
	l := (*FloatImmutableList)(nil).NewWithAll(sample_float64)

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

func TestFloatImmutableListSort(t *testing.T) {
	l := (*FloatImmutableList)(nil).NewWithAll(sample_float64)
	newList := l.Sorted(func(i, j float64) bool {
		return i < j
	})

	var v float64
	newList.Each(func(element float64) {
		if element < v {
			t.Fail()
		} else {
			v = element
		}
	})
}
