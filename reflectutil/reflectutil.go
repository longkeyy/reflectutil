package reflectutil

import (
	"fmt"
	"reflect"
	"strings"
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
	sType := sVal.Type()

	if sVal.Kind() != reflect.Struct {
		return fmt.Errorf("destination must be a pointer to a struct")
	}

	for i := 0; i < sVal.NumField(); i++ {
		field := sVal.Field(i)
		fieldType := sType.Field(i)

		if !field.CanSet() {
			continue
		}

		tag := getTag(fieldType)
		if tag == "" {
			tag = fieldType.Name
		}

		value, exists := m[tag]
		if !exists {
			continue
		}

		if err := setField(field, value); err != nil {
			return fmt.Errorf("error setting field %s: %v", fieldType.Name, err)
		}
	}

	return nil
}

func getTag(field reflect.StructField) string {
	tags := []string{"mapstructure", "json", "yaml", "xml"}
	for _, tagName := range tags {
		if tag := field.Tag.Get(tagName); tag != "" {
			return strings.Split(tag, ",")[0] // 处理可能的选项，如 `json:"name,omitempty"`
		}
	}
	return ""
}

func setField(field reflect.Value, value interface{}) error {
	val := reflect.ValueOf(value)

	if field.Type() == val.Type() {
		field.Set(val)
		return nil
	}

	if val.Type().ConvertibleTo(field.Type()) {
		field.Set(val.Convert(field.Type()))
		return nil
	}

	if field.Kind() == reflect.Ptr {
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		return setField(field.Elem(), value)
	}

	return fmt.Errorf("cannot set field of type %v with value of type %v", field.Type(), val.Type())
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
