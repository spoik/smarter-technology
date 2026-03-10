package main

import (
	"errors"
	"fmt"
)

func main() {
	width := 10.0
	height := 10.0
	length := 10.0
	mass := 10.0
	stack, err := SortPackage(width, height, length, mass)

    if err != nil {
        fmt.Printf("There was an issue sorting the package: %s\n", err)
        return
    }

    fmt.Println("The package should be added to the", stack, "stack.")
}

func SortPackage(width, height, length, mass float64) (string, error) {
	err := validateInput(width, height, length, mass)

	if err != nil {
		return "", err
	}

	heavy := isHeavy(mass)
	bulky := isBulky(width, height, length)

	if heavy && bulky {
		return "REJECTED", nil
	}

	if heavy || bulky {
		return "SPECIAL", nil
	}

	return "STANDARD", nil
}

func validateInput(width, height, length, mass float64) error {
	if width <= 0 {
		return errors.New("Width must be greater than 0.")
	}

	if height <= 0 {
		return errors.New("Height must be greater than 0.")
	}

	if length <= 0 {
		return errors.New("Length must be greater than 0.")
	}

	if mass <= 0 {
		return errors.New("Mass must be greater than 0.")
	}

	return nil
}

func isBulky(width, height, length float64) bool {
	if width >= 150 || height >= 150 || length >= 150 {
		return true
	}

	volume := width * height * length
	if volume >= 1000000 {
		return true
	}

	return false
}

func isHeavy(mass float64) bool {
	return mass >= 20
}
