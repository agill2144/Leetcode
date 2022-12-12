/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if lists == nil || len(lists) == 0 {return nil}
    mn := &minHeap{nodes: []*ListNode{}}
    for i := 0; i < len(lists); i++ {
        if lists[i] == nil {continue}
        heap.Push(mn, lists[i])
    }
    
    out := &ListNode{Val: 0}
    tail := out
    
    for mn.Len() != 0 {
        popped := heap.Pop(mn).(*ListNode)
        tail.Next = popped
        tail = tail.Next
        if popped.Next != nil {
            heap.Push(mn, popped.Next)
            popped.Next = nil
        }
    }
    return out.Next
}

type minHeap struct {
    nodes []*ListNode
}

func(m *minHeap) Push(x interface{}){
    m.nodes = append(m.nodes, x.(*ListNode))
}
func(m *minHeap) Pop() interface{} {
    out := m.nodes[len(m.nodes)-1]
    m.nodes = m.nodes[:len(m.nodes)-1]
    return out
}

func(m *minHeap) Swap(i, j int) {m.nodes[i],m.nodes[j]=m.nodes[j],m.nodes[i]}
func(m *minHeap) Len() int {return len(m.nodes)}
func(m *minHeap) Less(i, j int) bool {
    return m.nodes[i].Val < m.nodes[j].Val
}