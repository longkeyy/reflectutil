// File: main.go
package main

import (
	"fmt"
	"github.com/longkeyy/reflectutil/reflectutil"
)

func main() {
	fmt.Println("reflectutil Usage Examples")
	fmt.Println("==========================")

	// Example 1: CopyMatchingFields
	fmt.Println("\n1. CopyMatchingFields Example:")
	src := struct {
		Name string
		Age  int
	}{
		Name: "John Doe",
		Age:  30,
	}

	dst := struct {
		Name string
		Age  int
	}{}

	err := reflectutil.CopyMatchingFields(src, &dst)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Source: %+v\n", src)
		fmt.Printf("Destination after copy: %+v\n", dst)
	}

	// Example 2: MapToStructByFieldName
	fmt.Println("\n2. MapToStructByFieldName Example:")
	m := map[string]interface{}{
		"Name": "Jane Doe",
		"Age":  25,
	}

	s := struct {
		Name string
		Age  int
	}{}

	err = reflectutil.MapToStructByFieldName(m, &s)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Map: %+v\n", m)
		fmt.Printf("Struct after conversion: %+v\n", s)
	}

	// Example 3: StructToMapByFieldName
	fmt.Println("\n3. StructToMapByFieldName Example:")
	person := struct {
		Name    string
		Age     int
		IsAdmin bool
	}{
		Name:    "Alice",
		Age:     35,
		IsAdmin: true,
	}

	resultMap, err := reflectutil.StructToMapByFieldName(person)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Original struct: %+v\n", person)
		fmt.Printf("Resulting map: %+v\n", resultMap)
	}
}
