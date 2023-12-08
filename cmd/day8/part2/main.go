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

	keysEndingInA := []string{}
	for k := range nodeMap {
		if k[2] == 'A' {
			keysEndingInA = append(keysEndingInA, k)
		}
	}

	fmt.Println(keysEndingInA)

	moveToReachZ := []int{}
	for _, k := range keysEndingInA {
		m, _ := moves(k, 'Z', nodeMap, order)
		moveToReachZ = append(moveToReachZ, m)	
	}

	fmt.Println(moveToReachZ)

	result := LCM(moveToReachZ[0], moveToReachZ[1], moveToReachZ[2:]...)
	fmt.Println(result)
}

func constructMap(input []string) map[string]node {
	replacer := strings.NewReplacer("=", "", "(", "", ")", "", ",", "")
	nodeMap := map[string]node{}
	for i := 0; i < len(input); i++ {
		cleaned := replacer.Replace(input[i])
		vals := strings.Fields(cleaned)

		nodeMap[vals[0]] = node{left: vals[1], right: vals[2]}
	}

	return nodeMap
}

func moves(start string, endingIn rune, nodeMap map[string]node, order []string) (int, string) {
	moves := 0
	i := 0
	currentNode := nodeMap[start]
	var newNode string
	for {
		direction := order[i]
		if direction == "L" {
			newNode = currentNode.left
		} else {
			newNode = currentNode.right
		}
		if newNode[2] == byte(endingIn) {
			moves++
			break
		}

		if i+1 < len(order) {
			i++
		} else {
			i = 0
		}
		moves++
		currentNode = nodeMap[newNode]
	}
	return moves, newNode
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
