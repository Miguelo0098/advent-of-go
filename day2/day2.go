package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseSet(set string) map[string]int {
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

func isValidSet(set map[string]int) bool  {
	return set["green"] <= MAX_GREEN && set["red"] <= MAX_RED && set["blue"] <= MAX_BLUE
}

func parseGame(gameRecord string) (sets []string, gameId int) {
	splittedStr := regexp.MustCompile(`\:`).Split(gameRecord, 2)
	
	gameIdString := regexp.MustCompile(`\d+`).FindString(splittedStr[0])
	value, _ :=strconv.ParseInt(gameIdString, 0, 64)

	sets = regexp.MustCompile(`\;`).Split(splittedStr[1], 100)
	gameId = int(value)

	return 	

}

func getValidGameValue(gameRecord string) int  {
	sets, gameId := parseGame(gameRecord)
	
	for i := 0; i < len(sets); i++ {
		if !isValidSet(parseSet(sets[i])) {
			return 0	
		}
	}
	
	return gameId 
}

func getGamePower(gameRecord string) int {
	sets, _ := parseGame(gameRecord)
	minSet := map[string]int{"green": 1, "blue": 1, "red": 1} 
	
	for i := 0; i < len(sets); i++ {
		set := parseSet(sets[i])
		minSet["green"] = max(minSet["green"], set["green"])
		minSet["red"] = max(minSet["red"], set["red"])
		minSet["blue"] = max(minSet["blue"], set["blue"])
	}
	
	return minSet["green"] * minSet["red"] * minSet["blue"] 
}

func getStringsFromFile(fileName string) []string {
	dat, err := os.ReadFile(fileName)
	check(err)

	re := regexp.MustCompile(`.*\n`)

	res := re.FindAllString(string(dat), -1)

	return res
}

func main() {
	sum := 0

	fileName := flag.String("file", "data/data", "file name with the data to consume")
	power := flag.Bool("power", false, "calculate games power instead")

	flag.Parse()

	gameRecords := getStringsFromFile(*fileName)

	for i := 0; i < len(gameRecords); i++ {
		if *power {
		sum += getGamePower(gameRecords[i])	
		}else{
			sum += getValidGameValue(gameRecords[i])
		}
	}

	fmt.Println(sum)
}
