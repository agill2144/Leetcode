func topKFrequent(words []string, k int) []string {
    freq := map[string]int{}
    for i := 0; i < len(words); i++ {
        freq[words[i]]++
    }
    deduped := []string{}
    sort.Strings(deduped)
    for k, _ := range freq {deduped = append(deduped, k)}
    targetIdx := len(deduped)-k  
    left := 0
    right := len(deduped)-1
    for left <= right {
        pivot := right
        nsf := left
        for i := left; i < pivot; i++ {
            iVal := deduped[i]
            pivotVal := deduped[pivot]
            iFreq := freq[iVal]
            pivotFreq := freq[pivotVal]
            if iFreq == pivotFreq {
                //The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
                // so then shouldn't this be == -1 ?
                // no, because we want our final quick select partition to contain the smaller word!
                // our final partition is towards the end of the array, so if we move iVal to left, we have lost an answer, when this iVal < pivotVal, we want iVal to be in our partition ( i.e the right side )
                // we actually want to move lexicographically larger words towards the beginning during partitioning and smaller ones towards the end of the array 
                if strings.Compare(iVal, pivotVal) == 1 {
                    deduped[nsf], deduped[i] = deduped[i], deduped[nsf]
                    nsf++
                }
            } else if iFreq < pivotFreq {
                deduped[nsf], deduped[i] = deduped[i], deduped[nsf]
                nsf++
            }
        }
        deduped[nsf], deduped[pivot] = deduped[pivot], deduped[nsf]
        if nsf == targetIdx {break}
        if targetIdx > nsf {
            left = nsf+1
        } else {
            right = nsf-1
        }
    }
    out := deduped[targetIdx:]
    sort.Slice(out, func(i, j int)bool{
        if freq[out[i]] == freq[out[j]] {
            //The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
            return strings.Compare(out[i], out[j]) == -1
        }
        return freq[out[i]] > freq[out[j]]
    })
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
