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

	for gameIdx, game := range games {
		distanceToBeat := (*game)["Distance"]
		timeAllowed := (*game)["Time"]

		numWaysWeCanWin := 0

		for i := 1; i <= timeAllowed; i++ {
			chargeTime := i
			timeRemainingAfterCharge := timeAllowed - chargeTime
			distanceFromCharge := timeRemainingAfterCharge * chargeTime
			if distanceFromCharge > distanceToBeat {
				numWaysWeCanWin++
				fmt.Printf("Game %d could charge %d to beat %d going %d\n",
					gameIdx,
					chargeTime,
					distanceToBeat,
					distanceFromCharge,
				)
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

		values := strings.Split(labelSplit[1], " ")

		gameIdx := 0
		for _, val := range values {
			if val == "" {
				continue
			}

			// init game - index is important
			if len(games) < gameIdx+1 {
				games = append(games, &Game{})
			}

			num := intOrPanic(val)

			game := games[gameIdx]
			(*game)[label] = num

			// not using loop index because we will hit many subsequent empty values
			gameIdx++
		}
	}
}

func intOrPanic(numAsString string) int {
	num, err := strconv.Atoi(numAsString)
	if err != nil {
		panic("cannot make int")
	}

	return num
}
