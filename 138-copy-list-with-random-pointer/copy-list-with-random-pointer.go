/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
/*
    approach: no extra space, in-place copy, with 3 passes
    #1
    - while traversing the LL
    - create a copy of current node
    - and make it part of the current LL 
        - insert the copy in between curr and next node
    #2
    - connect the orig random with their dupes
        - random dupes are at random.next
    
    #3
    - extract out the duplicated LL
*/
func copyRandomList(head *Node) *Node {
    if head == nil {return nil}

    curr := head
    for curr != nil {
        next := curr.Next
        cp := &Node{Val: curr.Val}
        curr.Next = cp
        cp.Next = next
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

    curr = head
    // |---|
    // 1 1 2 2-3-3-4-4-5-5
    //   |___|
    // c   n
    dummy := &Node{Val: 0}
    tail := dummy
    for curr != nil {
        next := curr.Next.Next
        dupe := curr.Next
        curr.Next = next
        dupe.Next = nil
        tail.Next = dupe
        tail = tail.Next
        curr = next        
    }
    return dummy.Next

}
/*
    approach: hashmap
    - we can take 2 passes
        1. create a copy first and create a mapping between original to copy node
        2. then take another pass to create the random ptrs
    tc = o(2n)
    sc = o(n)

    - we can do this in 1 pass too
        - while traversing, make copies of both nodes
        - curr and random
    tc = o(n)
    sc = o(n)

*/
// func copyRandomList(head *Node) *Node {
//     if head == nil {return head}
//     origToCp := map[*Node]*Node{}
//     dummy := &Node{Val: head.Val}
//     tail := dummy
//     curr := head
//     for curr != nil {
//         next := curr.Next
//         node := origToCp[curr]
//         if node == nil {
//             node = &Node{Val: curr.Val}
//             origToCp[curr] = node
//         }
//         tail.Next = node
//         tail = tail.Next
//         if curr.Random != nil {
//             randCp := origToCp[curr.Random]
//             if randCp == nil {
//                 randCp = &Node{Val: curr.Random.Val}
//                 origToCp[curr.Random] = randCp
//             }
//             node.Random = randCp
//         }
//         curr = next
//     }
//     return dummy.Next
// }