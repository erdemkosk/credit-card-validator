![Logo](https://i.pinimg.com/736x/35/a7/d3/35a7d34dfd78c89cee2b80e0f1769cae.jpg)


# Luhn Algorithm

The Luhn algorithm, also known as the "modulus 10" or "mod 10" algorithm, is a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers, IMEI numbers, and more.

 [What is Luhn_algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm)

## How it Works

The algorithm verifies the validity of a given number through the following steps:

1. Starting from the rightmost digit, double the value of every second digit.
2. If the resulting doubled value is greater than 9, subtract 9 from it.
3. Sum all the digits together.
4. If the total sum is divisible by 10 without a remainder, the number is valid according to the Luhn algorithm.

Usage
This Go program demonstrates the application of the Luhn algorithm to validate credit card numbers and generate random valid credit card numbers.

How to Use
1. The generateCreditCard function creates a random 16-digit credit card number according to the Luhn algorithm.

2. The CheckCreditCardIsValid function checks the validity of a given credit card number using the Luhn algorithm.

Example 

```bash
  func main() {
	fmt.Println("-------------Lunch Algorithm test-------------")

	success, _ := CheckCreditCardIsValid("5388466965157906")

	fmt.Println("This 5388466965157906 credit card is Valid? ", success)

	generatedCreditCard := generateCreditCard()

	val, _ := CheckCreditCardIsValid(generatedCreditCard)

	fmt.Println("This "+generatedCreditCard+" genereted with system ! This credit card is Valid? ", val)
}
```
    

