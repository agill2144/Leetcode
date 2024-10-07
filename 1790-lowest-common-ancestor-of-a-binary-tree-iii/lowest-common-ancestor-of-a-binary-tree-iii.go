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
    if p == nil || q == nil {return nil}
    getH := func(t *Node) int {
        ch := 0
        for t != nil {
            t = t.Parent
            ch++
        }
        return ch 
    }
    pH := getH(p)
    qH := getH(q)
    for qH > pH { 
        q = q.Parent
        qH--
    }
    for pH > qH {
        p = p.Parent
        pH--
    }
    for p != q {
        p = p.Parent
        q = q.Parent
    }
    return p
}