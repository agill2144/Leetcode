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
        if r == nil {return 0}
        h := 0
        for r != nil {
            r = r.Parent
            h++
        }
        return h
    }
    pH := getH(p)
    qH := getH(q)
    for pH > qH {
        p = p.Parent
        pH--
    }
    for qH > pH {
        q = q.Parent
        qH--
    }
    for p != q {
        p = p.Parent
        q = q.Parent
    }
    return p
}