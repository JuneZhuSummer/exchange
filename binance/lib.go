package binance

import (
	"errors"
	"reflect"
	"strings"
)

const (
	TagJson = "json"
)

var (
	ErrorKind = errors.New("invalid kind")
)

// StructToMap 结构体转为map[string]any
func StructToMap(in any, tagName string) (map[string]any, error) {
	out := make(map[string]any)

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		// 非结构体返回错误提示
		return nil, ErrorKind
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)

		tagValue := field.Tag.Get(tagName)
		if tagValue != "" && tagValue != "-" && !isBlank(v.Field(i)) {
			if strings.Contains(tagValue, ",") {
				tagValue = strings.Split(tagValue, ",")[0] //忽略逗号后的内容
			}
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	default:
		return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
	}
}
