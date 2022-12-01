package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sum(s []int) int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	buckets := strings.Split(string(content), "\n\n")
	totalPerBucket := make([]int, len(buckets))

	for i, bucket := range buckets {
		for _, v := range strings.Split(bucket, "\n") {
			c, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			totalPerBucket[i] += c
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totalPerBucket)))
	fmt.Printf("\nMost calories packed by elf: %d\nTotal of top three elves: %d\r\n", totalPerBucket[0], sum(totalPerBucket[:3]))
}
