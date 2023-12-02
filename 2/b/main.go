package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Id        int
	Raw       string
	MaxCounts []*ColorMax
}

func (g *Game) GetColor(color string) *ColorMax {
	for _, cm := range g.MaxCounts {
		if cm.Name == color {
			return cm
		}
	}

	return nil
}

func (g *Game) PrintAllColors() {
	for _, c := range g.MaxCounts {
		fmt.Printf("%v", c)
	}
}

func (g *Game) ColorPower() int {
	product := 1
	for _, color := range g.MaxCounts {
		product *= color.MaxAtOnce
	}
	return product
}

type ColorMax struct {
	Name      string
	MaxAtOnce int
}

func (cm *ColorMax) AreAtLeast(n int) {
	if n > cm.MaxAtOnce {
		cm.MaxAtOnce = n
	}
}

var games []*Game

func main() {
	loadAndParseGames()

	sum := 0
	for _, game := range games {
		game.PrintAllColors()
		sum += game.ColorPower()
		fmt.Printf("\n%d\n", sum)
	}

	fmt.Printf("sum=%d\n", sum)
}

func loadAndParseGames() {
	rawGames := loadGames()
	parseManyRawGames(rawGames)
}

func loadGames() []string {
	rawGames := []string{}

	input, err := os.Open("../input.txt")
	if err != nil {
		panic("could not read input")
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		rawGames = append(rawGames, scanner.Text())
	}

	return rawGames
}

func parseManyRawGames(rawGames []string) {
	for _, rawGame := range rawGames {
		games = append(games, parseRawGame(rawGame))
	}
}

func parseRawGame(rawGame string) *Game {
	game := Game{}

	id := strings.Split(rawGame, " ")[1]
	idAsInt, err := strconv.Atoi(strings.Split(id, ":")[0])
	if err != nil {
		panic("bad game id")
	}

	game.Raw = rawGame

	game.Id = idAsInt

	reveals := strings.Split(rawGame, ": ")[1]
	for _, reveal := range strings.Split(reveals, "; ") {
		colors := strings.Split(reveal, ", ")

		for _, color := range colors {
			count := strings.Split(color, " ")[0]
			countAsInt, err := strconv.Atoi(count)
			if err != nil {
				panic("bad count")
			}

			name := strings.Split(color, " ")[1]

			if game.GetColor(name) != nil {
				game.GetColor(name).AreAtLeast(countAsInt)
			} else {
				game.MaxCounts = append(game.MaxCounts, &ColorMax{
					Name:      name,
					MaxAtOnce: countAsInt,
				})
			}
		}
	}

	return &game
}
