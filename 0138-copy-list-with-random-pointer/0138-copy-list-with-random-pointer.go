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
    srcToCopy := map[*Node]*Node{}
    out := &Node{Val:0}
    tail := out
    curr := head
    for curr != nil {
        newNode := &Node{Val: curr.Val}
        srcToCopy[curr] = newNode
        tail.Next = newNode
        tail = tail.Next
        curr = curr.Next
    }
    
    curr = head
    for curr != nil {
        if curr.Random != nil {
            src := srcToCopy[curr]
            target := srcToCopy[curr.Random]
            src.Random = target
        }
        curr = curr.Next
    }
    return out.Next
    
}