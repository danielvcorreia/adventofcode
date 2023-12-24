package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse() ([]int, []int) {
	s := bufio.NewScanner(os.Stdin)
	times := []int{}
	distances := []int{}

	for s.Scan() {
		line := s.Text()

		if strings.Contains(line, "Time") {
			timeValues := strings.Fields(strings.Split(line, ":")[1])

			for _, value := range timeValues {
				time, _ := strconv.Atoi(value)
				times = append(times, time)
			}
		} else if strings.Contains(line, "Distance") {
			distValues := strings.Fields(strings.Split(line, ":")[1])

			for _, value := range distValues {
				distance, _ := strconv.Atoi(value)
				distances = append(distances, distance)
			}
		} else {
			continue
		}
	}
	return times, distances
}

func partOne(times []int, distances []int) int {
	sumPossibilities := 1

	for i := 0; i < len(times); i++ {
		possibleDistances := 0

		for velocity := 1; velocity < times[i]; velocity++ {
			timeLeft := times[i] - velocity
			dist := timeLeft * velocity

			if dist > distances[i] {
				possibleDistances += 1
			}
		}
		sumPossibilities *= possibleDistances
	}
	return sumPossibilities
}

func partTwo(times []int, distances []int) int {
	strActualTime := ""
	strActualDistance := ""
	possibleDistances := 0
	sumPossibilities := 1

	for i := range times {
		strActualTime += strconv.Itoa(times[i])
		strActualDistance += strconv.Itoa(distances[i])
	}
	actualTime, _ := strconv.Atoi(strActualTime)
	actualDistance, _ := strconv.Atoi(strActualDistance)

	for velocity := 1; velocity < actualTime; velocity++ {
		timeLeft := actualTime - velocity
		dist := timeLeft * velocity

		if dist > actualDistance {
			possibleDistances += 1
		}
	}
	sumPossibilities *= possibleDistances
	return sumPossibilities
}

func main() {
	times, distances := parse()

	fmt.Println(partOne(times, distances))
	fmt.Println(partTwo(times, distances))
}
