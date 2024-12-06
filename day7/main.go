package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleLinePointTwo(line string) int {
	splitValues := strings.Split(line, ": ")
	if len(splitValues) != 2 {
		log.Fatalf("Invalid line format: %s", line)
	}

	testResult, err := strconv.Atoi(splitValues[0])
	check(err)

	numStrs := strings.Fields(splitValues[1])
	numbers := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		numbers[i], err = strconv.Atoi(numStr)
		check(err)
	}

	results := []int{numbers[0]}
	for _, value := range numbers[1:] {
		results = addMultiplyConcatenate(results, value)
	}

	for _, result := range results {
		if result == testResult {
			return testResult
		}
	}

	return 0
}

func handleLinePointOne(line string) int {
	splitValues := strings.Split(line, ": ")
	if len(splitValues) != 2 {
		log.Fatalf("Invalid line format: %s", line)
	}

	testResult, err := strconv.Atoi(splitValues[0])
	check(err)

	numStrs := strings.Fields(splitValues[1])
	numbers := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		numbers[i], err = strconv.Atoi(numStr)
		check(err)
	}

	totals := []int{numbers[0]}
	numbers = numbers[1:]

	for _, value := range numbers {
		totals = addAndMultiply(totals, value)
	}

	for _, total := range totals {
		if total == testResult {
			return testResult
		}
	}

	return 0
}

func addMultiplyConcatenate(results []int, value int) []int {
	newResults := make([]int, 0, len(results)*3)

	for _, current := range results {
		sum := current + value
		newResults = append(newResults, sum)

		product := current * value
		newResults = append(newResults, product)

		concatenated := concatenate(current, value)
		newResults = append(newResults, concatenated)
	}

	return append(results, newResults...)
}

func addAndMultiply(totals []int, value int) []int {
	newTotals := make([]int, 0, len(totals)*2)

	for _, total := range totals {
		newTotals = append(newTotals, total+value)

		newTotals = append(newTotals, total*value)
	}

	return append(totals, newTotals...)
}

func concatenate(a, b int) int {
	concatStr := strconv.Itoa(a) + strconv.Itoa(b)
	result, err := strconv.Atoi(concatStr)
	if err != nil {
		log.Fatalf("Error concatenating %d and %d", a, b)
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result1, result2 int
	for scanner.Scan() {
		line := scanner.Text()
		result1 += handleLinePointOne(line)
		result2 += handleLinePointTwo(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Result One: ", result1)
	fmt.Println("Result Two: ", result2)
}
