# reflectutil

reflectutil is a Go package that provides utility functions for working with structs and maps using reflection.

## Installation

To install the reflectutil package, use the following command:

```
go get github.com/longkeyy/reflectutil
```

## Usage

Import the package in your Go code:

```go
import "github.com/longkeyy/reflectutil/reflectutil"
```

### Available Functions and Examples

1. `CopyMatchingFields(src, dst interface{}) error`
   - Copies fields from src to dst if they have the same name and type.

   Example:
   ```go
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

   err := reflectutil.CopyMatchingFields(src, &dst)
   if err != nil {
       fmt.Printf("Error: %v\n", err)
   } else {
       fmt.Printf("Destination after copy: %+v\n", dst)
   }
   // Output: Destination after copy: {Name:John Doe Age:30}
   ```

2. `MapToStructByFieldName(m map[string]interface{}, s interface{}) error`
   - Copies values from a map to a struct based on field names.

   Example:
   ```go
   type Person struct {
       Name string
       Age  int
   }

   m := map[string]interface{}{
       "Name": "Jane Doe",
       "Age":  25,
   }

   p := Person{}

   err := reflectutil.MapToStructByFieldName(m, &p)
   if err != nil {
       fmt.Printf("Error: %v\n", err)
   } else {
       fmt.Printf("Person after conversion: %+v\n", p)
   }
   // Output: Person after conversion: {Name:Jane Doe Age:25}
   ```

3. `StructToMapByFieldName(s interface{}) (map[string]interface{}, error)`
   - Converts a struct to a map using field names as keys.

   Example:
   ```go
   type Employee struct {
       Name    string
       Age     int
       IsAdmin bool
   }

   e := Employee{Name: "Alice", Age: 35, IsAdmin: true}

   resultMap, err := reflectutil.StructToMapByFieldName(e)
   if err != nil {
       fmt.Printf("Error: %v\n", err)
   } else {
       fmt.Printf("Resulting map: %+v\n", resultMap)
   }
   // Output: Resulting map: map[Age:35 IsAdmin:true Name:Alice]
   ```

## Testing

To run the tests for the reflectutil package:

```
go test github.com/longkeyy/reflectutil/reflectutil
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the terms of the LICENSE file included in the repository.