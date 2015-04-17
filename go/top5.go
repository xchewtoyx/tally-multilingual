package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "sort"
    "strings"
    "./counter"
)

func read_numbers(filename string) []string {
    var numberlist []string

    infile, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer infile.Close()

    buffer := bufio.NewReader(infile)
    for {
        line, err := buffer.ReadString('\n')
        if err == io.EOF {
            break
        }
        numberlist = append(numberlist, strings.Trim(line, "\n"))
    }

    return numberlist
}

func main() {
    tally := make(map[string]int)

    numbers := read_numbers(os.Args[1])
    for i:= 0; i < len(numbers); i++ {
        tally[numbers[i]] += 1
    }

    counts := counter.Counter{}
    counts = counts.LoadMap(tally)
    sort.Sort(sort.Reverse(counts))

	for i, c := range counts {
        if i > 4 {
            break
        }
        fmt.Printf("%4s %d\n", c.Key, c.Count)
    }
}
