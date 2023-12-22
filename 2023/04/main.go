package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	id       int
	Wnumbers map[int]bool
	numbers  []int
	points   int
}

func parse() map[int]Card {
	s := bufio.NewScanner(os.Stdin)
	cards := make(map[int]Card)

	for s.Scan() {
		line := s.Text()
		card := Card{
			id:       0,
			Wnumbers: make(map[int]bool),
			numbers:  []int{},
			points:   0,
		}
		_, err := fmt.Sscanf(line, "Card %d", &card.id)

		if err != nil {
			panic(err)
		}

		numbers := strings.Split(strings.Split(line, ":")[1], " | ")

		for _, field := range strings.Fields(numbers[0]) {
			wNumber, _ := strconv.Atoi(field)

			if err != nil {
				panic(err)
			}
			card.Wnumbers[wNumber] = true
		}

		for _, field := range strings.Fields(numbers[1]) {
			number, _ := strconv.Atoi(field)

			if err != nil {
				panic(err)
			}
			card.numbers = append(card.numbers, number)
		}

		cards[card.id] = card
	}

	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return cards
}

func partOne(cards map[int]Card) int {
	points := 0

	for _, card := range cards {
		nHits := -1

		for _, number := range card.numbers {
			_, exists := card.Wnumbers[number]

			if exists {
				nHits++
			}
		}

		card.points = int(math.Pow(2, float64(nHits)))
		points += card.points
	}

	return points
}

func partTwo(cards map[int]Card) int {
	nScratchcards := make(map[int]int)
	keys := make([]int, 0)
	sum := 0

	for k, _ := range cards {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, id := range keys {
		hits := 0
		nScratchcards[id] += 1

		for i := 0; i < len(cards[id].numbers); i++ {
			_, exists := cards[id].Wnumbers[cards[id].numbers[i]]

			if exists {
				hits++
			}
		}

		for j := 0; j < hits; j++ {
			nScratchcards[id+1+j] += nScratchcards[id]
		}
	}

	for _, copies := range nScratchcards {
		sum += copies
	}
	return sum
}

func main() {
	cards := parse()

	fmt.Printf("%d\n", partOne(cards))
	fmt.Printf("%d\n", partTwo(cards))
}
