/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
    mh := &minHeap{items: []int64{}}
    q := []*TreeNode{root}
    levels := 0
    
    for len(q) != 0 {
        qSize := len(q)
        var levelSum int64 = 0
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            levelSum += int64(dq.Val)
            if dq.Left != nil {
                q = append(q, dq.Left)
            }
            if dq.Right != nil {
                q = append(q, dq.Right)
            }
            qSize--
        }
        levels++
        heap.Push(mh, levelSum)
        if mh.Len() > k {
            heap.Pop(mh)
        }
    }
    if k > levels {return -1}
    return int64(mh.items[0])
}


/**
    type Interface interface{
        sort.Interface
        Pop()
        Push()
    }
**/

type minHeap struct {
    items []int64
}
func (m *minHeap) Len() int {return len(m.items)}
func (m *minHeap) Swap(i, j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
func (m *minHeap) Less(i, j int) bool { return m.items[i] < m.items[j] }
func (m *minHeap) Push(x interface{}) {m.items = append(m.items, x.(int64))}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}
