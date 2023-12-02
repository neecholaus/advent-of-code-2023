package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("could not open file")
		return
	}

	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := allNums(line)

		if len(nums) < 1 {
			continue
		}

		calibrationValue, err := strconv.Atoi(fmt.Sprintf("%d%d", nums[0], nums[len(nums)-1]))
		if err != nil {
			fmt.Println("error making calibration value")
		}
		fmt.Printf("%d, %d, %s, %v\n", sum, calibrationValue, line, nums)
		sum += calibrationValue
	}

	fmt.Println(sum)
}

func allNums(line string) []int {
	matchedIndexes := map[int]int{}
	for _, spelling := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
		r := regexp.MustCompile(spelling)
		idxs := r.FindAllStringIndex(line, -1)
		if idxs == nil {
			continue
		}

		num, err := getIntFromStrRep(spelling)
		if err != nil {
			fmt.Printf("no int from spelled out (%s)\n", spelling)
			continue
		}

		for _, idx := range idxs {
			matchedIndexes[idx[0]] = num
		}
	}

	var sortedIndexes []int
	for key := range matchedIndexes {
		sortedIndexes = append(sortedIndexes, key)
	}
	sort.Ints(sortedIndexes)

	var sortedNums []int
	for _, idx := range sortedIndexes {
		sortedNums = append(sortedNums, matchedIndexes[idx])
	}

	return sortedNums
}

func getIntFromStrRep(match string) (int, error) {
	if len(match) < 2 {
		num, _ := strconv.Atoi(match)
		return num, nil
	}

	switch match {
	case "one":
		return 1, nil
	case "two":
		return 2, nil
	case "three":
		return 3, nil
	case "four":
		return 4, nil
	case "five":
		return 5, nil
	case "six":
		return 6, nil
	case "seven":
		return 7, nil
	case "eight":
		return 8, nil
	case "nine":
		return 9, nil
	}
	return 0, errors.New("could not determine number")
}
