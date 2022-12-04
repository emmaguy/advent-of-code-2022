package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readInput("/Users/emma/code/advent-of-code-2022/day2/input.txt")

	total := playGame(input)
	log.Printf("total: %d", total)
}

func playGame(input []string) int {
	var total = 0
	for i := range input {
		opponentPlays := input[i][0]
		youPlay := input[i][2]

		roundScore := 0
		if opponentPlays == 'A' { // Rock
			if youPlay == 'X' { // Rock
				roundScore += 3 + 1 // Draw + points for rock
			} else if youPlay == 'Y' { // Paper
				roundScore += 6 + 2 // Win + points for paper
			} else if youPlay == 'Z' { // Scissors
				roundScore += 3 // Loss, points for scissors
			} else {
				log.Fatal(fmt.Sprintf("Unexpected character, %x", opponentPlays))
			}
		} else if opponentPlays == 'B' { // Paper
			if youPlay == 'X' { // Rock
				roundScore += 0 + 1 // Loss, points for rick
			} else if youPlay == 'Y' { // Paper
				roundScore += 3 + 2 // Draw + points for paper
			} else if youPlay == 'Z' { // Scissors
				roundScore += 6 + 3 // Win + points for scissors
			} else {
				log.Fatal(fmt.Sprintf("Unexpected character, %x", opponentPlays))
			}
		} else if opponentPlays == 'C' { // Scissors
			if youPlay == 'X' { // Rock
				roundScore += 6 + 1 // Win + points for rock
			} else if youPlay == 'Y' { // Paper
				roundScore += 0 + 2 // Loss, points for paper
			} else if youPlay == 'Z' { // Scissors
				roundScore += 3 + 3 // Draw + points for scissors
			} else {
				log.Fatal(fmt.Sprintf("Unexpected character, %x", opponentPlays))
			}
		}

		total += roundScore
		log.Printf("total: %d", total)
	}
	return total
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
	return input
}
