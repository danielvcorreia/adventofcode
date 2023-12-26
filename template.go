package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := parse(os.Stdin)

	fmt.Println(partOne(a))
	fmt.Println(partTwo(a))
}

func partOne(a int) int {
	return a
}

func partTwo(a int) int {
	return a
}

func parse(file *os.File) int {
	s := bufio.NewScanner(file)

	for s.Scan() {
		line := s.Text()
		fmt.Println(line)
	}

	return 0
}
