package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Logic(data *os.File) {
	scanner := bufio.NewScanner(data)
	result1 := 0
	result2 := 0
	enabled := true

	// Regex per rilevare `do()`, `don't()` e `mul(x, y)`
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)

	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
			} else if match[0] == "don't()" {
				enabled = false
			} else if len(match) > 2 && enabled {
				val1, err1 := strconv.Atoi(match[1])
				val2, err2 := strconv.Atoi(match[2])
				if err1 != nil || err2 != nil {
					continue
				}
				result1 += val1 * val2
				result2 += val1 * val2
			} else {
				val1, err1 := strconv.Atoi(match[1])
				val2, err2 := strconv.Atoi(match[2])
				if err1 != nil || err2 != nil {
					continue
				}
				result1 += val1 * val2
			}
		}
	}

	fmt.Println("Result One: ", result1)
	fmt.Println("Result Two: ", result2)
}

func ReadFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	Logic(file)
}

func main() {
	ReadFile("input.txt")
}
