/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

func BoolImmutableMap_NewWithKeyValue(key bool, value bool) BoolImmutableMap {
	return BoolImmutableMap{
		key: value,
	}
}

func BoolImmutableMap_NewWithAllKeyValues(values map[bool]bool) (newMap BoolImmutableMap) {
	newMap = BoolImmutableMap{}
	for k, v := range values {
		newMap[k] = v
	}
	return
}

func (m BoolImmutableMap) NewWithoutKey(key bool) (newMap BoolImmutableMap) {
	newMap = BoolImmutableMap{}
	for k, v := range m {
		if key != k {
			newMap[k] = v
		}
	}
	return
}

func (m BoolImmutableMap) NewWithoutAllKeys(keys []bool) (newMap BoolImmutableMap) {
	newMap = BoolImmutableMap{}
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

func (m BoolImmutableMap) Select(predicate func(key bool, value bool) bool) (newMap BoolImmutableMap) {
	newMap = BoolImmutableMap{}
	for k, v := range m {
		if predicate(k, v) {
			newMap[k] = v
		}
	}
	return
}

func (m BoolImmutableMap) Reject(predicate func(key bool, value bool) bool) (newMap BoolImmutableMap) {
	newMap = BoolImmutableMap{}
	for k, v := range m {
		if predicate(k, v) == false {
			newMap[k] = v
		}
	}
	return
}

func (m BoolImmutableMap) Tap(procedure func(value bool)) {
	for _, v := range m {
		procedure(v)
	}
}

func (m BoolImmutableMap) Partition(predicate func(key bool, value bool) bool) (selected, rejected []bool) {
	selected = []bool{}
	rejected = []bool{}
	for k, v := range m {
		if predicate(k, v) {
			selected = append(selected, v)
		} else {
			rejected = append(rejected, v)
		}
	}
	return
}
