# Luhn Algorithm

The Luhn algorithm, also known as the "modulus 10" or "mod 10" algorithm, is a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers, IMEI numbers, and more.

## How it Works

The algorithm verifies the validity of a given number through the following steps:

1. Starting from the rightmost digit, double the value of every second digit.
2. If the resulting doubled value is greater than 9, subtract 9 from it.
3. Sum all the digits together.
4. If the total sum is divisible by 10 without a remainder, the number is valid according to the Luhn algorithm.