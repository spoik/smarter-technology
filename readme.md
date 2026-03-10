# Smarter Technology Technical Screen
## How to use this package
```go
// Call SortPackage with a width, height, length, and mass to get the stack the package should be dispatched to. Width, height, and length are floats in centimeters. Mass is a float in kilograms.
//
// If there is no error, err will be nil and stack will either be the "REJECTED", "SPECIAL", or "STANDARD" strings.
//
// If there is an error, err will contain the error and stack will be the empty string.
width := 10.0
height := 10.0
length := 10.0
mass := 10.0
stack, err := SortPackage(width, height, length, mass)
```

## How to run the tests
Run `go run ./...` from the command line to run all the tests.
