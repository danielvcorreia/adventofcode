package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func partOne() int {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		line := s.Text()
		fmt.Println(line)

	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return 0
}

func partTwo() int {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		line := s.Text()
		fmt.Println(line)

		
	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return 0
}
