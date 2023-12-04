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

	gamesPile := []game{}
	gamesPile = append(gamesPile, games...)
	for i := 0; i < len(gamesPile); i++ {
		game := gamesPile[i]
		count := countMatchingNumbers(game)
		for j := 0; j < count; j++ {
			idx := game.gameId + j
			gamesPile = append(gamesPile, games[idx])
		}
	}

	fmt.Println(len(gamesPile))
}

func parseGame(input string) game {
	gameParts := strings.Split(input, ": ")
	game := parseGameId(gameParts[0])
	parseGameNumbers(&game, gameParts[1])
	return game
}

func parseGameId(input string) game {
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
