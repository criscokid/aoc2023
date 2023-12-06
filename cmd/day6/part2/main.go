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

	race := getRace(lines)
	fmt.Println(race)

	fmt.Println(findPossibleWins(race))
}

func getRace(input []string) race {
	timeParts := strings.Split(input[0], ":")
	time := strings.ReplaceAll(timeParts[1], " ", "")
	distanceParts := strings.Split(input[1], ":")
	distance := strings.ReplaceAll(distanceParts[1], " ", "")

	timeVal, err := strconv.Atoi(time)
	if err != nil {
		log.Fatal(err)
	}

	distanceVal, err := strconv.Atoi(distance)
	if err != nil {
		log.Fatal(err)
	}

	return race{time: timeVal, distance: distanceVal}
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
