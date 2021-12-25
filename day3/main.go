package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type biCnt struct {
	ones, zeroes int
}

func main() {
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	gammaCounts := make([]biCnt, len(input[0]))
	for _, in := range input {
		for i, char := range in {
			if char == '1' {
				gammaCounts[i].ones++
			} else {
				gammaCounts[i].zeroes++
			}
		}
	}

	var gamma, epsilon int
	for i, cnt := range gammaCounts {
		if cnt.ones == cnt.zeroes {
			log.Fatalf("counts equal!")
		}
		if cnt.ones > cnt.zeroes {
			gamma += int(math.Pow(2.0, float64(len(gammaCounts)-i-1)))
		} else {
			epsilon += int(math.Pow(2.0, float64(len(gammaCounts)-i-1)))
		}
	}

	fmt.Printf("gamma: %d, epsilon: %d, power consumption: %d\n", gamma, epsilon, gamma*epsilon)

	oxygenCandidates, co2Candidates := input, input
	for i := 0; i < len(input[0]); i++ {
		if len(oxygenCandidates) > 1 {
			var ocOnes, ocZeroes []string
			for _, oc := range oxygenCandidates {
				if oc[i] == '1' {
					ocOnes = append(ocOnes, oc)
				} else {
					ocZeroes = append(ocZeroes, oc)
				}
			}
			if len(ocOnes) >= len(ocZeroes) {
				oxygenCandidates = ocOnes
			} else {
				oxygenCandidates = ocZeroes
			}
		}
		if len(co2Candidates) > 1 {
			var ccOnes, ccZeros []string
			for _, cc := range co2Candidates {
				if cc[i] == '1' {
					ccOnes = append(ccOnes, cc)
				} else {
					ccZeros = append(ccZeros, cc)
				}
			}
			if len(ccOnes) < len(ccZeros) {
				co2Candidates = ccOnes
			} else {
				co2Candidates = ccZeros
			}
		}
	}

	oxygenRating, err := strconv.ParseInt(oxygenCandidates[0], 2, 64)
	if err != nil {
		log.Fatalf("%v", err)
	}
	co2Rating, err := strconv.ParseInt(co2Candidates[0], 2, 64)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("oxygen: %d, co2: %d, life support: %d\n", oxygenRating, co2Rating, oxygenRating*co2Rating)
}
