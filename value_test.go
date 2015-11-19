package cfg

import (
	"reflect"
	"testing"
)

func TestNewValue(t *testing.T) {
	v := NewValue(BOOL, true)
	t.Log(reflect.ValueOf(v.Bool))
	AssertEqual(t, v.Bool, true, "Bool 测试出错")
	m := make(map[string]*Value)
	m["test"] = &Value{Bool: true}
	v = NewValue(OBJECT, m)
	AssertEqual(t, m["test"].Bool, true, "object 测试出错")
}
