package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var directions = []struct {
	dx, dy int
}{
	{0, -1}, // Up (^)
	{1, 0},  // Right (>)
	{0, 1},  // Down (v)
	{-1, 0}, // Left (<)
}

func ParseMap(lines []string) ([][]rune, int, int, int) {
	var grid [][]rune
	var guardX, guardY, direction int

	for y, line := range lines {
		row := []rune(line)
		for x, char := range row {
			if char == '^' {
				guardX, guardY, direction = x, y, 0
				row[x] = '.'
			} else if char == '>' {
				guardX, guardY, direction = x, y, 1
				row[x] = '.'
			} else if char == 'v' {
				guardX, guardY, direction = x, y, 2
				row[x] = '.'
			} else if char == '<' {
				guardX, guardY, direction = x, y, 3
				row[x] = '.'
			}
		}
		grid = append(grid, row)
	}
	return grid, guardX, guardY, direction
}

func SimulateGuardPart1(grid [][]rune, startX, startY, startDirection int) int {
	visited := make(map[string]bool)
	x, y, direction := startX, startY, startDirection

	markVisited := func(x, y int) {
		key := fmt.Sprintf("%d,%d", x, y)
		visited[key] = true
	}

	markVisited(x, y)

	for {
		dx, dy := directions[direction].dx, directions[direction].dy
		nx, ny := x+dx, y+dy

		if ny < 0 || ny >= len(grid) || nx < 0 || nx >= len(grid[ny]) {
			break
		}

		if grid[ny][nx] == '#' {
			direction = (direction + 1) % 4
		} else {
			x, y = nx, ny
			markVisited(x, y)
		}
	}

	return len(visited)
}

func SimulateGuardPart2(grid [][]rune, startX, startY, startDirection int) bool {
	visited := make(map[string]bool)
	x, y, direction := startX, startY, startDirection

	for {
		stateKey := fmt.Sprintf("%d,%d,%d", x, y, direction)
		if visited[stateKey] {
			return true
		}
		visited[stateKey] = true

		dx, dy := directions[direction].dx, directions[direction].dy
		nx, ny := x+dx, y+dy

		if ny < 0 || ny >= len(grid) || nx < 0 || nx >= len(grid[ny]) {
			break
		}

		if grid[ny][nx] == '#' {
			direction = (direction + 1) % 4
		} else {
			x, y = nx, ny
		}
	}

	return false
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

	grid, startX, startY, startDirection := ParseMap(lines)

	distinctPositions := SimulateGuardPart1(grid, startX, startY, startDirection)
	fmt.Println("Part One:", distinctPositions)

	possiblePositions := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != '.' || (x == startX && y == startY) {
				continue
			}

			grid[y][x] = '#'

			if SimulateGuardPart2(grid, startX, startY, startDirection) {
				possiblePositions++
			}

			grid[y][x] = '.'
		}
	}

	fmt.Println("Part Two:", possiblePositions)
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
