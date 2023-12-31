package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/criscokid/aoc2023/internal/fileinput"
)

var diceInBag = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type game struct {
	gameId int
	draws  []drawSet
}

type drawSet struct {
	dice []visibleDraw
}

type visibleDraw struct {
	color string
	count int
}

type bag struct {
	contents map[string]int
}

func main() {
	file_path := "input.txt"
	lines, err := fileinput.ReadLines(file_path)
	if err != nil {
		log.Fatal(err)
	}

	games := []game{}
	sum := 0
	for _, line := range lines {
		game := parseGame(line)
		games = append(games, game)
		bag := findMinimumBagForGame(game)
		sum += bag.power()
	}

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

func parseDrawSet(input string) drawSet {
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
	diceParts := strings.Split(strings.TrimSpace(input), " ")
	color := strings.TrimSpace(diceParts[1])
	count, err := strconv.Atoi(strings.TrimSpace(diceParts[0]))
	if err != nil {
		log.Fatal(err.Error())
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

func findMinimumBagForGame(game game) bag {
	content := map[string]int{
		"green": 0,
		"blue": 0,
		"red": 0,
	}

	for _, ds := range game.draws {
		for _, d := range ds.dice {
			currentVal := content[d.color]
			if d.count > currentVal {
				content[d.color] = d.count
			}
		}
	}

	return bag{contents: content}
}

func (b bag) power() int {
	return b.contents["green"] * b.contents["blue"] * b.contents["red"]
}
