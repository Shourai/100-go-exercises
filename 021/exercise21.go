package main

import (
	"100-golang-exercises/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type direction struct {
	vertical   float64
	horizontal float64
}

func main() {
	input := utils.GetMultiLineInput()
	var coordinate direction

	for _, c := range input {
		d := strings.Split(c, " ")
		n, _ := strconv.ParseFloat(d[1], 64)
		if d[0] == "UP" {
			coordinate.vertical += n
		} else if d[0] == "DOWN" {
			coordinate.vertical -= n
		} else if d[0] == "RIGHT" {
			coordinate.horizontal += n
		} else if d[0] == "LEFT" {
			coordinate.horizontal -= n
		}
	}

	calculateDistance(coordinate)
}

func calculateDistance(coordinate direction) {
	distance := math.Sqrt(math.Pow(coordinate.horizontal, 2) + math.Pow(coordinate.vertical, 2))
	fmt.Println(math.Round(distance))
}
