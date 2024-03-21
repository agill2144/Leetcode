/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

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
func copyRandomList(head *Node) *Node {
    if head == nil  {return head}
    oldToNew := map[*Node]*Node{}
    out := &Node{Val: -1}
    tail := out
    curr := head
    for curr != nil {
        newNode := &Node{Val: curr.Val}
        tail.Next = newNode
        tail = tail.Next
        oldToNew[curr] = newNode
        curr = curr.Next
    }

    curr = head
    for curr != nil {
        if curr.Random != nil {
            random := curr.Random
            deepCp := oldToNew[curr]
            deepCpRandom := oldToNew[random]
            deepCp.Random = deepCpRandom
        }
        curr = curr.Next
    }
    return out.Next
}