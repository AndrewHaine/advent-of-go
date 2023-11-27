package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type SolutionDate struct {
	Day  int
	Year int
}

func main() {
	var day = 0
	var year = 0

	flag.IntVar(&day, "day", 1, "Day of the month; 1-25")
	flag.IntVar(&year, "year", 2023, "Year; 2015+")
	flag.Parse()

	if day == 0 {
		log.Fatal("Please specify a day")
	}

	if day > 25 {
		log.Fatalf("Day %d is out of range", day)
	}

	if year == 0 {
		log.Fatal("Please specify a year")
	}

	date := SolutionDate{day, year}
	dir := generateDir(date)

	// Create a directory for the day
	fsDir := filepath.Join(".", dir)
	os.MkdirAll(fsDir, os.ModePerm)

	// If any of the files exist already, throw an error we don't want to overwrite them
	templates := []string{"input_test.txt", "input.txt", "solutions.go", "solutions_test.go"}
	for _, template := range templates {
		dest := filepath.Join(fsDir, template)
		if _, err := os.Stat(dest); err == nil {
			log.Fatalf("Error creating %s, file already exists", template)
		}

		src := filepath.Join(".", "cmd", "generators", "day", "templates", template)

		data, readErr := os.ReadFile(src)
		checkErr(readErr)

		writeErr := os.WriteFile(dest, data, 0644)
		checkErr(writeErr)
	}
}

func generateDir(date SolutionDate) string {
	dayString := strconv.Itoa(date.Day)

	if date.Day < 10 {
		dayString = "0" + strconv.Itoa(date.Day)
	}

	return strconv.Itoa(date.Year) + "/" + dayString
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
