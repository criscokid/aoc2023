package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/criscokid/aoc2023/internal/fileinput"
)

type race struct {
	time     int
	distance int
}

func main() {
	file_path := "input.txt"
	lines, err := fileinput.ReadLines(file_path)
	if err != nil {
		log.Fatal(err)
	}

	races := getRaces(lines)
	fmt.Println(races)
	wins := 1
	for _, r := range races {
		possibleWins := findPossibleWins(r)
		wins *= possibleWins
	}
	fmt.Println(wins)
}

func getRaces(input []string) []race {
	races := []race{}
	timeParts := strings.Split(input[0], ":")
	times := strings.Fields(timeParts[1])
	distanceParts := strings.Split(input[1], ":")
	distances := strings.Fields(distanceParts[1])

	for i, v := range times {
		timeVal, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		distanceVal, err := strconv.Atoi(distances[i])
		if err != nil {
			log.Fatal(err)
		}

		races = append(races, race{time: timeVal, distance: distanceVal})
	}

	return races
}

func findPossibleWins(input race) int {
	min := input.time / 2
	max := input.time / 2
	
	for {
		newMin := min - 1
		timeRemaining := input.time - newMin
		if newMin*timeRemaining > input.distance {
			min = newMin
			continue
		}

		break
	}

	for {
		newMax := max + 1
		timeRemaining := input.time - newMax
		if newMax*timeRemaining > input.distance {
			max = newMax
			continue
		}

		break
	}
	
	fmt.Printf("min %d, max %d\n", min, max)
	return max - min + 1
}
