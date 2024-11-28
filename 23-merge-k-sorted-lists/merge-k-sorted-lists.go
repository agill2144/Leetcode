/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    mn := &minHeap{nodes: []*ListNode{}}
    for i := 0; i < len(lists); i++ {
        if lists[i] == nil {continue}
        heap.Push(mn, lists[i])
    }
    head := &ListNode{Val: 0}
    tail := head
    for mn.Len() != 0 {
        top := heap.Pop(mn).(*ListNode)
        next := top.Next
        top.Next = nil
        tail.Next = top
        tail = tail.Next
        if next != nil { heap.Push(mn, next) }
    }
    return head.Next
}


type minHeap struct {
	nodes []*ListNode
}

func (m *minHeap) Less(i, j int) bool {
	return m.nodes[i].Val < m.nodes[j].Val
}
func (m *minHeap) Swap(i, j int) {
	m.nodes[i], m.nodes[j] = m.nodes[j], m.nodes[i]
}
func (m *minHeap) Len() int {
	return len(m.nodes)
}
func (m *minHeap) Push(x interface{}) {
	m.nodes = append(m.nodes, x.(*ListNode))
}
func (m *minHeap) Pop() interface{} {
	out := m.nodes[len(m.nodes)-1]
	m.nodes = m.nodes[:len(m.nodes)-1]
	return out
}
