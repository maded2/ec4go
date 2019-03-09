/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

func IntImmutableMap_NewWithKeyValue(key int, value int) IntImmutableMap {
	return IntImmutableMap{
		key: value,
	}
}

func IntImmutableMap_NewWithAllKeyValues(values map[int]int) (newMap IntImmutableMap) {
	newMap = IntImmutableMap{}
	for k, v := range values {
		newMap[k] = v
	}
	return
}

func (m IntImmutableMap) NewWithoutKey(key int) (newMap IntImmutableMap) {
	newMap = IntImmutableMap{}
	for k, v := range m {
		if key != k {
			newMap[k] = v
		}
	}
	return
}

func (m IntImmutableMap) NewWithoutAllKeys(keys []int) (newMap IntImmutableMap) {
	newMap = IntImmutableMap{}
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

func (m IntImmutableMap) Select(predicate func(key int, value int) bool) (newMap IntImmutableMap) {
	newMap = IntImmutableMap{}
	for k, v := range m {
		if predicate(k, v) {
			newMap[k] = v
		}
	}
	return
}

func (m IntImmutableMap) Reject(predicate func(key int, value int) bool) (newMap IntImmutableMap) {
	newMap = IntImmutableMap{}
	for k, v := range m {
		if predicate(k, v) == false {
			newMap[k] = v
		}
	}
	return
}

func (m IntImmutableMap) Tap(procedure func(value int)) {
	for _, v := range m {
		procedure(v)
	}
}

func (m IntImmutableMap) Partition(predicate func(key int, value int) bool) (selected, rejected []int) {
	selected = []int{}
	rejected = []int{}
	for k, v := range m {
		if predicate(k, v) {
			selected = append(selected, v)
		} else {
			rejected = append(rejected, v)
		}
	}
	return
}
