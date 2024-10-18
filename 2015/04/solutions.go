package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"strconv"

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
	secretKey := aoc.ParseInput(partInput)

	number := 0

	for {
		indexString := strconv.Itoa(number)
		var buffer bytes.Buffer
		buffer.WriteString(secretKey)
		buffer.WriteString(indexString)

		hasher := md5.New()
		hasher.Write(buffer.Bytes())

		md5String := hex.EncodeToString(hasher.Sum(nil))
		md5Prefix := md5String[:5];

		if (md5Prefix == "00000") {
			break
		}

		number++
	}

	return strconv.Itoa(number)
}

func part2(partInput string) string {
	secretKey := aoc.ParseInput(partInput)

	number := 0

	for {
		indexString := strconv.Itoa(number)
		var buffer bytes.Buffer
		buffer.WriteString(secretKey)
		buffer.WriteString(indexString)

		hasher := md5.New()
		hasher.Write(buffer.Bytes())

		md5String := hex.EncodeToString(hasher.Sum(nil))
		md5Prefix := md5String[:6];

		if (md5Prefix == "000000") {
			break
		}

		number++
	}

	return strconv.Itoa(number)
}
