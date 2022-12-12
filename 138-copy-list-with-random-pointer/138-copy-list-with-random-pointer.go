/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    
    if head == nil {
        return head
    }
    
    curr := head
    for curr != nil {
        next := curr.Next
        newNode := &Node{Val: curr.Val}
        curr.Next = newNode
        newNode.Next = next
        curr = next
    }
    
    
    curr = head
    for curr != nil && curr.Next != nil {
        next := curr.Next.Next
        if curr.Random != nil {
            curr.Next.Random = curr.Random.Next
        }
        curr = next
    }
    
    out := &Node{Val: 0}
    curr = head
    tail := out
    for curr != nil && curr.Next != nil {
        next := curr.Next.Next
        deepCp := curr.Next
        deepCp.Next = nil
        curr.Next = next
        tail.Next = deepCp
        tail = tail.Next
        curr = next
    }
    return out.Next
}