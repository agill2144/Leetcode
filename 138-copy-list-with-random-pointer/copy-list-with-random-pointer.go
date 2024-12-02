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
    origToCp := map[*Node]*Node{}
    dummy := &Node{Val: head.Val}
    tail := dummy
    curr := head
    for curr != nil {
        next := curr.Next
        node := origToCp[curr]
        if node == nil {
            node = &Node{Val: curr.Val}
            origToCp[curr] = node
        }
        tail.Next = node
        tail = tail.Next
        if curr.Random != nil {
            randCp := origToCp[curr.Random]
            if randCp == nil {
                randCp = &Node{Val: curr.Random.Val}
                origToCp[curr.Random] = randCp
            }
            node.Random = randCp
        }
        curr = next
    }
    return dummy.Next
}