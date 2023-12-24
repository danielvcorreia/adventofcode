package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sD struct {
	destination int
	source      int
}

func parse() ([]int, map[int]map[sD]int, int) {
	s := bufio.NewScanner(os.Stdin)
	almanac := make(map[int]map[sD]int)
	seeds := []int{}
	nMaps := -1

	for s.Scan() {
		line := s.Text()

		if line == "" {
			continue
		}

		if strings.Contains(line, "seeds") {
			strSeeds := strings.Fields(strings.Split(line, ": ")[1])

			for _, strSeed := range strSeeds {
				seed, err := strconv.Atoi(strSeed)

				if err != nil {
					panic(err)
				}
				seeds = append(seeds, seed)
			}
			continue
		}

		if strings.Contains(line, "map") {
			nMaps++
			almanac[nMaps] = make(map[sD]int)
			continue
		}

		values := strings.Fields(line)
		destination, _ := strconv.Atoi(values[0])
		source, _ := strconv.Atoi(values[1])
		length, _ := strconv.Atoi(values[2])

		almanac[nMaps][sD{destination, source}] = length
	}

	return seeds, almanac, nMaps + 1
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)

	for i := range a {
		a[i] = min + i
	}
	return a
}

func partOne(seeds []int, almanac map[int]map[sD]int, nMaps int) int {
	locations := []int{}

	for i := 0; i < len(seeds); i++ {
		location := seeds[i]

		for j := 0; j < nMaps; j++ {
			keys := []sD{}

			for k := range almanac[j] {
				keys = append(keys, k)
			}

			for _, key := range keys {
				if location > key.source && location < key.source+almanac[j][key] {
					location = (location - key.source) + key.destination
					break
				}
			}
		}
		locations = append(locations, location)
	}

	sort.Ints(locations)
	return locations[0]
}

func partTwo(seeds []int, almanac map[int]map[sD]int, nMaps int) int {
	locations := []int{}
	actualSeeds := make(map[int][]int)

	for i := 0; i < len(seeds)/2; i++ {
		rangeSeeds := makeRange(seeds[i*2], seeds[i*2]+seeds[i*2+1]-1)
		actualSeeds[i] = rangeSeeds
	}

	for _, value := range actualSeeds {
		for i := 0; i < len(value); i++ {
			location := value[i]

			for j := 0; j < nMaps; j++ {
				keys := []sD{}

				for k := range almanac[j] {
					keys = append(keys, k)
				}

				for _, key := range keys {
					if location >= key.source && location < key.source+almanac[j][key] {
						location = (location - key.source) + key.destination
						break
					}
				}
			}
			locations = append(locations, location)
		}
	}
	sort.Ints(locations)
	return locations[0]
}

func main() {
	seeds, almanac, nMaps := parse()

	fmt.Println(partOne(seeds, almanac, nMaps))
	fmt.Println(partTwo(seeds, almanac, nMaps))
}
