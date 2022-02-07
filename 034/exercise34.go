package main

import (
	"fmt"
	"math"
)

func main() {
	squaredMap()

}

func squaredMap() {
	var sqMap = make(map[float64]float64)
	for i := 1.0; i <= 20; i++ {
		sqMap[i] = math.Pow(i, 2)
	}

	fmt.Println(sqMap)
}
