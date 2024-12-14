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
        if p1.Parent == nil {p1 = q} else {p1 = p1.Parent}
        if p2.Parent == nil {p2 = p} else {p2 = p2.Parent}
    }
    return p1
}