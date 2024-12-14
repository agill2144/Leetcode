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
            h++
            r = r.Parent
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