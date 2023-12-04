package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var lines []*Line

type Line struct {
	Raw string
}

func (l *Line) GetSymbols() []*Symbol {
	r := regexp.MustCompile(`[^\d.]`)
	idxRanges := r.FindAllStringIndex(l.Raw, -1)

	var syms []*Symbol

	for _, idxRange := range idxRanges {
		sym := Symbol{
			Raw:        l.Raw[idxRange[0]:idxRange[1]],
			IndexStart: idxRange[0],
			Line:       l,
		}
		syms = append(syms, &sym)
	}

	return syms
}

func (l *Line) GetNumbers() []*Number {
	r := regexp.MustCompile(`(\d+)`)
	idxRanges := r.FindAllStringIndex(l.Raw, -1)

	var nums []*Number

	for _, idxRange := range idxRanges {
		num := Number{
			Raw:        l.Raw[idxRange[0]:idxRange[1]],
			IndexStart: idxRange[0],
			IndexEnd:   idxRange[1],
			Line:       l,
		}

		nums = append(nums, &num)
	}

	return nums
}

type Number struct {
	Raw        string
	IndexStart int
	IndexEnd   int
	Line       *Line
}

type Symbol struct {
	Raw        string
	IndexStart int
	Line       *Line
}

func main() {
	loadLines()

	totalSum := 0

	for idx, l := range lines {
		// Make list of symbols on adjacent lines
		var adjacentSymbols []*Symbol
		for _, adjacentLine := range getAdjacentLines(idx) {
			adjacentSymbols = append(adjacentSymbols, adjacentLine.GetSymbols()...)
		}

		nums := l.GetNumbers()
		var numsWithAdjacentSymbol []*Number

		for _, num := range nums {
			adjacentStart := num.IndexStart - 1
			if adjacentStart < 0 {
				adjacentStart = 0
			}
			adjacentEnd := num.IndexEnd

			// Is there an adjacent symbol?
			for _, sym := range adjacentSymbols {
				if sym.IndexStart >= adjacentStart && sym.IndexStart <= adjacentEnd {
					numsWithAdjacentSymbol = append(numsWithAdjacentSymbol, num)

					numAsInt, err := strconv.Atoi(num.Raw)
					if err != nil {
						panic("cannot get num as int")
					}
					totalSum += numAsInt
					break
				}
			}
		}
	}

	fmt.Printf("sum=%d\n", totalSum)
}

func loadLines() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic("could not open file")
	}

	s := bufio.NewScanner(file)

	for s.Scan() {
		line := Line{
			Raw: s.Text(),
		}

		lines = append(lines, &line)
	}
}

func getAdjacentLines(lineIdx int) []*Line {
	var aLines []*Line

	if lineIdx > 0 {
		aLines = append(aLines, lines[lineIdx-1])
	}

	aLines = append(aLines, lines[lineIdx])

	if lineIdx < len(lines)-1 {
		aLines = append(aLines, lines[lineIdx+1])
	}

	return aLines
}
