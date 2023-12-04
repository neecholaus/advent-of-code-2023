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

func (l *Line) GetSymbolsOfType(raw string) []*Symbol {
	var matchSyms []*Symbol
	for _, sym := range l.GetSymbols() {
		if sym.Raw == raw {
			matchSyms = append(matchSyms, sym)
		}
	}
	return matchSyms
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

func (n *Number) AsInt() int {
	a, err := strconv.Atoi(n.Raw)
	if err != nil {
		panic("cannot get num as int")
	}

	return a
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
		var numsOnAdjacentLines []*Number
		for _, adjacentLine := range getAdjacentLines(idx) {
			numsOnAdjacentLines = append(numsOnAdjacentLines, adjacentLine.GetNumbers()...)
		}

		for _, sym := range l.GetSymbolsOfType("*") {
			var adjacentNums []*Number

			for _, num := range numsOnAdjacentLines {
				// if more than two symbol is excluded
				if len(adjacentNums) > 2 {
					adjacentNums = nil
					break
				}

				if sym.IndexStart >= num.IndexStart-1 && sym.IndexStart <= num.IndexEnd {
					adjacentNums = append(adjacentNums, num)
				}
			}

			if len(adjacentNums) == 2 {
				gearRatio := adjacentNums[0].AsInt() * adjacentNums[1].AsInt()
				totalSum += gearRatio
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
