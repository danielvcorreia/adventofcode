package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse() int {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		line := s.Text()
		fmt.Println(line)
	}

	return 0
}

func partOne(a int) int {
	return a
}

func partTwo(a int) int {
	return a
}

func main() {
	a := parse()

	fmt.Println(partOne(a))
	fmt.Println(partTwo(a))
}
