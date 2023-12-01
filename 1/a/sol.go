package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("could not open file")
		return
	}

	defer file.Close()

	calibrationValues := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := allNums(line)
		calibrationValue, err := strconv.Atoi(fmt.Sprintf("%d%d", nums[0], nums[len(nums)-1]))
		if err != nil {
			fmt.Println("error making calibration value")
		}
		calibrationValues = append(calibrationValues, calibrationValue)
	}

	sum := 0
	for _, num := range calibrationValues {
		fmt.Printf("%d - %d\n", sum, num)
		sum += num
	}

	fmt.Println(sum)
}

func allNums(line string) []int {
	nums := []int{}

	r := regexp.MustCompile(`\d`)
	matches := r.FindAllString(line, -1)
	if matches == nil {
		return nums
	}

	for _, char := range matches {
		numConversion, err := strconv.Atoi(char)
		if err != nil {
			fmt.Printf("error converting (%s) to int\n", char)
		}
		nums = append(nums, numConversion)
	}

	return nums
}
