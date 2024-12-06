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
        curr.Next = &Node{Val: curr.Val}
        curr.Next.Next = next
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
    dummy := &Node{Val: 0}
    tail := dummy 
    curr = head
    for curr != nil {
        next := curr.Next.Next
        r := curr.Next
        r.Next = nil
        curr.Next = next
        tail.Next = r
        tail = tail.Next
        curr = next
    }
    return dummy.Next
}

// func copyRandomList(head *Node) *Node {
//     if head == nil {return nil}
//     m := map[*Node]*Node{}
//     dummy := &Node{Val:0}
//     tail := dummy
//     curr := head
//     for curr != nil {
//         next := curr.Next
//         node := m[curr]
//         if node == nil {
//             node = &Node{Val: curr.Val}
//             m[curr] = node
//         }
//         tail.Next = node
//         tail = tail.Next
//         if curr.Random != nil {
//             randCp := m[curr.Random]
//             if randCp == nil {
//                 randCp = &Node{Val:curr.Random.Val}
//                 m[curr.Random] = randCp
//             }
//             node.Random = randCp
//         }
//         curr = next
//     }
//     return dummy.Next
// }