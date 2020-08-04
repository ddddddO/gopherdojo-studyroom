package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	limitTime = 60 * time.Second
)

func main() {
	flag.DurationVar(&limitTime, "l", limitTime, "specify time limit duration")
	flag.Parse()

	endQ := make(chan struct{})
	go run(endQ)

	after := time.After(limitTime)
END:
	for {
		select {
		case <-after:
			fmt.Println("time up!")
			break END
		case <-endQ:
			fmt.Println("clear!")
			break END
		}
	}

	fmt.Printf("game result: %d\n", count)
}

var words = []string{
	"apple",
	"banana",
	"cherry",
	"dorian",
	"peach",
}

var (
	stdin = bufio.NewScanner(os.Stdin)
	count = 0
)

func run(endQ chan<- struct{}) {
	for _, word := range words {
		fmt.Printf("      %s      \n", word)
		for stdin.Scan() {
			inputWord := stdin.Text()
			if inputWord == word {
				count++
				break
			} else {
				fmt.Println("bad...")
			}
		}
	}
	endQ <- struct{}{}
}
