package main

import (
	"100-golang-exercises/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := utils.GetUserInput()
	slice := strings.Split(input, ",")

	sliceOfOddNumbers := make([]string, 0)

	for _, i := range slice {
		num, _ := strconv.Atoi(i)
		if num%2 != 0 {
			squared := strconv.Itoa(num * num)
			sliceOfOddNumbers = append(sliceOfOddNumbers, squared)
		}
	}

	fmt.Println(strings.Join(sliceOfOddNumbers, ","))
}
