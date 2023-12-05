package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func getSetMap(set string) map[string]int {
	balls := regexp.MustCompile(`\,`).Split(set, 3)

	setMap := make(map[string]int)

	for i := 0; i < len(balls); i++ {
		color := regexp.MustCompile(`red|blue|green`).FindString(balls[i])
		amount := regexp.MustCompile(`\d+`).FindString(balls[i])

		amountVal, _ := strconv.ParseInt(amount, 0, 64)

		setMap[color] = int(amountVal)
	}

	return setMap
}

func main() {

	gameRecord := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	splittedStr := regexp.MustCompile(`\:`).Split(gameRecord, 2)

	gameId := splittedStr[0]
	setData := splittedStr[1]

	setsArray := regexp.MustCompile(`\;`).Split(setData, 100)

	for i := 0; i < len(setsArray); i++ {
		fmt.Println(getSetMap(setsArray[i]))
	}

	fmt.Println(gameId)
}
