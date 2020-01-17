package main

import (
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	ans := ""

	phoneNumber := args[0]

	format := args[1]

	onlyX := strings.ReplaceAll(format, "-", "")

	if len(onlyX) != len(phoneNumber) {
		panic("Wrong input x character must be equal input number")
	}

	checkIndex := 0

	for i := 0; i < len(format); i++ {
		if strings.ToLower(string(format[i])) == "x" {
			ans += string(phoneNumber[i-checkIndex])
		} else {
			checkIndex += 1
			ans += "-"
		}
	}

	print(ans)
}
