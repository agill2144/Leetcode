/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // k =2
 // [3,2,1]
// tc = o(n * logk)
// sc = o(k) for heap + o(n) for recursive stack
func closestKValues(root *TreeNode, target float64, k int) []int {
    if root == nil || k == 0 {return nil}
    type tmpNode struct {
        val int
        diff float64
    }
    tmp := []tmpNode{}
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        dfs(r.Left)
        tmp = append(tmp, tmpNode{r.Val, abs(target-float64(r.Val))})
        if len(tmp) > k {
            if tmp[0].diff > tmp[len(tmp)-1].diff {
                // drop the 0th element
                tmp = tmp[1:]
            } else {
                // drop the last element
                tmp = tmp[:len(tmp)-1]
            }
        }
        dfs(r.Right)
    }
    dfs(root)
    out := []int{}
    for i := 0; i < len(tmp); i++ {
        out = append(out, tmp[i].val)
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