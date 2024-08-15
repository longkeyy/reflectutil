package reflectutil

import (
	"reflect"
	"testing"
)

func TestCopyMatchingFields(t *testing.T) {
	type Source struct {
		Name string
		Age  int
	}

	type Destination struct {
		Name string
		Age  int
	}

	src := Source{Name: "John Doe", Age: 30}
	dst := Destination{}

	err := CopyMatchingFields(src, &dst)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(src, Source(dst)) {
		t.Errorf("Structs are not equal. Expected %+v, got %+v", src, dst)
	}
}

func TestMapToStructByFieldName(t *testing.T) {
	type TestStruct struct {
		Name string
		Age  int
	}

	m := map[string]interface{}{
		"Name": "Jane Doe",
		"Age":  25,
	}

	s := TestStruct{}

	err := MapToStructByFieldName(m, &s)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if s.Name != "Jane Doe" || s.Age != 25 {
		t.Errorf("Struct fields not set correctly. Expected {Jane Doe 25}, got %+v", s)
	}
}

func TestStructToMapByFieldName(t *testing.T) {
	type TestStruct struct {
		Name string
		Age  int
	}

	s := TestStruct{Name: "Alice", Age: 35}

	m, err := StructToMapByFieldName(s)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := map[string]interface{}{
		"Name": "Alice",
		"Age":  35,
	}

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("Maps are not equal. Expected %+v, got %+v", expected, m)
	}
}
