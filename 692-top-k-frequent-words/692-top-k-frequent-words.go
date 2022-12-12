
import "container/heap"

type element struct {
	val         string
	occurrences int
}
type maxHeap struct {
	items []*element
}
func (m *maxHeap) Less(i, j int) bool {
	if m.items[i].occurrences == m.items[j].occurrences {
		return m.items[i].val < m.items[j].val
	}
	return m.items[i].occurrences > m.items[j].occurrences
}
func (m *maxHeap) Len() int {return len(m.items)}
func (m *maxHeap) Swap(i, j int) {m.items[i],m.items[j] = m.items[j], m.items[i]}
func (m *maxHeap) Push(x interface{}) {m.items = append(m.items, x.(*element))}
func (m *maxHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}


func topKFrequent(words []string, k int) []string {
	if words == nil || len(words) == 0 {
		return nil
	}
	seen := map[string]int{}
	for _, e := range words {
		numTimes, ok := seen[e]
		if !ok {
			seen[e] = 1
			continue
		}
		seen[e] = numTimes+1
	}
	mx := &maxHeap{items: []*element{}}
	for w, times := range seen {
		i := &element{val: w,occurrences: times}
		heap.Push(mx, i)
	}
	out := []string{}
	for i := 0; i < k; i++ {
		out = append(out, heap.Pop(mx).(*element).val)
	}
	return out
}

