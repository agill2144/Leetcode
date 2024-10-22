/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
    if root == nil {return 0}
    mn := &minHeap{items: []int64{}}    
    q := []*TreeNode{root}
    for len(q) != 0 {
        qSize := len(q) 
        var levelSum int64
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            levelSum += int64(dq.Val)
            if dq.Left != nil {q = append(q, dq.Left)}
            if dq.Right != nil {q = append(q, dq.Right)}
            qSize--
        }
        heap.Push(mn, levelSum)
        if mn.Len() > k {heap.Pop(mn)}
    }
    if mn.Len() < k {return -1}
    return mn.items[0]
}

func max(x, y int64) int64 {
    if x > y {return x}
    return y
}

type minHeap struct {
    items []int64
}

func (m *minHeap) Less(i, j int) bool {return m.items[i] < m.items[j]}
func (m *minHeap) Swap(i, j int) { m.items[i], m.items[j] = m.items[j], m.items[i]}
func (m *minHeap) Len() int {return len(m.items)}
func (m *minHeap) Push(x interface{}) {m.items = append(m.items, x.(int64))}
func (m *minHeap) Pop()interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}
