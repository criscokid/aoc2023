package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/criscokid/aoc2023/internal/fileinput"
)

type node struct {
	left  string
	right string
}

func main() {
	file_path := "input.txt"
	lines, err := fileinput.ReadLines(file_path)
	if err != nil {
		log.Fatal(err)
	}

	order := strings.Split(lines[0], "")
	fmt.Println(order)

	nodeMap := constructMap(lines[2:])
	fmt.Printf("node map: %v\n", nodeMap)

	i := 0
	currentNode := nodeMap["AAA"]
	moves := 0
	for {
		direction := order[i]
		var newNode string
		if direction == "L" {
			newNode = currentNode.left
		} else {
			newNode = currentNode.right
		}
		if newNode == "ZZZ" {
			moves++
			break
		}
		
		if i + 1 < len(order) {
			i++
		} else {
			i = 0
		}
		moves++
		currentNode = nodeMap[newNode]
	}

	fmt.Println(moves)
}

func constructMap(input []string) (map[string]node) {
	replacer := strings.NewReplacer("=", "", "(", "", ")", "", ",", "")
	nodeMap := map[string]node{}
	for i := 0; i < len(input); i++ {
		cleaned := replacer.Replace(input[i])
		vals := strings.Fields(cleaned)


		nodeMap[vals[0]] = node{left: vals[1], right:vals[2]}	
	}

	return nodeMap
}
