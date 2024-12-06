package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"unicode"
)

type Coord struct {
	x, y int
}

type InfiniteGrid map[Coord]rune

func (c Coord) Translate(factor int) Coord {
	return Coord{c.x * factor, c.y * factor}
}

func (c Coord) Move(dx, dy int) Coord {
	return Coord{c.x + dx, c.y + dy}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func parseGrid(file string) InfiniteGrid {
	grid := make(InfiniteGrid)
	data, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, ch := range line {
			grid[Coord{x, y}] = ch
		}
		y++
	}
	return grid
}

func getAntennasByFrequency(grid InfiniteGrid) map[rune][]Coord {
	antennas := make(map[rune][]Coord)
	for coord, ch := range grid {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			antennas[ch] = append(antennas[ch], coord)
		}
	}
	return antennas
}

func allPairs(coords []Coord) [][2]Coord {
	var pairs [][2]Coord
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			pairs = append(pairs, [2]Coord{coords[i], coords[j]})
		}
	}
	return pairs
}

func contains(grid InfiniteGrid, coord Coord) bool {
	_, ok := grid[coord]
	return ok
}

func appendWhile(start Coord, dx, dy int, grid InfiniteGrid) []Coord {
	var result []Coord
	curr := start
	for {
		curr = curr.Move(dx, dy)
		if !contains(grid, curr) {
			break
		}
		result = append(result, curr)
	}
	return result
}

func part1(grid InfiniteGrid) int {
	antennas := getAntennasByFrequency(grid)
	resonants := make(map[Coord]bool)

	for _, locs := range antennas {
		for _, pair := range allPairs(locs) {
			a, b := pair[0], pair[1]
			antinode1 := a.Translate(2).Move(-b.x, -b.y)
			antinode2 := b.Translate(2).Move(-a.x, -a.y)

			if contains(grid, antinode1) {
				resonants[antinode1] = true
			}
			if contains(grid, antinode2) {
				resonants[antinode2] = true
			}
		}
	}

	return len(resonants)
}

func part2(grid InfiniteGrid) int {
	antennas := getAntennasByFrequency(grid)
	resonants := make(map[Coord]bool)

	for _, locs := range antennas {
		for _, antenna := range locs {
			resonants[antenna] = true
		}
	}

	for _, locs := range antennas {
		for _, pair := range allPairs(locs) {
			a, b := pair[0], pair[1]
			dx := b.x - a.x
			dy := b.y - a.y
			divisor := gcd(int(math.Abs(float64(dx))), int(math.Abs(float64(dy))))
			stepX := dx / divisor
			stepY := dy / divisor

			forward := appendWhile(b, stepX, stepY, grid)
			for _, f := range forward {
				resonants[f] = true
			}

			backward := appendWhile(a, -stepX, -stepY, grid)
			for _, b := range backward {
				resonants[b] = true
			}
		}
	}

	return len(resonants)
}

func main() {
	grid := parseGrid("input.txt")

	result1 := part1(grid)
	result2 := part2(grid)

	fmt.Printf("Result One: %d\n", result1)
	fmt.Printf("Result Two: %d\n", result2)
}
