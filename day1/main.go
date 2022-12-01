package main

import (
	"bufio"
	"log"
	"os"
	"sort"
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
	sort.Slice(totalCalories, func(i, j int) bool {
		return totalCalories[i] > totalCalories[j]
	})

	log.Println(totalCalories[0] + totalCalories[1] + totalCalories[2])
}

func readInput(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
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
