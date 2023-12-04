package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var lines []*Line

type Line struct {
	Raw           string
	Id            int
	WinningNums   []int
	NumsYouHave   []int
	MatchingCount int
	Copies        int
}

func (l *Line) ParseRaw() {
	r := regexp.MustCompile(`Card\s+(?P<id>\d+):\s(?P<winning>[\d\s]+)\|\s(?P<have>[\d\s]+)`)
	matches := r.FindStringSubmatch(l.Raw)

	// id
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		panic("cannot convert id")
	}
	l.Id = id

	// winning nums
	for _, num := range strings.Split(matches[2], " ") {
		if num == "" {
			continue
		}
		numAsInt, err := strconv.Atoi(num)
		if err != nil {
			panic("cannot convert winning num")
		}
		l.WinningNums = append(l.WinningNums, numAsInt)
	}

	// nums we have
	for _, num := range strings.Split(matches[3], " ") {
		if num == "" {
			continue
		}
		numAsInt, err := strconv.Atoi(num)
		if err != nil {
			panic("cannot convert num we have")
		}
		l.NumsYouHave = append(l.NumsYouHave, numAsInt)
	}
}

func (l *Line) CountMatches() int {
	l.MatchingCount = 0

	for _, num := range l.NumsYouHave {
		if slices.Contains(l.WinningNums, num) {
			l.MatchingCount += 1
		}
	}

	return l.MatchingCount
}

func main() {
	loadLines()

	totalSum := 0

	for _, l := range lines {
		l.CountMatches()

		fmt.Printf("line %d has %d matching and %d copies\n", l.Id, l.MatchingCount, l.Copies)

		for i := 0; i < l.MatchingCount; i++ {
			lines[l.Id+i].Copies += 1 + l.Copies
			fmt.Printf("-- adding copy to %d; now %d copies\n", lines[l.Id+i].Id, lines[l.Id+i].Copies)
		}

		totalSum += l.Copies + 1
		fmt.Printf("adding=%d sumSoFar=%d\n", l.Copies+1, totalSum)
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
		line.ParseRaw()

		lines = append(lines, &line)
	}
}
