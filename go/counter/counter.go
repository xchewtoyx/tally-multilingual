package counter

type count struct {
    Key string
    Count int
}

type Counter []count

func LoadMap(source map[string]int) *Counter {
    counts := make(Counter, len(source))
    i := 0
    for k, v := range source {
        counts[i] = count{Key: k, Count: v}
        i++
    }
    return &counts
}

func (self Counter) Len() int {
    return len(self)
}

func (self Counter) Less(i, j int) bool {
    return self[i].Count < self[j].Count
}

func (self Counter) Swap(i, j int) {
    self[i], self[j] = self[j], self[i]
}
