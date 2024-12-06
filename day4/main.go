package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countXMAS(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	isXMAS := func(grid [][]rune, i, j int) bool {
		return grid[i][j] == 'A' &&
			((grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S') ||
				(grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M')) &&
			((grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S') ||
				(grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M'))
	}

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if isXMAS(grid, i, j) {
				count++
			}
		}
	}

	return count
}

func isMatch(segment, match []rune) bool {
	if len(segment) != len(match) {
		return false
	}
	for i := range match {
		if segment[i] != match[i] {
			return false
		}
	}
	return true
}

func reverse(runes []rune) []rune {
	n := len(runes)
	reversed := make([]rune, n)
	for i, r := range runes {
		reversed[n-1-i] = r
	}
	return reversed
}

func Logic(data *os.File) {

	scanner := bufio.NewScanner(data)
	result1 := 0
	result2 := 0

	var matrix [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	toMatch1 := "XMAS"
	matchRunes1 := []rune(toMatch1)

	// ->
	for _, row := range matrix {
		for j := 0; j <= len(row)-len(matchRunes1); j++ {
			if isMatch(row[j:j+len(matchRunes1)], matchRunes1) {
				result1++
			}
		}
	}

	// ←
	for _, row := range matrix {
		for j := len(row) - 1; j >= len(matchRunes1)-1; j-- {
			if isMatch(row[j-len(matchRunes1)+1:j+1], reverse(matchRunes1)) {
				result1++
			}
		}
	}

	// ↓
	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i <= len(matrix)-len(matchRunes1); i++ {
			var verticalSegment []rune
			for k := 0; k < len(matchRunes1); k++ {
				verticalSegment = append(verticalSegment, matrix[i+k][j])
			}
			if isMatch(verticalSegment, matchRunes1) {
				result1++
			}
		}
	}

	// ↑
	for j := 0; j < len(matrix[0]); j++ {
		for i := len(matrix) - 1; i >= len(matchRunes1)-1; i-- {
			var verticalSegment []rune
			for k := 0; k < len(matchRunes1); k++ {
				verticalSegment = append(verticalSegment, matrix[i-k][j])
			}
			if isMatch(verticalSegment, matchRunes1) {
				result1++
			}
		}
	}

	// ↘
	for i := 0; i <= len(matrix)-len(matchRunes1); i++ {
		for j := 0; j <= len(matrix[0])-len(matchRunes1); j++ {
			var diagonalSegment []rune
			for k := 0; k < len(matchRunes1); k++ {
				diagonalSegment = append(diagonalSegment, matrix[i+k][j+k])
			}
			if isMatch(diagonalSegment, matchRunes1) {
				result1++
			}
		}
	}

	// ↙
	for i := 0; i <= len(matrix)-len(matchRunes1); i++ {
		for j := len(matchRunes1) - 1; j < len(matrix[0]); j++ {
			var diagonalSegment []rune
			for k := 0; k < len(matchRunes1); k++ {
				diagonalSegment = append(diagonalSegment, matrix[i+k][j-k])
			}
			if isMatch(diagonalSegment, matchRunes1) {
				result1++
			}
		}
	}

	// ↖
	for i := len(matrix) - 1; i >= len(matchRunes1)-1; i-- {
		for j := len(matrix[0]) - 1; j >= len(matchRunes1)-1; j-- {
			var diagonalSegment []rune
			for k := 0; k < len(matchRunes1); k++ {
				diagonalSegment = append(diagonalSegment, matrix[i-k][j-k])
			}
			if isMatch(diagonalSegment, matchRunes1) {
				result1++
			}
		}
	}

	// ↗
	for i := len(matrix) - 1; i >= len(matchRunes1)-1; i-- {
		for j := 0; j <= len(matrix[0])-len(matchRunes1); j++ {
			var diagonalSegment []rune
			for k := 0; k < len(matchRunes1); k++ {
				diagonalSegment = append(diagonalSegment, matrix[i-k][j+k])
			}
			if isMatch(diagonalSegment, matchRunes1) {
				result1++
			}
		}
	}

	result2 = countXMAS(matrix)

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
