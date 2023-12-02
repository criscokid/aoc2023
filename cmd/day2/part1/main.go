package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/criscokid/aoc2023/internal/fileinput"
	"github.com/criscokid/aoc2023/internal/stringutils"
)

var diceInBag = map[string]int {
	"red": 12,
	"green": 13,
	"blue": 14,
}

type game struct {
	gameId int
	draws []drawSet	
}

type drawSet struct {
	dice []visibleDraw
}

type visibleDraw struct {
	color string
	count int
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
		fmt.Println(game)
	}
	
	sum := sumPossibleGameIds(games)
	fmt.Println(sum)
}

func parseGame(input string) game {
	gameSplits := strings.Split(input, ":")
	game := newGameFrom(gameSplits[0])

	drawSets := strings.Split(gameSplits[1], ";")
	for _, v := range drawSets {
		drawSet := parseDrawSet(v)
		game.draws = append(game.draws, drawSet)
	}

	return game
}

func newGameFrom(gameIdString string) game {
	parts := strings.Split(gameIdString, " ")
	gameId, err := strconv.Atoi(strings.TrimSpace(parts[1]))

	if err != nil {
		log.Fatal(err.Error())
	}

	return game{gameId: gameId, draws: make([]drawSet, 0)}
}

func parseDrawSet(input string) (drawSet){
	diceGroups := strings.Split(input, ",")
	visibleDraws := make([]visibleDraw, 0)
	for _, v := range diceGroups {
		if len(strings.TrimSpace(v)) == 0 {
			continue
		}
		visibleDraw := parseVisibleDice(v)
		visibleDraws = append(visibleDraws, visibleDraw)
	}

	return drawSet{dice: visibleDraws}
}

func parseVisibleDice(input string) visibleDraw {
	diceParts := stringutils.TrimAndSplit(input, " ")
	color := strings.TrimSpace(diceParts[1])
	count, err := strconv.Atoi(strings.TrimSpace(diceParts[0]))
	if err != nil {
		log.Fatal(err. Error())
	}
	return visibleDraw{count: count, color: color}
}

func sumPossibleGameIds(games []game) int {
	sum := 0
	for _, g := range games {
		addToSet := true
		drawSet:
		for _, ds := range g.draws {
			for _, d := range ds.dice {
				if d.count > diceInBag[d.color] {
					addToSet = false
					break drawSet
				}
			}
		}

		if addToSet {
			sum += g.gameId
		}
	}

	return sum
}
