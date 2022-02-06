package main

import (
	"100-golang-exercises/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	input := utils.GetUserInput()
	slice := strings.Split(input, ",")

	checkPassword(slice)

}

func checkPassword(input []string) {
	for _, password := range input {

		regexLowercase, _ := regexp.Compile("[a-z]+")
		matchLowercase := regexLowercase.MatchString(password)

		regexUppercase, _ := regexp.Compile("[A-Z]+")
		matchUppercase := regexUppercase.MatchString(password)

		regexNumber, _ := regexp.Compile("[0-9]+")
		matchNumber := regexNumber.MatchString(password)

		regexSymbol, _ := regexp.Compile("[#@!]+")
		matchSymbol := regexSymbol.MatchString(password)

		length := len(password)

		if matchLowercase && matchUppercase && matchNumber && matchSymbol && length >= 6 && length <= 12 {
			fmt.Print(password)
		}
	}

}
