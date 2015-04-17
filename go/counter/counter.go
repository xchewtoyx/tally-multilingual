package counter

type Count struct {
    Key string
    Count int
}

type Counter []Count

func (p Counter) LoadMap(source map[string]int) Counter {
    for k, v := range source {
        p = append(p, Count{Key: k, Count: v})
    }
    return p
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
