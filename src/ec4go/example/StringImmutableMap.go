/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

func StringImmutableMap_NewWithKeyValue(key string, value string) StringImmutableMap {
	return StringImmutableMap{
		key: value,
	}
}

func StringImmutableMap_NewWithAllKeyValues(values map[string]string) (newMap StringImmutableMap) {
	newMap = StringImmutableMap{}
	for k, v := range values {
		newMap[k] = v
	}
	return
}

func (m StringImmutableMap) NewWithoutKey(key string) (newMap StringImmutableMap) {
	newMap = StringImmutableMap{}
	for k, v := range m {
		if key != k {
			newMap[k] = v
		}
	}
	return
}

func (m StringImmutableMap) NewWithoutAllKeys(keys []string) (newMap StringImmutableMap) {
	newMap = StringImmutableMap{}
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

func (m StringImmutableMap) Select(predicate func(key string, value string) bool) (newMap StringImmutableMap) {
	newMap = StringImmutableMap{}
	for k, v := range m {
		if predicate(k, v) {
			newMap[k] = v
		}
	}
	return
}

func (m StringImmutableMap) Reject(predicate func(key string, value string) bool) (newMap StringImmutableMap) {
	newMap = StringImmutableMap{}
	for k, v := range m {
		if predicate(k, v) == false {
			newMap[k] = v
		}
	}
	return
}

func (m StringImmutableMap) Tap(procedure func(value string)) {
	for _, v := range m {
		procedure(v)
	}
}

func (m StringImmutableMap) Partition(predicate func(key string, value string) bool) (selected, rejected []string) {
	selected = []string{}
	rejected = []string{}
	for k, v := range m {
		if predicate(k, v) {
			selected = append(selected, v)
		} else {
			rejected = append(rejected, v)
		}
	}
	return
}
