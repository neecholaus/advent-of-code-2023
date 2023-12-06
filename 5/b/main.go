package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var seedRanges [][]int
var maps []*Map

func main() {
	/**
	 *
	 * Slow, but it works.
	 * Didn't even use channels for async.
	 * Could have done that.
	 * Would have done that, if I was getting paid.
	 *
	 */

	loadInput()

	lowestLocation := 0

	for _, seedRange := range seedRanges {
		for i := seedRange[0]; i < seedRange[0]+seedRange[1]; i++ {
			seed := i
			sourceNum := seed
			for _, m := range maps {
				for _, mi := range m.MapIndexes {
					if sourceNum >= mi.SourceStart && sourceNum < mi.SourceStart+mi.Range {
						offset := sourceNum - mi.SourceStart
						sourceNum = mi.DestinationStart + offset
						break
					}
				}
			}

			if lowestLocation == 0 || lowestLocation > sourceNum {
				lowestLocation = sourceNum
			}
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
			processSeeds(scanner)
		}

		if strings.Contains(line, "map") {
			processMap(scanner)
		}
	}
}

func processSeeds(scanner *bufio.Scanner) {
	vals := strings.Split(scanner.Text(), ": ")
	if len(vals) < 2 {
		panic("cannot get seeds")
	}

	var r []int

	for _, v := range strings.Split(vals[1], " ") {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic("str to int")
		}
		r = append(r, num)

		// range hydrated, add seed list
		if len(r) >= 2 {
			seedRanges = append(seedRanges, r)
			r = []int{}
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
