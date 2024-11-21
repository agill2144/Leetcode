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
    p1, p2 := p, q
    for p1 != p2 {
        if p1 != nil {p1 = p1.Parent} else {p1 = q}
        if p2 != nil {p2 = p2.Parent} else {p2 = p}
    }
    return p1
}

// func lowestCommonAncestor(p *Node, q *Node) *Node {
//     getH := func(r *Node) int {
//         h := 0
//         for r != nil {
//             h++
//             r = r.Parent
//         }
//         return h
//     }
//     // o(h)
//     ph := getH(p)
//     // o(h)
//     qh := getH(q)
//     // o(h)
//     for ph > qh {p = p.Parent; ph--}
//     // o(h)
//     for qh > ph {q = q.Parent; qh--}
//     for p != q {
//         p = p.Parent
//         q = q.Parent
//     }
//     return p
// }