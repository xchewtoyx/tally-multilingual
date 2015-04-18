package counter

import (
    "bufio"
    "io"
    "log"
    "os"
    "strings"
)

type NumberList chan string

func (self NumberList) ReadNumbers(buffer *bufio.Reader) {
    defer close(self)
    for {
        line, err := buffer.ReadString('\n')
        if err == io.EOF {
            break
        }
        self <- strings.Trim(line, "\n")
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
