package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBulky(t *testing.T) {
	type testCase struct {
		description    string
		width          float64
		height         float64
		length         float64
		expectedResult bool
	}

	testCases := []testCase{
		{
			description:    "Not bulky",
			width:          10,
			height:         10,
			length:         10,
			expectedResult: false,
		},
		{
			description:    "Bulky width",
			width:          150,
			height:         10,
			length:         10,
			expectedResult: true,
		},
		{
			description:    "Bulky height",
			width:          10,
			height:         150,
			length:         10,
			expectedResult: true,
		},
		{
			description:    "Bulky length",
			width:          10,
			height:         10,
			length:         150,
			expectedResult: true,
		},
		{
			description:    "Bulky Volume",
			width:          100,
			height:         100,
			length:         100,
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			t.Parallel()

			result := isBulky(
				testCase.width,
				testCase.height,
				testCase.length,
			)

			assert.Equal(t, testCase.expectedResult, result)
		})
	}
}

func TestIsHeavy(t *testing.T) {
	type testCase struct {
		description    string
		mass           float64
		expectedResult bool
	}

	testCases := []testCase{
		{
			description:    "Not heavy",
			mass:           19.999,
			expectedResult: false,
		},
		{
			description:    "Heavy",
			mass:           20,
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			t.Parallel()
			result := isHeavy(testCase.mass)

			assert.Equal(t, testCase.expectedResult, result)
		})
	}
}

func TestSortPackage(t *testing.T) {
	type testCase struct {
		description    string
		width          float64
		height         float64
		length         float64
		mass           float64
		expectedResult string
	}

	testCases := []testCase{
		{
			description:    "Standard",
			width:          1,
			height:         1,
			length:         1,
			mass:           1,
			expectedResult: "STANDARD",
		},
		{
			description:    "Heavy but not bulky",
			width:          1,
			height:         1,
			length:         1,
			mass:           20,
			expectedResult: "SPECIAL",
		},
		{
			description:    "Bulky but not heavy",
			width:          100,
			height:         100,
			length:         100,
			mass:           19,
			expectedResult: "SPECIAL",
		},
		{
			description:    "Bulky and heavy",
			width:          100,
			height:         100,
			length:         100,
			mass:           20,
			expectedResult: "REJECTED",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			t.Parallel()

			result, err := SortPackage(
				testCase.width,
				testCase.height,
				testCase.length,
				testCase.mass,
			)

			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedResult, result)
		})
	}
}

func TestSortPackageInvalidParams(t *testing.T) {
	type testCase struct {
		description    string
		width          float64
		height         float64
		length         float64
		mass           float64
		expectedErrMsg string
	}

	testCases := []testCase{
		{
			description:    "Width is 0",
			width:          0,
			height:         1,
			length:         1,
			mass:           1,
			expectedErrMsg: "Width must be greater than 0.",
		},
		{
			description:    "Height is 0",
			width:          1,
			height:         0,
			length:         1,
			mass:           1,
			expectedErrMsg: "Height must be greater than 0.",
		},
		{
			description:    "Length is 0",
			width:          1,
			height:         1,
			length:         0,
			mass:           1,
			expectedErrMsg: "Length must be greater than 0.",
		},
		{
			description:    "Mass is 0",
			width:          1,
			height:         1,
			length:         1,
			mass:           0,
			expectedErrMsg: "Mass must be greater than 0.",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			t.Parallel()

			_, err := SortPackage(
				testCase.width,
				testCase.height,
				testCase.length,
				testCase.mass,
			)

			assert.EqualError(t, err, testCase.expectedErrMsg)
		})
	}
}
