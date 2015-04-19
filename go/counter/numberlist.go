package counter

import (
    "bufio"
    "io"
    "log"
    "os"
    "strings"
)

type NumberList chan string

func chomp(r rune) bool { return r == '\n' }

func (self NumberList) ReadLines(infile io.Reader) {
    buffer := bufio.NewReader(infile)
    defer close(self)
    for {
        line, err := buffer.ReadString('\n')
        if err == io.EOF {
            break
        }
        self <- strings.TrimFunc(line, chomp)
    }
}

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
