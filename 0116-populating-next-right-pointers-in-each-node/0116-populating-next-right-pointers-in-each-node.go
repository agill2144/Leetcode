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
    if root == nil {return nil}
    l := root
    for l != nil {
        i := l
        for i != nil {
            if i.Left != nil && i.Right != nil {
                i.Left.Next = i.Right
            }
            if i.Next != nil {
                if i.Next.Left != nil {
                    if i.Right != nil {
                        i.Right.Next = i.Next.Left
                    } else if i.Left != nil {
                        i.Left.Next = i.Next.Left
                    }
                } else if i.Next.Right != nil {
                    if i.Right != nil {
                        i.Right.Next = i.Next.Right
                    } else if i.Left != nil {
                        i.Left.Next = i.Next.Right
                    }
                }
            }
            i = i.Next
        }
        if l.Left != nil {
            l = l.Left
        } else {
            l = l.Right
        }
    }
    return root
}

// approach: BFS
// time = o(n)
// space = o(maxWidthOfTree) for queue
// func connect(root *Node) *Node {
//     if root == nil {
//         return root
//     }
    
//     q := []*Node{root}
//     for len(q) != 0 {
//         qSize := len(q)
//         for qSize != 0 {
//             dq := q[0]
//             q = q[1:]
//             if qSize > 1 {
//                 dq.Next = q[0]
//             }
//             if dq.Left != nil {
//                 q = append(q, dq.Left)
//             }
//             if dq.Right != nil {
//                 q = append(q, dq.Right)
//             }
//             qSize--
//         }
//     }
//     return root
// }