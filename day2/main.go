package day2

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readInput("/Users/emma/code/advent-of-code-2022/day2/sample.txt")
	points, err := points(input[0][0])
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("points %d")
}

func points(character byte) (int, error) {
	if character == 'A' {
		return 1, nil
	} else if character == 'B' {
		return 2, nil
	} else if character == 'C' {
		return 3, nil
	}
	return 0, errors.New(fmt.Sprintf("Unexpected points %x", character))
}

// 0 if you lost, 3 if the round was a draw, and 6 if you won

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
