# Smarter Technology Technical Screen
`main.go` contains all the logic. `SortPackage()` is a function that will sort packages based on a package's width, height, length, and mass. There is a `main()` function with a width, height, length and mass that can be changed to easily play with `SortPackage()`. This can be run by installing the dependencies with `go download` and then running `go run main.go` to run `main()`.

`main_test.go` contains all the unit tests. These can be run by running `go run ./...`.

## How to use `SortPackage()`
Call SortPackage with a width, height, length, and mass to get the stack the package should be dispatched to. Width, height, and length are floats in centimeters that represents the package's dimensions. Mass is a float in kilograms that represents the package's weight.

If there is no error, err will be `nil` and stack will either be the `"REJECTED"`, `"SPECIAL"`, or `"STANDARD"` string.

If there is an error, err will be an `error` instance and stack will be the empty string.
```go
width := 10.0
height := 10.0
length := 10.0
mass := 10.0
stack, err := SortPackage(width, height, length, mass)
```
