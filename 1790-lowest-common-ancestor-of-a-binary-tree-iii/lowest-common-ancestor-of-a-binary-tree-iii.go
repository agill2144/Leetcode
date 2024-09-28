/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

// p = number of nodes from root to p
// q = number of nodes from root to q
// time = o(p) + o(q)
// space = o(p)
func lowestCommonAncestor(p *Node, q *Node) *Node {
    
    root := p
    for root.Parent != nil {root = root.Parent}

    var getH func(r, target *Node, h int) int
    getH = func(r, target *Node, h int) int {
        // base
        if r == nil {return -1}

        // logic
        if r == target {return h}
        left := getH(r.Left, target, h+1)
        if left != -1 {return left}
        return getH(r.Right, target, h+1)
    }
    pH := getH(root, p,0)
    qH := getH(root, q,0)
    for pH > qH && p != nil {
        p = p.Parent
        pH--
    }
    for qH > pH && q != nil {
        q = q.Parent
        qH--
    }
    for p != q {
        p = p.Parent
        q = q.Parent
    }
    return p
}