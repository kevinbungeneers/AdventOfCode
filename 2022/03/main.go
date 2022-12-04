package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(strings.Split(string(content), "\n")))

	re, err := regexp.Compile(".*\n.*\n.*\n?")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partTwo(re.FindAllString(string(content), -1)))
}

func partOne(s []string) int {
	priority := 0
	for _, line := range s {
		half := len(line) / 2
		commonChars := findCommonChars(line[0:half], line[half:])
		priority += calculatePriority(commonChars)
	}

	return priority
}

func partTwo(s []string) int {
	priority := 0
	for _, grouped := range s {
		parts := strings.Split(grouped, "\n")
		commonChars := findCommonChars(parts[0], parts[1])
		commonChars = findCommonChars(string(commonChars), parts[2])
		priority += calculatePriority(commonChars)
	}

	return priority
}

func findCommonChars(a string, b string) []rune {
	checkedChars := make(map[rune]bool)
	commonChars := make([]rune, 0)

	for _, char := range a {
		if strings.Contains(b, string(char)) && checkedChars[char] == false {
			checkedChars[char] = true
			commonChars = append(commonChars, char)
		}
	}

	return commonChars
}

func calculatePriority(r []rune) int {
	sum := 0
	for _, r2 := range r {
		if unicode.IsUpper(r2) { // 65 - 90
			sum += int(r2 - 38)
		} else if unicode.IsLower(r2) { // 97 - 122
			sum += int(r2 - 96)
		}
	}

	return sum
}
