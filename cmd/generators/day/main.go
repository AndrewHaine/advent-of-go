package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/manifoldco/promptui"
)

type SolutionDate struct {
	Day  int
	Year int
}

func main() {
	yearSelectTemplates := &promptui.SelectTemplates{
		Active:   "\U0001F385 {{ . | bold }}",
		Inactive: "   {{ . }}",
		Selected: "\U0001F385 Year: {{ . | bold | faint }}",
	}

	yearPrompt := promptui.Select{
		Label:     "Select Year",
		Templates: yearSelectTemplates,
		Items:     []int{2024, 2023, 2022, 2021, 2020, 2019, 2018, 2017, 2016, 2015},
		Size:      10,
	}

	_, year, yearErr := yearPrompt.Run()

	if yearErr != nil {
		log.Fatalf("Invalid year %v\n", yearErr)
	}

	daySelectTemplates := &promptui.SelectTemplates{
		Active:   "\U0001F385 {{ . | bold }}",
		Inactive: "   {{ . }}",
		Selected: "\U0001F385 Day: {{ . | bold | faint }}",
	}

	dayPrompt := promptui.Select{
		Label:     "Select Day",
		Templates: daySelectTemplates,
		Items:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
		Size:      10,
	}

	_, day, dayErr := dayPrompt.Run()

	if dayErr != nil {
		log.Fatalf("Invalid day sleected %v\n", dayErr)
	}

	dayValue, _ := strconv.Atoi(day)
	yearValue, _ := strconv.Atoi(year)

	date := SolutionDate{dayValue, yearValue}
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
