package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"andrewhaine/advent-of-go/util/mathutil"
	"slices"
	"sort"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type PageNumberMapping map[string][]string

type Update []string

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
	sections := strings.Split(parsed, "\n\n")

	pageNumberMapping := generatePageNumberMapping(sections[0])
	updates := generatePageNumberUpdates(sections[1])

	validUpdates := []Update{}

	for _, update := range updates {
		valid := isUpdateValid(update, pageNumberMapping)
		if valid {
			validUpdates = append(validUpdates, update)
		}
	}

	answer := getSumOfMiddlePages(validUpdates)
	return strconv.Itoa(answer)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	sections := strings.Split(parsed, "\n\n")

	pageNumberMapping := generatePageNumberMapping(sections[0])
	updates := generatePageNumberUpdates(sections[1])

	invalidUpdates := []Update{}

	for _, update := range updates {
		valid := isUpdateValid(update, pageNumberMapping)

		if !valid {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	fixedUpdates := []Update{}

	for _, invalidUpdate := range invalidUpdates {
		fixedUpdates = append(fixedUpdates, fixUpdate(invalidUpdate, pageNumberMapping))
	}

	answer := getSumOfMiddlePages(fixedUpdates)

	return strconv.Itoa(answer)
}

/*
Generate a mapping of page numbers against those which should appear after it
*/
func generatePageNumberMapping(ruleString string) PageNumberMapping {
	var mapping PageNumberMapping = make(PageNumberMapping)
	rules := strings.Split(ruleString, "\n")

	for _, rule := range rules {
		ruleParts := strings.Split(rule, "|")
		mapping[ruleParts[0]] = append(mapping[ruleParts[0]], ruleParts[1])
	}

	return mapping
}

func generatePageNumberUpdates(updateString string) [][]string {
	updates := [][]string{}
	updateRows := strings.Split(updateString, "\n")

	for _, row := range updateRows {
		updates = append(updates, strings.Split(row, ","))
	}

	return updates
}

func isUpdateValid(update Update, pageNumberMapping PageNumberMapping) bool {
	valid := true

	for i, pageNumber := range update {
		pagesWhichShouldNotAppearBefore := pageNumberMapping[pageNumber]
		for _, page := range pagesWhichShouldNotAppearBefore {
			if slices.Contains(update[:i], page) {
				valid = false
				break
			}
		}

		if !valid {
			break
		}
	}
	return valid
}

func getSumOfMiddlePages(updates []Update) int {
	middleInts := []int{}

	for _, update := range updates {
		updateLen := len(update) / 2
		middleOfUpdate := update[updateLen]

		middleInt, _ := strconv.Atoi(middleOfUpdate)
		middleInts = append(middleInts, middleInt)
	}

	return mathutil.SumInts(middleInts)
}

func fixUpdate(update Update, pageNumberMapping PageNumberMapping) Update {
	sort.Slice(update, func(i, j int) bool {
		left := update[i]
		right := update[j]
		rule, found := pageNumberMapping[left]

		if found && slices.Contains(rule, right) {
			return true
		}

		return false
	})

	return update
}
