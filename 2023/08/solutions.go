package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"regexp"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type Node struct {
	name  string
	left  string
	right string
}

func main() {
	part := aoc.PartFlag()

	if part == 1 {
		aoc.PrintSolution(part1(input))
	} else {
		aoc.PrintSolution(part2(input))
	}
}

func part1(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	instructions, nodes := getInstructionsAndNodesFromInput(parsed)

	nodeDictionary := map[string]Node{}

	for _, node := range nodes {
		nodeDictionary[node.name] = node
	}

	stepCount := 0
	instructionIndex := 0
	currentNode := nodeDictionary["AAA"]

	for currentNode.name != "ZZZ" {
		instruction := instructions[instructionIndex]
		var newNodeName string
		if instruction == 'L' {
			newNodeName = currentNode.left
		} else {
			newNodeName = currentNode.right
		}
		newNode := nodeDictionary[newNodeName]
		stepCount += 1

		if instructionIndex == len(instructions)-1 {
			instructionIndex = 0
		} else {
			instructionIndex += 1
		}

		currentNode = newNode
	}

	return strconv.Itoa(stepCount)
}

// Brute-force method
// Answer is in the trillions so come back in a week!

// func part2(partInput string) string {
// 	parsed := aoc.ParseInput(partInput)
// 	instructions, nodes := getInstructionsAndNodesFromInput(parsed)

// 	nodeDictionary := map[string]Node{}
// 	startNodes := []Node{}

// 	for _, node := range nodes {
// 		nodeDictionary[node.name] = node
// 		if node.name[len(node.name)-1] == 'A' {
// 			startNodes = append(startNodes, node)
// 		}
// 	}

// 	stepCount := 0
// 	instructionIndex := 0
// 	endReached := false
// 	currentNodes := startNodes

// 	for !endReached {
// 		endReached = true
// 		instruction := instructions[instructionIndex]
// 		newNodes := []Node{}

// 		for _, node := range currentNodes {
// 			var newNodeName string
// 			if instruction == 'L' {
// 				newNodeName = node.left
// 			} else {
// 				newNodeName = node.right
// 			}
// 			newNode := nodeDictionary[newNodeName]
// 			newNodes = append(newNodes, newNode)
// 			if newNodeName[len(newNodeName)-1] != 'Z' {
// 				endReached = false
// 			}
// 		}

// 		stepCount += 1

// 		if instructionIndex == len(instructions)-1 {
// 			instructionIndex = 0
// 		} else {
// 			instructionIndex += 1
// 		}

// 		currentNodes = newNodes
// 	}

// 	return strconv.Itoa(stepCount)
// }

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	instructions, nodes := getInstructionsAndNodesFromInput(parsed)

	nodeDictionary := map[string]Node{}
	startNodes := []Node{}

	for _, node := range nodes {
		nodeDictionary[node.name] = node
		if strings.HasSuffix(node.name, "A") {
			startNodes = append(startNodes, node)
		}
	}

	nodeCounts := []int{}

	for _, node := range startNodes {
		instructionIndex := 0
		nodeStepCount := 0
		endReached := false
		currentNode := node

		for !endReached {
			instruction := instructions[instructionIndex]
			var newNodeName string
			if instruction == 'L' {
				newNodeName = currentNode.left
			} else {
				newNodeName = currentNode.right
			}

			if strings.HasSuffix(newNodeName, "Z") {
				endReached = true
			}

			currentNode = nodeDictionary[newNodeName]

			nodeStepCount += 1

			if instructionIndex == len(instructions)-1 {
				instructionIndex = 0
			} else {
				instructionIndex += 1
			}
		}
		nodeCounts = append(nodeCounts, nodeStepCount)
	}

	stepCount := lowestCommonMultiple(nodeCounts[0], nodeCounts[1], nodeCounts...)

	return strconv.Itoa(stepCount)
}

func getInstructionsAndNodesFromInput(input string) (instructions string, nodes []Node) {
	split := strings.Split(input, "\n\n")
	instructions = split[0]

	nodeRows := strings.Split(split[1], "\n")
	nodes = []Node{}

	for _, row := range nodeRows {
		nodeRegex := regexp.MustCompile(`([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)`)
		parsedRow := nodeRegex.FindStringSubmatch(row)
		node := Node{parsedRow[1], parsedRow[2], parsedRow[3]}
		nodes = append(nodes, node)
	}

	return
}

func highestCommonFactor(a int, b int) int {
	for b != 0 {
		c := b
		b = a % b
		a = c
	}
	return a
}

func lowestCommonMultiple(a int, b int, numbers ...int) int {
	result := a * b / highestCommonFactor(a, b)

	for i := 0; i < len(numbers); i++ {
		result = lowestCommonMultiple(result, numbers[i])
	}

	return result
}
