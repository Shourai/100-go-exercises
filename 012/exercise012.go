package main

import (
	"fmt"
	"strings"
)

func main() {

	list := []string{}

	for i := 1000; i <= 3000; i++ {
		if i%2 == 0 {
			list = append(list, fmt.Sprint(i))
		}
	}

	numbers := strings.Join(list, ",")
	fmt.Println(numbers)

}
