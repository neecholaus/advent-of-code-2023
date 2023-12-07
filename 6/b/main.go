package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game map[string]int

var games []*Game

func main() {
	loadInput()

	numWaysToWinEachGame := []int{}

	for _, game := range games {
		distanceToBeat := (*game)["Distance"]
		timeAllowed := (*game)["Time"]

		numWaysWeCanWin := 0

		for i := 1; i <= timeAllowed; i++ {
			chargeTime := i
			timeRemainingAfterCharge := timeAllowed - chargeTime
			distanceFromCharge := timeRemainingAfterCharge * chargeTime
			if distanceFromCharge > distanceToBeat {
				numWaysWeCanWin++
				//fmt.Printf("Game %d could charge %d to beat %d going %d\n",
				//	gameIdx,
				//	chargeTime,
				//	distanceToBeat,
				//	distanceFromCharge,
				//)
			}
		}

		numWaysToWinEachGame = append(numWaysToWinEachGame, numWaysWeCanWin)
	}

	fmt.Printf("waysToBeachEachGame=%+v\n", numWaysToWinEachGame)
	product := 1
	for _, num := range numWaysToWinEachGame {
		product *= num
	}

	fmt.Printf("result=%d\n", product)
}

func loadInput() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic("cannot open input")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		labelSplit := strings.Split(line, ": ")
		label := labelSplit[0]

		gameIdx := 0

		singleVal := strings.Replace(labelSplit[1], " ", "", -1)

		if len(games) == 0 {
			games = append(games, &Game{})
		}

		num := intOrPanic(singleVal)

		game := games[gameIdx]
		(*game)[label] = num
	}
}

func intOrPanic(numAsString string) int {
	num, err := strconv.Atoi(numAsString)
	if err != nil {
		panic("cannot make int")
	}

	return num
}
