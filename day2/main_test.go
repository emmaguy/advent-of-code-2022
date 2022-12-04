package main

import "testing"

func TestDay2_Part1(t *testing.T) {
	input := []string{
		"A Y",
		"B X",
		"C Z",
	}

	expected := 15
	actual := playGame(input)

	if actual != expected {
		t.Fatal("got", actual, "wanted", expected)
	}
}
