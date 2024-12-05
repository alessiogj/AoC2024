package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func insertOrdered(slice []int, val int) []int {
	pos := 0
	for i, v := range slice {
		if v > val {
			pos = i
			break
		}
		pos = len(slice)
	}
	slice = append(slice[:pos], append([]int{val}, slice[pos:]...)...)
	return slice
}

func countTimesInOrderedList(slice []int, val int) int {
	count := 0
	for _, v := range slice {
		if v == val {
			count++
		}
		if v > val {
			break
		}
	}
	return count
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func Logic(data *os.File) {

	left := []int{}
	right := []int{}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), "   ")
		first, err := strconv.Atoi(res[0])
		if err != nil {
			log.Fatal(err)
		}
		second, err := strconv.Atoi(res[1])
		if err != nil {
			log.Fatal(err)
		}
		left = insertOrdered(left, first)
		right = insertOrdered(right, second)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var result1 int
	var result2 int

	for i := 0; i < len(left); i++ {
		result1 += absDiffInt(left[i], right[i])
		result2 += left[i] * countTimesInOrderedList(right, left[i])
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
