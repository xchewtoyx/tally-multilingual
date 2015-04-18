package main

import (
	"./counter"
	"fmt"
	"os"
	"sort"
)

func main() {
	numbers := make(counter.NumberList)
	tally := make(map[string]int)

	go numbers.ReadNumbersFromFile(os.Args[1])

	for number := range numbers {
		tally[number] += 1
	}

	counts := counter.LoadMap(tally)
	sort.Sort(sort.Reverse(counts))

	for i, c := range *counts {
		if i > 4 {
			break
		}
		fmt.Printf("%4s %d\n", c.Key, c.Count)
	}
}
