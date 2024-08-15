package reflectutil

import (
	"fmt"
	"reflect"
)

func CopyMatchingFields(src, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	if dstVal.Kind() != reflect.Ptr || dstVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("destination must be a pointer to a struct")
	}

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if srcVal.Kind() != reflect.Struct {
		return fmt.Errorf("source must be a struct or a pointer to a struct")
	}

	dstElem := dstVal.Elem()
	srcType := srcVal.Type()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		srcFieldName := srcType.Field(i).Name

		dstField := dstElem.FieldByName(srcFieldName)
		if !dstField.IsValid() || !dstField.CanSet() {
			continue
		}

		if srcField.Type() == dstField.Type() {
			dstField.Set(srcField)
		}
	}

	return nil
}

func MapToStructByFieldName(m map[string]interface{}, s interface{}) error {
	sVal := reflect.ValueOf(s).Elem()

	if sVal.Kind() != reflect.Struct {
		return fmt.Errorf("destination must be a pointer to a struct")
	}

	for key, value := range m {
		field := sVal.FieldByName(key)
		if !field.IsValid() || !field.CanSet() {
			continue
		}

		val := reflect.ValueOf(value)
		if field.Type() == val.Type() {
			field.Set(val)
		}
	}

	return nil
}

func StructToMapByFieldName(s interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("source must be a struct or a pointer to a struct")
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.CanInterface() {
			result[typ.Field(i).Name] = field.Interface()
		}
	}

	return result, nil
}
