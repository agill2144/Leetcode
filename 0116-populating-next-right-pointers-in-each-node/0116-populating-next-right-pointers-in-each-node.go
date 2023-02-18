/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    r := root
    for root != nil {
        cl := root
        nl := root.Left
        for cl != nil {
            if cl.Left != nil && cl.Right != nil {
                cl.Left.Next = cl.Right
            }
            if cl.Next != nil && cl.Right != nil {
                cl.Right.Next = cl.Next.Left
            }
            cl = cl.Next
        }
        root = nl
    }
    return r
}

/*
    approach: dfs
    - at a root node, connect left->right
    - if at a node we have a next ptr, then make the cross tree connection
    time: o(n)
    space: o(h)
*/
// func connect(root *Node) *Node {
//     var dfs func(r *Node)
//     dfs = func(r *Node) {
//         // base
//         if r == nil {return}
//         if r.Left == nil && r.Right == nil {return}
        
//         // logic
//         r.Left.Next = r.Right
//         if r.Next != nil {
//             r.Right.Next = r.Next.Left
//         }
//         dfs(r.Left)
//         dfs(r.Right)
//     }
//     dfs(root)
//     return root
// }


/*
    approach: bfs
    - go level by level
    - for each dq'd node, if the top of queue is not the last node in that level, make a next ptr connection to that top node
    time: o(n)
    space: o(maxWidth) or o(n) in a skewed tree
*/
// func connect(root *Node) *Node {
//     if root == nil {return root}
//     q := []*Node{root}
//     for len(q) != 0 {
//         qSize := len(q)
//         for i := 0; i < qSize; i++ {
//             dq := q[0]
//             q = q[1:]
//             if i != qSize-1 {
//                 dq.Next = q[0]
//             }
//             if dq.Left != nil {
//                 q = append(q, dq.Left)
//             }
//             if dq.Right != nil {
//                 q = append(q, dq.Right)
//             }
//         }
//     }
//     return root
// }