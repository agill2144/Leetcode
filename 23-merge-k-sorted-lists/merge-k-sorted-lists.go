/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    mn := &minHeap{items: []*ListNode{}}
    for i := 0; i < len(lists); i++ {
        if lists[i] == nil {continue}
        heap.Push(mn, lists[i])
    }
    head := &ListNode{Val: 0}
    tail := head
    for mn.Len() != 0 {
        top := heap.Pop(mn).(*ListNode)
        if top.Next != nil {heap.Push(mn, top.Next); top.Next = nil}
        tail.Next = top
        tail = tail.Next
    }
    return head.Next
}


type minHeap struct {
	items []*ListNode
}

func (m *minHeap) Less(i, j int) bool {
	return m.items[i].Val < m.items[j].Val
}
func (m *minHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *minHeap) Len() int {
	return len(m.items)
}
func (m *minHeap) Push(x interface{}) {
	m.items = append(m.items, x.(*ListNode))
}
func (m *minHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
