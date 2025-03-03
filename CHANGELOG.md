# Change Log

All notable changes to the "go-syntax" extension will be documented in this file.

## [0.8.1]

- Supports the comparable type as `entity.name.type.comparable.go`
- Fixes minor bugs related to generic types in double parentheses and structs

## [0.8.0]

- Supports constant variables

## [0.7.9]

- Fixes minor bugs related to functions ([#17](https://github.com/worlpaker/go-syntax/issues/17))

## [0.7.8]

- Supports generic alias types

## [0.7.7]

- Fixes minor bugs related to struct fields ([#16](https://github.com/worlpaker/go-syntax/issues/16))

## [0.7.6]

- Fixes a minor bug in map type

## [0.7.5]

- Fixes a minor bug in switch type

## [0.7.4]

- Fixes a minor bug in interface type

## [0.7.3]

- Fixes property variables (`variable.other.property.field.go` is deprecated)

## [0.7.2]

- Fixes minor bugs in struct field hover and generic type

## [0.7.1]

- Updates `README.md`

## [0.7.0]

- Fixes minor bugs related to multi-line struct tags

## [0.6.9]

- Fixes minor bugs in make built-in function

## [0.6.8]

- Fixes struct field when hovering with the mouse

## [0.6.7]

- Improves constants definitions ([PR 10](https://github.com/worlpaker/go-syntax/pull/10) by [@butuzov](https://github.com/butuzov))

## [0.6.6]

- Fixes minor bugs in struct and single types

## [0.6.5]

- Fixes a minor bug related to map types in functions

## [0.6.4]

- Fixes a minor bug in struct type

## [0.6.3]

Fixes:

- Minor bugs in struct type
- Interface type in functions and generics

## [0.6.2]

- Fixes a minor bug in functions inline

## [0.6.1]

Fixes:

- Minor bugs in interface types and make built-in function
- Early highlighting of variables after control keywords (before formatting with gofmt)
- Early highlighting of types after function declaration (before formatting with gofmt)

## [0.6.0]

- Fixes variables after case keyword in the switch statement
- Fixes variables and types after control keywords
- Removes unnecessary sections

## [0.5.9]

- Fixes a minor bug in double parentheses type
- Improves performance for large files, especially those with long lines of support functions

## [0.5.8]

- Enhances types after the case keyword in the switch type statement

## [0.5.7]

- Refactors map types
- Enhances generics types
- Fixes variable assignment with generic types
- Fixes support functions and functions inline
- Fixes minor bugs in types

## [0.5.6]

Fixes:

- Minor bugs related to chan keyword in variable assignment
- Property field variables and slice index variables with operators

## [0.5.5]

Fixes:

- Minor bugs in types related to double parentheses
- Property field variables with comparison operators
- Slice index variables when using arithmetic operators

## [0.5.4]

- Refactors variable assignment and other struct/interface expressions
- Optimizes performance for faster opening of large files, especially those with long lines of variables, parameters, and types
- Implements stress testing to ensure stability and reliability
- Fixes variables

## [0.5.3]

Updates `.vscodeignore` to exclude unnecessary files and folders (`examples`, `test`)

## [0.5.2]

Fixes early highlighting of variables after control keywords (before formatting with gofmt)

## [0.5.1]

Enhances struct property variables (See [#4](https://github.com/worlpaker/go-syntax/issues/4)):

- struct initialization property variables field will be scoped as `"variable.other.property.field.go"`

Fixes:

- minor bugs related to struct types within function declaration
- switch/select case variables with support functions
- early highlighting of struct declaration and variables after control keywords (before formatting with gofmt)
- type declaration with generics on multi-lines

## [0.5.0]

- Enhances function and struct declaration
- Refactors function and generic parameter types
- Supports double parentheses types
- Corrects keyword operator behaviors
- Optimizes code
- Improves documentation
- Removes unnecessary sections
- Adds more tests
- Fixes import declaration with semicolon (before formatting with gofmt)
- Fixes minor bugs in variable declaration, switch case variables, multi type declaration, type assertion, and function in-line
- Fixes bugs related to struct types within function parameters, struct declaration, map types, and variables

## [0.4.8]

- Fixes function declaration
- Fixes generic types in the new built-in function
- Fixes chan in make built-in function, struct type fields, variables, and function parameter types
- Enhances function in-line with support for multi return types
- Improves chan with slices for types

## [0.4.7]

- Fixes early highlighting of variables after control keywords (before formatting with gofmt)

## [0.4.6]

- Fixes minor bugs in new built-in function

## [0.4.5]

- Fixes bugs related to property variables from third party imports
- Fixes minor bugs in map type

## [0.4.4]

- Fixes minor bugs in function declaration

## [0.4.3]

- Fixes minor bugs in struct type

## [0.4.2]

- Fixes multi type declaration

## [0.4.1]

- Fixes minor bugs in generic functions

## [0.4.0]

- Fixes struct type
- Corrects map type with functions
- Updates make and new built-in functions
- Improves function declarations and function inline
- Enhances support for generics in support functions

## [0.3.9]

- Fixes type assertion in switch statements when using support functions

## [0.3.8]

- Fixes minor bugs in type and function declaration

## [0.3.7]

- Fixes pre-highlighting of the single type field in struct before formatting with gofmt

## [0.3.6]

- Enhances generic support for types and functions

## [0.3.5]

- Enhances pre-highlighting of variables after control keywords before formatting with gofmt

## [0.3.4]

- Fixes pre-highlighting of variables after control keywords before formatting with gofmt

## [0.3.3]

- Fixes the preview of variables when hovering with the cursor for the tooltip

## [0.3.2]

- Fixes bugs related to types when declaring variables using the var keyword

## [0.3.1]

- Enhances generic type support for map types

## [0.3.0]

- Updates `README.md`

## [0.2.40]

- Improves generic type support with map when declaring variables using the var keyword
- Adds new snap commands for test

## [0.2.39]

- Enhances generic type support with map for the built-in make function

## [0.2.38]

- Improves generic type support for the built-in make function

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
