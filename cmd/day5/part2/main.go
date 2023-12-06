package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/criscokid/aoc2023/internal/fileinput"
	"github.com/criscokid/aoc2023/internal/mathutils"
)

type seed struct {
	start  int
	stride int
}

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

	totalVals := mathutils.SumSlice(seeds, func(s seed) int {
		return s.stride
	})

	fmt.Printf("Total seeds: %d\n", totalVals)

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

	
	num_jobs := len(seeds)
	jobs := make(chan seed, num_jobs)
	results := make(chan int, num_jobs)

	for i := 0; i < num_jobs; i++ {
		go func(jobs <-chan seed, results chan<- int) {
			for s := range jobs {
				loc := findLowestLocationForSeed(maps, s)
				results <- loc
			}
		}(jobs, results)
	}

	for _, v := range seeds {
		jobs <- v	
	}
	close(jobs)

	min := 0
	for a := 0; a < num_jobs; a++ {
		loc := <-results
		if min == 0 || loc < min {
			min = loc
		}
	}

	fmt.Println(min)
}

func parseSeeds(input string) []seed {
	input = strings.ReplaceAll(input, "seeds:", "")
	seeds := []seed{}
	seedValues := strings.Fields(input)
	for i := 0; i < len(seedValues); i += 2 {
		start := seedValues[i]
		stride := seedValues[i+1]

		startVal, err := strconv.Atoi(start)
		if err != nil {
			log.Fatal(err)
		}

		strideVal, err := strconv.Atoi(stride)
		if err != nil {
			log.Fatal(err)
		}

		seeds = append(seeds, seed{start: startVal, stride: strideVal})
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

func findLocation(maps [][]srcDestMap, val int) int {
	currentLocation := val
	for _, m := range maps {
		idx := slices.IndexFunc(m, func(sdm srcDestMap) bool {
			return sdm.IsInMapRange(currentLocation)
		})
		if idx == -1 {
			continue
		}

		currentLocation = m[idx].FindDest(currentLocation)
	}

	return currentLocation
}

func findLowestLocationForSeed(maps [][]srcDestMap, s seed) int {
	minLoc := 0
	max := s.start + s.stride
	for i := s.start; i <= max; i++ {
		loc := findLocation(maps, i)
		if minLoc == 0 || loc < minLoc {
			minLoc = loc
		}
	}
	return minLoc
}
