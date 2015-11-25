// value
package cfg

import (
	"errors"
	"reflect"
)

type Value struct {
	Bool      bool              //bool类型
	Int       int64             //整形
	String    string            // 字符串类型
	Double    float64           // 浮点数类型
	Array     []*Value          // 数组类型 []  弃用
	Object    map[string]*Value //对象类型 {} 弃用
	ValueType int               // 值类型
}

func (v *Value) GetBool() bool {
	AssertEqual(v.ValueType, BOOL, "Not Bool")
	return v.Bool
}

func (v *Value) GetString() string {
	AssertEqual(v.ValueType, STRING, "Not String")
	return v.String
}

func NewValue(valueType int, v interface{}) *Value {
	value := &Value{
		ValueType: valueType,
	}
	t := reflect.TypeOf(v).Kind()
	switch valueType {
	case BOOL:
		if t == reflect.Bool {
			value.Bool = v.(bool)
		} else {
			panic(errors.New("数值类型非 bool 类型,而是 " + reflect.TypeOf(v).String()))
		}
	case INT:
		if t >= reflect.Int || t <= reflect.Uint64 { // 只要是整数全部转为 int64 避免麻烦
			value.Int = v.(int64)
		} else {
			panic(errors.New("数值类型非 int 类型,而是 " + reflect.TypeOf(v).String()))
		}
	case STRING:
		if t == reflect.String {
			value.String = v.(string)
		} else {
			panic(errors.New("数值类型非 string 类型,而是 " + reflect.TypeOf(v).String()))
		}
	case DOUBLE:
		if t >= reflect.Float32 || t <= reflect.Float64 {
			value.Double = v.(float64)
		} else {
			panic(errors.New("数值类型非 float 类型,而是 " + reflect.TypeOf(v).String()))
		}
	case ARRAY:
		if t == reflect.Array {
			value.Array = append(value.Array, v.([]*Value)...)
		} else {
			panic(errors.New("数值类型非 array 类型,而是 " + reflect.TypeOf(v).String()))
		}
	case OBJECT:
		if t == reflect.Map {
			//			value.Int = v.(int64)
			value.Object = make(map[string]*Value)
			for key, vlu := range v.(map[string]*Value) {
				value.Object[key] = vlu
			}
		} else {
			panic(errors.New("数值类型非 object 类型,而是 " + reflect.TypeOf(v).String()))
		}
	}
	return value
}
