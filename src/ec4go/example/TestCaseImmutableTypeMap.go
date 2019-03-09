/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */
package example

func TestCaseImmutableTypeMap_NewWithKeyValue(key int, value *TestCaseType) TestCaseImmutableTypeMap {
	return TestCaseImmutableTypeMap{
		key: value,
	}
}

func TestCaseImmutableTypeMap_NewWithAllKeyValues(values map[int]*TestCaseType) (newMap TestCaseImmutableTypeMap) {
	newMap = TestCaseImmutableTypeMap{}
	for k, v := range values {
		newMap[k] = v
	}
	return
}

func (m TestCaseImmutableTypeMap) NewWithoutKey(key int) (newMap TestCaseImmutableTypeMap) {
	newMap = TestCaseImmutableTypeMap{}
	for k, v := range m {
		if key != k {
			newMap[k] = v
		}
	}
	return
}

func (m TestCaseImmutableTypeMap) NewWithoutAllKeys(keys []int) (newMap TestCaseImmutableTypeMap) {
	newMap = TestCaseImmutableTypeMap{}
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

func (m TestCaseImmutableTypeMap) Select(predicate func(key int, value *TestCaseType) bool) (newMap TestCaseImmutableTypeMap) {
	newMap = TestCaseImmutableTypeMap{}
	for k, v := range m {
		if predicate(k, v) {
			newMap[k] = v
		}
	}
	return
}

func (m TestCaseImmutableTypeMap) Reject(predicate func(key int, value *TestCaseType) bool) (newMap TestCaseImmutableTypeMap) {
	newMap = TestCaseImmutableTypeMap{}
	for k, v := range m {
		if predicate(k, v) == false {
			newMap[k] = v
		}
	}
	return
}

func (m TestCaseImmutableTypeMap) Tap(procedure func(value *TestCaseType)) {
	for _, v := range m {
		procedure(v)
	}
}

func (m TestCaseImmutableTypeMap) Partition(predicate func(key int, value *TestCaseType) bool) (selected, rejected []*TestCaseType) {
	selected = []*TestCaseType{}
	rejected = []*TestCaseType{}
	for k, v := range m {
		if predicate(k, v) {
			selected = append(selected, v)
		} else {
			rejected = append(rejected, v)
		}
	}
	return
}
