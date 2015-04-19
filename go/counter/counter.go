package counter

type count struct {
    Key string
    Count int
}

type Counter struct {
    Counts []count
    KeyIndex map[string]int
}

func New() Counter {
    var self Counter
    self.KeyIndex = make(map[string]int)
    return self
}

func NewFromMap(source map[string]int) Counter {
    var self Counter
    self.Counts = make([]count, len(source))
    self.KeyIndex = make(map[string]int)
    i := 0
    for key, value := range source {
        self.Counts[i] = count{Key: key, Count: value}
        self.KeyIndex[key] = i
        i++
    }
    return self
}

func (self *Counter) Tally(key string) {
    if val, ok := self.KeyIndex[key]; ok {
        self.Counts[val].Count += 1
    } else {
        self.KeyIndex[key] = len(self.Counts)
        self.Counts = append(self.Counts, count{Key: key, Count: 1})
    }
}

// The following methods implement the sort.Interface interface
func (self Counter) Len() int {
    return len(self.Counts)
}

func (self Counter) Less(i, j int) bool {
    return self.Counts[i].Count < self.Counts[j].Count
}

func (self Counter) Swap(i, j int) {
    iKey := self.Counts[i].Key
    jKey := self.Counts[j].Key
    self.Counts[i], self.Counts[j] = self.Counts[j], self.Counts[i]
    self.KeyIndex[iKey], self.KeyIndex[jKey] = j, i
}
