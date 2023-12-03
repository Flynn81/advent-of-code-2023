package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//https://adventofcode.com/2023/day/2

func main() {
	fmt.Println(sumIDs("input"))
}

func sumIDs(filename string) int {

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

	total := 0
	for _, line := range lines {
		total = total + parseLine(line)
	}
	return total
}

func parseLine(line string) int {
	id := extractFirstNumber(line)

	redCubes := 12
	greenCubes := 13
	blueCubes := 14

	game := strings.Split(line, ":")
	reveals := strings.Split(game[1], ";")
	red := true
	blue := true
	green := true
	for _, reveal := range reveals {
		cubes := strings.Split(reveal, ",")
		for _, color := range cubes {
			n := extractFirstNumber(color)
			if strings.Contains(color, "red") {
				if n > redCubes {
					red = false
				}
			} else if strings.Contains(color, "blue") {
				if n > blueCubes {
					blue = false
				}
			} else if strings.Contains(color, "green") {
				if n > greenCubes {
					green = false
				}
			}
		}
	}
	
	if red && green && blue {
		return id
	}
	return 0
}

func extractFirstNumber(s string) int {
	regex := regexp.MustCompile("[0-9]+")
	return convertStringToInt(string(regex.FindAllString(s, -1)[0]))
}

func convertStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}