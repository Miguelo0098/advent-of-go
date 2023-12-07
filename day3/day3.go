package main

import (
	"flag"
	"fmt"

	"github.com/miguelo0098/advent-of-go/utils"
)

func main()  {

	fileName := flag.String("file", "data/data", "file name with the data to consume")

	flag.Parse()

	lines := utils.GetStringsFromFile(*fileName)

	for i := 0; i < len(lines); i++ {
		fmt.Print(lines[i])
	}

}