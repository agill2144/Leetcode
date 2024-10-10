func topKFrequent(words []string, k int) []string {
    freq := map[string]int{}
    for i := 0; i < len(words); i++ {
        freq[words[i]]++
    }
    mn := &minHeap{items: []*item{}}
    for word, count := range freq {
        heap.Push(mn, &item{word, count}) 
        if mn.Len() > k {
            heap.Pop(mn)
        }
    }
    out := []string{}
    for mn.Len() != 0 {
        out = append(out, heap.Pop(mn).(*item).word)
    }
    for i := 0; i < len(out)/2; i++ {
        out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
    }
    return out
}


type item struct {
    word string
    freq int
}

type minHeap struct {
	items []*item
}

func (m *minHeap) Less(i, j int) bool {
    if m.items[i].freq == m.items[j].freq {
        // The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
        return strings.Compare(m.items[i].word, m.items[j].word) == 1
    }
    return m.items[i].freq < m.items[j].freq
}
func (m *minHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *minHeap) Len() int {
	return len(m.items)
}
func (m *minHeap) Push(x interface{}) {
	m.items = append(m.items, x.(*item))
}
func (m *minHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
