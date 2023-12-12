// to further reduce time, we only care about 2 max numbers
// maintain them in 2 vars
// time = o(n)
// space = o(1)
func maxProduct(nums []int) int {
    first := math.MinInt64
    sec := math.MinInt64
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        if num > first {
            sec = first
            first = num
        } else if num > sec {
            sec = num
        }
    }
    return (first-1) * (sec-1)
}

// want to reduce sort time, consider heap
// we only care about 2 elements, therefore heap time becomes o(nlog2) ; o(nlog)
// space = o(2)
// func maxProduct(nums []int) int {
//     min := &minHeap{items: []int{}}
//     for i := 0; i < len(nums); i++ {
//         heap.Push(min, nums[i])
//         if min.Len() > 2 {
//             heap.Pop(min)
//         }
//     }
//     return (min.items[0]-1) * (min.items[1]-1)
// }

// type minHeap struct{
//     items []int
// }
// func (m *minHeap) Less(i, j int)bool{
//     return m.items[i] < m.items[j]
// }
// func (m *minHeap) Len() int {return len(m.items)}
// func (m *minHeap) Swap(i, j int) {
//     m.items[i], m.items[j] = m.items[j], m.items[i]
// }
// func (m *minHeap) Push(x interface{}) {m.items = append(m.items, x.(int))}
// func (m *minHeap) Pop()interface{} {
//     out := m.items[len(m.items)-1]
//     m.items = m.items[:len(m.items)-1]
//     return out
// }

// we care about 2 max numbers
// time = o(nlogn)
// space = o(1)
// func maxProduct(nums []int) int {
//     sort.Ints(nums)
//     n1 := nums[len(nums)-1]
//     n2 := nums[len(nums)-2]
//     return (n1-1) * (n2-1)
// }