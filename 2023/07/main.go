package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type bidRank struct {
	bid      int
	handType int
	preRank  int
}

func parse(s *bufio.Scanner) map[string]bidRank {
	handsAndBids := make(map[string]bidRank)

	for s.Scan() {
		line := s.Text()
		handAndBidStr := strings.Split(line, " ")
		bid, _ := strconv.Atoi(handAndBidStr[1])
		handsAndBids[handAndBidStr[0]] = bidRank{bid, -1, -1}
	}
	return handsAndBids
}

func bestHand(hits []int, countJ int) []int {
	if len(hits) == 0 {
		return []int{5}
	}
	for i := 0; i < countJ; i++ {
		hits[len(hits)-1] += 1
	}
	return hits
}

func partOne(handsValues map[string]bidRank) int {
	cards := map[string]int{"2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "T": 9, "J": 10, "Q": 11, "K": 12, "A": 13}
	types := map[int]int{5: 7, 14: 6, 23: 5, 113: 4, 122: 3, 1112: 2, 11111: 1}
	typeArrays := make(map[int][]int)
	mapRanks := make(map[int]int)
	keys := make([]int, 0)
	rank := 1
	totalWinnings := 0

	for hand, values := range handsValues {
		hits := []int{}
		handDecider := 0
		handType := 0

		for _, handCard := range hand {
			cardRank := cards[string(handCard)]
			handDecider = handDecider*100 + cardRank
		}

		for possibleCard := range cards {
			count := strings.Count(hand, possibleCard)

			if count != 0 {
				hits = append(hits, count)
			}
		}
		sort.Ints(hits)

		for _, hit := range hits {
			handType = handType*10 + hit
		}
		typeArrays[types[handType]] = append(typeArrays[types[handType]], handDecider)
		handsValues[hand] = bidRank{values.bid, types[handType], handDecider}
	}

	for k := range typeArrays {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		sort.Ints(typeArrays[k])
		for i := 0; i < len(typeArrays[k]); i++ {
			mapRanks[typeArrays[k][i]] = rank
			rank++
		}
	}

	for _, values := range handsValues {
		totalWinnings += values.bid * mapRanks[values.preRank]
	}

	return totalWinnings
}

func partTwo(handsValues map[string]bidRank) int {
	cards := map[string]int{"2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "T": 9, "J": 10, "Q": 11, "K": 12, "A": 13}
	types := map[int]int{5: 7, 14: 6, 23: 5, 113: 4, 122: 3, 1112: 2, 11111: 1}
	typeArrays := make(map[int][]int)
	mapRanks := make(map[int]int)
	keys := make([]int, 0)
	rank := 1
	totalWinnings := 0

	for hand, values := range handsValues {
		hits := []int{}
		handDecider := 0
		handType := 0
		countJ := 0

		for _, handCard := range hand {
			cardRank := cards[string(handCard)]
			handDecider = handDecider*100 + cardRank
		}

		for possibleCard := range cards {
			count := strings.Count(hand, possibleCard)

			if count != 0 {
				if possibleCard == "J" {
					countJ += count
				} else {
					hits = append(hits, count)
				}
			}
		}
		sort.Ints(hits)
		hits = bestHand(hits, countJ)

		for _, hit := range hits {
			handType = handType*10 + hit
		}
		typeArrays[types[handType]] = append(typeArrays[types[handType]], handDecider)
		handsValues[hand] = bidRank{values.bid, types[handType], handDecider}
	}

	for k := range typeArrays {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		sort.Ints(typeArrays[k])
		for i := 0; i < len(typeArrays[k]); i++ {
			mapRanks[typeArrays[k][i]] = rank
			rank++
		}
	}

	for _, values := range handsValues {
		totalWinnings += values.bid * mapRanks[values.preRank]
	}

	return totalWinnings
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	handsAndBids := parse(s)

	fmt.Println(partOne(handsAndBids))
	fmt.Println(partTwo(handsAndBids))
}
