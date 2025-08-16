# DeepCopy Go

A Go package that provides deep copying functionality for complex data structures, including nested structs, pointers, slices, maps, and circular references.

## Features

- Deep copies of complex Go data structures
- Handles circular references correctly
- Preserves the structure and values of nested types
- Supports all basic types, structs, slices, arrays, and maps
- Properly handles pointer references
- Preserves the relationship between shared references


## Usage

```go

// For any value, create a deep copy
old := 1
new, err := deepcopy.DeepCopy(old)
if err != nil {
    // handle error
}
newValue := new.(int)// Assert the type
```


## How It Works

The package uses reflection to traverse the entire object graph and create new copies of all values. It handles:

- Basic types (int, string, bool, etc.)
- Pointers (including circular references)
- Slices and arrays
- Maps
- Structs (including unexported fields)
- Nested combinations of the above

## Testing

Run the tests with:

```bash
go test -v
```

The test suite includes cases for:
- Basic type copying
- Nested structs
- Circular references
- Shared references
- Pointer handling
- Map and slice copying

## Limitations

- Performance: Reflection-based copying is slower than manual copying
- Some edge cases with unexported fields might not be handled
- Complex types like channels and functions are not supported

## License

[MIT](LICENSE)
