package main

import (
	"100-golang-exercises/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.GetUserInput()
	slice := strings.Split(input, " ")
	getNetAmount(slice)
}

func getNetAmount(input []string) {
	netAmount := 0

	for idx, transaction := range input {
		if transaction == "D" {
			amount, _ := strconv.Atoi(input[idx+1])
			netAmount += amount
		} else if transaction == "W" {
			amount, _ := strconv.Atoi(input[idx+1])
			netAmount -= amount
		}
	}

	fmt.Println(netAmount)

}
