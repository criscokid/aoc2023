package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/criscokid/aoc2023/internal/fileinput"
)

type srcDestMap struct {
	srcMin  int
	srcMax  int
	destMin int
}

func main() {
	file_path := "input.txt"
	lines, err := fileinput.ReadLines(file_path)
	if err != nil {
		log.Fatal(err)
	}

	maps := [][]srcDestMap{}

	seeds := parseSeeds(lines[0])

	tempMaps := []srcDestMap{}
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		line = strings.TrimSpace(line)
		if line == "" || line[len(line)-1] == ':' {
			if len(tempMaps) > 0 {
				maps = append(maps, tempMaps)
				tempMaps = []srcDestMap{}
			}
			continue
		}

		parsedMap := parseMapRow(line)
		tempMaps = append(tempMaps, parsedMap)
	}

	if len(tempMaps) > 0 {
		maps = append(maps, tempMaps)
	}
	

	locations := []int{}
	for _, v := range seeds {
		currentLocation := v
		for _, m := range maps {
			idx := slices.IndexFunc(m , func(sdm srcDestMap) bool {
				return sdm.IsInMapRange(currentLocation)
			})
			if idx == -1 {
				continue
			}

			currentLocation = m[idx].FindDest(currentLocation)
		}

		locations = append(locations, currentLocation)
	}

	fmt.Println(locations)
	min := slices.Min(locations)
	fmt.Println(min)
}

func parseSeeds(input string) []int {
	input = strings.ReplaceAll(input, "seeds:", "")
	seeds := []int{}
	for _, v := range strings.Fields(input) {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		seeds = append(seeds, val)
	}

	return seeds
}

func parseMapRow(input string) srcDestMap {
	values := []int{}
	for _, v := range strings.Fields(input) {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		values = append(values, val)
	}

	return srcDestMap{srcMin: values[1], srcMax: values[1] + values[2], destMin: values[0]}
}

func (sdm srcDestMap) IsInMapRange(val int) bool {
	return val >= sdm.srcMin && val <= sdm.srcMax 
}

func (sdm srcDestMap) FindDest(val int) int {
	offset := val - sdm.srcMin
	return sdm.destMin + offset
}
