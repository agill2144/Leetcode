/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
/*
    approach: level order using BFS
    - At each level,
    - For each node we dq, we will set the dq.Next to top queue.peek 
        - Unless its the last node of the level ( farthest right )
    - Which means we will take a queue size
    
    Potential Follow up:
    Can you wrap the right most node and connect to next level left most node?
    Basically next-right-pointer for a right most node will be the next level left most node
    - Ans: yes, take out processing queue level by level and blindly take the dq'd item from 
        our queue and connect the dq'd to q[0]th element
        - we WILL have to store the childs first before processing the dq'd item
    
    time: o(n)
    space: o(maxWidthOfTree)
*/
// func connect(root *Node) *Node {
//     if root == nil {
//         return nil
//     }
//     q := []*Node{root}
//     for len(q) != 0 {
//         qSize := len(q)
//         for i := 0; i < qSize; i++ {
//             dq := q[0]
//             q = q[1:]
//             if i < qSize-1 {
//                 dq.Next = q[0]
//             }
//             if dq.Left != nil {q = append(q, dq.Left)}
//             if dq.Right != nil {q = append(q, dq.Right)}
//         }
//     }
//     return root
// }


/*

    This approach is best for
    
    Follow-up:
    You may only use constant extra space.
    The recursive approach is not accepted ( implicit stack is not accepted )

    approach : BFS level order traversal without any extra space ( no queue )
    - Using two pointers, one pointing to start of level ( left most node )
    - and another "curr" be used to go next element in the current level ( like a fake queue )
    - We are processing the nodes level by level from left to right
        - level ptr == level by level
        - curr ptr = used to go next node in a level
    - Once we start with a root level ,
        - this parent has access to both of its left and right child
        - then parent.Left.Next = parent.Right
            1(parent)(current)(level)
           / \
          2---3
         - Once a curr ptr hits a nil, we tell our level ptr to go to next level = level.Left
                  1
                 / \
(level)(current)2---3
               /\   /\
              4  5 6  7
        - now current ptr can repeat and connect its child in its subtree ( 4-5 )
                  1
                 / \
(level)(current)2---3
               /\   /\
              4--5 6  7
        - now because curr has a next ref set, we can connect nodes accross 2 subtrees ( 5-6 )
        - how? curr.next != nil ? curr.right.next = curr.next.left : otherwise-do-nothing
                  1
                 / \
(level)(current)2---3
               /\   /\
              4--5-6  7
        - now because curr has a next ref set, we can go to next element within this level
                  1
                 / \
         (level)2---3(current)
               /\   /\
              4--5-6  7
    
        - now current ptr can repeat and connect its child in its subtree ( 6-7 )
                  1
                 / \
         (level)2---3(current)
               /\   /\
              4--5-6--7
        - now current hits nil and we go to next level
                   1
                  / \
                 2---3
                /\   /\
(level)(current)4--5-6--7
*/
func connect(root *Node) *Node {
    var curr *Node = root
    var level *Node = root
    for level != nil && level.Left != nil {
        curr = level
        for curr != nil {
            curr.Left.Next = curr.Right
            // the and is unecessary because we are given a PERFECT binary tree
            if curr.Next != nil && curr.Right != nil {
                curr.Right.Next = curr.Next.Left
            }
            curr = curr.Next
        }
        level = level.Left
    }
    return root
}


/*
    approach: DFS using top-down recursion

    This approach is best for
    
    Follow-up:
    You may only use constant extra space.
    The recursive approach is fine. You may assume implicit stack space does not count as extra space for this problem.


    - process childs from parent node itself
    - A parent can connect its two childs
    - A parent can also connect its right child if the parent has a next node set
    - So...
    - Pretend like we are on a parent node and this parent has 2 childs
    - Then parent.Left.Next = parent.Right
        1
       / \
      2---3
    - Then if Parent has a next node set, then parent.Right.Next to parent.Next.Left
    - For this, pretend like we are node 2 and 2 is the parent which already has a next connection because the previous recursion set it.
        1
       / \
      2---3
     /\   /\
    4--5-6  7
    
    time: o(n)
    space: o(h) - for recursive stack
*/

// func connect(root *Node) *Node {
//     var dfs func(r *Node)
//     dfs = func(r *Node) {
//         // base
//         if r == nil {
//             return
//         }
//         // logic
//         if r.Left != nil && r.Right != nil {
//             r.Left.Next = r.Right
//         }
//         if r.Next != nil && r.Right != nil {
//             r.Right.Next = r.Next.Left
//         }
//         dfs(r.Left)
//         dfs(r.Right)
//     }
    
//     dfs(root)
//     return root
// }