package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

//https://adventofcode.com/2023/day/1

func main() {
	readFile, err := os.Open("input")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	lines := []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	fmt.Println(sumValues(lines)) //54953
}

func sumValues(lines []string) int {
	total := 0
	for _, line := range lines {
		total = total + decodeLine(line)
	}
	return total
}

func decodeLine(line string) int {
	// line = strings.ToUpper(line)
	// line = strings.ReplaceAll(line, "ONE", "O1E")
	// line = strings.ReplaceAll(line, "TWO", "T2O")
	// line = strings.ReplaceAll(line, "THREE", "TH3EE")
	// line = strings.ReplaceAll(line, "FOUR", "FO4R")
	// line = strings.ReplaceAll(line, "FIVE", "FI5E")
	// line = strings.ReplaceAll(line, "SIX", "S6X")
	// line = strings.ReplaceAll(line, "SEVEN", "SE7EN")
	// line = strings.ReplaceAll(line, "EIGHT", "EI8HT")
	// line = strings.ReplaceAll(line, "NINE", "N9NE")
	regex := regexp.MustCompile("[0-9]+")
	numbers := regex.FindAllString(line, -1)
	if numbers == nil {
		return 0
	}
	length := len(numbers)
	first := string(numbers[0][0])
	last := ""
	if length == 1 {
		if(len(numbers[0]) > 1) {
			last = string(numbers[0][len(numbers[length-1])-1])
		}
	} else {
		last = string(numbers[length-1][len(numbers[length-1])-1])
	}
	if last == "" {
		last = first
	}
	return convertStringToInt(first + last)
}

func convertStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}