package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseRules(lines []string) map[int][]int {
	rules := make(map[int][]int)
	for _, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			continue
		}
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])
		rules[from] = append(rules[from], to)
	}
	return rules
}

func IsValidUpdate(update []int, rules map[int][]int) bool {
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}
	for from, tos := range rules {
		for _, to := range tos {
			fromPos, fromExists := position[from]
			toPos, toExists := position[to]
			if fromExists && toExists && fromPos >= toPos {
				return false
			}
		}
	}
	return true
}

func FindMiddle(update []int) int {
	return update[len(update)/2]
}

func CorrectOrder(update []int, rules map[int][]int) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	for _, page := range update {
		graph[page] = []int{}
		inDegree[page] = 0
	}
	for from, tos := range rules {
		for _, to := range tos {
			if _, existsFrom := graph[from]; existsFrom {
				if _, existsTo := graph[to]; existsTo {
					graph[from] = append(graph[from], to)
					inDegree[to]++
				}
			}
		}
	}

	// Topological sort
	var queue []int
	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}
	var sorted []int
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		sorted = append(sorted, current)
		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return sorted
}

func Logic(data *os.File) {
	scanner := bufio.NewScanner(data)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var rulesSection []string
	var updatesSection [][]int
	readingRules := true
	for _, line := range lines {
		if line == "" {
			readingRules = false
			continue
		}
		if readingRules {
			rulesSection = append(rulesSection, line)
		} else {
			pages := strings.Split(line, ",")
			var update []int
			for _, page := range pages {
				num, _ := strconv.Atoi(page)
				update = append(update, num)
			}
			updatesSection = append(updatesSection, update)
		}
	}

	rules := ParseRules(rulesSection)

	sumCorrect := 0
	sumFixed := 0
	for _, update := range updatesSection {
		if IsValidUpdate(update, rules) {
			sumCorrect += FindMiddle(update)
		} else {
			corrected := CorrectOrder(update, rules)
			sumFixed += FindMiddle(corrected)
		}
	}

	fmt.Println("Result One: ", sumCorrect)
	fmt.Println("Result Two: ", sumFixed)
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
