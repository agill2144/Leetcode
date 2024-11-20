/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func lowestCommonAncestor(p *Node, q *Node) *Node {
    getH := func(r *Node) int {
        h := 0
        for r != nil {
            r = r.Parent
            h++
        }
        return h
    }
    ph := getH(p)
    qh := getH(q)
    for ph > qh {p = p.Parent; ph--}
    for qh > ph {q = q.Parent; qh--}
    for p != q {
        p = p.Parent
        q = q.Parent
    }
    return p
}

// func lowestCommonAncestor(p *Node, q *Node) *Node {
//     if p == nil || q == nil {return nil}
//     pPaths := []*Node{}
//     qPaths := []*Node{}
//     // h = height of tree
//     // p and q are some leaf nodes
//     // p = o(h)
//     // q = o(q)
//     // ph = height from p to root
//     // qh = height from q to root
//     // time = o(ph) + o(qh) + o(ph+qh)
//     // space = o(ph) + o(qh)
//     for p != nil {pPaths = append(pPaths, p); p = p.Parent}
//     for q != nil {qPaths = append(qPaths, q); q = q.Parent}
//     p1, p2 := len(pPaths)-1, len(qPaths)-1
//     var out *Node
//     for p1 >= 0 && p2 >= 0 {
//         if pPaths[p1] != qPaths[p2] {
//             break
//         }
//         out = pPaths[p1]
//         p1--
//         p2--
//     }
//     return out
// }