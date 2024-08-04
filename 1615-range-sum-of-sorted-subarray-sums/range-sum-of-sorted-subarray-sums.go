func rangeSum(nums []int, n int, left int, right int) int {
    count := 0
    mn := &minHeap{items: [][]int{}}
    for i := 0; i < len(nums); i++ {
        item := []int{nums[i], i}
        heap.Push(mn, item)
    }
    tmp := []int{}
    for count < right {
        item := heap.Pop(mn).([]int) // [sum, idx]
        sum := item[0]
        idx := item[1]
        tmp = append(tmp, sum)
        count++
        if idx + 1 < len(nums) {
            heap.Push(mn, []int{sum+nums[idx+1], idx+1})
        }
    }
    sum := 0
    mod := 1000000007
    for i := left-1; i < right; i++ {sum = (sum + tmp[i]) % mod }
    return sum

}

type minHeap struct {
	items [][]int // [sum, idx]
}

func (m *minHeap) Less(i, j int) bool {
	return m.items[i][0] < m.items[j][0]
}
func (m *minHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *minHeap) Len() int {
	return len(m.items)
}
func (m *minHeap) Push(x interface{}) {
	m.items = append(m.items, x.([]int))
}
func (m *minHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}


// the brute force works....
// func rangeSum(nums []int, n int, left int, right int) int {
//     sums := []int{}
//     for i := 0; i < n; i++ {
//         s := 0
//         for j := i; j < n; j++ {
//             s+=nums[j]
//             sums = append(sums, s)
//         }
//     }
//     sort.Ints(sums)
//     ans := 0
//     mod := 1000000007
//     for i := left-1; i < right; i++ {
//         ans = (ans + sums[i]) % mod
//     }
//     return ans
// }