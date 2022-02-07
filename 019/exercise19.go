package main

import (
	"100-golang-exercises/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type person struct {
	name  string
	age   int
	score int
}

func main() {
	input := utils.GetSingleLineInput()
	slice := strings.Split(input, " ")

	var people []person

	for _, p := range slice {
		var person person
		info := strings.Split(p, ",")
		person.name = info[0]
		person.age, _ = strconv.Atoi(info[1])
		person.score, _ = strconv.Atoi(info[2])
		people = append(people, person)
	}

	sort.Slice(people, func(i, j int) bool {
		if people[i].name < people[j].name {
			return true
		} else if people[i].name > people[j].name {
			return false
		} else if people[i].age > people[j].age {
			return true
		}
		return people[i].score > people[j].score
	})

	fmt.Println(people)

}
