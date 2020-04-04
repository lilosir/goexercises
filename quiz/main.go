package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			line[0],
			strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer")
	timeLimit := flag.Int("limit", 10, "a time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %s\n", *csvFileName))
	}

	data, err := csv.NewReader(file).ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to read file: %s\n", *csvFileName))
	}
	problems := parseLines(data)

	timer := time.NewTimer(time.Second * time.Duration(*timeLimit))
	score := 0

problemLoop:
	for i, p := range problems {
		fmt.Printf("#%d: %s", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			// fmt.Printf("\nyour result: %d/%d\n", score, len(problems))
			break problemLoop
		case answer := <-answerCh:
			if strings.TrimSpace(answer) == p.a {
				score++
			}
		}
	}
	fmt.Printf("\nyour result: %d/%d", score, len(problems))
}
