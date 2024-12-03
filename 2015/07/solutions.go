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

type WireSignalMap map[string]string

func main() {
	part := aoc.PartFlag()

	if part == 1 {
		aoc.PrintSolution(part1(input, "a"))
	} else {
		aoc.PrintSolution(part2(input))
	}
}

func part1(partInput string, wire string) string {
	parsed := aoc.ParseInput(partInput)
	definitions := strings.Split(parsed, "\n")

	wireSignalMap := generateWireSignalMapFromDefinitions(definitions)
	signal := getSignalForWire(wire, &wireSignalMap)

	return signal
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	definitions := strings.Split(parsed, "\n")

	firstWireSignalMap := generateWireSignalMapFromDefinitions(definitions)
	signalA := getSignalForWire("a", &firstWireSignalMap)

	secondWireSignalMap := generateWireSignalMapFromDefinitions(definitions)
	secondWireSignalMap["b"] = signalA

	signalAPart2 := getSignalForWire("a", &secondWireSignalMap)

	return signalAPart2
}

func generateWireSignalMapFromDefinitions(definitions []string) WireSignalMap {
	wireSignalMap := WireSignalMap{}
	wireSignalRegex := regexp.MustCompile(`(.*) -> ([a-z]+)`)

	for _, definition := range definitions {
		matches := wireSignalRegex.FindStringSubmatch(definition)

		wireSignalMap[matches[2]] = matches[1]
	}

	return wireSignalMap
}

func getSignalForWire(wire string, wireSignalMap *WireSignalMap) string {
	signalValue := (*wireSignalMap)[wire]

	if regexp.MustCompile(`^\d+$`).MatchString(signalValue) {
		return signalValue
	}

	var signal uint16

	signalOperationRegex := regexp.MustCompile(`([a-z]+|\d+)?\s?(AND|OR|NOT|LSHIFT|RSHIFT)?\s?([a-z]+|\d+)?`)
	parts := signalOperationRegex.FindStringSubmatch(signalValue)
	left, action, right := parts[1], parts[2], parts[3]

	leftIsNumerical := regexp.MustCompile(`\d+`).MatchString(left)
	rightIsNumerical := regexp.MustCompile(`\d+`).MatchString(right)

	if left != "" && !leftIsNumerical {
		left = getSignalForWire(left, wireSignalMap)
	}

	if right != "" && !rightIsNumerical {
		right = getSignalForWire(right, wireSignalMap)
	}

	switch action {
	case "AND":
		signal = stringToUint16(left) & stringToUint16(right)
	case "OR":
		signal = stringToUint16(left) | stringToUint16(right)
	case "NOT":
		signal = ^stringToUint16(right)
	case "LSHIFT":
		signal = stringToUint16(left) << stringToUint16(right)
	case "RSHIFT":
		signal = stringToUint16(left) >> stringToUint16(right)
	case "":
		signal = stringToUint16(left)
	}

	signalString := strconv.Itoa(int(signal))

	(*wireSignalMap)[wire] = signalString

	return signalString
}

func stringToUint16(inputString string) uint16 {
	inputAsInt, _ := strconv.Atoi(inputString)
	return uint16(inputAsInt)
}
