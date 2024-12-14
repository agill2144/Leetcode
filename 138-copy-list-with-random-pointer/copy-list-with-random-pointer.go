/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {return head}
    curr := head
    for curr != nil {
        next := curr.Next
        newNode := &Node{Val:curr.Val}
        curr.Next = newNode
        newNode.Next = next
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
    dummy := &Node{Val:0}
    tail := dummy
    curr = head
    for curr != nil {
        next := curr.Next.Next
        tail.Next = curr.Next
        tail = tail.Next
        curr.Next = next
        curr = next
    }
    return dummy.Next
}