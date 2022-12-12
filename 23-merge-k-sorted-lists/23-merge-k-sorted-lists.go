/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
Anything kth type question == see if heap can be of the answer
if heap is the answer, then sorting can also be another answer
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {return nil}
	
    mh := &minHeap{items: []*ListNode{}}
    for i := 0; i < len(lists); i++ {
        if lists[i] != nil {
            heap.Push(mh, lists[i])
        }
    }
    
	out := &ListNode{Val: 0}
	tail := out
	for mh.Len() != 0 {
        top := heap.Pop(mh).(*ListNode)
        if top.Next != nil { heap.Push(mh, top.Next) }
        top.Next = nil
        tail.Next = top
        tail = tail.Next
	}
	return out.Next
}

type minHeap struct {
	items []*ListNode
}
func (m *minHeap) Swap(i, j int) { m.items[i],m.items[j] = m.items[j], m.items[i]}
func (m *minHeap) Less(i, j int) bool {return m.items[i].Val < m.items[j].Val}
func (m *minHeap) Len() int {return len(m.items)}
func (m *minHeap) Push(x interface{}) {m.items = append(m.items, x.(*ListNode))}
func (m *minHeap) Pop()interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
