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

Variable declaration (without initialization)
"zero value" is assigned to uninitialized variabled
```
var greeting string  // zero value is empty string ""
```
Variable initialization, now we assign a value
```
greeting = "Hello friend"
```
Can declare multiple variables on the same line
```
var firstName, lastName string
firstName = "John"
lastName = "Doe"
```
Short variable delcaration and initialization at the same time
```
email := "test@test.com"
```
Redundant but acceptable syntax
```
var year int = 2025
```
Acceptable and not redundant but less clean than `:=`
```
var month = "May"
```
Note: can declare/initialize variable outside of function

## Constants

Constants are immutable and cannot be reassigned
```
const HOST string = "localhost"
```
The variable type may be inferenced
```
const PORT = 8080
```
Note: can declare/initialize const outside of function

## Enum