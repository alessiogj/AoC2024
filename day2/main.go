package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafeDistance(val1, val2 int) bool {
	return absDiffInt(val1, val2) >= 1 && absDiffInt(val1, val2) <= 3
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func Logic(data *os.File) {

	scanner := bufio.NewScanner(data)
	result1 := 0
	result2 := 0
	for scanner.Scan() {

		values := []int{}

		res := strings.Split(scanner.Text(), " ")

		for _, v := range res {
			val, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			values = append(values, val)
		}

		isInc := values[0] > values[1]

		isValid := true

		for i := 0; i < len(values)-1; i++ {
			if isInc {
				if values[i] < values[i+1] {
					isValid = false
				}
			} else {
				if values[i] > values[i+1] {
					isValid = false
				}
			}
			if !isSafeDistance(values[i], values[i+1]) {
				isValid = false
			}
		}

		if isValid {
			result1++
			result2++
		} else {
			for j := 0; j < len(values); j++ {

				values1 := make([]int, len(values))
				copy(values1, values)

				values1 = append(values1[:j], values1[j+1:]...)

				isInc := values1[0] > values1[1]

				isValid := true

				for i := 0; i < len(values1)-1; i++ {
					if isInc {
						if values1[i] < values1[i+1] {
							isValid = false
							break
						}
					} else {
						if values1[i] > values1[i+1] {
							isValid = false
							break
						}
					}
					if !isSafeDistance(values1[i], values1[i+1]) {
						isValid = false
						break
					}
				}

				if isValid {
					result2++
					break
				}
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
