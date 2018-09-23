package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Q string
	A string
}

func main() {
	// `csv` flag to choose csv file to load in command line
	csvFileName := flag.String("csv", "problem.csv", "a csv file in the format of 'question, answer'")
	// `limit` flag to set quiz time limit (in second)
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	// parse flag dari argumen command line
	flag.Parse()

	// open file csv, reference csvFileName using pointer
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	// read csv file
	r := csv.NewReader(file)
	// reads all remaining record from r to end of file,
	// to array 2 dimension (contain item from csv file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	// memasukkan
	problems := parseLines(lines)

	// NewTimer -> create a new timer
	// Duration -> to enable create new timer if NewTimer time is already stored in variable
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	reader := bufio.NewReader(os.Stdin)

	// counter to correct answer
	correct := 0
	// problemloop:

	// Loop and print question
	for i, p := range problems {
		// Print question
		fmt.Printf("Problem #%d: %s = ", i+1, p.Q)

		// Create answer channel
		answerCh := make(chan string)

		// create go routine to accept answer & to not block timer
		go func() {
			// var answer string

			// // get answer from command line
			// fmt.Scanf("%s\n", &answer)
			// fmt.Println(answer)
			// answer = strings.Replace(answer, "\n", "", -1)
			// answer = strings.ToLower(answer)
			// answer = strings.TrimSpace(answer)

			// if answer == p.A {
			// 	fmt.Println("ok")
			// }

			// read and return a strin
			answer, _ := reader.ReadString('\n')

			// return answer to answser channel
			answerCh <- answer
		}()

		// make goroutine wait on multiple operations
		select {

		// jika timer sudah habis
		case <-timer.C:

			// cetak hasil
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))

			// kelular dari problem
			return
			// fmt.Println()
			// break problemloop

		// jika menerima answer
		case answer := <-answerCh:
			// jika answer = jawaban asli, tambah nilai correct
			var newAnswer string
			newAnswer = strings.Replace(answer, "\n", "", -1)
			newAnswer = strings.ToLower(newAnswer)
			newAnswer = strings.TrimSpace(newAnswer)
			if newAnswer == strings.ToLower(p.A) {
				correct++
			}

			// if answer == p.A {
			// 	correct++
			// }
		}

	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []Problem {
	// buat slice dengan tipa struct : Problem, dengan jumlah sesuai panjang array lines
	ret := make([]Problem, len(lines))

	// perulangan untuk mengisi slice ret dengan question dan problem
	for i, line := range lines {
		// isi nilai array dari index 0, dengan struct Problem
		ret[i] = Problem{
			// question(Q) diisi dengan indeks 0 dari arr line

			Q: line[0],
			// answer(A) diisi dengan indeks 1 dari arr line
			A: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// function for exit from the program
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
