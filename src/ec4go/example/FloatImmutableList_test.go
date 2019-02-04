/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

import (
	"testing"
)

func TestSort(t *testing.T) {
	l := (*FloatImmutableList)(nil).NewWithAll([]float64{5.0, 4.0, 3.0, 2.0, 1.0})
	newList := l.Sorted(func(i, j float64) bool {
		return i < j
	})

	v := 0.0
	newList.Each(func(element float64) {
		if element < v {
			t.Fail()
		} else {
			v = element
		}
	})
}
