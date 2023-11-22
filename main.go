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

	generatedCreditCard := GenerateCreditCard()

	val, _ := CheckCreditCardIsValid(generatedCreditCard)

	fmt.Println("This "+generatedCreditCard+" genereted with system ! This credit card is Valid? ", val)

}

func SumSlice(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}

func LunhAlgorithmChecker(creditCardNumber string) ([]int, []int) {

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

func GenerateCreditCard() string {
	var result strings.Builder

	for i := 0; i < 15; i++ {
		result.WriteString(fmt.Sprintf("%d", rand.Intn(9)))
	}

	oddOrderValues, evenOrderValues := LunhAlgorithmChecker(result.String())

	sumOfVal := SumSlice(oddOrderValues) + SumSlice(evenOrderValues)

	validatorDigit := 10 - sumOfVal%10

	result.WriteString(fmt.Sprintf("%d", validatorDigit))

	return result.String()
}

func CheckCreditCardIsValid(creditCardNumber string) (bool, error) {
	if len(creditCardNumber) != 16 {
		return false, errors.New("Credit cart mus be 16 Digit")
	}

	oddOrderValues, evenOrderValues := LunhAlgorithmChecker(creditCardNumber)

	sumOfValues := SumSlice(oddOrderValues) + SumSlice(evenOrderValues)

	return sumOfValues%10 == 0, nil
}
