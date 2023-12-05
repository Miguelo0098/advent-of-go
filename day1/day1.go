package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getNumberSum(str string) int {
	re := regexp.MustCompile(`\d`)
	numStrArray := re.FindAllString(str, -1)

	if len(numStrArray) == 0 {
		return 0
	}

	first, _ := strconv.ParseInt(numStrArray[0], 0, 64)
	last, _ := strconv.ParseInt(numStrArray[len(numStrArray)-1], 0, 64)

	return int(first*10 + last)
}

func getStringsFromFile(fileName string) []string {
	dat, err := os.ReadFile(fileName)
	check(err)

	re := regexp.MustCompile(`.*\n`)

	return re.FindAllString(string(dat), -1)
}

func main() {
	sum := 0
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No file provided as argument (example: go run day1.go <file-with-data>)")
		os.Exit(42069)
	}

	stringArray := getStringsFromFile(args[0])

	for i := 0; i < len(stringArray); i++ {
		sum += getNumberSum(stringArray[i])
	}

	fmt.Println(sum)
}
