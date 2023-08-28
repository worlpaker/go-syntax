# Change Log

All notable changes to the "go-syntax" extension will be documented in this file.

## [0.2.37]

- Fixes a bug related to function types with var keyword

## [0.2.36]

- Fixes bug related to returning multiple params in functions in-line

## [0.2.35]

- Fixes a small bug related to generics in functions in-line

## [0.2.34]

- Fixes minor bugs in function declarations
- Enhances code readability
- Removes unnecessary sections

## [0.2.33]

- Fixes a small bug in function declaration

## [0.2.32]

- Fixes a small bug in make built-in function

## [0.2.31]

- Fixes small bugs related to types within function declaration

## [0.2.30]

- Fixes a small bug in new built-in function

## [0.2.29]

- Fixes minor bugs related to variable types within functions

- Enhances generic support in functions

- Improves type support for the built-in make function

- Introduces new built-in functions: min, max, and clear

## [0.2.28]

- Fixes small bugs in variables

## [0.2.27]

- Fixes small bugs in interface types

## [0.2.26]

- Fixes a small bug in interface types

## [0.2.25]

- Fixes small bugs in switch type assertion

## [0.2.24]

- Fixes small bugs in variable types

## [0.2.23]

- Fixes a small bug in variable types

## [0.2.22]

- Fixes small bugs in all types

## [0.2.21]

- Fixes small bugs in generic types

## [0.2.20]

Starting from now, the extension will provide significantly enhanced support for variables compared to previous versions.

Details:

- Variable parameters in functions and generics will be scoped as "variable.parameter.go"

- Variable names in structs and struct initialization will be scoped as "variable.other.property.go"

- Variable names in imports will be scoped as "variable.other.import.go"

- Loop label names will be scoped as "variable.other.label.go"

## [0.2.19]

Add support for:

- Variable parameters in functions

- Variable parameters in generics

- Variable names in struct

- Variable names in struct initialization

- Label loop names

## [0.2.18]

- Enhances generic support in all types

## [0.2.17]

- Enhances generic support in function declarations

## [0.2.16]

- Adds a new icon

## [0.2.15]

- Fixes a small bug related to structs in function parameters

## [0.2.14]

- Adds type support to the new keyword (to instantiate a struct)

## [0.2.13]

- Fixes a small bug in import

## [0.2.12]

- Adds generic support to support functions

- Fixes small bugs (interface types, type assertions)

## [0.2.11]

- Fixes small bug in map types

## [0.2.10]

- Fixes small bugs in struct fields

## [0.2.9]

- Fixes incorrect highlighting for multidimensional slice type in struct's fields

## [0.2.8]

- Fixes small bugs in function declaration

## [0.2.7]

- Fixes different colors (storage types) in support functions with a new solution

## [0.2.6]

- Fixes decrementing operator highlighting in for loop
- Fixes a small bug that occurred in the v0.2.5

## [0.2.5]

- Fixes different colors (storage types) in support functions

## [0.2.4]

- Adds generic support to the struct

## [0.2.3]

- Stable Version

## [0.2.2]

- Fixes unnecessary highlighting (control keywords)

## [0.2.1]

- Fixes unnecessary highlighting (other struct expressions, other type names)

## [0.2.0]

- Stable Version

## [0.1.1]

- Fixes other type names declaration in function

## [0.1.0]

- First version of Go Syntax
