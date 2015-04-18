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

func (self NumberList) ReadNumbers(buffer *bufio.Reader) {
    defer close(self)
    for {
        line, err := buffer.ReadString('\n')
        if err == io.EOF {
            break
        }
        self <- strings.TrimFunc(line, chomp)
    }
}

func (self NumberList) ReadNumbersFromFile(filename string) {
    infile, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer infile.Close()
    buffer := bufio.NewReader(infile)
    self.ReadNumbers(buffer)
}
