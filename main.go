package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println("-------------Lunch Algorithm test-------------")

	success, _ := CheckCreditCardIsValid("5388466965157906")

	fmt.Println("This 5388466965157906 credit card is Valid? ", success)

	generatedCreditCard := generateCreditCard()

	val, _ := CheckCreditCardIsValid(generatedCreditCard)

	fmt.Println("This "+generatedCreditCard+" genereted with system ! This credit card is Valid? ", val)

}

func sumSlice(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}

func lunhAlgorithmChecker(creditCardNumber string) ([]int, []int) {

	var oddOrderValues []int
	var evenOrderValues []int

	for i := 0; i < len(creditCardNumber); i++ {
		if i%2 == 0 {
			val := int(creditCardNumber[i]-'0') * 2

			if val > 9 {
				firstDigit := val / 10
				secondDigit := val % 10

				oddOrderValues = append(oddOrderValues, int(firstDigit+secondDigit))
			} else {
				oddOrderValues = append(oddOrderValues, int(val))
			}
		} else {
			evenOrderValues = append(evenOrderValues, int(creditCardNumber[i]-'0'))
		}
	}

	return oddOrderValues, evenOrderValues
}

func generateCreditCard() string {
	var result strings.Builder

	for i := 0; i < 15; i++ {
		result.WriteString(fmt.Sprintf("%d", rand.Intn(9)))
	}

	oddOrderValues, evenOrderValues := lunhAlgorithmChecker(result.String())

	sumOfVal := sumSlice(oddOrderValues) + sumSlice(evenOrderValues)

	validatorDigit := 10 - sumOfVal%10

	result.WriteString(fmt.Sprintf("%d", validatorDigit))

	return result.String()
}

func CheckCreditCardIsValid(creditCardNumber string) (bool, error) {
	if len(creditCardNumber) != 16 {
		return false, errors.New("Credit cart mus be 16 Digit")
	}

	oddOrderValues, evenOrderValues := lunhAlgorithmChecker(creditCardNumber)

	sumOfValues := sumSlice(oddOrderValues) + sumSlice(evenOrderValues)

	return sumOfValues%10 == 0, nil
}
