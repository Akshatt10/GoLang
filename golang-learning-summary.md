# GoLang Learning Summary

This document summarizes the key concepts and features learned in Go (Golang) so far.

---

## 1. Basics

- **Variables:** Declared using `var` or `:=` for short declaration.
- **Data Types:** int, float64, string, bool, arrays, slices, maps, structs.
- **Constants:** Declared with `const`.
- **Functions:** Defined with `func`, can return multiple values.

## 2. Control Structures

- **If/Else:** Standard conditional statements.
- **Switch:** Used for multiple conditions.
- **For Loop:** The only loop in Go, can be used as while or traditional for.

## 3. Arrays, Slices, and Maps

- **Arrays:** Fixed size, same type elements.
- **Slices:** Dynamic size, reference type, more commonly used than arrays.
- **Maps:** Key-value pairs, similar to dictionaries in other languages.

## 4. Structs

- Custom data types grouping multiple fields.
- Used to model real-world entities.

## 5. Pointers

- Hold memory addresses of variables.
- Useful for modifying values in functions.

## 6. Functions

- Support for multiple return values.
- Variadic functions (accept variable number of arguments).
- Anonymous functions and closures.

## 7. Methods and Interfaces

- Methods: Functions with a receiver argument.
- Interfaces: Define behavior, implemented implicitly.

## 8. Packages and Imports

- Code organized into packages.
- Use `import` to include standard or custom packages.

## 9. Error Handling

- Errors are values, handled explicitly.
- Use `error` type and `errors.New()` or `fmt.Errorf()`.

## 10. Concurrency

- Goroutines: Lightweight threads managed by Go runtime.
- Channels: Used for communication between goroutines.

---

## Example Code

```go
package main

import "fmt"

func main() {
    var message string = "Hello, Go!"
    fmt.Println(message)
}
```

---

## Additional Notes

- Go uses garbage collection.
- Strongly typed and statically compiled.
- Simple syntax, designed for readability and efficiency.

---
