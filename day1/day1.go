package main

import (
	"flag"
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

func parseNumber(str string) int {

	numMap := map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	num, err := strconv.ParseInt(str, 0, 64)

	if err != nil {
		num = numMap[str]
	}

	return int(num)
}

func getNumberSum(str string, verbose bool) int {
	var regexString string
	if verbose {
		regexString = `(\d|one|two|three|four|five|six|seven|eight|nine)`

	} else {
		regexString = `\d`
	}

	re := regexp.MustCompile(regexString)

	// this helps to catch cases like 'oneight' or 'sevenine'
	// Go regex does not support lookarounds, which will make this way easier
	fixedStr := re.ReplaceAllStringFunc(str, func(s string) string {
		return fmt.Sprintf("%s%s", s, s[len(s)-1:])
	})

	numStrArray := re.FindAllString(fixedStr, -1)

	if len(numStrArray) == 0 {
		return 0
	}

	res := parseNumber(numStrArray[0])*10 + parseNumber(numStrArray[len(numStrArray)-1])

	return res
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
	verbose := flag.Bool("verbose", true, "considers 'one', 'two', 'three'... as digits")

	flag.Parse()

	stringArray := getStringsFromFile(*fileName)

	for i := 0; i < len(stringArray); i++ {
		sum += getNumberSum(stringArray[i], *verbose)
	}

	fmt.Println(sum)
}
