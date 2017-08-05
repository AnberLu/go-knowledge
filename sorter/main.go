package main

import (
	"fmt"
	"sort"
)

func main() {
	readFile := "./unsorted.bat"
	writeFile := "./sorted.bat"
	values, err := ReadValues(readFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	sort.Ints(values)
	WriteValues(values, writeFile)
}
