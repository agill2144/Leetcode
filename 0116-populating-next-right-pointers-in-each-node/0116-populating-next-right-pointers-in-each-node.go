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
    var dfs func(r *Node)
    dfs = func(r *Node) {
        // base
        if r == nil {return}
        if r.Left == nil && r.Right == nil {return}
        
        // logic
        r.Left.Next = r.Right
        if r.Next != nil {
            r.Right.Next = r.Next.Left
        }
        dfs(r.Left)
        dfs(r.Right)
    }
    dfs(root)
    return root
}


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