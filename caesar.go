package main

import (
	"os"
	"fmt"
	"strconv"
	"github.com/williamolojede/cs50-golang"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("One cli argument required")
		os.Exit(1)
	}

	key, _ := strconv.ParseUint(args[0], 10, 8)


	if key > 26 {
		key = key % 26
	}

	plaintext := cs50.GetString("plaintext: ")

	c := []byte(plaintext)

	for ix, ch := range c {
		newChar := uint8(key) + ch
		if 'a' <= ch && ch <= 'z' {
			if newChar > 'z' {
				newChar = 'a' + ((uint8(key) - 1) - ('z' - ch))
			}
			c[ix] = newChar
		} else if 'A' <= ch && ch <= 'Z' {
			if newChar > 'Z' {
				newChar = 'A' + ((uint8(key) - 1) - ('Z' - ch))
			}
			c[ix] = newChar
		}
	}

	ciphertext := string(c)

	fmt.Printf("ciphertext: %s\n", ciphertext)
}