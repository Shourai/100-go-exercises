package main

import "fmt"

func main() {

	var array []int

	for i := 2300; i <= 3200; i++ {
		if i%7 == 0 && i%5 != 0 {
			array = append(array, i)
		}
	}

	fmt.Println(array)
}
