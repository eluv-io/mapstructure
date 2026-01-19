package mapstructure

import (
	"reflect"
	"strings"
	"sync"
)

var (
	structTypeCache sync.Map // map[cacheKey]*structType
)

type structType struct {
	fields []*structField
}

type cacheKey struct {
	t       reflect.Type
	tagName string
}

func getStructType(t reflect.Type, tagName string) *structType {
	key := cacheKey{t, tagName}
	if v, _ := structTypeCache.Load(key); v != nil {
		return v.(*structType)
	}

	var fields []*structField
	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)
		fields = append(fields, &structField{
			field: fieldType,
			tags:  strings.Split(fieldType.Tag.Get(tagName), ","),
		})
	}

	ret := &structType{fields: fields}
	structTypeCache.Store(key, ret)

	return ret
}

// A structField describes a single field in a struct.
type structField struct {
	field reflect.StructField
	tags  tags
}

type tags []string

func (t tags) containsAfterTag(s string) bool {
	for i := 1; i < len(t); i++ {
		if t[i] == s {
			return true
		}
	}
	return false
}
func (t tags) omitEmpty() bool {
	return t.containsAfterTag("omitempty")
}
func (t tags) squash() bool {
	return t.containsAfterTag("squash")
}
