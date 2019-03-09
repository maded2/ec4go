/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

func FloatImmutableMap_NewWithKeyValue(key float64, value float64) FloatImmutableMap {
	return FloatImmutableMap{
		key: value,
	}
}

func FloatImmutableMap_NewWithAllKeyValues(values map[float64]float64) (newMap FloatImmutableMap) {
	newMap = FloatImmutableMap{}
	for k, v := range values {
		newMap[k] = v
	}
	return
}

func (m FloatImmutableMap) NewWithoutKey(key float64) (newMap FloatImmutableMap) {
	newMap = FloatImmutableMap{}
	for k, v := range m {
		if key != k {
			newMap[k] = v
		}
	}
	return
}

func (m FloatImmutableMap) NewWithoutAllKeys(keys []float64) (newMap FloatImmutableMap) {
	newMap = FloatImmutableMap{}
skip:
	for k, v := range m {
		for _, k2 := range keys {
			if k2 == k {
				continue skip
			}
		}
		newMap[k] = v
	}
	return
}

func (m FloatImmutableMap) Select(predicate func(key float64, value float64) bool) (newMap FloatImmutableMap) {
	newMap = FloatImmutableMap{}
	for k, v := range m {
		if predicate(k, v) {
			newMap[k] = v
		}
	}
	return
}

func (m FloatImmutableMap) Reject(predicate func(key float64, value float64) bool) (newMap FloatImmutableMap) {
	newMap = FloatImmutableMap{}
	for k, v := range m {
		if predicate(k, v) == false {
			newMap[k] = v
		}
	}
	return
}

func (m FloatImmutableMap) Tap(procedure func(value float64)) {
	for _, v := range m {
		procedure(v)
	}
}

func (m FloatImmutableMap) Partition(predicate func(key float64, value float64) bool) (selected, rejected []float64) {
	selected = []float64{}
	rejected = []float64{}
	for k, v := range m {
		if predicate(k, v) {
			selected = append(selected, v)
		} else {
			rejected = append(rejected, v)
		}
	}
	return
}
