package cfg

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, left, right interface{}, err string) {
	if reflect.TypeOf(left).Kind() != reflect.TypeOf(right).Kind() {
		t.Error(err)
	}
	switch reflect.TypeOf(left).Kind() {
	case reflect.Bool:
		if reflect.ValueOf(left).Bool() != reflect.ValueOf(right).Bool() {
			t.Error(err)
		}
	}
}

func TestNewValue(t *testing.T) {
	v := NewValue(BOOL, true)
	t.Log(reflect.ValueOf(v.Bool))
	AssertEqual(t, v.Bool, true, "Bool 测试出错")
	m := make(map[string]*Value)
	m["test"] = &Value{Bool: true}
	v = NewValue(OBJECT, m)
	AssertEqual(t, m["test"].Bool, true, "object 测试出错")
}
