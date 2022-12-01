package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readInput("/Users/emma/code/advent-of-code-2022/day1/input1.txt")

	var index = 0
	var totalCalories = make([]int, len(input)) // Wasteful
	for i := range input {
		row := input[i]
		if row == "" {
			index++
		} else {
			calories, err := strconv.Atoi(row)
			if err != nil {
				log.Fatal(err)
			}
			totalCalories[index] += calories
		}
	}

	log.Println(findMaxCalories(totalCalories))
}

func findMaxCalories(totalCalories []int) int {
	var max = totalCalories[0]
	for i := range totalCalories {
		if totalCalories[i] > max {
			max = totalCalories[i]
		}
	}
	return max
}

func readInput(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println(input)
	return input
}
