package main

import (
	"os"
	"fmt"
	"regexp"
	"github.com/williamolojede/cs50-golang"
	"strings"
)

func main() {
	args := os.Args[1:]

	// ensure program is ran with only one arguments
	if len(args) != 1 {
		fmt.Println("Usage: ./crack hash")
		os.Exit(1)
	}

	// ensure all character in the key string is an alphabet
	isOnlyAlpha, _ := regexp.MatchString("^[a-zA-Z]+$", args[0])
	if !isOnlyAlpha {
		fmt.Println("Usage: ./crack hash")
		os.Exit(1)
	}

	// convert key upper case since [a-zA-z] is 0-25
	key := strings.ToUpper(args[0])
	keyIndex := 0
	plaintext := cs50.GetString("plaintext:  ")
	c := []byte(plaintext)

	for ix, ch :=  range c {
		if keyIndex > (len(key) - 1) {
			keyIndex = 0
		}
		// the shift for each character is its distance from 'A'
		shift := key[keyIndex] - 'A'
		newChar := shift + ch
		if 'a' <= ch && ch <= 'z' {
			if newChar > 'z' {
				newChar = 'a' + ((shift - 1) - ('z' - ch))
			}
			c[ix] = newChar
			keyIndex += 1
		} else if 'A' <= ch && ch <= 'Z' {
			if newChar > 'Z' {
				newChar = 'A' + ((shift - 1) - ('Z' - ch))
			}
			c[ix] = newChar
			keyIndex += 1
		}
	}
	ciphertext := string(c)
	fmt.Printf("ciphertext: %s\n", ciphertext)
}