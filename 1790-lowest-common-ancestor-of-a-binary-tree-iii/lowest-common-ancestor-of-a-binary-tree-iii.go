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
        if p1.Parent != nil {p1 = p1.Parent} else {p1 = q}
        if p2.Parent != nil {p2 = p2.Parent} else {p2 = p}
    }
    return p1
}

// func lowestCommonAncestor(p *Node, q *Node) *Node {
//     set := map[*Node]bool{}
//     for p != nil {
//         set[p] = true
//         p = p.Parent
//     }
//     for !set[q] {
//         q = q.Parent
//     }
//     return q
// }