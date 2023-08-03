/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
    curr := root
    for curr != nil {
        next := curr.Next
        if curr.Child != nil {
            childHead := flatten(curr.Child)
            curr.Next = childHead
            childHead.Prev = curr
            curr.Child = nil
            for childHead.Next != nil {
                childHead = childHead.Next
            }
            // now childHead is at the tail end of the LL
            if next != nil {
                childHead.Next = next
                next.Prev = childHead
            }
        }
        curr = next
    }
    return root
}