/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestKValues(root *TreeNode, target float64, k int) []int {
    if root == nil || k == 0 {return nil}   
    mx := &maxHeap{items: []*heapNode{}}
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        node := &heapNode{node:r.Val,diff:abs(target-float64(r.Val))}
        heap.Push(mx, node)
        if mx.Len() > k {
            heap.Pop(mx)
        }
        dfs(r.Left)
        dfs(r.Right)
    }
    dfs(root)
    out := []int{}
    for mx.Len() > 0 {
        out = append(out, heap.Pop(mx).(*heapNode).node)
    }
    return out
}

func abs(x float64) float64 {
    if x < 0 {return x * -1.0}
    return x
}


type heapNode struct {
    node int
    diff float64
}

type maxHeap struct {
	items []*heapNode
}

func (m *maxHeap) Less(i, j int) bool {
	return m.items[i].diff > m.items[j].diff
}
func (m *maxHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *maxHeap) Len() int {
	return len(m.items)
}
func (m *maxHeap) Push(x interface{}) {
	m.items = append(m.items, x.(*heapNode))
}
func (m *maxHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}