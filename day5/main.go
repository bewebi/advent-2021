package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	diagonals = flag.Bool("diagonals", false, "wherer or not to process diagonal lines")
)
type line struct {
	x1, y1, x2, y2 int64
}

func main() {
	flag.Parse()
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()

	var lines []line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		coordStrs := strings.Split(scanner.Text(), " -> ")
		c1Strs := strings.Split(coordStrs[0], ",")
		c2Strs := strings.Split(coordStrs[1], ",")
		x1, err := strconv.ParseInt(c1Strs[0], 10, 64)
		if err != nil {
			log.Fatalf("%v", err)
		}
		y1, err := strconv.ParseInt(c1Strs[1], 10, 64)
		if err != nil {
			log.Fatalf("%v", err)
		}
		x2, err := strconv.ParseInt(c2Strs[0], 10, 64)
		if err != nil {
			log.Fatalf("%v", err)
		}
		y2, err := strconv.ParseInt(c2Strs[1], 10, 64)
		if err != nil {
			log.Fatalf("%v", err)
		}
		lines = append(lines, line{x1,y1,x2,y2})
	}

	plot := map[string]int{}
	for i, l := range lines {
		_ = i
		switch {
		case l.x1 == l.x2:
			var lowY, highY int64
			if l.y1 < l.y2 {
				lowY, highY = l.y1, l.y2
			} else {
				lowY, highY = l.y2, l.y1
			}
			for i := lowY; i <= highY; i++ {
				plot[fmt.Sprintf("%d,%d", l.x1, i)]++
			}
		case l.y1 == l.y2:
			var lowX, highX int64
			if l.x1 < l.x2 {
				lowX, highX = l.x1, l.x2
			} else {
				lowX, highX = l.x2, l.x1
			}
			for i := lowX; i <= highX; i++ {
				plot[fmt.Sprintf("%d,%d", i, l.y1)]++
			}
		default:
			if !*diagonals {
				continue
			}
			xSlope, ySlope := int64(1), int64(1)
			if l.x1 > l.x2 {
				xSlope = -1
			}
			if l.y1 > l.y2 {
				ySlope = -1
			}
			for i := int64(0); i <= int64(math.Abs(float64(l.x1-l.x2))); i++ {
				point := fmt.Sprintf("%d,%d", l.x1+(i*xSlope), l.y1+(i*ySlope))
				plot[point]++
			}
		}
	}
	var intersectCount int
	for _, cnt := range plot {
		if cnt > 1 {
			intersectCount++
		}
	}

	fmt.Printf("intersection point count: %d", intersectCount)
}
