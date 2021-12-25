package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()

	var input []int
	scanner := bufio.NewScanner(file)
	incs := 0
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("%v", err)
		}
		input = append(input, i)
	}

	prev := math.MaxInt32
	for _, cur := range input {
		if cur > prev {
			incs++
		}
		prev = cur
	}
	fmt.Printf("there were %d increases\n", incs)

	prev = math.MaxInt32
	threeIncs := 0
	for i, cur := range input {
		if i < 2 {
			continue
		}
		threeSum := input[i-2] + input[i-1] + cur
		if threeSum > prev {
			threeIncs++
		}
		prev = threeSum
	}
	fmt.Printf("there were %d three sum increases\n", threeIncs)
}
