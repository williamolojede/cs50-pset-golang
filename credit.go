package main

import (
	"github.com/williamolojede/cs50-golang"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	card := cs50.GetString("Number: ")
	cardLength := len(card)

	var (
		cardArrayOfNumbers []int
		sum1, sum2 int
		)

	// test if card is a valid positive number string
	_, err := strconv.ParseUint(card, 10, 64)
	if err != nil {
		fmt.Println("INVALID")
		return
	}

	// test if the card is of required length
	if cardLength != 13 && cardLength != 15 && cardLength != 16 {
		fmt.Println("INVALID")
		return
	}

	// split the card into an array
	cardArrayOfStrings := strings.Split(card, "")

	// create an array of the card numbers
	for _, ch := range cardArrayOfStrings {
		num, _ := strconv.Atoi(ch)
		cardArrayOfNumbers = append(cardArrayOfNumbers, num)
	}

	/*
		STEP 0:
		Multiply every other digit by 2,
		starting with the number’s second-to-last digit,
		and then add those products' digits together.
	*/
	// grab every second number from the second to the last
	for i := cardLength - 2; i >= 0; i -= 2 {
		// multiply each by two
		times2 := cardArrayOfNumbers[i] * 2
		// get the first and second digit of each number
		firsDigit, secondDigit := times2 / 10, times2 % 10
		// checks if each number after been multiplied by 2 is a two digit number
		if  secondDigit != times2 {
			sum2 += firsDigit + secondDigit
		} else {
			sum2 += times2
		}
	}

	/*
		STEP 1:
		Add the sum to the sum of the digits that weren’t multiplied by 2.
	*/
	// sum up of every other number from the last
	for i := cardLength - 1; i >= 0; i -= 2 {
		sum1 += cardArrayOfNumbers[i]
	}


	// test if the last digit of sum1 + sum2 is 0
	if (sum1 + sum2) % 10 == 0 {
		amex := strings.HasPrefix(card, "37")
		visa := strings.HasPrefix(card, "4")
		mastercard := strings.HasPrefix(card, "53")

		if amex {
			fmt.Println("AMEX")
		} else if visa{
			fmt.Println("VISA")
		} else if mastercard {
			fmt.Println("MASTERCARD")
		} else {
			fmt.Println("INVALID")
		}
	} else {
		fmt.Println("INVALID")
	}
}