package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter binary sequence: ")
	input, _ := scanner.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	binaryList := strings.Split(input, ",")
	checkBinary(binaryList)

}

func checkBinary(list []string) {
	for _, binaryNumber := range list {
		number, _ := strconv.ParseInt(binaryNumber, 2, 64)
		if number%5 == 0 {
			fmt.Println(binaryNumber)
		}
	}
}
