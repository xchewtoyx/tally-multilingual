package main

import (
	"./counter"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func print_top5(counts *counter.Counter) {
	for i, c := range *counts {
		if i > 4 {
			break
		}
		fmt.Printf("%4s %d\n", c.Key, c.Count)
	}
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	numbers := make(counter.NumberList)
	tally := make(map[string]int)

	go numbers.ReadNumbersFromFile(flag.Args()[0])
	for number := range numbers {
		tally[number] += 1
	}

	counts := counter.LoadMap(tally)
	sort.Sort(sort.Reverse(counts))

	print_top5(counts)
}
