package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/criscokid/aoc2023/internal/fileinput"
)

type game struct {
	gameId         int
	winningNumbers []int
	numbersFound   []int
}

func main() {
	file_path := "input.txt"
	lines, err := fileinput.ReadLines(file_path)
	if err != nil {
		log.Fatal(err)
	}

	games := []game{}
	for _, line := range lines {
		game := parseGame(line)
		games = append(games, game)
	}

	sum := 0
	for _, game := range games {
		count := countMatchingNumbers(game)
		points := 0
		if count >= 1 {
			points = 1
		}
		for i := 0; i < count-1; i++ {
			points *= 2
		}
		sum += points
	}
	fmt.Println(sum)
}

func parseGame(input string) game {
	gameParts := strings.Split(input, ": ")
	game := parseGameId(gameParts[0])
	parseGameNumbers(&game, gameParts[1])
	return game
}

func parseGameId(input string) game {
	fmt.Println(input)
	input = strings.ReplaceAll(input, "Card", " ")
	input = strings.TrimSpace(input)
	id, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return game{gameId: id}
}

func parseGameNumbers(game *game, input string) {
	numberParts := strings.Split(input, " | ")
	fields := strings.Fields(numberParts[1])
	fmt.Println(fields)
	for _, v := range fields {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		game.numbersFound = append(game.numbersFound, num)
	}

	fields = strings.Fields(numberParts[0])
	for _, v := range fields {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		game.winningNumbers = append(game.winningNumbers, num)
	}
	// foundNumReader := strings.NewReader(numberParts[1])
	// for {
	// 	num := make([]byte, 2)
	// 	_, err := foundNumReader.Read(num)
	// 	if err != nil {
	// 		break
	// 	}
	//
	// 	numStr := strings.TrimSpace(string(num))
	// 	val, err := strconv.Atoi(numStr)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	game.numbersFound = append(game.numbersFound, val)
	//
	// 	space := make([]byte,1)
	// 	_, err = foundNumReader.Read(space)
	// 	if err != nil {
	// 		break
	// 	}
	// }

	// winningNumReader := strings.NewReader(numberParts[0])
	// for {
	// 	num := make([]byte, 2)
	// 	_, err := winningNumReader.Read(num)
	// 	if err != nil {
	// 		break
	// 	}
	//
	// 	numStr := strings.TrimSpace(string(num))
	// 	val, err := strconv.Atoi(numStr)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	game.winningNumbers = append(game.winningNumbers, val)
	//
	// 	space := make([]byte, 1)
	// 	_, err = winningNumReader.Read(space)
	// 	if err != nil {
	// 		break
	// 	}
	// }
}

func countMatchingNumbers(game game) int {
	count := 0
	for _, v := range game.numbersFound {
		if slices.Index(game.winningNumbers, v) > -1 {
			count += 1
		}
	}

	return count
}
