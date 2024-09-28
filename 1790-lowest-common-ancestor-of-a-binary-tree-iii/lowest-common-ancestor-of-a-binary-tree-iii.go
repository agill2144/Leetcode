/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

// no space solu
func lowestCommonAncestor(p *Node, q *Node) *Node {    
    root := p
    for root.Parent != nil {root = root.Parent}
    pH := -1
    qH := -1
    var getH func(r *Node, h int)
    getH = func(r *Node, h int) {
        // base
        if r == nil {return}
        if pH != -1 && qH != -1 {return }

        // logic
        if r == p {pH = h}
        if r == q {qH = h}
        getH(r.Left, h+1)
        getH(r.Right, h+1)
    }
    getH(root, 0)
    for pH > qH && p != nil {p = p.Parent; pH--}
    for qH > pH && q != nil {q = q.Parent; qH--}
    for p != q { p = p.Parent; q = q.Parent }
    return p
}