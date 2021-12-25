package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	value int
}

func main() {
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()

	var input []instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("%v", err)
		}
		input = append(input, instruction{parts[0],val})
	}

	var depth, distance int
	for _, in := range input {
		switch in.direction {
		case "forward":
			distance += in.value
		case "up":
			depth -= in.value
		case "down":
			depth += in.value
		}
	}
	fmt.Printf("distance: %d, depth: %d, multiplied: %d\n", distance, depth, distance*depth)

	var aim, aimDepth int
	for _, in := range input {
		switch in.direction {
		case "forward":
			aimDepth += aim*in.value
		case "up":
			aim -= in.value
		case "down":
			aim += in.value
		}
	}
	fmt.Printf("distance: %d, *aimed* depth: %d, multiplied: %d\n", distance, aimDepth, distance*aimDepth)
}
