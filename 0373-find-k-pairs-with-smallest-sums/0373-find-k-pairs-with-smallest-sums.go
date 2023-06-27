func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    pq := &minHeap{items: [][]int{ {nums1[0]+nums2[0], 0,0 } }}
    
    out := [][]int{}
    visited := map[[2]int]struct{}{[2]int{0,0}: struct{}{}}
    for pq.Len() != 0 && len(out) < k {
        top := heap.Pop(pq).([]int)
        i := top[1];
        j := top[2];
        out = append(out, []int{nums1[i], nums2[j]})
        if i+1 < len(nums1) {
            if _, ok := visited[[2]int{i+1, j}]; !ok {
                heap.Push(pq, []int{nums1[i+1]+nums2[j], i+1,j})
                visited[[2]int{i+1, j}] = struct{}{}
            }
        }
        if j+1 < len(nums2) {
            if _, ok := visited[[2]int{i,j+1}]; !ok {
                heap.Push(pq, []int{nums1[i]+nums2[j+1], i,j+1})
                visited[[2]int{i,j+1}] = struct{}{}
            }
        }
        
    }
    return out
}


/*
    pkg: container/heap
    
    type Interface interface {
        sort.Interface
        Push(x T)
        Pop() T
    }
*/
type minHeap struct {
    items  [][]int// [ [sum, num1, num2], [sum, num1, num2] ]
}
func (m *minHeap) Len() int {return len(m.items)}
func (m *minHeap) Swap(i, j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
func (m *minHeap) Less(i, j int) bool {return m.items[i][0] < m.items[j][0] }
func (m *minHeap) Push(x interface{}) {m.items = append(m.items, x.([]int))}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}

