package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var seeds []int
var maps []*Map

func main() {
	loadInput()

	lowestLocation := 0

	for _, seed := range seeds {
		sourceNum := seed
		for _, m := range maps {
			miMatched := false
			for _, mi := range m.MapIndexes {
				if sourceNum >= mi.SourceStart && sourceNum < mi.SourceStart+mi.Range {
					offset := sourceNum - mi.SourceStart
					fmt.Printf("seed=%d matchedSource=%d offset=%d sourceStart=%d offsetDest=%d on map %s\n",
						seed,
						sourceNum,
						offset,
						mi.SourceStart,
						mi.DestinationStart+offset,
						m.RawName,
					)

					sourceNum = mi.DestinationStart + offset
					miMatched = true
					break
				}
			}

			if !miMatched {
				fmt.Printf("seed=%d no match on map %s\n", seed, m.RawName)
			}
		}

		fmt.Printf("seed=%d location=%d\n", seed, sourceNum)
		if lowestLocation == 0 || lowestLocation > sourceNum {
			lowestLocation = sourceNum
		}
	}

	fmt.Printf("lowest location=%d", lowestLocation)
}

func loadInput() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic("could not open input")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "seeds") {
			vals := strings.Split(line, ": ")
			if len(vals) < 2 {
				panic("cannot get seeds")
			}
			for _, v := range strings.Split(vals[1], " ") {
				num, err := strconv.Atoi(v)
				if err != nil {
					panic("str to int")
				}
				seeds = append(seeds, num)
			}
		}

		if strings.Contains(line, "map") {
			processMap(scanner)
		}
	}
}

func processMap(scanner *bufio.Scanner) {
	m := Map{
		RawName: scanner.Text(),
	}

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		// we are on a buffer line between maps
		if len(parts) < 3 {
			break
		}

		mi := MapIndex{}
		mi.DestinationStart = mi.intOrPanic(parts[0])
		mi.SourceStart = mi.intOrPanic(parts[1])
		mi.Range = mi.intOrPanic(parts[2])

		m.MapIndexes = append(m.MapIndexes, &mi)

		if scanner.Text() == "\n" {
			break
		}
	}

	maps = append(maps, &m)
}
