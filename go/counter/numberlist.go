package counter

import (
    "bufio"
    "io"
    "log"
    "os"
)

type NumberList chan string

func (self NumberList) ScanLines(infile io.Reader) {
    scanner := bufio.NewScanner(infile)
    defer close(self)
    for scanner.Scan() {
        self <- scanner.Text()
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func (self NumberList) ReadNumbersFromFile(filename string) {
    infile, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer infile.Close()
    self.ScanLines(infile)
}
