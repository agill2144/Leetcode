/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
/*
    approach: no extra space
    - when being asked to do things without extra space, the choices are
    - use output 
    - or
    - use input
    - to keep track of whatever we were tracking in the extra space
    - output linkedlist does not help, so we can try to use input LL
    - 1st pass; create deep cp of each next to their orig node ( orig -> deepCp -> orig -> deepCp )
    - 2nd pass; go over the modified LL ( jump by twice, skipping deepCp nodes )
        - if a orig node has a random ptr
        - then we know the deepCp of this orig node is right next to it
        - we also know the deepCp of the random node
        - therefore curr.Next.Random = curr.Random.Next
    - Finally take out the copy nodes into its own LL and restore the original LL

    time; o(n)
    space = o(1)
*/
func copyRandomList(head *Node) *Node {
    if head == nil {return head}
    curr := head
    for curr != nil {
        next := curr.Next
        newNode := &Node{Val: curr.Val}
        curr.Next = newNode
        newNode.Next = next
        curr = next
    }

    curr = head
    for curr != nil {
        if curr.Random != nil {
            curr.Next.Random = curr.Random.Next
        }
        curr = curr.Next.Next
    }

    curr = head
    out := &Node{Val:0}
    tail := out
    for curr != nil && curr.Next != nil {
        next := curr.Next.Next
        deepCp := curr.Next
        deepCp.Next = nil
        tail.Next = deepCp
        tail = tail.Next
        curr.Next = next
        curr = next
    }
    return out.Next

}

/*
    approach: same as below with hashmap
    - 1 pass
    - check if curr node exists in our map
        - if it does not, create it, link it with our result tail
        - save it in our hashmap
    - then check if curr node has a random node
        - check if that random node already exists
        - if it does not
            - create it, save it in our hashmap
        - link our new deepCp.Random with the random node

    time = o(n)
    space = o(n)
*/
// func copyRandomList(head *Node) *Node {
//     if head == nil {return head}
//     oldToNew := map[*Node]*Node{}
//     curr := head
//     out := &Node{Val: -1}
//     tail := out
//     for curr != nil {
//         next := curr.Next
        
//         // does new node exists ?
//         // could happen if random node was created first
//         node, ok := oldToNew[curr]
//         if !ok {
//             node = &Node{Val: curr.Val}
//             oldToNew[curr] = node
//         }
//         tail.Next = node
//         tail = tail.Next

//         // does this node have a random ptr ?
//         if curr.Random != nil {
//             randomNode, ok := oldToNew[curr.Random]
//             if !ok {
//                 randomNode = &Node{Val: curr.Random.Val}
//                 oldToNew[curr.Random] = randomNode
//             }
//             node.Random = randomNode
//         }

//         curr = next
//     }
//     return out.Next
// }

/*
    approach: use hashmap
    - while creating a new linkedlist
    - keep a mapping of old node to new node in a hashmap
    - then take another pass over the input LL
    - if a node has a random ptr, then find the same node and its random ptr in the map
    - connect the random ptr

    time = o(2n)
    space = o(n)
*/
// func copyRandomList(head *Node) *Node {
//     if head == nil  {return head}
//     oldToNew := map[*Node]*Node{}
//     out := &Node{Val: -1}
//     tail := out
//     curr := head
//     for curr != nil {
//         newNode := &Node{Val: curr.Val}
//         tail.Next = newNode
//         tail = tail.Next
//         oldToNew[curr] = newNode
//         curr = curr.Next
//     }

//     curr = head
//     for curr != nil {
//         if curr.Random != nil {
//             random := curr.Random
//             deepCp := oldToNew[curr]
//             deepCpRandom := oldToNew[random]
//             deepCp.Random = deepCpRandom
//         }
//         curr = curr.Next
//     }
//     return out.Next
// }