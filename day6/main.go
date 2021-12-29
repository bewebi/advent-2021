package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()

	fishCounts := make([]int, 9)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fishStrs := strings.Split(scanner.Text(), ",")
		for _, s := range fishStrs {
			fish, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatalf("%v", err)
			}
			fishCounts[fish]++
		}
	}

	for i := 0; i < 256; i++ {
		newFish := make([]int, 9)
		for d := 1; d < len(fishCounts); d++ {
			newFish[d-1] = fishCounts[d]
		}
		newFish[6] += fishCounts[0]
		newFish[8] += fishCounts[0]

		fishCounts = newFish

		if i == 79 {
			totalCount := 0
			for _, cnt := range fishCounts {
				totalCount += cnt
			}

			fmt.Printf("fish counts after 80 days: %v; total fish: %d\n", fishCounts, totalCount)
		}
	}

	totalCount := 0
	for _, cnt := range fishCounts {
		totalCount += cnt
	}

	fmt.Printf("fish counts after 256 days: %v; total fish: %d\n", fishCounts, totalCount)
}
