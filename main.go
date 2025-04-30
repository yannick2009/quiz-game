package main

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	DEFAULT_TIME_SECONDS = 30 // Default time limit for the game

	PATH_FLAG_DESCRIPTION  = "Specify csv file path"
	LIMIT_FLAG_DESCRIPTION = "Specify the time limit for completing the game"

	CSV_NOT_FOUND = "We didn't find your csv file"
	WIN_MSG       = "Perfect Score! All answers correct!"
	GAME_OVER_MSG = "Incorrect! Final Score: %d/%d\n"
	TIME_UP_MSG   = "\nTime's Up! Score: %d/%d\n"
)

var (
	//go:embed assets/problems.csv
	fileByte []byte
	//go:embed assets/ascii.txt
	asciiDraw string
)

func main() {
	path := flag.String("path", "", PATH_FLAG_DESCRIPTION)
	limit := flag.Uint("limit", (DEFAULT_TIME_SECONDS), LIMIT_FLAG_DESCRIPTION)

	flag.Parse()

	var score uint
	questions, total := readCSV(*path)

	fmt.Print(asciiDraw)

	timer := time.NewTimer(time.Duration(*limit) * time.Second)

	go func() {
		select {
		case <-timer.C:
			fmt.Printf(TIME_UP_MSG, score, total)
			os.Exit(0)
		}
	}()

	for k, v := range questions {
		var res string
		fmt.Printf("%v = ", k)
		fmt.Fscan(os.Stdin, &res)

		if strings.TrimSpace(res) == v {
			score++
			continue
		}
		fmt.Printf(GAME_OVER_MSG, score, total)
		return
	}

	if !timer.Stop() {
		<-timer.C
	}

	fmt.Println(WIN_MSG)
}

// checkFileExists checks if the file exists at the given path
// and returns true if it does, false otherwise.
func checkFileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// readCSV reads a CSV file from the given path or uses the embedded file
// if no path is provided. It returns a map of questions and answers
func readCSV(path string) (map[string]string, int) {
	var content []byte = fileByte
	if path != "" {
		if checkFileExists(path) {
			f, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			content = f
		} else {
			fmt.Println(CSV_NOT_FOUND)
			os.Exit(0)
		}
	}

	r := csv.NewReader(bytes.NewReader(content))
	records, err := r.ReadAll()

	if err != nil {
		panic(err)
	}

	lenght := len(records)
	result := make(map[string]string, lenght)

	for _, v := range records {
		result[v[0]] = v[1]
	}

	return result, lenght
}
