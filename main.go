package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Lunch Algorithm test")

	success, _ := LunhAlgorithmChecker("5388466965157906")

	fmt.Println((success))

}

func sumSlice(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}

func LunhAlgorithmChecker(creditCardNumber string) (bool, error) {
	if len(creditCardNumber) != 16 {
		return false, errors.New("Credit cart mus be 16 Digit")
	}

	var singleValues []int
	var secondValues []int

	for i := 0; i < len(creditCardNumber); i++ {
		if i%2 == 0 {
			val := int(creditCardNumber[i]-'0') * 2

			if val > 9 {
				firstDigit := val / 10
				secondDigit := val % 10

				singleValues = append(singleValues, int(firstDigit+secondDigit))
			} else {
				singleValues = append(singleValues, int(val))
			}
		} else {
			secondValues = append(secondValues, int(creditCardNumber[i]-'0'))
		}

	}

	sumOfValues := sumSlice(singleValues) + sumSlice(secondValues)

	return sumOfValues%10 == 0, nil

}
