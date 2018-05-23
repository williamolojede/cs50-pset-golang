package main

import (
	"fmt"

	cs50 "github.com/williamolojede/cs50-golang"
)

func printloop(start, end int, str string) {
	for i := start; i < end; i++ {
		fmt.Print(str)
	}
}
func main() {
	x := cs50.GetInt("Prompt: ")

	for x < 0 || x > 23 {
		x = cs50.GetInt("Retry: ")
	}

	for r := 1; r <= x; r++ {
		printloop(0, x-r, " ")
		printloop(0, r, "#")
		printloop(0, 2, " ")
		printloop(0, r, "#")

		fmt.Println()
	}
}
