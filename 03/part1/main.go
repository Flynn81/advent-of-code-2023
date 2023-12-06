package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//https://adventofcode.com/2023/day/3

func main() {
	fmt.Println(sum("input"))
}

func sum(filename string) int {

	readFile, err := os.Open(filename)
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

	return parseSchematic(lines)
}

func parseSchematic(s []string) int {
	total := 0
	for r:=0; r<len(s); r++ {
		strNum := ""
		adjacent := false
		for c := 0; c<len(s[r]); c++ {
			if isNumber(string(s[r][c])) {
				strNum = strNum + string(s[r][c])
				//check for adjacent symbol
				if r-1 > 0 && string(s[r-1][c]) != "." && !isNumber(string(s[r-1][c])) {
					adjacent = true					
				} else if r+1 < len(s) && string(s[r+1][c]) != "." && !isNumber(string(s[r+1][c])) {
					adjacent = true					
				} else if c-1> 0 && string(s[r][c-1]) != "." && !isNumber(string(s[r][c-1])) {
					adjacent = true					
				} else if c+1 < len(s[r]) && string(s[r][c+1]) != "." && !isNumber(string(s[r][c+1])) {
					adjacent = true					
				} else if r-1 > 0 && c+1 < len(s[r]) && string(s[r-1][c+1]) != "." && !isNumber(string(s[r-1][c+1])) {
					adjacent = true					
				} else if r-1 > 0 && c-1 > 0 && string(s[r-1][c-1]) != "." && !isNumber(string(s[r-1][c-1])) {
					adjacent = true					
				} else if r+1 < len(s) && c+1 < len(s[r]) && string(s[r+1][c+1]) != "." && !isNumber(string(s[r+1][c+1])) {
					adjacent = true					
				} else if r+1 < len(s) && c-1 > 0 && string(s[r+1][c-1]) != "." && !isNumber(string(s[r+1][c-1])) {
					adjacent = true					
				}
			} else {
				if strNum != "" {
					if adjacent {
						total += convertStringToInt(strNum)
					}
					strNum = ""
					adjacent = false
				}
			}
			if c+1 == len(s[r]) {
				if strNum != "" {
					if adjacent {
						total += convertStringToInt(strNum)
					}
				}
			}
		}
	}
	return total
}

func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func convertStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}