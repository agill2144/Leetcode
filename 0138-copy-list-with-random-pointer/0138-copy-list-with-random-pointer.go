/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

// time: o(n)
// space: o(1)
func copyRandomList(head *Node) *Node {
    if head == nil {return nil}
    // 1-2-3-4-5
    deepCpMap := map[*Node]*Node{}
    deepCp := &Node{Val: 0}
    tail := deepCp
    curr := head
    for curr != nil {
        next := curr.Next
        
        
        newNode, ok := deepCpMap[curr]
        if !ok {
            newNode = &Node{Val: curr.Val}
            deepCpMap[curr] = newNode
        }
        
        if curr.Random != nil {
            randomNode, ok := deepCpMap[curr.Random]
            if !ok {
                randomNode = &Node{Val: curr.Random.Val}
                deepCpMap[curr.Random] = randomNode
            }
            newNode.Random = randomNode
        }

        tail.Next = newNode
        tail = tail.Next
        curr = next
    }
    return deepCp.Next
}