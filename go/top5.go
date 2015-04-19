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

func print_top5(counts counter.Counter) {
	for i, c := range counts.Counts {
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
	tally := counter.New()

	go numbers.ReadNumbersFromFile(flag.Args()[0])
	for number := range numbers {
		tally.Tally(number)
	}

	sort.Sort(sort.Reverse(tally))

	print_top5(tally)
}
