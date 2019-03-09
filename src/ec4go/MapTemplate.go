package main

var immutableMapTemplate = `/*
* CODE GENERATED AUTOMATICALLY WITH github.com/maded2/ec4go
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/
package {{.PackageName}}

func {{.TypeName}}_NewWithKeyValue(key {{.KeyType}}, value {{.ValueType}}) {{.TypeName}} {
	return {{.TypeName}}{
		key: value,
	}
}

func {{.TypeName}}_NewWithAllKeyValues(values map[{{.KeyType}}]{{.ValueType}}) (newMap {{.TypeName}}) {
	newMap = {{.TypeName}}{}
	for k, v := range values {
		newMap[k] = v
	}
	return
}

func (m {{.TypeName}}) NewWithoutKey(key {{.KeyType}}) (newMap {{.TypeName}}) {
	newMap = {{.TypeName}}{}
	for k, v := range m {
		if key != k {
			newMap[k] = v
		}
	}
	return
}

func (m {{.TypeName}}) NewWithoutAllKeys(keys []{{.KeyType}}) (newMap {{.TypeName}}) {
	newMap = {{.TypeName}}{}
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

func (m {{.TypeName}}) Select(predicate func(key {{.KeyType}}, value {{.ValueType}}) bool) (newMap {{.TypeName}}) {
	newMap = {{.TypeName}}{}
	for k, v := range m {
		if predicate(k, v) {
			newMap[k] = v
		}
	}
	return
}

func (m {{.TypeName}}) Reject(predicate func(key {{.KeyType}}, value {{.ValueType}}) bool) (newMap {{.TypeName}}) {
	newMap = {{.TypeName}}{}
	for k, v := range m {
		if predicate(k, v) == false {
			newMap[k] = v
		}
	}
	return
}

func (m {{.TypeName}}) Tap(procedure func(value {{.ValueType}})) {
	for _, v := range m {
		procedure(v)
	}
}

func (m {{.TypeName}}) Partition(predicate func(key {{.KeyType}}, value {{.ValueType}}) bool) (selected, rejected []{{.ValueType}}) {
	selected = []{{.ValueType}}{}
	rejected = []{{.ValueType}}{}
	for k, v := range m {
		if predicate(k, v) {
			selected = append(selected, v)
		} else {
			rejected = append(rejected, v)
		}
	}
	return
}


`
