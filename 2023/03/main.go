package main

import (
	"bufio"
	"fmt"
	"os"
)

type grid struct{ i, j int }

func parse() (map[grid]byte, int) {
	s := bufio.NewScanner(os.Stdin)
	m := make(map[grid]byte)
	size := 0

	for i := 0; s.Scan(); i++ {
		line := s.Text()

		for j := 0; j < len(line); j++ {
			m[grid{i, j}] = line[j]
		}

		size = len(line)
	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return m, size
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func valid(puzzle map[grid]byte, s grid, e grid) bool {
	for i := s.i - 1; i <= e.i+1; i++ {
		for j := s.j - 1; j <= e.j+1; j++ {
			currentByte, exists := puzzle[grid{i, j}]

			if exists && !isDigit(currentByte) && currentByte != '.' {
				return true
			}
		}
	}
	return false
}

func getPart(puzzle map[grid]byte, c grid) (int, int) {
	index := grid{c.i, c.j - 1}
	part := 0
	var sDigitIndex grid
	var eDigitIndex grid

	for isDigit(puzzle[index]) {
		index = grid{index.i, index.j - 1}
	}
	sDigitIndex = grid{index.i, index.j + 1}

	index = grid{c.i, c.j + 1}
	for isDigit(puzzle[index]) {
		index = grid{index.i, index.j + 1}
	}
	eDigitIndex = grid{index.i, index.j - 1}

	for j := sDigitIndex.j; j <= eDigitIndex.j; j++ {
		part = part*10 + int(puzzle[grid{c.i, j}]-'0')
	}

	return part, eDigitIndex.j
}

func gearRatio(puzzle map[grid]byte, c grid) int {
	parts := []int{}

	for i := c.i - 1; i <= c.i+1; i++ {
		for j := c.j - 1; j <= c.j+1; j++ {
			currentByte, exists := puzzle[grid{i, j}]

			if exists && isDigit(currentByte) {
				part, jCurrent := getPart(puzzle, grid{i, j})
				parts = append(parts, part)
				j = jCurrent
			}

			if len(parts) > 2 {
				return 0
			}
		}
	}

	if len(parts) < 2 {
		return 0
	} else {
		return parts[0] * parts[1]
	}
}

func partOne(puzzle map[grid]byte, size int) int {
	sum := 0

	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < size; j++ {
			byteIndex := grid{i, j}
			currentByte := puzzle[byteIndex]

			if isDigit(currentByte) {
				s_index := grid{i, j}
				e_index := grid{0, 0}
				number := 0

				for isDigit(currentByte) {
					number = number*10 + int(currentByte-'0')
					e_index = byteIndex
					j++
					byteIndex = grid{i, j}
					currentByte = puzzle[byteIndex]
				}

				if valid(puzzle, s_index, e_index) {
					sum += number
				}
			}
		}
	}

	return sum
}

func partTwo(puzzle map[grid]byte, size int) int {
	sum := 0

	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < size; j++ {
			byteIndex := grid{i, j}
			currentByte := puzzle[byteIndex]

			if currentByte == '*' {
				sum += gearRatio(puzzle, byteIndex)
			}
		}
	}

	return sum
}

func main() {
	puzzle, size := parse()

	fmt.Printf("%d\n", partOne(puzzle, size))
	fmt.Printf("%d\n", partTwo(puzzle, size))
}
