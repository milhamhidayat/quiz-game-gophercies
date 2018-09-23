package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Q string
	A string
}

func main() {
	csvFileName := flag.String("csv", "problem.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	randomize := flag.Bool("random", false, "randomize questions")

	flag.Parse()

	csvFile, err := loadCsv(*csvFileName)
	if err != nil {
		exit(err)
	}

	csvQA, err := readCsv(csvFile)
	if err != nil {
		exit(err)
	}

	if *randomize {
		csvQA = shuffle(csvQA)
	}

	problems := parseLines(csvQA)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	reader := bufio.NewReader(os.Stdin)
	correct := 0
	done := make(chan bool, 1)

	go func() {
		for i, p := range problems {
			fmt.Printf("Problem #%d: %s = ", i+1, p.Q)

			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(strings.ToLower(strings.Replace(answer, "\n", "", -1)))
			if answer == strings.ToLower(p.A) {
				correct++
			}
		}
		done <- true
	}()

	select {
	case <-done:
		fmt.Printf("\nGood Job!")
	case <-timer.C:
		fmt.Printf("\nTime is up")
	}

	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
}

func loadCsv(CsvFile string) (*os.File, error) {
	file, err := os.Open(CsvFile)

	if err != nil {
		errMsg := errors.New(fmt.Sprintf("failed to open the CSV file: %s", CsvFile))
		return nil, errMsg
	}

	return file, nil
}

func readCsv(csvData *os.File) ([][]string, error) {
	r := csv.NewReader(csvData)
	csvQA, err := r.ReadAll()

	if err != nil {
		errMsg := errors.New(fmt.Sprintf("failed to read the provided CSV file"))
		return nil, errMsg
	}

	return csvQA, nil

}

func shuffle(lines [][]string) [][]string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i1 := 0; i1 < len(lines); i1++ {
		i2 := r.Intn(len(lines) - 1)

		lines[i1], lines[i2] = lines[i2], lines[i1]
	}

	return lines
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))

	for i, line := range lines {
		ret[i] = Problem{

			Q: line[0],
			A: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func exit(msg error) {
	fmt.Println(msg)
	os.Exit(1)
}
