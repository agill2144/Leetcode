/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {return nil}
    // 1-2-3-4-5
    curr := head
    for curr != nil {
        next := curr.Next
        dupe := &Node{Val: curr.Val}
        curr.Next = dupe
        dupe.Next = next
        curr = next
    }

    curr = head
    for curr != nil {
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