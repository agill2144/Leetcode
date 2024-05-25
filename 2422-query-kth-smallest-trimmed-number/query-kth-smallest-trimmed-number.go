
func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
    n := len(nums)
    out := []int{}

    for _, query := range queries {
        kth := query[0]
        keepLast := query[1]
        mh := &maxHeap{}
        heap.Init(mh)

        for j := 0; j < n; j++ {
            num := nums[j]
            subStr := num[len(num)-keepLast:]
            heap.Push(mh, &item{value: subStr, index: j})
            if mh.Len() > kth {
                heap.Pop(mh)
            }
        }

        result := heap.Pop(mh).(*item)
        out = append(out, result.index)
    }

    return out
}

type item struct {
    value string
    index int
}

type maxHeap []*item

func (m maxHeap) Len() int           { return len(m) }
func (m maxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m maxHeap) Less(i, j int) bool {
    if m[i].value == m[j].value {
        return m[i].index > m[j].index
    }
    return m[i].value > m[j].value
}

func (m *maxHeap) Push(x interface{}) {
    *m = append(*m, x.(*item))
}

func (m *maxHeap) Pop() interface{} {
    old := *m
    n := len(old)
    x := old[n-1]
    *m = old[0 : n-1]
    return x
}
