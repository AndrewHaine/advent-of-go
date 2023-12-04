package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"regexp"
	"slices"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

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
	cards := strings.Split(parsed, "\n")
	rowRegex := regexp.MustCompile(`Card\s+\d+\:\s([\d+|\s+]+)\|\s([\d+|\s+]+)`)
	total := 0

	for _, card := range cards {
		matches := rowRegex.FindStringSubmatch(card)
		winningNumbers, ourNumbers := matches[1], matches[2]
		winning, ours := stringToNumberList(winningNumbers), stringToNumberList(ourNumbers)

		cardTotal := 0

		for _, number := range ours {
			if slices.Contains(winning, number) {
				if cardTotal < 1 {
					cardTotal = 1
				} else {
					cardTotal = cardTotal * 2
				}
			}
		}

		total += cardTotal
	}

	return strconv.Itoa(total)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	cards := strings.Split(parsed, "\n")
	rowRegex := regexp.MustCompile(`Card\s+(\d+)\:\s([\d+|\s+]+)\|\s([\d+|\s+]+)`)

	cardQuantities := map[int]int{}

	for _, card := range cards {
		matches := rowRegex.FindStringSubmatch(card)
		cardNumberString, winningNumbers, ourNumbers := matches[1], matches[2], matches[3]
		winning, ours := stringToNumberList(winningNumbers), stringToNumberList(ourNumbers)

		matchCount := 0

		for _, number := range ours {
			if slices.Contains(winning, number) {
				matchCount += 1
			}
		}

		cardNumber, _ := strconv.Atoi(cardNumberString)
		cardQuantities[cardNumber] += 1

		currentCardQuantity := cardQuantities[cardNumber]

		// Start at the last card we need to increment (current card number + number of matches)
		// Increment by how many of the current card we have, working backwards until we get to the next card
		lastPrizeCardNumber := cardNumber + matchCount
		for cardToIncrement := lastPrizeCardNumber; cardToIncrement > cardNumber; cardToIncrement -= 1 {
			if cardToIncrement > len(cards) {
				continue
			}

			cardQuantities[cardToIncrement] += currentCardQuantity
		}
	}

	total := 0

	for _, cardQuantity := range cardQuantities {
		total += cardQuantity
	}

	return strconv.Itoa(total)
}

func stringToNumberList(numbersString string) []int {
	trimmed := strings.Trim(numbersString, " ")
	numberStrings := strings.Split(strings.ReplaceAll(trimmed, "  ", " "), " ")
	numbers := []int{}
	for _, numberString := range numberStrings {
		number, _ := strconv.Atoi(numberString)
		numbers = append(numbers, number)
	}
	return numbers
}
