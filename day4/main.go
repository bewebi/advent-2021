package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type spot struct {
	value int64
	called bool
}

type board struct {
	spots [][]*spot
	won bool
}

func main() {
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var calls []int64
	for _, s := range strings.Split(scanner.Text(), ",") {
		call, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatalf("%v", err)
		}
		calls = append(calls, call)
	}

	var boards []*board
	for scanner.Scan() {
		rowStr := scanner.Text()
		if rowStr == "" {
			boards = append(boards, &board{})
			continue
		}
		var row []*spot
		for _, numStr := range strings.Split(rowStr, " ") {
			if numStr == "" {
				continue
			}
			num, err := strconv.ParseInt(numStr, 10, 64)
			if err != nil {
				log.Fatalf("%v", err)
			}
			row = append(row, &spot{num, false})
		}
		boards[len(boards)-1].spots = append(boards[len(boards)-1].spots, row)
	}

	var firstWinner, lastWinner *board
	var firstWinningCall, lastWinningCall int64
	for _, call := range calls {
		for _, b := range boards {
			if b.won {
				continue
			}
			b.mark(call)
			if b.winner() {
				b.won = true
				if firstWinner == nil {
					firstWinner = b
					firstWinningCall = call
				}
				lastWinner = b
				lastWinningCall = call
			}
		}
	}

	fmt.Printf("first winning board: %v, first winning call: %d, score: %d\n", firstWinner, firstWinningCall, firstWinner.uncalledSum()*firstWinningCall)
	fmt.Printf("last winning board: %v, last winning call: %d, score: %d\n", lastWinner, lastWinningCall, lastWinner.uncalledSum()*lastWinningCall)
}

func (b *board) mark(i int64) {
	for _, col := range b.spots {
		for _, s := range col {
			if s.value == i {
				s.called = true
				return
			}
		}
	}
}

func (b *board) winner() bool {
	for _, col := range b.spots {
		colWin := true
		for _, s := range col {
			if !s.called {
				colWin = false
				break
			}
		}
		if colWin {
			return colWin
		}
	}
	for i := range b.spots {
		rowWin := true
		for _, col := range b.spots {
			if !col[i].called {
				rowWin = false
				break
			}
		}
		if rowWin {
			return rowWin
		}
	}
	return false
}

func (b *board) uncalledSum() int64 {
	var sum int64
	for _, col := range b.spots {
		for _, s := range col {
			if !s.called {
				sum += s.value
			}
		}
	}
	return sum
}