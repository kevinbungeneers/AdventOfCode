package main

// Each possible move is represented using a binary value:
// A, rock:    00000000
// B, paper:   00000001
// C, scissor: 00000010
// We do a bitwise OR with 4 (0100) to our own move, so we can differentiate between the two moves.

// So, for instance (opponent v me):
// rock v paper => 0 (0000) vs 5 (0100 OR 0001)
// 5 - 0 = 0
//
// paper v rock => 1 (0001) vs 4 (0100 OR 0000)
// 4 - 1 = 3
//
// scissors v scissors => 2 (0010) vs 6 (0100 OR 0010)
// 6 - 2 = 4
//
// rock vs scissors => 0 (0000) vs 6 (0100 OR 0100)
// 6 - 0 = 6
//
// From this pattern we can conclude that in case of a draw, the outcome is 4. If the outcome is divisible by 3,
// it means that I've lost. Any other outcome means that I've won.

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var valueMap = map[string]int{
	"A": 0, // 000
	"B": 1, // 001
	"C": 2, // 010

	"X": 0, // 000
	"Y": 1, // 001
	"Z": 2, // 010
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scorePartOne := 0
	scorePartTwo := 0
	for _, round := range strings.Split(string(content), "\n") {
		choices := strings.Split(round, " ")
		scorePartOne += calcPartOne(choices)
		scorePartTwo += calcPartTwo(choices)
	}

	fmt.Printf("Score for part one: %d\nScore for part two: %d\n", scorePartOne, scorePartTwo)
}

func game(myChoice int, opponentChoice int) int {
	match := (myChoice | 4) - opponentChoice
	if match == 4 {
		return myChoice + 1 + 3
	} else if match%3 == 0 {
		return myChoice + 1
	} else {
		return myChoice + 1 + 6
	}
}

func calcPartOne(choices []string) int {
	return game(valueMap[choices[1]], valueMap[choices[0]])
}

func calcPartTwo(choices []string) int {
	var myChoice int
	opponentChoice := valueMap[choices[0]]

	switch choices[1] {
	case "X": // Lose
		myChoice = opponentChoice - 1
		if myChoice < 0 {
			myChoice = 2
		}
	case "Z": // Win
		myChoice = opponentChoice + 1
		if myChoice > 2 {
			myChoice = 0
		}
	case "Y": // Draw
		myChoice = opponentChoice
	}

	return game(myChoice, opponentChoice)
}
