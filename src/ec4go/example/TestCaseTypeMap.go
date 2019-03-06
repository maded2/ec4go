package example

type TestCaseTypeMap struct {
	values map[int]*TestCaseType
}

func NewWithKeyValue(key int, value *TestCaseType) *TestCaseTypeMap {
	return &TestCaseTypeMap{
		values: map[int]*TestCaseType{
			key: value,
		},
	}
}

func NewWithAllKeyValues(keyValues map[int]*TestCaseType) *TestCaseTypeMap {
	return &TestCaseTypeMap{
		values: keyValues,
	}
}

func (m *TestCaseTypeMap) NewWithoutKey(key int) (newMap *TestCaseTypeMap) {
	newMap = &TestCaseTypeMap{
		values: map[int]*TestCaseType{},
	}
	for k, v := range m.values {
		if key != k {
			newMap.values[k] = v
		}
	}
	return
}

func (m *TestCaseTypeMap) NewWithoutAllKeys(keys []int) (newMap *TestCaseTypeMap) {
	newMap = &TestCaseTypeMap{
		values: map[int]*TestCaseType{},
	}
skip:
	for k, v := range m.values {
		for _, k2 := range keys {
			if k2 == k {
				continue skip
			}
		}
		newMap.values[k] = v
	}
	return
}

//MutableMapm *TestCaseTypeMap toMap();

//ImmutableSetMultimap<V, K> flip();

func (m *TestCaseTypeMap) Select(predicate func(key int, value *TestCaseType) bool) (newMap *TestCaseTypeMap) {
	newMap = &TestCaseTypeMap{
		values: map[int]*TestCaseType{},
	}
	for k, v := range m.values {
		if predicate(k, v) {
			newMap.values[k] = v
		}
	}
	return
}

func (m *TestCaseTypeMap) Reject(predicate func(key int, value *TestCaseType) bool) (newMap *TestCaseTypeMap) {
	newMap = &TestCaseTypeMap{
		values: map[int]*TestCaseType{},
	}
	for k, v := range m.values {
		if predicate(k, v) == false {
			newMap.values[k] = v
		}
	}
	return
}

func (m *TestCaseTypeMap) Tap(procedure func(values *TestCaseType)) {
	for _, v := range m.values {
		procedure(v)
	}
}

func (m *TestCaseTypeMap) Partition(predicate func(key int, value *TestCaseType) bool) (selected, rejected []*TestCaseType) {
	selected = []*TestCaseType{}
	rejected = []*TestCaseType{}
	for k, v := range m.values {
		if predicate(k, v) {
			selected = append(selected, v)
		} else {
			rejected = append(rejected, v)
		}
	}
	return
}
