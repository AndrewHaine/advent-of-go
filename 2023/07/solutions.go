package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"andrewhaine/advent-of-go/util/maputil"
	"slices"
	"sort"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

var cardScores = []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

var cardScoresWithJoker = []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

var handRanks = map[string]int{
	"fiveOfKind":  7,
	"fourOfKind":  6,
	"fullHouse":   5,
	"threeOfKind": 4,
	"twoPair":     3,
	"onePair":     2,
	"highCard":    1,
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
	rows := strings.Split(parsed, "\n")

	sortedRows := sortRows(rows, false)

	score := 0

	for i, row := range sortedRows {
		_, bid := splitRow(row)
		score += bid * (i + 1)
	}

	return strconv.Itoa(score)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	rows := strings.Split(parsed, "\n")

	sortedRows := sortRows(rows, true)

	score := 0

	for i, row := range sortedRows {
		_, bid := splitRow(row)
		score += bid * (i + 1)
	}

	return strconv.Itoa(score)
}

func splitRow(row string) (hand string, bid int) {
	split := strings.Split(row, " ")
	hand = split[0]
	bid, _ = strconv.Atoi(split[1])
	return
}

func sortRows(rows []string, includeJokers bool) []string {
	sort.Slice(rows, func(i, j int) bool {

		iHand, _ := splitRow(rows[i])
		jHand, _ := splitRow(rows[j])

		iScore := getHandRank(iHand, includeJokers)
		jScore := getHandRank(jHand, includeJokers)

		if iScore != jScore {
			return iScore < jScore
		}

		iLess := false

		for ii, iCard := range iHand {
			jHandCards := []rune(jHand)
			var iCardScore int
			var jCardScore int
			if includeJokers {
				iCardScore = slices.Index(cardScoresWithJoker, iCard)
				jCardScore = slices.Index(cardScoresWithJoker, jHandCards[ii])
			} else {
				iCardScore = slices.Index(cardScores, iCard)
				jCardScore = slices.Index(cardScores, jHandCards[ii])
			}
			if iCardScore != jCardScore {
				iLess = iCardScore < jCardScore
				break
			}
		}

		return iLess
	})

	return rows
}

func getHandRank(hand string, includeJokers bool) int {
	handCardQuantities := map[rune]int{}

	for _, card := range hand {
		handCardQuantities[card] += 1
	}

	jokerQuantity := handCardQuantities['J']

	if len(handCardQuantities) == 1 {
		return handRanks["fiveOfKind"]
	}

	if len(handCardQuantities) == 5 {
		if includeJokers && jokerQuantity == 1 {
			return handRanks["onePair"]
		}

		return handRanks["highCard"]
	}

	cardQuantities := maputil.Values(handCardQuantities)

	if len(handCardQuantities) == 2 {
		if slices.Contains(cardQuantities, 4) {
			if includeJokers && (jokerQuantity == 1 || jokerQuantity == 4) {
				return handRanks["fiveOfKind"]
			}
			return handRanks["fourOfKind"]
		}

		if includeJokers && (jokerQuantity == 2 || jokerQuantity == 3) {
			return handRanks["fiveOfKind"]
		}

		return handRanks["fullHouse"]
	}

	if len(handCardQuantities) == 3 {
		if slices.Contains(cardQuantities, 3) {
			if includeJokers && (jokerQuantity == 1 || jokerQuantity == 3) {
				return handRanks["fourOfKind"]
			}

			return handRanks["threeOfKind"]
		}

		if includeJokers && jokerQuantity == 2 {
			return handRanks["fourOfKind"]
		}

		if includeJokers && jokerQuantity == 1 {
			return handRanks["fullHouse"]
		}

		return handRanks["twoPair"]
	}

	if includeJokers && (jokerQuantity == 1 || jokerQuantity == 2) {
		return handRanks["threeOfKind"]
	}

	return handRanks["onePair"]
}
