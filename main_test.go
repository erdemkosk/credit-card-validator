package main

import (
	"testing"
)

func TestSumSlice(t *testing.T) {
	testCases := []struct {
		slice          []int
		expectedResult int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{4, 5, 6}, 15},
		{[]int{0, 0, 0}, 0},
	}

	for _, testCase := range testCases {
		result := SumSlice(testCase.slice)
		if result != testCase.expectedResult {
			t.Errorf("Expected: %v, happened: %v", testCase.expectedResult, result)
		}
	}
}

func TestLunhAlgorithmChecker(t *testing.T) {
	testCases := []struct {
		creditCardNumber string
		expectedOddSum   []int
		expectedEvenSum  []int
	}{
		{
			creditCardNumber: "5388466965157906",
			expectedOddSum:   []int{1, 7, 8, 3, 3, 2, 5, 0},
			expectedEvenSum:  []int{3, 8, 6, 9, 5, 5, 9, 6},
		},
		{
			creditCardNumber: "1234567890123456",
			expectedOddSum:   []int{2, 6, 1, 5, 9, 2, 6, 1},
			expectedEvenSum:  []int{2, 4, 6, 8, 0, 2, 4, 6},
		},
	}

	for _, testCase := range testCases {
		oddSum, evenSum := LunhAlgorithmChecker(testCase.creditCardNumber)

		for i, val := range oddSum {
			if val != testCase.expectedOddSum[i] {
				t.Errorf("Expected Odd Sum: %v, happened: %v", testCase.expectedOddSum, oddSum)
				break
			}
		}

		for i, val := range evenSum {
			if val != testCase.expectedEvenSum[i] {
				t.Errorf("Expected Even Sum: %v, happened: %v", testCase.expectedEvenSum, evenSum)
				break
			}
		}
	}
}

func TestCheckCreditCardIsValid(t *testing.T) {
	testCases := []struct {
		creditCardNumber string
		expectedResult   bool
		expectedError    bool
	}{
		{"5388466965157906", true, false},
		{"1234", false, true},
	}

	for _, testCase := range testCases {
		result, err := CheckCreditCardIsValid(testCase.creditCardNumber)
		if (err != nil) != testCase.expectedError {
			t.Errorf("Expected error: %v", testCase.expectedError)
		}

		if result != testCase.expectedResult {
			t.Errorf("Expected: %v, happened: %v", testCase.expectedResult, result)
		}
	}
}
