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
    if root == nil {return root}
    r := root
    for r != nil {
        curr := r
        for curr != nil {
            if curr.Left != nil && curr.Right != nil {
                curr.Left.Next = curr.Right
            }
            if curr.Right != nil && curr.Next != nil {
                if curr.Next.Left != nil  {
                    curr.Right.Next = curr.Next.Left
                } else if curr.Next.Right != nil {
                    curr.Right.Next = curr.Next.Right
                }
            }
            curr = curr.Next
        }
        if r.Left != nil {
            r = r.Left
        } else {
            r = r.Right
        }
    }
    return root
}
// func connect(root *Node) *Node {
//     if root == nil {return root}
// 	q := []*Node{root}
//     for len(q) != 0 {
//         qSize := len(q)
//         for qSize != 0 {
//             dq := q[0]
//             q = q[1:]
//             qSize--
//             if qSize > 0 && len(q) != 0 {
//                 dq.Next = q[0]
//             }
//             if dq.Left != nil {q = append(q, dq.Left)}
//             if dq.Right != nil {q = append(q, dq.Right)}
//         }
//     }
//     return root
// }