/*
    have heap create all subarrays sum
    - heap contains [sum, idx]
        - where sum is the sum of subarray ending at idx
        - does not matter where it started from
    - push 1 size subarray into a minHeap first
        - their sum will always be their own value
        - [1,2,3,4] -> [1], [2], [3], [4] are the 1 size subarrays
        - into heap -> [1,0], [2,1], [3,2], [4,3]
        - [sum, idx]
    - This is minHeap, meaning the smallest sum is at the top
    - therefore pop and process the top element
    - and increase the popped subarray size
        - popped = [1,0] ; [$sum, $idx]
        - expand this subarray by including the value of next idx 
        - new subarray sum = [3, 1] ; 
            - prevSum = 1; prevIdx = 0
            - nextIdx value = 2, nextIdx = 1
            - new pair = [$prevSum+$nextVal , $prevIdx+1]
            - therefore push this pair [3,1] back into the minHeap
    - we keep processing until we have popped and seen $right count of items
    - we can do the above ^ by maintaining a counter starting at 0
    - so keep processing heap while count < right
    - while processing heap, we can either push sums into a new tmp array
    - or check whether this item is in the range of left and right idx
        - count >= left and count <= right
        - if yes, add to our ans sum, do the mod    
    - the latter is better
    - time = o(n^2 logn)
    - space = o(n)
*/
func rangeSum(nums []int, n int, left int, right int) int {
    count := 0
    mn := &minHeap{items: [][]int{}}
    for i := 0; i < len(nums); i++ {
        item := []int{nums[i], i}
        heap.Push(mn, item)
    }
    ans := 0
    mod := 1000000007
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
        if count >= left && count <= right {ans = (ans+sum)% mod }
    }
    return ans

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