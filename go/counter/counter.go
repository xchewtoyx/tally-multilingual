package counter

type Count struct {
    Key string
    Count int
}

type Counter []Count

func LoadMap(source map[string]int) *Counter {
    counts := make(Counter, len(source))
    i := 0
    for k, v := range source {
        counts[i] = Count{Key: k, Count: v}
        i++
    }
    return &counts
}

func (p Counter) Len() int {
    return len(p)
}

func (p Counter) Less(i, j int) bool {
    return p[i].Count < p[j].Count
}

func (p Counter) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
