package utils

import (
	"os"
	"regexp"
)

func GetStringsFromFile(fileName string) []string {
	dat, err := os.ReadFile(fileName)
	
	Check(err)

	re := regexp.MustCompile(`.*\n`)

	res := re.FindAllString(string(dat), -1)

	return res
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}