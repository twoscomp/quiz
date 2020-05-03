package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	// Parse command-lin
	dur := flag.Int("dur", 30, "duration for the quiz in number of seconds")
	flag.Parse()
	args := flag.Args()

	// Read csv containing questions and answers
	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	bReader := bufio.NewReader(os.Stdin)

	total := len(records)
	cnt := 1
	correct := 0

	// Ask the user if they're ready
	fmt.Print("Hit 'enter' to start the quiz...")
	bReader.ReadString('\n')

	// Start a timer in a separate process
	timer := time.NewTimer(time.Duration(*dur) * time.Second)
	go func() {
		<-timer.C
		fmt.Println("\nTimes up!")
		fmt.Printf("%d of %d Correct, %d Incorrect\n", correct, total, total-correct)
		os.Exit(0)
	}()

	// Ask questions and read from stdin
	for _, record := range records {
		fmt.Printf("Question %d: %s? ", cnt, record[0])

		answer, err := bReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		answer = strings.Trim(answer, "\n")
		if answer == record[1] {
			fmt.Println("Correct")
			correct++
		}
		cnt++
	}
	fmt.Printf("%d of %d Correct, %d Incorrect\n", correct, total, total-correct)
}
