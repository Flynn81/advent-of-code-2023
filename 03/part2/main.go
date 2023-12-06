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

type number struct {
	value int
	row int
	startColumn int
	endColumn int
}

func parseSchematic(s []string) int {
	total := 0
	numbers := []number{}

	for r:=0; r<len(s); r++ {
		strNum := ""
		for c := 0; c<len(s[r]); c++ {
			if isNumber(string(s[r][c])) {
				strNum = strNum + string(s[r][c])
				
				// if r-1 > 0 && string(s[r-1][c]) != "." && !isNumber(string(s[r-1][c])) {
				// 	adjacent = true					
				// } else if r+1 < len(s) && string(s[r+1][c]) != "." && !isNumber(string(s[r+1][c])) {
				// 	adjacent = true					
				// } else if c-1> 0 && string(s[r][c-1]) != "." && !isNumber(string(s[r][c-1])) {
				// 	adjacent = true					
				// } else if c+1 < len(s[r]) && string(s[r][c+1]) != "." && !isNumber(string(s[r][c+1])) {
				// 	adjacent = true					
				// } else if r-1 > 0 && c+1 < len(s[r]) && string(s[r-1][c+1]) != "." && !isNumber(string(s[r-1][c+1])) {
				// 	adjacent = true					
				// } else if r-1 > 0 && c-1 > 0 && string(s[r-1][c-1]) != "." && !isNumber(string(s[r-1][c-1])) {
				// 	adjacent = true					
				// } else if r+1 < len(s) && c+1 < len(s[r]) && string(s[r+1][c+1]) != "." && !isNumber(string(s[r+1][c+1])) {
				// 	adjacent = true					
				// } else if r+1 < len(s) && c-1 > 0 && string(s[r+1][c-1]) != "." && !isNumber(string(s[r+1][c-1])) {
				// 	adjacent = true					
				// }
			} else {
				if strNum != "" {
					numbers = append(numbers, number{convertStringToInt(strNum),r,c-len(strNum),c})
					strNum = ""
				}
			}
			if c+1 == len(s[r]) {
				if strNum != "" {
					numbers = append(numbers, number{convertStringToInt(strNum),r,c-len(strNum),c})
				}
			}
		}
	}

	for r:=0; r<len(s); r++ {
		for c := 0; c<len(s[r]); c++ {
			if string(s[r][c]) == "*" {
				first := -1
				second := -1
				for _, n := range numbers {
					if n.value== 755 || n.value == 598 || n.value == 664{
						fmt.Println(n.value,c,">=",n.startColumn-1," && ",c,"<=",n.endColumn+1)
					}
					match := false
					if n.row == r {
						if n.startColumn == c+1 {
							match = true
						} else if n.endColumn == c-1 {
							match = true
						}
					} else if n.row + 1 == r || n.row - 1 == r {
						if c >= n.startColumn-1 && c <= n.endColumn {
							match = true
						}
					}
					if match {
						if first == -1 {
							first = n.value
						} else if second == -1 {
							second = n.value
						} else {
							first = -2
							second = -2
							fmt.Println("third number :(", n)
						}
					}
				}
				if first >= 0 && second >= 0 {
					fmt.Println("gear ", first, second, (first * second))
					total += first * second
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