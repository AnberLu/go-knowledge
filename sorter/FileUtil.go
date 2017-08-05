package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadValues(inFile string) (values []int, err error) {
	file, err := os.Open(inFile)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := reader.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("tow long")
			return
		}

		str := string(line)
		if str == "" {
			break
		}
		value, err2 := strconv.Atoi(str)
		if err2 != nil {
			err = err2
			return
		}

		values = append(values, value)
	}
	return
}

func WriteValues(values []int, outFile string) error {
	file, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}
