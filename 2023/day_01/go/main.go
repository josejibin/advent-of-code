package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var digitsRe = regexp.MustCompile("[0-9]+")

func getTotalCalibrated(fn string) int {
	// read a file and return a string
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	s := 0
	for scanner.Scan() {
		s += getCalibrationVal(scanner.Text())
	}
	return s
}

func getCalibrationVal(s string) int {
	matches := digitsRe.FindAllString(s, -1)
	if len(matches) == 0 {
		// lets panic
		panic("No digits found in string")
	}
	lastElm := matches[len(matches)-1]

	val := string(matches[0][0]) + string(lastElm[len(lastElm)-1])
	i, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	fmt.Println(getTotalCalibrated("../input.txt"))
}
