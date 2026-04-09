# Section 4 Summary

## Print

`Println` adds newline character at end by default and spaces between args
```
fmt.Println(true, false)
```
`Printf` offers string formatting using format specifiers to substitute variables into a string
e.g. %d for integers, %s for strings, %f for float, %t for bool, %v for default format, %T for type, %p for pointer
```
fmt.Printf("Hello, %s version %.2f\n", name, version)
```
*Primary formatters*:
| Specifier | Description |
|-----------|-------------|
| `%+v`     | Like `%v`; for structs, includes field names |
| `%#v`     | Go-syntax style (types/literals where it applies) |
| `%8v`     | Width allocated |
| `%-v`     | Left-align, makes sense with width allocated |
| `%0d`     | Zero-pad for numbers |
Many more types exist for integers, floats, strings / byte slices, and pointers.

## Variables

## Constants

## Enum