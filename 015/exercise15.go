package main

import (
	"100-golang-exercises/utils"
	"fmt"
	"strconv"
)

func main() {
	input := utils.GetUserInput()
	integer, _ := strconv.Atoi(input)
	compute(integer)

}

func compute(num int) {
	fmt.Print(1*num + 11*num + 111*num + 1111*num)
}
